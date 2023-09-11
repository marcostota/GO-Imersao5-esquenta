package process_transaction

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_entity "github.com/marcostota/imersao5esquenta/entity/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenItIsValid(t *testing.T) {
	input := TransactionDTOInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
