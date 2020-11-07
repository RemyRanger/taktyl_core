package controllers

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"taktyl.com/m/src/api/rpc"
	"taktyl.com/m/src/models"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Server : server structure
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize : init router and database link
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Event{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// RunREST : launching server
func (server *Server) RunREST(addr string) {
	fmt.Println("Listening to port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// RunGRPC : launching server
func (server *Server) RunGRPC() {
	fmt.Println("Listening to port 50051")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	rpc.RegisterEventServiceServer(s, &Server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
