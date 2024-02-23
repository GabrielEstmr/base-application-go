package main_usecases_beans

import (
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_beans_factories "baseapplicationgo/main/usecases/beans/Factories"
	main_usecases_factories "baseapplicationgo/main/usecases/factories"
	main_usecases_lockers "baseapplicationgo/main/usecases/lockers"
	"sync"
)

var once sync.Once

var usecaseBeans *UsecaseBeans

type UsecaseBeans struct {
	CreateNewUser                                           *main_usecases.CreateNewUser
	FindUserById                                            *main_usecases.FindUserById
	FindUsersByFilter                                       *main_usecases.FindUsersByFilter
	CreateTransactionAmqpEvent                              *main_usecases.CreateTransactionAmqpEvent
	PersistTransaction                                      *main_usecases.PersistTransaction
	CreateEmail                                             *main_usecases.CreateEmail
	CreateEmailBodyFactory                                  *main_usecases_factories.CreateEmailBodyFactory
	SendEmailGatewayFactory                                 *main_usecases_factories.SendEmailGatewayFactory
	CreateEmailBodyForWelcomeEmail                          *main_usecases.CreateEmailBodyForWelcomeEmail
	FindEmailsByFilter                                      *main_usecases.FindEmailsByFilter
	CreateEmailBodySendAndPersistAsSent                     *main_usecases.CreateEmailBodySendAndPersistAsSent
	ReprocessEmailEvent                                     *main_usecases.ReprocessEmailEvent
	SendEmailEventsToReprocess                              *main_usecases.SendEmailEventsToReprocess
	CreateEmailFallback                                     *main_usecases.CreateEmailFallback
	EnableInternalProviderUser                              *main_usecases.EnableInternalProviderUser
	CreateSessionByIdentityProvider                         *main_usecases.CreateSessionByIdentityProvider
	EnableExternalProviderUser                              *main_usecases.EnableExternalProviderUser
	CreateInternalAuthUserPasswordChangeRequest             *main_usecases.CreateInternalAuthUserPasswordChangeRequest
	ChangeInternalProviderUserPassword                      *main_usecases.ChangeInternalProviderUserPassword
	CreateInternalProviderUserSession                       *main_usecases.CreateInternalProviderUserSession
	RefreshSession                                          *main_usecases.RefreshSession
	EndSession                                              *main_usecases.EndSession
	AtomicLockedCreateNewUser                               *main_usecases_lockers.AtomicLockedCreateNewUser
	AtomicLockedEnableExternalProviderUser                  *main_usecases_lockers.AtomicLockedEnableExternalProviderUser
	AtomicLockedEnableInternalProviderUser                  *main_usecases_lockers.AtomicLockedEnableInternalProviderUser
	AtomicLockedCreateInternalAuthUserPasswordChangeRequest *main_usecases_lockers.AtomicLockedCreateInternalAuthUserPasswordChangeRequest
	AtomicLockedChangeInternalProviderUserPassword          *main_usecases_lockers.AtomicLockedChangeInternalProviderUserPassword
	AtomicLockedCreateExternalAuthProviderUserSession       *main_usecases_lockers.AtomicLockedCreateExternalAuthProviderUserSession
}

func GetUsecaseBeans() *UsecaseBeans {
	once.Do(func() {
		if usecaseBeans == nil {
			usecaseBeans = subscriptUsecaseBeans()
		}
	})
	return usecaseBeans
}

func subscriptUsecaseBeans() *UsecaseBeans {
	return &UsecaseBeans{
		CreateNewUser:                                           main_usecases_beans_factories.NewCreateNewUserBean().Get(),
		FindUserById:                                            main_usecases_beans_factories.NewFindUserByIdBean().Get(),
		FindUsersByFilter:                                       main_usecases_beans_factories.NewFindUsersByFilterBean().Get(),
		CreateTransactionAmqpEvent:                              main_usecases_beans_factories.NewCreateTransactionAmqpEventBean().Get(),
		PersistTransaction:                                      main_usecases_beans_factories.NewPersistTransactionBean().Get(),
		CreateEmail:                                             main_usecases_beans_factories.NewCreateEmailBean().Get(),
		CreateEmailBodyFactory:                                  main_usecases_beans_factories.NewCreateEmailBodyFactoryBean().Get(),
		SendEmailGatewayFactory:                                 main_usecases_beans_factories.NewSendEmailGatewayFactoryBean().Get(),
		CreateEmailBodyForWelcomeEmail:                          main_usecases_beans_factories.NewCreateEmailBodyForWelcomeEmailBean().Get(),
		FindEmailsByFilter:                                      main_usecases_beans_factories.NewFindEmailsByFilterBean().Get(),
		CreateEmailBodySendAndPersistAsSent:                     main_usecases_beans_factories.NewCreateEmailBodySendAndPersistAsSentBeanFactory().Get(),
		ReprocessEmailEvent:                                     main_usecases_beans_factories.NewReprocessEmailEventBean().Get(),
		SendEmailEventsToReprocess:                              main_usecases_beans_factories.NewSendEmailEventsToReprocessBean().Get(),
		CreateEmailFallback:                                     main_usecases_beans_factories.NewCreateEmailFallbackBean().Get(),
		EnableInternalProviderUser:                              main_usecases_beans_factories.NewEnableInternalProviderUserBean().Get(),
		CreateSessionByIdentityProvider:                         main_usecases_beans_factories.NewCreateSessionByIdentityProviderBean().Get(),
		EnableExternalProviderUser:                              main_usecases_beans_factories.NewEnableExternalProviderUserBean().Get(),
		CreateInternalAuthUserPasswordChangeRequest:             main_usecases_beans_factories.NewCreateInternalAuthUserPasswordChangeRequestBean().Get(),
		ChangeInternalProviderUserPassword:                      main_usecases_beans_factories.NewChangeInternalProviderUserPasswordBean().Get(),
		CreateInternalProviderUserSession:                       main_usecases_beans_factories.NewCreateInternalProviderUserSessionBean().Get(),
		RefreshSession:                                          main_usecases_beans_factories.NewRefreshSessionBean().Get(),
		EndSession:                                              main_usecases_beans_factories.NewEndSessionBean().Get(),
		AtomicLockedCreateNewUser:                               main_usecases_beans_factories.NewAtomicLockedCreateNewUser().Get(),
		AtomicLockedEnableExternalProviderUser:                  main_usecases_beans_factories.NewAtomicLockedEnableExternalProviderUserBean().Get(),
		AtomicLockedEnableInternalProviderUser:                  main_usecases_beans_factories.NewAtomicLockedEnableInternalProviderUserBeanFactory().Get(),
		AtomicLockedCreateInternalAuthUserPasswordChangeRequest: main_usecases_beans_factories.NewAtomicLockedCreateInternalAuthUserPasswordChangeRequestBean().Get(),
		AtomicLockedChangeInternalProviderUserPassword:          main_usecases_beans_factories.NewAtomicLockedChangeInternalProviderUserPasswordBean().Get(),
		AtomicLockedCreateExternalAuthProviderUserSession:       main_usecases_beans_factories.NewAtomicLockedCreateExternalAuthProviderUserSessionBeanFactory().Get(),
	}
}
