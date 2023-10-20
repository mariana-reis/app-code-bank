package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/mariana-reis/app.code.bank/application/usecase"
	"github.com/mariana-reis/app.code.bank/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         &pixRepository,
	}

	return transactionUseCase
}
