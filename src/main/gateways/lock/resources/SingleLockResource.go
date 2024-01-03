package main_gateways_lock_resources

import (
	"baseapplicationgo/main/domains/lock"
	"github.com/go-redsync/redsync/v4"
)

type SingleLockResource struct {
	lock *redsync.Mutex
}

func NewSingleLockResource(lock *redsync.Mutex) *SingleLockResource {
	return &SingleLockResource{lock: lock}
}

func (this *SingleLockResource) Lock() error {
	return this.lock.Lock()
}

func (this *SingleLockResource) Extend() (bool, error) {
	return this.lock.Extend()
}

func (this *SingleLockResource) Unlock() (bool, error) {
	return this.lock.Unlock()
}

func (this *SingleLockResource) ToDomain() lock.SingleLock {
	return *lock.NewSingleLock(this.lock)
}
