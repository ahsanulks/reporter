package repository

import (
	"context"
	"daily_report/internal/domain"
	"reflect"
	"testing"
)

func TestAccountRepo_GetReportBeforeEOD(t *testing.T) {
	zeroAccounts := []*domain.Account{}
	accounts := []*domain.Account{
		{
			ID:           99,
			Name:         "asep",
			Age:          11,
			PrevBalance:  456,
			AvgBalance:   789,
			FreeTransfer: 21,
		},
	}
	accounts[0].SetBalance(123)

	type fields struct {
		InputFile  string
		OutputFile string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Account
		wantErr bool
	}{
		{
			name: "error file not found",
			fields: fields{
				InputFile: "testabc.csv",
			},
			want:    zeroAccounts,
			wantErr: true,
		},
		{
			name: "invalid delimiter",
			fields: fields{
				InputFile: "test_data/invalid_delimiter.csv",
			},
			want:    zeroAccounts,
			wantErr: true,
		},
		{
			name: "empty file",
			fields: fields{
				InputFile: "test_data/empty_file.csv",
			},
			want:    zeroAccounts,
			wantErr: false,
		},
		{
			name: "valid file",
			fields: fields{
				InputFile: "test_data/example.csv",
			},
			want:    accounts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := AccountRepo{
				InputFile:  tt.fields.InputFile,
				OutputFile: tt.fields.OutputFile,
			}
			got, err := ac.GetReportBeforeEOD(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountRepo.GetReportBeforeEOD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountRepo.GetReportBeforeEOD() = %v, want %v", got, tt.want)
			}
		})
	}
}
