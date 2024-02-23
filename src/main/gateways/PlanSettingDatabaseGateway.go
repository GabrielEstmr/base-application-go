package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type PlanSettingDatabaseGateway interface {
	Save(ctx context.Context, planSetting main_domains.PlanSetting) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException)
	FindById(ctx context.Context, id string) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException)
	FindByPlanTypeAndHasEndDate(ctx context.Context, planType main_domains_enums.PlanType, hasEndDate bool) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException)
	Update(ctx context.Context, planSetting main_domains.PlanSetting) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException)
	FindByFilter(ctx context.Context, filter main_domains.FindPlanSettingFilter, pageable main_domains.Pageable) (
		main_domains.Page, main_domains_exceptions.ApplicationException)
}
