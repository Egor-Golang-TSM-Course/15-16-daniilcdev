package handlers

import (
	"encoding/json"
	"net/http"
	"path"
	"server-context/chiServer/services"
)

type UsersHandler struct {
	us *services.UsersService
}

func NewUsersHandler(us *services.UsersService) *UsersHandler {
	return &UsersHandler{us: us}
}

func (uh *UsersHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	all, _ := uh.us.GetAllUsers()

	p, err := json.Marshal(struct {
		Users []services.UserDto `json:"users"`
	}{
		Users: all,
	})

	if err != nil {
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}

	w.Write(p)
}

func (uh *UsersHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)

	user, err := uh.us.GetUser(services.UserId(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Write(p)
}

func (uh *UsersHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)

	userId, err := uh.us.CreateUser(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := json.Marshal(services.UserDto{
		Id:   userId,
		Name: name,
	})

	if err != nil {
		http.Error(w, "serialization error", http.StatusBadRequest)
		return
	}

	w.Write(p)
}
