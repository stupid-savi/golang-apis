package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/stupid-savi/golang-apis/internal/config"
)

func main() {
	// load the config
	confg := config.MustLoad()
	fmt.Println(confg.Env)

	fmt.Println("Welcome to Students APIs")

	//setup router

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Golang APIs"))
	})

	// setup server

	server := http.Server{
		Addr:    confg.Address,
		Handler: router,
	}

	successMsg := color.GreenString("Server is running at %s", confg.Address)
	slog.Info(successMsg)
	fmt.Printf("%s", successMsg)

	// For gracefully stopping the server

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {

			errMsg := color.RedString("Error starting the Server %s", err.Error())
			log.Fatalf("%s", errMsg)

		}
	}()

	<-done

	// after interupption now shutting down
	slog.Info("Shutting Down the Server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to Shutdown Server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
