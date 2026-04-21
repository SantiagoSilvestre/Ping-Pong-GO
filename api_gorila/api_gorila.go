package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var Users []User

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
		fmt.Fprintln(w, "Bem-vindo à API da Vovó")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Rotas disponíveis:")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "GET    /           -> Página inicial da API")
		fmt.Fprintln(w, "GET    /users      -> Lista todos os usuários")
		fmt.Fprintln(w, "GET    /users/{id} -> Busca um usuário pelo ID")
		fmt.Fprintln(w, "POST   /users      -> Cria um novo usuário")
		fmt.Fprintln(w, "PUT    /users/{id} -> Atualiza um usuário existente")
		fmt.Fprintln(w, "DELETE /users/{id} -> Remove um usuário pelo ID")
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		for _, u := range Users {
			fmt.Fprintln(w, u)
		}
	}).Methods("GET")

	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		//converte para int
		userId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Fprintln(w, "Id do usuário inválido")
			return
		}
		if userId <= 0 {
			fmt.Fprintln(w, "user id deve ser maior que 0")
			return
		}
		user := getUserById(userId)
		if user.Id > 0 {
			fmt.Fprintln(w, user)
		} else {
			fmt.Fprintln(w, "Usuário não encontrado")
		}

	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		//converte para int
		userId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Fprintln(w, "Id do usuário inválido")
			return
		}
		if userId <= 0 {
			fmt.Fprintln(w, "user id deve ser maior que 0")
			return
		}
		isSuccessfull := deleteUserById(userId)
		if isSuccessfull {
			fmt.Fprintln(w, "Usuário deletado com sucesso")
		} else {
			fmt.Fprintln(w, "Usuário não deletado")
		}

	}).Methods("DELETE")
	r.HandleFunc("/users", createUser).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func getUserById(userId int) User {
	for _, u := range Users {
		if u.Id == userId {
			return u
		}
	}
	return User{}
}

func deleteUserById(userId int) bool {
	for i, u := range Users {
		if u.Id == userId {
			Users = append(Users[:i], Users[i+1:]...)
			return true
		}
	}
	return false
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Fprintln(w, "Erro ao ler JSON")
		return
	}

	if newUser.Id <= 0 {
		fmt.Fprintln(w, "ID inválido")
		return
	}

	Users = append(Users, newUser)

	fmt.Fprintln(w, "Usuário criado com sucesso")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// converter id da URL para int
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintln(w, "ID do usuário inválido")
		return
	}

	if userId <= 0 {
		fmt.Fprintln(w, "ID deve ser maior que 0")
		return
	}

	// receber JSON do body
	var updatedUser User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		fmt.Fprintln(w, "Erro ao ler JSON")
		return
	}

	// procurar usuário e atualizar
	for i, u := range Users {
		if u.Id == userId {
			// mantém o ID da URL como oficial
			updatedUser.Id = userId

			Users[i] = updatedUser

			fmt.Fprintln(w, "Usuário atualizado com sucesso")
			return
		}
	}

	fmt.Fprintln(w, "Usuário não encontrado")
}
