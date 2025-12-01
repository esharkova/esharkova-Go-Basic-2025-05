package main

import (
	app "TaskManagement/internal/appdb"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	application := app.New()

	application.Start()

	application.Stop()

}
