package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

//User usuarios
type User struct {
	Name  string
	Email string
}

//Users usuarios
var Users = []User{}

//UserDB mongo
var UserDB *UsersMONGO

func main() {

	UserDB = &UsersMONGO{
		Server: os.Getenv("DBAAS_MONGODB_ENDPOINT"),
	}
	UserDB.Connect()

	e := echo.New()

	e.GET("/", getAllUsers)
	e.GET("/:email", getUser)             // curl localhost:1323/manu@gmail.com --> mostra o usuario manuela, manu@gmail.com
	e.PUT("/users/:email", updateUser)    // curl -X PUT -F "name=JOE" localhost:1323/users/joe@gmail.com
	e.POST("/save", save)                 // curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
	e.DELETE("/users/:email", deleteUser) // curl -X DELETE localhost:1323/users/manu@gmail.com

	e.Logger.Fatal(e.Start(":8888"))
}

//ListUsers lista os usuarios
func ListUsers() ([]User, error) {
	return UserDB.getAllUsers()
}

//GetUserByEmail lista o usuario pelo email
func GetUserByEmail(email string) (User, error) {
	return UserDB.getUser(email)
}

//AddUser adiciona usuarios
func AddUser(user User) error {
	_, err := GetUserByEmail(user.Email)
	if err == nil { // se o erro for igual nil, entao ele encontrou um usuario. Entao nao deixa adicionar
		return fmt.Errorf("User already exists")
	}
	return UserDB.save(user)
}

//RemoveUser remove usuarios
func RemoveUser(user User) error {
	return UserDB.deleteUser(user)
}

//UpdateUserName atualiza o nome
func UpdateUserName(user User, name string) error {
	user = User{name, user.Email}
	return UserDB.updateUser(user, name)
}
