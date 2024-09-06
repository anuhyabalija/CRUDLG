package inputhandlers

import (
	"CRUDLG/server"
	"bufio"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		reader *bufio.Reader
		s      *server.Server
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.reader, tt.args.s)
		})
	}
}

func TestReadUser(t *testing.T) {
	type args struct {
		reader *bufio.Reader
		s      *server.Server
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadUser(tt.args.reader, tt.args.s)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		reader *bufio.Reader
		s      *server.Server
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUser(tt.args.reader, tt.args.s)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		reader *bufio.Reader
		s      *server.Server
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUser(tt.args.reader, tt.args.s)
		})
	}
}

func TestListUsers(t *testing.T) {
	type args struct {
		s *server.Server
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListUsers(tt.args.s)
		})
	}
}
