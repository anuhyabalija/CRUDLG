package server

import (
	models "CRUDLG/models"
	user "CRUDLG/proto"
	"context"
	"reflect"
	"sync"
	"testing"
)

func TestServer_CreateUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
		DB                             *models.DB
		greetMutex                     sync.Mutex
		greetMap                       map[string]int
	}
	type args struct {
		ctx context.Context
		req *user.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.CreateUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				DB:                             tt.fields.DB,
				greetMutex:                     tt.fields.greetMutex,
				greetMap:                       tt.fields.greetMap,
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
		DB                             *models.DB
		greetMutex                     sync.Mutex
		greetMap                       map[string]int
	}
	type args struct {
		ctx context.Context
		req *user.GetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.GetUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				DB:                             tt.fields.DB,
				greetMutex:                     tt.fields.greetMutex,
				greetMap:                       tt.fields.greetMap,
			}
			got, err := s.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
		DB                             *models.DB
		greetMutex                     sync.Mutex
		greetMap                       map[string]int
	}
	type args struct {
		ctx context.Context
		req *user.UpdateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.UpdateUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				DB:                             tt.fields.DB,
				greetMutex:                     tt.fields.greetMutex,
				greetMap:                       tt.fields.greetMap,
			}
			got, err := s.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
		DB                             *models.DB
		greetMutex                     sync.Mutex
		greetMap                       map[string]int
	}
	type args struct {
		ctx context.Context
		req *user.DeleteUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.DeleteUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				DB:                             tt.fields.DB,
				greetMutex:                     tt.fields.greetMutex,
				greetMap:                       tt.fields.greetMap,
			}
			got, err := s.DeleteUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ListUsers(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
		DB                             *models.DB
		greetMutex                     sync.Mutex
		greetMap                       map[string]int
	}
	type args struct {
		ctx context.Context
		req *user.ListUsersRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.ListUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				DB:                             tt.fields.DB,
				greetMutex:                     tt.fields.greetMutex,
				greetMap:                       tt.fields.greetMap,
			}
			got, err := s.ListUsers(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GreetUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
		DB                             *models.DB
		greetMutex                     sync.Mutex
		greetMap                       map[string]int
	}
	type args struct {
		ctx context.Context
		req *user.GreetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.GreetUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				DB:                             tt.fields.DB,
				greetMutex:                     tt.fields.greetMutex,
				greetMap:                       tt.fields.greetMap,
			}
			got, err := s.GreetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GreetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GreetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
