package login

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	valid := (credentials.UserName == "admin" && credentials.Password == "admin")

	if !valid {
		http.Error(w, "Incorrect credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := GenerateToken(credentials.UserName)

	if err != nil {
		http.Error(w, "Failed to Generate the Token", http.StatusInternalServerError)
		log.Println("Error Generating token", err)
		return
	}

	response := map[string]string{"token": tokenString}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GenerateToken(userName string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	singedToken, err := token.SignedString([]byte("your_secret_key"))

	if err != nil {
		return "", err
	}

	return singedToken, nil
}
