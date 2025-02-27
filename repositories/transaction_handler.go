package repositories

import (
	"context"
	"database/sql"
)

// BeginTransaction starts a database transaction for both repositories
func BeginTransaction(runnersRepository *RunnersRepository,
	resultsRepository *ResultsRepository) error {
	ctx := context.Background()
	transaction, err := resultsRepository.dbHandler.
		BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	runnersRepository.transaction = transaction
	resultsRepository.transaction = transaction
	return nil
}

// RollbackTransaction reverts the changes made during the transaction
func RollbackTransaction(runnersRepository *RunnersRepository,
	resultsRepository *ResultsRepository) error {
	transaction := runnersRepository.transaction
	runnersRepository.transaction = nil
	resultsRepository.transaction = nil
	return transaction.Rollback()
}

// CommitTransaction saves the changes made during the transaction
func CommitTransaction(runnersRepository *RunnersRepository,
	resultsRepository *ResultsRepository) error {
	transaction := runnersRepository.transaction
	runnersRepository.transaction = nil
	resultsRepository.transaction = nil
	return transaction.Commit()
}
