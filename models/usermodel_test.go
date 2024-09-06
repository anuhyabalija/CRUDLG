package models

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name    string
		want    *DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTable(t *testing.T) {
	type args struct {
		db *sql.DB
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
			if err := CreateTable(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	type args struct {
		db    *sql.DB
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
			if err := InsertUser(tt.args.db, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStoreUserProfile(t *testing.T) {
	type args struct {
		db          *sql.DB
		email       string
		displayName string
		avatarURL   string
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
			if err := StoreUserProfile(tt.args.db, tt.args.email, tt.args.displayName, tt.args.avatarURL); (err != nil) != tt.wantErr {
				t.Errorf("StoreUserProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	type args struct {
		db    *sql.DB
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserByEmail(tt.args.db, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUserByEmail(t *testing.T) {
	type args struct {
		db    *sql.DB
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
			if err := DeleteUserByEmail(tt.args.db, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListUsers(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListUsers(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
