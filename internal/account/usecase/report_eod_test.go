package usecase

import (
	"context"
	"testing"
)

func TestAccountUsecase_EndOfTheDay(t *testing.T) {
	errReadCtx := context.WithValue(context.Background(), "need_error", "read")
	errWriteCtx := context.WithValue(context.Background(), "need_error", "write")
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "error when get data data",
			args:    args{errReadCtx},
			wantErr: true,
		},
		{
			name:    "error when write data",
			args:    args{errWriteCtx},
			wantErr: true,
		},
		{
			name:    "success",
			args:    args{context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeRepo := new(fakeRepo)
			au := &AccountUsecase{
				repo:      fakeRepo,
				threadNum: new(thread),
			}
			if err := au.EndOfTheDay(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("AccountUsecase.EndOfTheDay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
