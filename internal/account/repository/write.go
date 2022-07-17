package repository

import (
	"context"
	"daily_report/internal/domain"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

func (ar AccountRepo) WriteReportAfterEOD(ctx context.Context, accounts []*domain.Account) error {
	// create the output file
	file, err := os.Create(ar.OutputFile)
	if err != nil {
		return errors.New("can't create output file")
	}
	defer file.Close()

	// create csv writer
	writer := csv.NewWriter(file)
	writer.Comma = delimiterCSV
	defer writer.Flush()

	// write the header to file that we already create
	header := []string{"id", "Nama", "Age", "Balanced", "No 2b Thread-No", "No 3 Thread-No", "Previous Balanced", "Average Balanced", "No 1 Thread-No", "Free Transfer", "No 2a Thread-No"}
	writer.Write(header)

	// write every account become to each row on csv
	for _, account := range accounts {
		record := []string{
			strconv.Itoa(int(account.ID)),
			account.Name,
			strconv.Itoa(int(account.Age)),
			strconv.Itoa(int(account.Balance())),
			toString(account.ThreadNum2b),
			toString(account.ThreadNum3),
			strconv.Itoa(int(account.PrevBalance)),
			strconv.FormatFloat(account.AvgBalance, 'f', -1, 64),
			toString(account.ThreadNum1),
			strconv.Itoa(int(account.FreeTransfer)),
			toString(account.ThreadNum2a),
		}

		writer.Write(record)
	}
	return err
}

// toString will convert if the value is default or 0 will be empty string not "0" string
func toString(num uint16) string {
	if num == 0 {
		return ""
	}
	return strconv.Itoa(int(num))
}
