package server

import (
	"context"
	"database/sql"

	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	models "CRUDLG/models"
	user "CRUDLG/proto"
	"CRUDLG/utils"
)

type Server struct {
	user.UnimplementedUserServiceServer
	DB         *models.DB
	greetMutex sync.Mutex
	greetMap   map[string]int
}

func (s *Server) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	// Check if user already exists
	existingUser, err := models.GetUserByEmail(s.DB.DB, req.Email)
	if err == nil {
		// Return the exisiting User details
		return &user.CreateUserResponse{
			User: &user.User{
				Id:         int32(existingUser.ID),
				Email:      existingUser.Email,
				Name:       existingUser.Name.String,
				AvatarURL:  existingUser.AvatarURL.String,
				GreetCount: 0,
			},
		}, nil
	} else if err != sql.ErrNoRows {
		// An unexpected error occurred
		return nil, status.Errorf(codes.Internal, "Failed to check existing user: %v", err)
	}

	// Fetch and store Gravatar profile
	err = utils.FetchAndStoreGravatarProfile(s.DB, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch and store Gravatar profile: %v", err)
	}

	// Insert the new user into the database
	err = models.InsertUser(s.DB.DB, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	// Retrieve the newly created user
	u, err := models.GetUserByEmail(s.DB.DB, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user: %v", err)
	}

	return &user.CreateUserResponse{
		User: &user.User{
			Id:         int32(u.ID),
			Email:      u.Email,
			Name:       u.Name.String,
			AvatarURL:  u.AvatarURL.String,
			GreetCount: 0,
		},
	}, nil
}

// Reading the created user details
func (s *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	u, err := models.GetUserByEmail(s.DB.DB, req.Email)
	if err != nil {
		// User profile that is not found
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User with email %s not found", req.Email)
		}
		// For no no user records
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user: %v", err)
	}

	return &user.GetUserResponse{
		User: &user.User{
			Id:        int32(u.ID),
			Email:     u.Email,
			Name:      u.Name.String,
			AvatarURL: u.AvatarURL.String,
		},
	}, nil
}

// Update user name for the fetched Gravatar profile
func (s *Server) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	err := models.StoreUserProfile(s.DB.DB, req.Email, req.Name, req.AvatarURL)
	if err != nil {
		return nil, status.Errorf(1, "Failed to update user: %v", err)
	}
	// Getting the user details through email
	u, err := models.GetUserByEmail(s.DB.DB, req.Email)
	if err != nil {
		return nil, status.Errorf(1, "Failed to retrieve user: %v", err)
	}

	return &user.UpdateUserResponse{
		User: &user.User{
			Id:        int32(u.ID),
			Email:     u.Email,
			Name:      u.Name.String,
			AvatarURL: u.AvatarURL.String,
		},
	}, nil
}

// Delete the user details
func (s *Server) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := models.DeleteUserByEmail(s.DB.DB, req.Email)
	if err != nil {
		return nil, status.Errorf(1, "Failed to delete user: %v", err)
	}

	return &user.DeleteUserResponse{Success: true}, nil
}

// For listing all the created users
func (s *Server) ListUsers(ctx context.Context, req *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	users, err := models.ListUsers(s.DB.DB)
	if err != nil {
		return nil, status.Errorf(1, "Failed to list users: %v", err)
	}

	var userList []*user.User
	for _, u := range users {
		userList = append(userList, &user.User{
			Id:        int32(u.ID),
			Email:     u.Email,
			Name:      u.Name.String,
			AvatarURL: u.AvatarURL.String,
		})
	}

	return &user.ListUsersResponse{Users: userList}, nil
}

// Function is used to greet when a user is created or when the user is read
func (s *Server) GreetUser(ctx context.Context, req *user.GreetUserRequest) (*user.GreetUserResponse, error) {
	email := req.Email
	// Fetch the user to greet
	_, err := models.GetUserByEmail(s.DB.DB, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User with email %s not found", email)
		}
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user: %v", err)
	}
	// For concurrent safety when multiple goroutines might access the greetMap
	s.greetMutex.Lock()
	defer s.greetMutex.Unlock()

	if s.greetMap == nil {
		s.greetMap = make(map[string]int)
	}

	// Check greeting limit
	if s.greetMap[email] >= 100 {
		return nil, status.Errorf(codes.ResourceExhausted, "User %s has been greeted 100 times", email)
	}

	// Increment the greet count
	s.greetMap[email]++

	return &user.GreetUserResponse{
		//Message: fmt.Sprintf("Hello, %s! Could this be anymore obvious", u.Name, s.greetMap[email]),
	}, nil
}
