package routes

import (
	"net/http"

	"github.com/basarrcan/NPCAI/services"
)

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	// // Parse the request body to get the new user data
	// newUserInput := &models.NewUserInput{}

	// _ = json.NewDecoder(r.Body).Decode(&newUserInput)

	// err := newUserInput.Validate()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// db := services.ConnectDB()

	// err = newUserInput.Save(db)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// RespondWithJSON(w, http.StatusOK, newUserInput)
	services.HandleGoogleLogin(w, r)
}
