package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_usecases "baseapplicationgo/main/usecases"
)

type PersistTransactionBean struct {
}

func NewPersistTransactionBean() *PersistTransactionBean {
	return &PersistTransactionBean{}
}

func (this *PersistTransactionBean) Get() *main_usecases.PersistTransaction {

	transactionRepository := main_gateways_mongodb_repositories.NewTransactionRepository()
	var transactionDatabaseGateway main_gateways.TransactionDatabaseGateway = main_gateways_mongodb.NewTransactionDatabaseGatewayImpl(*transactionRepository)

	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()

	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()

	return main_usecases.NewPersistTransaction(transactionDatabaseGateway, logsMonitoringGateway, distributedLockGateway)
}
