package handlers

import (
	"Rest/data"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type PatientsHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *data.PatientRepo
}

type UsersHandler struct {
	logger *log.Logger
	repo *data.UserRepo
}

// Injecting the logger makes this code much more testable.
func NewUsersHandler(l *log.Logger, r *data.UserRepo) *UsersHandler {
	return &UsersHandler{l, r}
}

func (u *UsersHandler) GetAllUsers(rw http.ResponseWriter, h* http.Request){
	u.logger.Println(h.Header.Get("username"))
	u.logger.Println(h.Header.Get("userType"))
	users, err := u.repo.GetAll()
	if err != nil {
		u.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}
	
	err = users.ToJSON(rw)
	u.logger.Println("rw")
	u.logger.Println(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Fatal("Unable to convert to json :", err)
		return 
	}
}

func (u *UsersHandler) InitTestDb(rw http.ResponseWriter, h* http.Request){
	err := u.repo.DropCollection()
	if err != nil {
		u.logger.Print("Database exception: ", err)
	}
	user :=  data.User{
		Username: "naz1",
		Password: "123",
		UserType: "regular",

	}
	u.repo.Insert(&user)
}

// func (p *PatientsHandler) GetPatientById(rw http.ResponseWriter, h *http.Request) {
// 	vars := mux.Vars(h)
// 	id := vars["id"]

// 	patient, err := p.repo.GetById(id)
// 	if err != nil {
// 		p.logger.Print("Database exception: ", err)
// 	}

// 	if patient == nil {
// 		http.Error(rw, "Patient with given id not found", http.StatusNotFound)
// 		p.logger.Printf("Patient with id: '%s' not found", id)
// 		return
// 	}

// 	err = patient.ToJSON(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
// 		p.logger.Fatal("Unable to convert to json :", err)
// 		return
// 	}
// }

func (u *UsersHandler) GetUsersByUsername(rw http.ResponseWriter, h *http.Request) {
	username := h.URL.Query().Get("username")

	users, err := u.repo.GetByUsername(username)
	if err != nil {
		u.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Fatal("Unable to convert to json: ", err)
		return
	}
}

func (u *UsersHandler) Login(rw http.ResponseWriter, h *http.Request){
	user := h.Context().Value(KeyProduct{}).(*data.User)
	retUser, err := u.repo.LoginUser(user.Username, user.Password)
	if err != nil {
		http.Error(rw, "Unable to login", http.StatusInternalServerError)
		u.logger.Fatal("Unable to login: ", err)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": retUser.Username,
		"userType": retUser.UserType,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(rw, "Unable to create token", http.StatusInternalServerError)
		u.logger.Fatal("Unable to create token: ", err)
		return
	}
	// rw.Write([]byte(tokenString))
	rw.Header().Set("Bearer",tokenString)
	rw.WriteHeader(http.StatusAccepted)
}

func (u *UsersHandler) PostUser(rw http.ResponseWriter, h *http.Request){
	user := h.Context().Value(KeyProduct{}).(*data.User)
	u.repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

func (u *UsersHandler) PatchUser(rw http.ResponseWriter, h *http.Request){
	vars := mux.Vars(h)
	id := vars["id"]
	user := h.Context().Value(KeyProduct{}).(*data.User)

	u.repo.Update(id, user)
	rw.WriteHeader(http.StatusOK)
}

func (p *PatientsHandler) AddPhoneNumber(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	var phoneNumber string
	d := json.NewDecoder(h.Body)
	d.Decode(&phoneNumber)

	p.repo.AddPhoneNumber(id, phoneNumber)
	rw.WriteHeader(http.StatusOK)
}

func (u *UsersHandler) DeleteUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	
	u.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (p *PatientsHandler) AddAnamnesis(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	anamnesis := h.Context().Value(KeyProduct{}).(*data.Anamnesis)

	p.repo.AddAnamnesis(id, anamnesis)
	rw.WriteHeader(http.StatusOK)
}

func (p *PatientsHandler) AddTherapy(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	therapy := h.Context().Value(KeyProduct{}).(*data.Therapy)

	p.repo.AddTherapy(id, therapy)
	rw.WriteHeader(http.StatusOK)
}

func (p *PatientsHandler) ChangeAddress(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	address := h.Context().Value(KeyProduct{}).(*data.Address)

	p.repo.UpdateAddress(id, address)
	rw.WriteHeader(http.StatusOK)
}

func (p *PatientsHandler) ChangePhone(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		http.Error(rw, "Unable to decode index", http.StatusBadRequest)
		p.logger.Fatal(err)
		return
	}

	var phoneNumber string
	d := json.NewDecoder(h.Body)
	d.Decode(&phoneNumber)

	p.repo.ChangePhone(id, index, phoneNumber)
	rw.WriteHeader(http.StatusOK)
}

func (p *PatientsHandler) Receipt(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	total, err := p.repo.Receipt(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalJson := map[string]float64{"total": total}

	e := json.NewEncoder(rw)
	e.Encode(totalJson)
}

func (p *PatientsHandler) Report(rw http.ResponseWriter, h *http.Request) {
	report, err := p.repo.Report()
	if err != nil {
		p.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(rw)
	e.Encode(report)
}

func (u *UsersHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request){
		user := &data.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *UsersHandler) MiddlewareLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request){
		user := &data.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *UsersHandler) MiddlewareAuth(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request){
		bearer := h.Header["Bearer"]
		if bearer != nil {

			tokenString := bearer[0]
			u.logger.Println(tokenString)
		
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secretKey := os.Getenv("SECRET_KEY")
			u.logger.Println(secretKey)
			return []byte(secretKey), nil
		})
		u.logger.Println("TOKEN: ")
		u.logger.Println(token)
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			u.logger.Println("Valid jwt")
			u.logger.Println(claims["username"], claims["userType"])
			h.Header.Set("username",claims["username"].(string))
			h.Header.Set("userType",claims["userType"].(string))
			next.ServeHTTP(rw, h)
			} else {
				u.logger.Println(err)
			}
		}
	})
}



// Solution: we added middlewares for Anamnesis, Therapy and Address objects
// func (p *PatientsHandler) MiddlewareAnamnesisDeserialization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		anamnesis := &data.Anamnesis{}
// 		err := anamnesis.FromJSON(h.Body)
// 		if err != nil {
// 			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
// 			p.logger.Fatal(err)
// 			return
// 		}

// 		ctx := context.WithValue(h.Context(), KeyProduct{}, anamnesis)
// 		h = h.WithContext(ctx)

// 		next.ServeHTTP(rw, h)
// 	})
// }

// func (p *PatientsHandler) MiddlewareTherapyDeserialization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		therapy := &data.Therapy{}
// 		err := therapy.FromJSON(h.Body)
// 		if err != nil {
// 			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
// 			p.logger.Fatal(err)
// 			return
// 		}

// 		ctx := context.WithValue(h.Context(), KeyProduct{}, therapy)
// 		h = h.WithContext(ctx)

// 		next.ServeHTTP(rw, h)
// 	})
// }

// func (p *PatientsHandler) MiddlewareAddressDeserialization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		address := &data.Address{}
// 		err := address.FromJSON(h.Body)
// 		if err != nil {
// 			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
// 			p.logger.Fatal(err)
// 			return
// 		}

// 		ctx := context.WithValue(h.Context(), KeyProduct{}, address)
// 		h = h.WithContext(ctx)

// 		next.ServeHTTP(rw, h)
// 	})
// }

func (u *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}