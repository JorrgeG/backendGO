package bd

import (
	"context"
	"time"

	"github.com/JorrgeG/backendGO/models"
)

func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //el defer se ejecuta al final de la rutina

	db := MongoCN.Database("twitter")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
