package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JorrgeG/backendGO/bd"
	"github.com/JorrgeG/backendGO/jwt"
	"github.com/JorrgeG/backendGO/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requeirdo ", 400)
		return
	}
	documento, exist := bd.IntentoLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "Usuario y/o Contraseña invalidos ", 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar general el Token correspondinete "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Como guardar una Cookie desde el Backend
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
