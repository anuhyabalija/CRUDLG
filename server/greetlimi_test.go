package server

import (
	models "CRUDLG/models"
	user "CRUDLG/proto"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
)

func TestGreetUserLimit(t *testing.T) {
	// Initialize database
	db, err := models.ConnectDB()
	assert.NoError(t, err)
	defer db.DB.Close()

	err = models.CreateTable(db.DB)
	assert.NoError(t, err)

	// Initialize server
	s := &Server{
		DB: db,
	}
	// Creating the scenario
	email := "anuhyabalija20@gmail.com"
	err = models.InsertUser(s.DB.DB, email)
	assert.NoError(t, err)

	// Open CSV file
	file, err := os.Create("greet_user_results.csv")
	assert.NoError(t, err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	err = writer.Write([]string{"attempt_id", "time_taken", "status_code"})
	assert.NoError(t, err)

	// Perform 500 greet attempts
	for i := 1; i <= 500; i++ {
		startTime := time.Now()

		req := &user.GreetUserRequest{Email: email}
		resp, err := s.GreetUser(context.Background(), req)
		fmt.Printf(resp.Message)

		duration := time.Since(startTime).Microseconds()
		var statusCode int

		if err != nil {
			// Map gRPC status codes to integers for CSV output
			st, ok := status.FromError(err)
			if ok {
				statusCode = int(st.Code())
			} else {
				statusCode = -1 // Unknown error
			}
		} else {
			statusCode = 200 // Success
		}

		// Write result to CSV
		record := []string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%d", duration),
			fmt.Sprintf("%d", statusCode),
		}
		err = writer.Write(record)
		assert.NoError(t, err)
	}
}
