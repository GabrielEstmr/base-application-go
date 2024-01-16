package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
)

type TransactionDatabaseGatewayImpl struct {
	transactionRepository main_gateways_mongodb_repositories.TransactionRepository
	spanGateway           main_gateways.SpanGateway
}

func NewTransactionDatabaseGatewayImpl(transactionRepository main_gateways_mongodb_repositories.TransactionRepository) *TransactionDatabaseGatewayImpl {
	return &TransactionDatabaseGatewayImpl{transactionRepository,
		main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *TransactionDatabaseGatewayImpl) Save(
	ctx context.Context,
	transaction main_domains.Transaction,
) (
	main_domains.Transaction,
	error,
) {
	span := this.spanGateway.Get(ctx, "TransactionDatabaseGatewayImpl-Save")
	defer span.End()

	userDocument := main_gateways_mongodb_documents.NewTransactionDocument(transaction)
	persistedTransactionDocument, err := this.transactionRepository.Save(span.GetCtx(), userDocument)
	if err != nil {
		return main_domains.Transaction{}, err
	}
	return persistedTransactionDocument.ToDomain(), nil
}
