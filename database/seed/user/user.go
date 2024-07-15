package main

import (
	"fmt"
	"log"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/database"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
	"github.com/AwahMarwah/Technical_Test_3/service/user"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var email, password string
	fmt.Print("email: ")
	if _, err := fmt.Scanln(&email); err != nil && err.Error() != common.UnexpectedNewline {
		log.Fatal(err)
	}
	fmt.Print("password: ")
	if _, err := fmt.Scanln(&password); err != nil && err.Error() != common.UnexpectedNewline {
		log.Fatal(err)
	}
	req := userModel.SeedReq{
		Email:    email,
		Password: password,
	}
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		log.Fatal(err)
	}
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	userService := user.NewService(userRepo.NewRepo(db.GormDb))
	if err = userService.Seed(&req); err != nil {
		log.Fatal(err)
	}
	log.Print(common.SuccessfullyCreated)
}
