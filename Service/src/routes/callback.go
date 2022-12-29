package routes

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/basarrcan/NPCAI/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {

	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config: %v\n", err)
	}

	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     config.GoogleApiClientID,
		ClientSecret: config.GoogleApiClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	oauthStateString := "datsguugleye"

	// Determine which provider the callback is from
	provider := r.FormValue("provider")
	var oauthConfig *oauth2.Config
	switch provider {
	case "google":
		oauthConfig = googleOauthConfig
	// case "twitter":
	// oauthConfig = twitterOauthConfig
	default:
		fmt.Println("Invalid OAuth 2.0 provider:", provider)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Use the token to retrieve the user's profile information
	var profileURL string
	switch provider {
	case "google":
		profileURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken
		// case "twitter":
		// 	profileURL = "https://api.twitter.com/1.1/account/verify_credentials.json"
	}
	resp, err := http.Get(profileURL)
	if err != nil {
		fmt.Printf("Failed to retrieve user profile: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// Use the data to display the user's profile information
	fmt.Fprintf(w, "User Info: %s\n", data)
}
