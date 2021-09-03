package app

import (
	"fmt"
	"log"
	"net/http"
	"yaml-to-openmetric/internal/config"
	handlers "yaml-to-openmetric/internal/delivery/http"
	"yaml-to-openmetric/internal/repository"
	"yaml-to-openmetric/internal/services"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	cfg := config.GetConfig(configPath)

	repos := repository.NewRepositories()


	services := services.NewServices(services.Deps{
		Repos:    repos,
	})

	router := handlers.NewHandlers(services).Init()

	s := &http.Server{
		Addr:           cfg.App.ServerAddr,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16, // 65536,
	}

	go waitQuit()

	fmt.Println("Server start")
	log.Fatal(s.ListenAndServe())
}

func waitQuit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Println("\r\nThe WEB-server is stopped. Signal:", quit)

	time.Sleep(time.Millisecond * 100)
	os.Exit(1)
}