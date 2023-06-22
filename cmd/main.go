package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ivan12093/VK-DB-PROJECT/internal/network"
	"github.com/ivan12093/VK-DB-PROJECT/internal/network/handlers"
	"github.com/ivan12093/VK-DB-PROJECT/internal/repository"
	"github.com/ivan12093/VK-DB-PROJECT/internal/usecase"
	"github.com/ivan12093/VK-DB-PROJECT/internal/utils/check"
	"github.com/jackc/pgx/v4/pgxpool"
)

const password = "ac322f35-e71e-47dd-a580-894bf3e6c460"

func main() {
	APIPort := "5000"
	DSN := fmt.Sprintf("host=localhost port=5432 user=forum_user password=%s dbname=forum sslmode=disable", password)
	APIAddr := fmt.Sprintf("0.0.0.0:%v", APIPort)

	_, err := check.GetInstance()
	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := pgxpool.ParseConfig(DSN)
	config.MaxConns = 2000
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
		return
	}

	gin.SetMode("release")

	repo := repository.NewRepository(db)
	usecases := usecase.NewUsecases(repo)
	handlrs := handlers.NewHandlers(usecases)

	e := network.InitRoutes(handlrs)
	err = e.Run(APIAddr)
	if err != nil {
		log.Fatal(err)
	}
}
