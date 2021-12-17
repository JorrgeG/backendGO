package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JorrgeG/backendGO/bd"
	"github.com/JorrgeG/backendGO/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertTweet(registro)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar ingresar el resgistro, reintente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
