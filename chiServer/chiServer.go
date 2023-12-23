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

	tasks := handlers.NewTasksHandler(services.NewTasksService())
	r.Get("/tasks*", tasks.GetAllTasksHandler)
	r.Post("/tasks*", tasks.CreateTaskHandler)
	r.Get("/tasks/{}", tasks.GetTaskByIdHandler)
	r.Get("/tasks/all*", tasks.GetUsersTasksHandler)

	http.ListenAndServe(":3000", r)
}
