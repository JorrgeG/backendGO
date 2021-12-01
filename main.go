package main

import (
	"log"

	"github.com/JorrgeG/backendGO/bd"
	"github.com/JorrgeG/backendGO/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexion a la BD")
		return
	}
	handlers.Manejadores()
}
