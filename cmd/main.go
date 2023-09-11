package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/marcostota/imersao5esquenta/adapter/repository"
	"github.com/marcostota/imersao5esquenta/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewTransactionRepositoryDB(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDTOInput{
		ID:        "1",
		AccountID: "1",
		Amount:    0,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err)
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))

}
