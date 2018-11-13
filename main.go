package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/mlabouardy/movies-restapi/config"
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

func main() { //passo 1: ler o config: criar um obj do tipo config e chamar o config.read
	//passo 2: pegar o obj config e usar ele pra setar o obj de conexao com o banco (db ?)
	//passo 3: chamar o connect do obj do banco

	c := config.Config{} // passo 1
	c.Read()

	UserDB = &UsersMONGO{
		Database: c.Database,
		Server:   c.Server,
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
