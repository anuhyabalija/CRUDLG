package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	database "CRUDLG/models"
)

// GravatarProfile represents the structure of the JSON response from Gravatar
type GravatarProfile struct {
	Entry []struct {
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"thumbnailUrl"`
	} `json:"entry"`
}

// GenerateMD5Hash generates an MD5 hash for the given email
// According to the Gravatar Developer document emailHash is passed as a parameter in the API call
func GenerateMD5Hash(email string) string {
	hasher := md5.New()
	hasher.Write([]byte(strings.ToLower(strings.TrimSpace(email))))
	return hex.EncodeToString(hasher.Sum(nil))
}

// FetchAndStoreGravatarProfile fetches the Gravatar profile and stores it in the database
func FetchAndStoreGravatarProfile(db *database.DB, email string) error {
	emailHash := GenerateMD5Hash(email)
	gravatarURL := fmt.Sprintf("https://www.gravatar.com/%s.json", emailHash)

	resp, err := http.Get(gravatarURL)
	if err != nil {
		return fmt.Errorf("failed to fetch Gravatar profile: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to retrieve profile: %s", resp.Status)
		// If profile not found, store user with default values
		return database.StoreUserProfile(db.DB, email, "", "")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var profile GravatarProfile
	err = json.Unmarshal(body, &profile)
	if err != nil {
		log.Printf("failed to parse JSON: %v", err)
		// If we can't parse the JSON, create a user with default values with name and avatarUrl being empty strings
		return database.StoreUserProfile(db.DB, email, "", "")
	}

	// If profile is fetched then use its data. Otherwise, use default values.
	displayName := ""
	avatarURL := ""
	if len(profile.Entry) > 0 {
		displayName = profile.Entry[0].DisplayName
		avatarURL = profile.Entry[0].AvatarURL
	}

	//log.Printf("Fetched profile - Email: %s, Name: %s, AvatarURL: %s", email, displayName, avatarURL)

	// Store the user profile in the database
	err = database.StoreUserProfile(db.DB, email, displayName, avatarURL)
	if err != nil {
		log.Printf("failed to store user profile: %v", err)
	}
	return err
}
