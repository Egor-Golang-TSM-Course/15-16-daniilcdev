package chiserver

import (
	"net/http"
	"server-context/chiServer/handlers"
	"server-context/chiServer/services"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartServer() {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(1 * time.Second))
	r.Use(middleware.Logger)

	users := handlers.NewUsersHandler(services.NewUsersService())
	r.Get("/users*", users.GetAllUsersHandler)
	r.Get("/users/{}", users.GetUserByIdHandler)
	r.Post("/users/{}", users.CreateUserHandler)

	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("tasks"))
	})

	http.ListenAndServe(":3000", r)
}
