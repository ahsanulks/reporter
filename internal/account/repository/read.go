package repository

import (
	"context"
	"daily_report/internal/domain"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

func (ac AccountRepo) GetReportBeforeEOD(ctx context.Context) ([]*domain.Account, error) {
	accounts := []*domain.Account{}
	// open the input csv file
	file, err := os.Open(ac.InputFile)
	if err != nil {
		return accounts, errors.New("unable to get input file")
	}
	defer file.Close()

	// create the reader of csv and replace default delimiter to ';'
	csvReader := csv.NewReader(file)
	csvReader.Comma = delimiterCSV
	// read the header first
	csvReader.Read()

	for {
		// read every line of csv
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return []*domain.Account{}, errors.New("unable to parse input file")
		}

		account := parseCSVtoAccount(record)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// parseCSVtoAccount will convert csv row to be domain account
func parseCSVtoAccount(record []string) *domain.Account {
	id, _ := strconv.ParseInt(record[0], 10, 64)
	age, _ := strconv.ParseInt(record[2], 10, 8)
	balance, _ := strconv.ParseInt(record[3], 10, 64)
	prevBalance, _ := strconv.ParseInt(record[4], 10, 64)
	avgbalance, _ := strconv.ParseFloat(record[5], 64)
	freeTransfer, _ := strconv.ParseInt(record[6], 10, 8)

	account := domain.Account{
		ID:           uint64(id),
		Name:         record[1],
		Age:          uint8(age),
		PrevBalance:  uint64(prevBalance),
		AvgBalance:   avgbalance,
		FreeTransfer: uint8(freeTransfer),
	}

	account.SetBalance(uint64(balance))
	return &account
}
