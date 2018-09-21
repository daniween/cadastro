package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

//Users usuarios
var Users = []User{}

func main() {

	AddUser(User{Name: "Stephanie Augusta", Email: "stephaug12@gmail.com"})
	AddUser(User{Name: "manuela", Email: "manu@gmail.com"})
	AddUser(User{Name: "pedro", Email: "pedro@gmail.com"})
	UpdateUserName(User{Name: "manuela", Email: "manu@gmail.com"}, "Manuela Augusta")
	UpdateUserName(User{Name: "pedro", Email: "pedro@gmail.com"}, "Pedro Fernandez")
	list, _ := ListUsers()
	fmt.Println(list)
	//AddUser(User{Name: "Stephanie Augusta", Email: "stephaug12@gmail.com"})
	GetUserByEmail("pedro@gmail.com")
	//RemoveUser(User{Name: "Manuela Augusta", Email: "manu@gmail.com"}) - quando tenta remover um usuario q foi atualizado, da erro.. ele remove mas tem uma msg
	RemoveUser(User{Name: "Stephanie Augusta", Email: "stephaug12@gmail.com"})

}

//ListUsers lista os usuarios
func ListUsers() ([]User, error) {

	return Users, nil
}

//GetUserByEmail lista o usuario pelo email
func GetUserByEmail(email string) (User, error) {

	for i, u := range Users {
		Users[i] = u

		if u.Email == email {
			fmt.Println("User:", u)
			return u, nil
		}
		//if u.Email != email { // nao da certo pq qnd ele pega o 1o usuario, ele ja cai aqui!!!
		//	fmt.Println("Email not found") // tem q achar um jeito de pesquisar TODOS, e se nenhum estivr na lista, retorna o email not found
		//	return User{}, fmt.Errorf("Email not Found")
		//}
	}

	for i, u := range Users {
		Users[i] = u

		if u.Email != email {
			fmt.Println("Email not found")
			return User{}, fmt.Errorf("Email not Found")
		}
	}
	return User{}, nil
}

//AddUser adiciona usuarios
func AddUser(user User) error {

	Users = append(Users, user)

	for i, u := range Users {
		Users[i] = u

		if user != u {
			return nil
		}
		if user == u { //se o usuario q ele tenta adicionar for igual a algum existente, retorna erro --> só funciona se eu adicionar o mesmo usuario depois de listar
			return fmt.Errorf("User already exists")
		}
	}
	return nil
}

//RemoveUser remove usuarios
func RemoveUser(user User) error {

	for i, u := range Users {
		Users[i] = u

		if u == user {
			copy(Users[i:], Users[i+1:]) //pega o user[i] e copia para o user [i+1](proximo)
			Users[len(Users)-1] = User{} // user[len(users)-1] é o ultimo termo do slice... o ultimo termo recebe vazio
			Users = Users[:len(Users)-1] // users vai acabar antes desse ultimo que ta vazio, excluindo esse termo selecionado
			fmt.Println(Users)
		}
	}

	for i, u := range Users { //para usuarios que nao existem
		Users[i] = u

		if u != user {
			fmt.Println("User:", user, "- não existe")
			return fmt.Errorf("")
		}
	}
	return nil
}

//UpdateUserName atualiza o nome
func UpdateUserName(user User, name string) error {

	//runtime.Breakpoint()
	for i, u := range Users {
		Users[i] = u

		if u == user {

			copy(Users[i:], Users[i+1:])
			Users[len(Users)-1] = User{}
			Users = Users[:len(Users)-1]

			newUserName := User{Name: name, Email: user.Email}
			AddUser(newUserName)
			return nil
		}
	}

	for i, u := range Users {
		Users[i] = u

		if u != user {
			return fmt.Errorf("Usuário não existe")
		}
	}
	return nil
}
