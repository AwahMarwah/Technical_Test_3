package main

import (
	"log"

	"github.com/AwahMarwah/Technical_Test_3/database"
	"github.com/AwahMarwah/Technical_Test_3/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = router.Run(db); err != nil {
		log.Fatal(err)
	}
}
