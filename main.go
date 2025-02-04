package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct{
	ID 			int 	`json:"id"`
	Completed 	bool	`json:"completed"`
	Body 		string	`json:"body"`
}

func main() {

	// if any changes are made we
	fmt.Println("Hello World")

	// var myName string = "Chinmay"
	// const mySecondName string = "Atharva"

	// myThirdName := "CK"

	// fmt.Println(myName)
	// fmt.Println(mySecondName)
	// fmt.Println(myThirdName)

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})
	
	// create a TODO
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // {id:0, completed:False, body: ""}

		if err := c.BodyParser(todo); err !=nil {
			return err
		} // take the data send by user and assign it to do

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error":"Todo body is requred"}) 
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Update a TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})

	// delete a TODO
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id{
				todos = append(todos[:i], todos[i+1:]...) // ... used to unpack the values
				return c.Status(200).JSON(fiber.Map{"Success": true})
			}
		}
 
		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})
	log.Fatal(app.Listen(":"+PORT))
	// till now if any changes are made we need to kill the program and rerun it
	// then installed air to change this 

	// use postman to test api calls 
	// create a new collection 
	// added a get connection request http://localhost:4000/


	 

}
