package main

import (
	"context"
	"keycut/keycut/app"
	"keycut/keycut/mux"
	"keycut/keycut/settings"
	"keycut/keycut/storage"
	"os"
	"strconv"

	"errors"
	"fmt"
	"net"
	"net/http"
)

func main() {
	dbHost := os.Getenv("DBHOST")
	dbPort, err := strconv.Atoi(os.Getenv("DBPORT"))

	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	addr := os.Getenv("ADDR")

	dbSettings := settings.DatabaseSettings{
		Host:   dbHost,
		Port:   dbPort,
		User:   dbUser,
		Dbname: dbName,
	}

	dbConnection, err := storage.NewDatabaseConnection(dbSettings)

	if err != nil {
		panic(err)
	}

	mux := mux.NewMux()

	defer dbConnection.Close()

	app := app.New(mux)

	ctx, cancelCtx := context.WithCancel(context.Background())

	serverOne := &http.Server{
		Addr:    addr,
		Handler: app.Mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
