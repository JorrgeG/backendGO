package handlers

import (
	"log"
	"net/http"
	"os"

	middlew "github.com/JorrgeG/backendGO/middleW"
	"github.com/JorrgeG/backendGO/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	//Rutas - Endpoints
	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routes.GraboTweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.ChequeoBD(middlew.ValidoJWT(routes.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
