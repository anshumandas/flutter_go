package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "github.com/flutter_go/app/database"
	"github.com/flutter_go/app/routers"
	"github.com/flutter_go/settings"
	"github.com/rs/cors"
)

func init() {
	settings.InitSettings()
	db.InitPostgre()
	db.InitRedis()
}

func main() {

	addr := fmt.Sprintf(":%s", settings.ServerSettings.Addr)
	port := addr
	router := cors.Default().Handler(routers.InitRouter())
	readTimeout := settings.ServerSettings.ReadTimeout
	writeTimeout := settings.ServerSettings.WriteTimeout
	mxHdrBytes := settings.ServerSettings.MaxHeaderBytes

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: mxHdrBytes,
	}

	var err error

	err = s.ListenAndServe()

	log.Printf("Server is running on port:  %s", port)

	go func() {
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds")
	}
	log.Println("Server exiting")

}
