package chapter3

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	t.Logf("bond to %q", listener.Addr())
}

//%v	默认格式				Printf("%v", 42) → 42
//%T	值类型					Printf("%T", 42) → int
//%%	字面百分号（不替换值）	 Printf("%%") → %
//%d	十进制整数				Printf("%d", 42) → 42
//%b	二进制					Printf("%b", 5) → 101
//%o	八进制					Printf("%o", 8) → 10
//%x	十六进制（小写）		 Printf("%x", 15) → f
//%X	十六进制（大写）		 Printf("%X", 15) → F
//%c	Unicode 字符			Printf("%c", 65) → A
//%U	Unicode 格式			Printf("%U", 'A') → U+0041
//%t	true/false				Printf("%t", true) → true
//%s	原始字符串				 Printf("%s", "Go") → Go
//%q	带双引号的转义字符串	  Printf("%q", "Go") → "Go"
//%x	十六进制（小写，每字节两字符）	Printf("%x", "Go") → 476f
//%X	十六进制（大写）	      Printf("%X", "Go") → 476F
//%p	指针地址（十六进制）	  Printf("%p", &x) → 0xc0000...
//%v	默认格式（字段值）	      Printf("%v", struct{X int}{1}) → {1}
//%+v	带字段名的默认格式	      Printf("%+v", struct{X int}{1}) → {X:1}
//%#v	Go 语法表示（可复制到代码）	Printf("%#v", struct{X int}{1}) → main.Struct{X:1}
