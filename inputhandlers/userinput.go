package inputhandlers

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	user "CRUDLG/proto"
	"CRUDLG/server"
)

func CreateUser(reader *bufio.Reader, s *server.Server) {
	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	req := &user.CreateUserRequest{Email: email}
	resp, err := s.CreateUser(context.TODO(), req)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	// After creating the user, greet the user
	greetUserReq := &user.GreetUserRequest{
		Email: email,
	}

	_, err = s.GreetUser(context.Background(), greetUserReq)
	if err != nil {
		fmt.Printf("Error greeting user: %v\n", err)
	} else {
		fmt.Printf("User %s slays!\n", email)
	}
	fmt.Printf("User created: ID=%d, Email=%s, Name=%s\n", resp.User.Id, resp.User.Email, resp.User.Name)
}

func ReadUser(reader *bufio.Reader, s *server.Server) {
	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	req := &user.GetUserRequest{Email: email}
	resp, err := s.GetUser(context.TODO(), req)
	if err != nil {
		fmt.Println("Error reading user:", err)
		return
	}
	// After creating the user, greet the user
	greetUserReq := &user.GreetUserRequest{
		Email: email,
	}
	_, err = s.GreetUser(context.Background(), greetUserReq)
	if err != nil {
		fmt.Printf("Error greeting user: %v\n", err)
	} else {
		fmt.Printf("User %s slays!\n", email)
	}
	fmt.Printf("User found: ID=%d, Email=%s, Name=%s, AvatarURL=%s\n",
		resp.User.Id, resp.User.Email, resp.User.Name, resp.User.AvatarURL)
}

func UpdateUser(reader *bufio.Reader, s *server.Server) {
	fmt.Print("Enter email of user to update: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter new name (leave blank to keep current): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter new avatar URL (leave blank to keep current): ")
	avatarURL, _ := reader.ReadString('\n')
	avatarURL = strings.TrimSpace(avatarURL)

	req := &user.UpdateUserRequest{Email: email, Name: name, AvatarURL: avatarURL}
	resp, err := s.UpdateUser(context.TODO(), req)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Printf("User updated: ID=%d, Email=%s, Name=%s, AvatarURL=%s\n",
		resp.User.Id, resp.User.Email, resp.User.Name, resp.User.AvatarURL)
}

func DeleteUser(reader *bufio.Reader, s *server.Server) {
	fmt.Print("Enter email of user to delete: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	req := &user.DeleteUserRequest{Email: email}
	resp, err := s.DeleteUser(context.TODO(), req)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	if resp.Success {
		fmt.Println("User successfully deleted.")
	} else {
		fmt.Println("Failed to delete user.")
	}
}

func ListUsers(s *server.Server) {
	req := &user.ListUsersRequest{}
	resp, err := s.ListUsers(context.TODO(), req)
	if err != nil {
		fmt.Println("Error listing users:", err)
		return
	}
	fmt.Println("Users:")
	for _, u := range resp.Users {
		fmt.Printf("ID=%d, Email=%s, Name=%s, AvatarURL=%s\n", u.Id, u.Email, u.Name, u.AvatarURL)
	}
}
