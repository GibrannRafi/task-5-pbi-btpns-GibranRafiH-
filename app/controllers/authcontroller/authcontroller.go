package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/config"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/helper"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	userInput := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	// checking user data :username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username  salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return

		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return

		}

	}

	// checking password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": " password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Proses Generate token or pembuatan jwt
	expTime := time.Now().Add(time.Minute * 5)
	claims := config.JWTclaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//Mendeklarasikan algoritma untuk signing

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login Berhasil!!"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func Register(w http.ResponseWriter, r *http.Request) {
	// Mengambil inputan JSON
	userInput := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	// Hash password menggunakan bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		response := map[string]string{"message": "Internal server error"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	userInput.Password = string(hashPassword)

	// Insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": "Internal server error"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Anda Telah Logout"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	// Decode data dari request body
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Cari user yang ingin di-update
	var existingUser models.User
	if err := models.DB.First(&existingUser, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to find user", http.StatusInternalServerError)
		return
	}

	// Update data user
	models.DB.Model(&existingUser).Updates(user) // Gunakan method Updates untuk update

	// Kirim response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Retrieve userID from the context
	userID := r.Context().Value("userID").(uint) // Cast to uint

	photos := &models.Photo{}
	if err := models.DB.Where("user_id = ?", userID).Find(photos).Error; err != nil {
		// Handle database query error
		http.Error(w, "Error fetching photos", http.StatusInternalServerError)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, photos)
}
