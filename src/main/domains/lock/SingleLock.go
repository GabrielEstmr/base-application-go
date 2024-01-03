package lock

import (
	"github.com/go-redsync/redsync/v4"
)

type SingleLock struct {
	lock *redsync.Mutex
}

func NewSingleLock(lock *redsync.Mutex) *SingleLock {
	return &SingleLock{lock: lock}
}

func (this *SingleLock) Lock() error {
	return this.lock.Lock()
}

func (this *SingleLock) Extend() (bool, error) {
	return this.lock.Extend()
}

func (this *SingleLock) Unlock() (bool, error) {
	return this.lock.Unlock()
}
