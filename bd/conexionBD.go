package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://practica1:practica1@cluster0.sp6bo.mongodb.net/test?authSource=admin&replicaSet=atlas-czyhbm-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

/* ConectarBD es la funcion que me permite conectar a la BD */
func ConectarBD() *mongo.Client {
	//Toma la conexion a la BD
	//Que son los context, es un entonrno de ejecucion, con espacio en memoria
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	//Si hubo un error en la BD
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa.")
	return client
}

/* ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
