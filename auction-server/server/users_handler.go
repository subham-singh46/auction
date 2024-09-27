package server

import (
	"errors"
	"net/http"

	"github.com/hemantsharma1498/auction/pkg/auth"
	"github.com/hemantsharma1498/auction/pkg/utils"
	"github.com/hemantsharma1498/auction/store/models"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	d := &LoginReq{}
	if err := utils.DecodeReqBody(r, d); err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	if !utils.ValidEmail(d.Email) {
		utils.WriteResponse(w, nil, errors.New("invalid email"), http.StatusBadRequest)
		return
	}
	users, err := s.store.GetUsersByEmail([]string{d.Email})

	if err != nil {
		utils.WriteResponse(w, err, nil, http.StatusBadRequest)
		return
	}

	decodedSalt, _ := utils.DecodeBase64(users[0].Salt)

	enteredPasswordHash := utils.CreateHash(d.Password, []byte(decodedSalt))
	if utils.EncodeBase64(enteredPasswordHash) != users[0].PwHash {
		utils.WriteResponse(w, errors.New("entered pasword and stored password don't match"), "entered pasword and stored password don't match", http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateJWT(users[0].UserID, d.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	res := &LoginRes{UserID: users[0].UserID, Token: token}
	utils.WriteResponse(w, nil, res, http.StatusOK)

}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	d := &SignUpReq{}
	if err := utils.DecodeReqBody(r, d); err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	if !utils.ValidEmail(d.Email) {
		utils.WriteResponse(w, nil, errors.New("invalid email"), http.StatusBadRequest)
		return
	}
	users, err := s.store.GetUsersByEmail([]string{d.Email})
	if err != nil && err.Error() != "no users found for the provided emails" {
		utils.WriteResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}
	res := &SignUpRes{}
	if len(users) > 0 {
		res.Msg = "User exists for the give email"
		utils.WriteResponse(w, nil, res, http.StatusBadRequest)
		return
	}

	salt, err := utils.GenerateSalt()
	if err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	pwHash := utils.CreateHash(d.Password, salt)
	user := &models.User{
		Name:   d.Name,
		Email:  d.Email,
		Mobile: d.Mobile,
		Salt:   utils.EncodeBase64(salt),
		PwHash: utils.EncodeBase64(pwHash),
	}
	userId, err := s.store.CreateUser(user)
	res.UserID = userId
	if err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateJWT(userId, d.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	res.Token = token
	utils.WriteResponse(w, nil, res, http.StatusOK)
}

func (s *Server) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	d := &UpdatePasswordReq{}
	if err := utils.DecodeReqBody(r, d); err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}

	if !utils.ValidEmail(d.Email) {
		utils.WriteResponse(w, errors.New("invalid email"), "invalid email", http.StatusBadRequest)
		return
	}
	salt, err := utils.GenerateSalt()
	if err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	pwHash := utils.CreateHash(d.NewPassword, salt)

	s.store.UpdatePassword(d.Email, utils.EncodeBase64(salt), utils.EncodeBase64(pwHash))

}
