package handlers

import (
	"encoding/json"
	"fmt"
	authdto "housy/dto/auth"
	dto "housy/dto/result"
	"housy/models"
	"housy/pkg/bcrypt"
	jwtToken "housy/pkg/jwt"
	"housy/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.SignUpRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Fullname:   request.Fullname,
		Username:   request.Username,
		Email:      request.Email,
		Password:   password,
		ListAsRole: request.ListAsRole,
		Gendre:     request.Gender,
		Phone:      request.Phone,
		Address:    request.Address,
		Image:      request.Image,
	}

	data, err := h.AuthRepository.SignUp(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	signUpResponse := authdto.SignUpResponse{
		Username: data.Username,
		Message:  "Succesfully Sign Up!",
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: signUpResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.SignInRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.SignIn(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	signInResponse := authdto.SignInResponse{
		Username: request.Username,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: signInResponse}
	json.NewEncoder(w).Encode(response)

}
