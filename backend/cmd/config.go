package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/internal/routes"
	"github.com/Alexsoup97/message-app/internal/service"
	"github.com/go-chi/chi/v5"
)

type IService interface{}

type router interface{}

type EndpointContainer struct {
	service IService
	router  chi.Router
	prefix  string
}

func setupEndpoints(db *db.Storage, router chi.Router) {

	endpointContainers := make([]EndpointContainer, 0)

	//User Service
	service := service.CreateUserService(db)
	route := routes.CreateUserRouter(service)
	endpointContainers = append(endpointContainers, EndpointContainer{
		service: service,
		router:  route,
		prefix:  "/user",
	})

	route = routes.CreateMessageRouter(db)
	endpointContainers = append(endpointContainers, EndpointContainer{
		service: service,
		router:  route,
		prefix:  "/messages",
	})

	//Mounting routers
	for _, endpoint := range endpointContainers {
		router.Mount("/api"+endpoint.prefix, endpoint.router)
	}
}

func frontendConfigure(router chi.Router) {

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}

	router.Handle("/*", spa)

}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
