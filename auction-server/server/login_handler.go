package server

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/hemantsharma1498/auction/pkg/auth"
	"github.com/hemantsharma1498/auction/store/models"
	"golang.org/x/crypto/argon2"
)

const (
	saltSize int    = 16
	sTime    uint32 = 6
	memory   uint32 = 32
	keyLen   uint32 = 32
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	fmt.Println(err, email)
	return err == nil
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	d := &LoginReq{}
	if err := decodeReqBody(r, d); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	if !valid(d.Email) {
		writeResponse(w, nil, errors.New("invalid email"), http.StatusBadRequest)
		return
	}
	users, err := s.store.GetUsersByEmail([]string{d.Email})

	if err != nil {
		writeResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	decodedSalt, _ := DecodeBase64(users[0].Salt)

	enteredPasswordHash := createHash(d.Password, []byte(decodedSalt))
	if EncodeBase64(enteredPasswordHash) != users[0].PwHash {
		writeResponse(w, errors.New("entered pasword and stored password don't match"), "entered pasword and stored password don't match", http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateJWT(users[0].UserID, d.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	d := &SignUpReq{}
	if err := decodeReqBody(r, d); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	if !valid(d.Email) {
		writeResponse(w, nil, errors.New("invalid email"), http.StatusBadRequest)
		return
	}
	_, err := s.store.GetUsersByEmail([]string{d.Email})
	if err != nil && err.Error() != "no users found for the provided emails" {
		writeResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	salt, err := generateSalt()
	if err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	pwHash := createHash(d.Password, salt)
	user := &models.User{
		Name:   d.Name,
		Email:  d.Email,
		Mobile: d.Mobile,
		Salt:   EncodeBase64(salt),
		PwHash: EncodeBase64(pwHash),
	}
	if err = s.store.CreateUser(user); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	writeResponse(w, nil, "account created successfully", http.StatusOK)
}

func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeBase64(encodedData string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodedData)
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func createHash(password string, salt []byte) []byte {
	hash := argon2.Key([]byte(password), salt, sTime, memory, 8, keyLen)
	return hash
}

func decodeReqBody(r *http.Request, d any) error {
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return err
	}
	return nil
}

func writeResponse(w http.ResponseWriter, err error, msg any, httpStatus int) error {
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(msg)
}
