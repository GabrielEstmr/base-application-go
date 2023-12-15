package main_gateways_mongodb

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	"fmt"
	"log/slog"
)

type CachedUserDatabaseGatewayImpl struct {
	userDatabaseGateway      main_gateways.UserDatabaseGateway
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway
	apLog                    *slog.Logger
}

func NewCachedUserDatabaseGatewayImpl(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway) *CachedUserDatabaseGatewayImpl {
	return &CachedUserDatabaseGatewayImpl{
		userDatabaseGateway:      userDatabaseGateway,
		userDatabaseCacheGateway: userDatabaseCacheGateway,
		apLog:                    main_configs_logs.GetLogConfigBean(),
	}
}

func (this *CachedUserDatabaseGatewayImpl) Save(user main_domains.User) (main_domains.User, error) {
	user, err := this.userDatabaseGateway.Save(user)
	if err != nil {
		return main_domains.User{}, err
	}
	go func() {
		_, err := this.userDatabaseCacheGateway.Save(user)
		if err != nil {
			this.apLog.Error(fmt.Sprintf("Error to save document into Redis. Document: User, Id: %s", user.Id))
		}
	}()
	return user, nil
}

func (this *CachedUserDatabaseGatewayImpl) FindById(id string) (main_domains.User, error) {
	cachedUser, err := this.userDatabaseCacheGateway.FindById(id)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindById(id)
	if err != nil {
		return main_domains.User{}, err
	}

	if !user.IsEmpty() {
		go func() {
			_, err := this.userDatabaseCacheGateway.Save(user)
			if err != nil {
				this.apLog.Error("Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	cachedUser, err := this.userDatabaseCacheGateway.FindByDocumentNumber(documentNumber)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindByDocumentNumber(documentNumber)
	if err != nil {
		return main_domains.User{}, err
	}

	if !user.IsEmpty() {
		go func() {
			_, err := this.userDatabaseCacheGateway.Save(user)
			if err != nil {
				this.apLog.Error("Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByFilter(
	filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	return this.userDatabaseGateway.FindByFilter(filter, pageable)
}
