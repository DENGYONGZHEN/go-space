package log

import (
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

var (
	offWidth uint64 = 4                   //offset 占4个字节
	posWidth uint64 = 8                   // position 占8个字节
	entWidth        = offWidth + posWidth //每条索引记录占12个字节
)

// 对单个索引文件的操作接口
type index struct {
	file *os.File    //索引文件，跟segment(日志段) 一一对应
	mmap gommap.MMap //映射后的内存区域，读写index不需要file read/write，直接操作内存
	size uint64      // 当前索引文件实际使用字节数
}

func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{
		file: f,
	}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size()) //如果是旧文件，size可能>0,如果是新文件，size =0

	//扩容文件到可以容纳最大索引的大小
	if err = os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}

	//然后使用mmap映射文件到内存。如果上面不扩容，就不能映射到足够大的内存，内存与文件大小一致
	if idx.mmap, err = gommap.Map(idx.file.Fd(), gommap.PROT_READ|gommap.PROT_WRITE, gommap.MAP_SHARED); err != nil {
		return nil, err
	}

	return idx, nil

}

func (i *index) Close() error {

	//确保内存中的内容写入磁盘
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	//确保缓冲区的内容写入磁盘
	if err := i.file.Sync(); err != nil {
		return err
	}
	//截取文件到实际占用的大小
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}
	if in == -1 { //都最后一条entry
		out = uint32((i.size / entWidth) - 1)
	} else {
		// 数组下标式索引方式（zero-based indexing）
		out = uint32(in)
	}
	pos = uint64(out) * entWidth
	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}
	out = enc.Uint32(i.mmap[pos : pos+offWidth])
	pos = enc.Uint64(i.mmap[pos+offWidth : pos+entWidth])
	return out, pos, nil
}

func (i *index) Write(off uint32, pos uint64) error {

	//确保mmap足够大(文件不会溢出)
	if uint64(len(i.mmap)) < i.size+entWidth {
		return io.EOF
	}

	//把新的entry写在当前size后面
	// 这是一个 切片操作，获取 mmap 中将要写入的位置
	// 切片在 Go 中是 引用类型（slice header 指向底层数组）
	// 所以对切片的修改会直接作用在 mmap 底层的内存，也就是文件映射区域
	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += uint64(entWidth)
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
