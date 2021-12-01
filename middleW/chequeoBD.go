package middleW

import (
	"net/http"

	"github.com/JorrgeG/backendGO/bd"
)

//un middleW recibe un paremetro y devuelve el mismo
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		//le entrego todo al proximo eslavon de la cadena
		next.ServeHTTP(w, r)
	}
}
