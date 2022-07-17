package repository

import (
	"context"
	"daily_report/internal/domain"
	"testing"
)

func TestAccountRepo_WriteReportAfterEOD(t *testing.T) {
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
	type args struct {
		accounts []*domain.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "invalid path output",
			fields: fields{
				OutputFile: "test_data/output/test.csv",
			},
			args:    args{accounts},
			wantErr: true,
		},
		{
			name: "success write output",
			fields: fields{
				OutputFile: "test_data/output_test.csv",
			},
			args:    args{accounts},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := AccountRepo{
				InputFile:  tt.fields.InputFile,
				OutputFile: tt.fields.OutputFile,
			}
			if err := ar.WriteReportAfterEOD(context.Background(), tt.args.accounts); (err != nil) != tt.wantErr {
				t.Errorf("AccountRepo.WriteReportAfterEOD() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_toString(t *testing.T) {
	type args struct {
		num uint16
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "more than 0",
			args: args{123},
			want: "123",
		},
		{
			name: "zero",
			args: args{0},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.args.num); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
