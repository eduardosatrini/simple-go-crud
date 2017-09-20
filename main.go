package main

import (
	"fmt"
	"log"
	"simple-crud-go/models"
)

func main() {
	db, err := models.Connection("root:@tcp(127.0.0.1:3306)/register")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := models.All(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range rows {
		fmt.Println(value.Name, "-", value.Email)
	}

	// data := map[string]string{
	// 	"name":  "John Doe II",
	// 	"email": "john2@mail.com",
	// }

	// id, err := models.Create(db, data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("id:", id)

	// id := 1
	// data := map[string]string{
	// 	"name":  "Joanna Doe",
	// 	"email": "joanna@mail.com",
	// }

	// err = models.Update(db, id, data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// row, err := models.Read(db, 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(row.Name)

	// err = models.Delete(db, 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
