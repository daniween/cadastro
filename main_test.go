package main

import (
	"testing"
)

func setup() {
	Users = []User{
		User{Name: "myself", Email: "me@example.com"},
	}
}

func teardown() {
	Users = []User{}
}

func TestListUsers(t *testing.T) {
	setup()
	defer teardown()

	users, err := ListUsers()
	if err != nil {
		t.Fatalf("expected err = nil, got %v\n", err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d\n", len(users))
	}
	if users[0].Email != "me@example.com" {
		t.Fatalf("expected user with email = me@example.com, got %q\n", users[0].Email)
	}
}

func TestGetUserByEmail(t *testing.T) {
	setup()
	defer teardown()

	user, err := GetUserByEmail("me@example.com")
	if err != nil {
		t.Fatalf("expected err = nil, got %v\n", err)
	}
	if user.Email != "me@example.com" {
		t.Fatalf("expected user with email = me@example.com, got %q\n", user.Email)
	}
}

func TestGetUserByEmailNotFound(t *testing.T) {
	setup()
	defer teardown()

	_, err := GetUserByEmail("unknown@example.com")
	if err == nil {
		t.Fatal("expected err != nil, got nil")
	}
}

func TestAddUser(t *testing.T) {
	setup()
	defer teardown()

	newUser := User{Name: "other user", Email: "other@example.com"}
	err := AddUser(newUser)
	if err != nil {
		t.Fatalf("expected err = nil, got %v\n", err)
	}
	user, err := GetUserByEmail(newUser.Email)
	if err != nil {
		t.Fatalf("expected err = nil, goot %v\n", err)
	}
	if user.Email != newUser.Email {
		t.Fatalf("expected user with email = %q got %q\n", newUser.Email, user.Email)
	}
}

func TestAddUserAlreadyExists(t *testing.T) {
	setup()
	defer teardown()

	user := User{Name: "myself", Email: "me@example.com"}
	err := AddUser(user)
	if err == nil {
		t.Fatal("expected err != nil, got nil")
	}
}

func TestRemoveUser(t *testing.T) {
	setup()
	defer teardown()

	myUser := User{Name: "myself", Email: "me@example.com"}
	err := RemoveUser(myUser)
	if err != nil {
		t.Fatalf("expected err = nil, got %v\n", err)
	}
	_, err = GetUserByEmail(myUser.Email)
	if err == nil {
		t.Fatal("expected err != nil - got nil") //erro nesse!!!
	}
}

func TestRemoveUserDoesNotExist(t *testing.T) {
	setup()
	defer teardown()

	user := User{Name: "unknown user", Email: "unknown@example.com"}
	err := RemoveUser(user)
	if err == nil {
		t.Fatal("expected err != nil, got nil")
	}
}

func TestUpdateUserName(t *testing.T) {
	setup()
	defer teardown()

	myUser := User{Name: "myself", Email: "me@example.com"}
	newName := "my new name"
	err := UpdateUserName(myUser, newName)
	if err != nil {
		t.Fatalf("expected err = nil - got %v\n", err)
	}
	user, err := GetUserByEmail(myUser.Email)
	if err != nil {
		t.Fatalf("expected err = nil, got %v\n", err)
	}
	if user.Name != newName {
		t.Fatalf("expected user with name = %q got %q\n", newName, user.Name)
	}

}

func TestUpdateUserNameDoesNotExist(t *testing.T) {
	setup()
	defer teardown()

	user := User{Name: "unknown user", Email: "unknown@example.com"}
	err := UpdateUserName(user, "new name")
	if err == nil {
		t.Fatal("expected err != nil, got nil")
	}
}
