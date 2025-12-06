package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/deng/go-space/networkProgrammingWithGo/chapter12/housework/v1"
	"github.com/deng/go-space/networkProgrammingWithGo/chapter12/protobuf"
)

type Rosie struct {
	dataFile string
	mu       sync.Mutex
	chores   []*housework.Chore
	housework.UnimplementedRobotMaidServer
}

func NewRosie(dataFile string) *Rosie {
	chores, err := load(dataFile)
	if err != nil {
		chores = make([]*housework.Chore, 0)
	}
	return &Rosie{
		dataFile: dataFile,
		chores:   chores,
	}
}

func load(dataFile string) ([]*housework.Chore, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return make([]*housework.Chore, 0), nil
	}

	df, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := df.Close(); err != nil {
			fmt.Printf("closing data file: %v", err)
		}
	}()

	return protobuf.Load(df)
}

func flush(dataFile string, chores []*housework.Chore) error {
	df, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := df.Close(); err != nil {
			fmt.Printf("closing data file: %v", err)
		}
	}()
	return protobuf.Flush(df, chores)
}

func (r *Rosie) Add(_ context.Context, chores *housework.Chores) (*housework.Response, error) {
	r.mu.Lock()
	r.chores = append(r.chores, chores.Chores...)
	flush(r.dataFile, r.chores)
	r.mu.Unlock()

	return &housework.Response{Message: "ok"}, nil
}

func (r *Rosie) Complete(_ context.Context, req *housework.CompleteRequest) (*housework.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.chores == nil || req.ChoreNumber < 1 || int(req.ChoreNumber) > len(r.chores) {
		return nil, fmt.Errorf("chore %d not found", req.ChoreNumber)
	}
	r.chores[req.ChoreNumber-1].Complete = true
	flush(r.dataFile, r.chores)
	return &housework.Response{Message: "ok"}, nil
}

func (r *Rosie) List(_ context.Context, _ *housework.Empty) (*housework.Chores, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.chores == nil {
		r.chores = make([]*housework.Chore, 0)
	}

	return &housework.Chores{Chores: r.chores}, nil
}
