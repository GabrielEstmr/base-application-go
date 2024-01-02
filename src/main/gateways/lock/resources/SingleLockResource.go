package main_gateways_lock_resources

import (
	main_domains "baseapplicationgo/main/domains"
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

func (this *SingleLockResource) ToDomain() main_domains.SingleLock {
	return *main_domains.NewSingleLock(*this)
}
