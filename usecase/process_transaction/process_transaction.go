package process_transaction

import "github.com/marcostota/imersao5esquenta/entity"

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDTOInput) (TransactionDTOOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.Amount = input.Amount
	transaction.AccountID = input.AccountID

	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}
	return p.approvedTransaction(transaction)
}

func (p *ProcessTransaction) approvedTransaction(transaction *entity.Transaction) (TransactionDTOOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	return TransactionDTOOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMessage: "",
	}, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDTOOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "reproved", invalidTransaction.Error())
	if err != nil {
		return TransactionDTOOutput{
			ID:           transaction.ID,
			Status:       "rejected",
			ErrorMessage: err.Error(),
		}, err
	}

	return TransactionDTOOutput{
		ID:           transaction.ID,
		Status:       "rejected",
		ErrorMessage: invalidTransaction.Error(),
	}, invalidTransaction
}
