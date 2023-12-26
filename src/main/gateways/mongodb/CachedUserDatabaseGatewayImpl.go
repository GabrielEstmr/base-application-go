package main_gateways_mongodb

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"fmt"
	"log/slog"
)

type CachedUserDatabaseGatewayImpl struct {
	userDatabaseGateway      main_gateways.UserDatabaseGateway
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway
	spanGateway              main_gateways.SpanGateway
	apLog                    *slog.Logger
}

func NewCachedUserDatabaseGatewayImpl(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway) *CachedUserDatabaseGatewayImpl {
	return &CachedUserDatabaseGatewayImpl{
		userDatabaseGateway:      userDatabaseGateway,
		userDatabaseCacheGateway: userDatabaseCacheGateway,
		spanGateway:              main_gateways_spans.NewSpanGatewayImpl(),
		apLog:                    main_configs_logs.GetLogConfigBean(),
	}
}

func (this *CachedUserDatabaseGatewayImpl) Save(ctx context.Context, user main_domains.User) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-Save")
	defer span.End()

	user, err := this.userDatabaseGateway.Save(span.GetCtx(), user)
	if err != nil {
		return main_domains.User{}, err
	}
	go func() {
		_, err := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
		if err != nil {
			this.apLog.Error(fmt.Sprintf("Error to save document into Redis. Document: User, Id: %s", user.Id))
		}
	}()
	return user, nil
}

func (this *CachedUserDatabaseGatewayImpl) FindById(ctx context.Context, id string) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindById")
	defer span.End()

	cachedUser, err := this.userDatabaseCacheGateway.FindById(span.GetCtx(), id)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindById(span.GetCtx(), id)
	if err != nil {
		return main_domains.User{}, err
	}

	if !user.IsEmpty() {
		go func() {
			_, err := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
			if err != nil {
				this.apLog.Error("Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindByDocumentNumber")
	defer span.End()

	cachedUser, err := this.userDatabaseCacheGateway.FindByDocumentNumber(span.GetCtx(), documentNumber)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindByDocumentNumber(span.GetCtx(), documentNumber)
	if err != nil {
		return main_domains.User{}, err
	}

	if !user.IsEmpty() {
		go func() {
			_, err := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
			if err != nil {
				this.apLog.Error("Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByFilter(ctx context.Context,
	filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	return this.userDatabaseGateway.FindByFilter(span.GetCtx(), filter, pageable)
}
