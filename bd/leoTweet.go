package bd

import (
	"context"
	"log"
	"time"

	"github.com/JorrgeG/backendGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweet(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar cuando termine la busqueda en la bd

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var resultado []*models.DevuelvoTweets
	condicion := bson.M{
		"userid": ID,
	}
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //trae los documentos ordenados desendente por fecha
	opciones.SetSkip((pagina - 1) * 20)                 //paginando y salteando

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	//TODO creo qun contexto vacio
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}
		resultado = append(resultado, &registro)
	}
	return resultado, true
}
