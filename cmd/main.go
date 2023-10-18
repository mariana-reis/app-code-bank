package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/mariana-reis/app.code.bank/application/grpc"
	"github.com/mariana-reis/app.code.bank/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
