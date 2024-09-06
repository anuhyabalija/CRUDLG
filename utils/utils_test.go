package utils

import (
	database "CRUDLG/models"
	"testing"
)

func TestGenerateMD5Hash(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMD5Hash(tt.args.email); got != tt.want {
				t.Errorf("GenerateMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchAndStoreGravatarProfile(t *testing.T) {
	type args struct {
		db    *database.DB
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FetchAndStoreGravatarProfile(tt.args.db, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("FetchAndStoreGravatarProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
