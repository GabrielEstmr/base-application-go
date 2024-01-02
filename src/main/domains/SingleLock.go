package main_domains

import (
	main_gateways_lock_resources "baseapplicationgo/main/gateways/lock/resources"
)

type SingleLock struct {
	lock main_gateways_lock_resources.SingleLockResource
}

func NewSingleLock(lockResource main_gateways_lock_resources.SingleLockResource) *SingleLock {
	return &SingleLock{lock: lockResource}
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
