package main

//User usuarios
type User struct {
	Name  string
	Email string
}

//Users usuarios
var Users = []User{}

//UserDB mongo
var UserDB *UsersMONGO

func init() {

	UserDB = &UsersMONGO{}
	UserDB.Connect()
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
	return UserDB.save(user)
}

//RemoveUser remove usuarios
func RemoveUser(user User) error {
	return UserDB.deleteUser(user)
}

//UpdateUserName atualiza o nome
func UpdateUserName(user User, name string) error {
	findUser, _ := GetUserByEmail(user.Email)
	newUser := User{name, findUser.Email}
	Users = append(Users, newUser)
	return UserDB.updateUser(newUser, name)
}
