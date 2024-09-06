package main

import (
	in "CRUDLG/inputhandlers"
	database "CRUDLG/models"
	user "CRUDLG/proto"
	server "CRUDLG/server"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	// Database connection and table creation
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	err = database.CreateTable(db.DB)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Create a server instance and register it with gRPC
	userServer := &server.Server{DB: db}
	user.RegisterUserServiceServer(grpcServer, userServer)

	// Start the gRPC server in a separate goroutine
	go startGRPCServer(grpcServer)

	// Handle user input
	reader := bufio.NewReader(os.Stdin)
	handleUserInput(reader, userServer)
	// If we reach here, the user has chosen to quit
	fmt.Println("Shutting down the server...")
	grpcServer.GracefulStop()
	fmt.Println("Server shut down. Goodbye!")
}

func startGRPCServer(grpcServer *grpc.Server) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func handleUserInput(reader *bufio.Reader, userServer *server.Server) {
	for {
		fmt.Println("\nCRUDL User Management System")
		fmt.Println("1. Create User")
		fmt.Println("2. Read User")
		fmt.Println("3. Update User")
		fmt.Println("4. Delete User")
		fmt.Println("5. List Users")
		fmt.Println("6. Quit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			in.CreateUser(reader, userServer)
		case "2":
			in.ReadUser(reader, userServer)
		case "3":
			in.UpdateUser(reader, userServer)
		case "4":
			in.DeleteUser(reader, userServer)
		case "5":
			in.ListUsers(userServer)
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
