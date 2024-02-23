package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"fmt"
)

type CachedUserDatabaseGatewayImpl struct {
	userDatabaseGateway      main_gateways.UserDatabaseGateway
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway
	spanGateway              main_gateways.SpanGateway
	logsMonitoringGateway    main_gateways.LogsMonitoringGateway
}

func NewCachedUserDatabaseGatewayImpl(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway) *CachedUserDatabaseGatewayImpl {
	return &CachedUserDatabaseGatewayImpl{
		userDatabaseGateway:      userDatabaseGateway,
		userDatabaseCacheGateway: userDatabaseCacheGateway,
		spanGateway:              main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway:    main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *CachedUserDatabaseGatewayImpl) Save(
	ctx context.Context,
	user main_domains.User,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-Save")
	defer span.End()

	user, err := this.userDatabaseGateway.Save(span.GetCtx(), user, options)
	if err != nil {
		return *new(main_domains.User), err
	}
	go func() {
		_, errC := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
		if errC != nil {
			this.logsMonitoringGateway.ERROR(
				span,
				fmt.Sprintf("Error to save document into Redis. Document: User, Id: %s", user.GetId()))
		}
	}()
	return user, nil
}

func (this *CachedUserDatabaseGatewayImpl) Update(
	ctx context.Context,
	user main_domains.User,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-Save")
	defer span.End()

	user, err := this.userDatabaseGateway.Update(span.GetCtx(), user, options)
	if err != nil {
		return *new(main_domains.User), err
	}
	go func() {
		_, errC := this.userDatabaseCacheGateway.Update(span.GetCtx(), user)
		if errC != nil {
			this.logsMonitoringGateway.ERROR(
				span,
				fmt.Sprintf("Error to update document into Redis. Document: User, Id: %s", user.GetId()))
		}
	}()
	return user, nil
}

func (this *CachedUserDatabaseGatewayImpl) FindById(
	ctx context.Context,
	id string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindById")
	defer span.End()

	cachedUser, err := this.userDatabaseCacheGateway.FindById(span.GetCtx(), id)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindById(span.GetCtx(), id, options)
	if err != nil {
		return *new(main_domains.User), err
	}

	if !user.IsEmpty() {
		go func() {
			_, err := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
			if err != nil {
				this.logsMonitoringGateway.ERROR(span, "Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByDocumentId(
	ctx context.Context,
	documentNumber string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindByDocumentNumber")
	defer span.End()

	cachedUser, err := this.userDatabaseCacheGateway.FindByDocumentId(span.GetCtx(), documentNumber)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindByDocumentId(span.GetCtx(), documentNumber, options)
	if err != nil {
		return *new(main_domains.User), err
	}

	if !user.IsEmpty() {
		go func() {
			_, errSC := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
			if errSC != nil {
				this.logsMonitoringGateway.ERROR(span, "Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByUserName(
	ctx context.Context,
	userName string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindByDocumentNumber")
	defer span.End()

	cachedUser, err := this.userDatabaseCacheGateway.FindByUserName(span.GetCtx(), userName)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindByUserName(span.GetCtx(), userName, options)
	if err != nil {
		return *new(main_domains.User), err
	}

	if !user.IsEmpty() {
		go func() {
			_, errSC := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
			if errSC != nil {
				this.logsMonitoringGateway.ERROR(span, "Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByEmail(
	ctx context.Context,
	email string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindByDocumentNumber")
	defer span.End()

	cachedUser, err := this.userDatabaseCacheGateway.FindByEmail(span.GetCtx(), email)
	if !cachedUser.IsEmpty() && err == nil {
		return cachedUser, nil
	}
	user, err := this.userDatabaseGateway.FindByEmail(span.GetCtx(), email, options)
	if err != nil {
		return *new(main_domains.User), err
	}

	if !user.IsEmpty() {
		go func() {
			_, errSC := this.userDatabaseCacheGateway.Save(span.GetCtx(), user)
			if errSC != nil {
				this.logsMonitoringGateway.ERROR(span, "Error to save in Redis")
			}
		}()
	}
	return user, err
}

func (this *CachedUserDatabaseGatewayImpl) FindByFilter(
	ctx context.Context,
	filter main_domains.FindUserFilter,
	pageable main_domains.Pageable,
	options main_domains.DatabaseOptions,
) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CachedUserDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	return this.userDatabaseGateway.FindByFilter(span.GetCtx(), filter, pageable, options)
}
