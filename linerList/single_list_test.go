package linerList

import (
	"testing"
	"math/rand"
	"time"
)

func TestSingleList_Insert(t *testing.T) {
	l := InitSingleList()
	rand.Seed(time.Now().Unix())
	for i:=0;i<100000 ;i++{
		l.Insert(i,rand.Int())
	}
}