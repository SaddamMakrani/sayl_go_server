package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	consts "main/const"
	"main/structs"
	"main/util"
	"main/validations"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	productsData, err := util.HttpRequest("GET", consts.PRODUCTS, nil)
	fmt.Println(err)

	var product *structs.Products
	err = json.NewDecoder(bytes.NewBuffer(productsData)).Decode(&product)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return err
	}

	return c.JSON(product)
}

func CreateUser(c *fiber.Ctx) error {
	var user structs.User
	err := json.Unmarshal(c.Request().Body(), &user)
	if err != nil {
		return err
	}
	if ok, err := validations.ValidateUser(user); ok && err == nil {
		customer := structs.Customer{
			User: user,
		}

		responseData, err := util.HttpRequest("POST", consts.CREATE_USER, customer)
		if err != nil {
			fmt.Println("Error reading the response:", err)
			return err
		}

		var userData *structs.CreatedUser
		err = json.NewDecoder(bytes.NewBuffer(responseData)).Decode(&userData)
		if err != nil {
			fmt.Println("Error while decoding response:", err)
			return err
		}
		return c.JSON(userData)

	} else {
		response := fiber.Map{
			"message": "Validation Error",
			"status":  "Error",
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
}
