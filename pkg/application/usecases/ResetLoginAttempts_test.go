package usecases

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/repositories"
	"testing"
)

func TestResetLoginAttemptsHandler_Execute(t *testing.T) {
	type fields struct {
		bucketRepository repositories.BucketRepository
		bucketCleaner    bucketCleaner
	}
	type args struct {
		useCase *ResetLoginAttempts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &resetLoginAttemptsHandler{
				bucketRepository: tt.fields.bucketRepository,
				bucketCleaner:    tt.fields.bucketCleaner,
			}
			if err := h.Execute(tt.args.useCase); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
