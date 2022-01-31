package main

import (
	"Products/pkg/api"
	"Products/pkg/data"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff"
)

var (
	host       = os.Getenv("POSTGRES_HOST_SERVER")
	port       = os.Getenv("POSTGRES_PORT_SERVER")
	user       = os.Getenv("POSTGRES_USER_SERVER")
	dbname     = os.Getenv("POSTGRES_DB_NAME_SERVER")
	password   = os.Getenv("POSTGRES_PASSWORD_SERVER")
	sslmode    = os.Getenv("POSTGRES_SSLMODE_SERVER")
	portServer = os.Getenv("SERVER_OUT_PORT")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	if password == "" {
		password = "password"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
	if portServer == "" {
		portServer = "8080"
	}
}

func main() {
	conn, err := initConnection()
	if err != nil {
		logrus.Fatal(err)
	}
	r := mux.NewRouter()
	userData := data.NewProductData(conn)
	api.InitConnectionToServer(r, *userData)
	r.Use(mux.CORSMethodMiddleware(r))
	listener, err := net.Listen("tcp", fmt.Sprint(":"+portServer))
	if err != nil {
		log.Fatal("Server Listen port...", err)
	}
	logrus.Info("Server works!")
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}

func initConnection() (*gorm.DB, error){
	var conn *gorm.DB
	back := config()
	var err error
	for {
		timeWait := back.NextBackOff()
		time.Sleep(timeWait)
		conn, err = data.GetConnection(host, port, user, dbname, password, sslmode)
		if err != nil {
			logrus.Error("we wait connect to redis, time: ", timeWait)
		} else {
			break
		}
	}
	return conn, err
}

func config() *backoff.ExponentialBackOff {
	back := backoff.NewExponentialBackOff()
	back.MaxInterval = 20 * time.Second
	back.Multiplier = 2
	return back
}