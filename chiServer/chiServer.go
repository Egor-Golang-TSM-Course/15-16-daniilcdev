package chiserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"server-context/chiServer/handlers"
	mw "server-context/chiServer/middleware"
	"server-context/chiServer/services"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartServer(parentContext context.Context, lifeTime time.Duration) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(1 * time.Second))
	r.Use(mw.RandomSleep(0.3, 2*time.Second))

	users := handlers.NewUsersHandler(
		services.NewUsersService(
			&services.InMemoryStorage[services.UserId, services.User]{}),
	)

	r.Get("/users*", users.GetAllUsersHandler)
	r.Get("/users/{}", users.GetUserByIdHandler)
	r.Post("/users/{}", users.CreateUserHandler)

	tasks := handlers.NewTasksHandler(
		services.NewTasksService(
			&services.InMemoryStorage[services.OwnerId, []services.Task]{}),
	)

	r.Get("/tasks*", tasks.GetAllTasksHandler)
	r.Post("/tasks*", tasks.CreateTaskHandler)
	r.Get("/tasks/{}", tasks.GetTaskByIdHandler)
	r.Get("/tasks/all*", tasks.GetUsersTasksHandler)

	fmt.Println("starting chi server on port 3000...")

	s := &http.Server{Addr: ":3000", Handler: r}

	ctx, cancel := context.WithTimeout(parentContext, lifeTime)
	defer cancel()
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	go func() {
		<-ctx.Done()
		err := s.Shutdown(ctx)
		if err != nil {
			fmt.Printf("error: server shutdown failed - %v\n", err)
		}
	}()

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("server shutdown: %v\n", err)
	}
}
