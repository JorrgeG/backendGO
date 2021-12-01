package bd

import "golang.org/x/crypto/bcrypt"

//encriptar contraseña
func EncriptarPassword(pass string) (string, error) {
	//el campo de la contraseña elevado a el costo en este caso 8, mientras mas grande mucho mejor pero mas demorado.
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
