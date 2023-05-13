package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	cfg "github.com/s-matke/abp/abp-back/api_gateway/startup/config"

	accommodationGw "github.com/s-matke/abp/abp-back/common/proto/accommodation_service"
	userGw "github.com/s-matke/abp/abp-back/common/proto/user_service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	fmt.Print("api_gateway -> NewServer")
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	fmt.Print("api_gateway -> initHandlers")
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, catalogueEmdpoint, opts)
	if err != nil {
		panic(err)
	}

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}

}

func (server *Server) Start() {
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000/**", "http://localhost:3000", "https://localhost:3000/**", "https://localhost:3000"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "Access-Control-Allow-Origin", "*"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handlers.CORS(headers, methods, origins)(server.mux)))
}
