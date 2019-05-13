package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/proximax-storage/faucet-backend"
	"github.com/proximax-storage/faucet-backend/db"
	"github.com/proximax-storage/faucet-backend/routes"
	"github.com/proximax-storage/faucet-backend/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	config, err := Faucet.LoadConfig()
	if err != nil {
		panic(err)
	}

	utils.Logger(0, "finished loading server config")

	utils.Logger(0, fmt.Sprintf("starting: %v", config.Server.Description))

	s, _ := utils.Marshal(config)
	fmt.Println(string(s))

	Faucet.InitClient()

	db.Init()

	go utils.Counter(config.Logging.ErrCtrl.MaxNumErr)

	gin.SetMode(gin.ReleaseMode)

	route := routes.NewRouter()

	srv := &http.Server{
		Addr:    config.FormatServer(),
		Handler: route,
	}

	go func() {
		// service connections
		utils.Logger(1, "listening API server on %v port[%v]", config.Server.Host, config.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Logger(2, "Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger(3, "Server Shutdown:", err)
		panic(nil)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		utils.Logger(1, "timeout of 5 seconds.")
	}
	utils.Logger(1, "Server exiting")
}
