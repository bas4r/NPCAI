package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/basarrcan/NPCAI/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Replace with your OAuth 2.0 client ID and client secret
const (
	clientID     = "YOUR_CLIENT_ID"
	clientSecret = "YOUR_CLIENT_SECRET"
)

var (
	// Replace with your OAuth 2.0 provider's authorization and token endpoints
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	// State is a random string used to protect against cross-site request forgery attacks

)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Load config
	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config: %v\n", err)
	}
	url := googleOauthConfig.AuthCodeURL(config.OAuthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
