package main

import (
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/server"
)

// @title Correção de Atividades API
// @version 1.0
// @description This is an example API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
