# CRUDL Management System

This README provides instructions for setting up and running the CRUDL Management System.

## Table of Contents
- [Setup](#setup)
- [Protoc Setup](#protoc-setup)
- [Run the Project](#run-the-project)

## Setup

To set up the project, run the following commands:

```bash
git clone https://github.com/anuhyabalija/CRUDLG.git
cd your-path/CRUDLG
go mod init 
go mod tidy 
go get "package name" # For any failed imports
```

## Protoc Setup

Follow these steps to set up protoc:

1. Find the protoc-gen-go-grpc and protoc-gen-go binaries:

   ```bash
   find $HOME/go/bin -name protoc-gen-go-grpc
   find $HOME/go/bin -name protoc-gen-go
   ```

2. Open your shell configuration file:

   ```bash
   nano ~/.zshrc
   ```

3. Add the following line to your shell configuration file and save it:

   ```bash
   export PATH=$PATH:your-path/go/bin # Path where go/bin is located
   ```

4. Open the `user.proto` file and add this line:

   ```protobuf
   option go_package = "./userservice";
   ```

5. Navigate to the directory containing the `user.proto` file and run:

   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
          --go-grpc_out=. --go-grpc_opt=paths=source_relative \
          user.proto
   ```

   This will generate the protoc files.

## Run the Project

To run the project, use the following command:

```bash
go run main.go
```

---

For any issues or questions, please open an issue in the GitHub repository.
