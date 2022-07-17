package main

import (
	"context"
	"daily_report/internal/account/repository"
	"daily_report/internal/account/usecase"
	"log"
)

func main() {
	inputFile := "./internal/account/repository/Before Eod.csv"
	outputFile := "./internal/account/repository/After Eod.csv"
	accountRepo := repository.NewAccountRepo(inputFile, outputFile)
	accountUsecase := usecase.NewAccountUsecase(accountRepo)

	err := accountUsecase.EndOfTheDay(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
