package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	consts "main/const"
	"main/structs"
	"main/util"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	productsData, err := util.HttpRequest("GET", consts.PRODUCTS, nil)
	fmt.Println(string(productsData), err)

	var product *structs.Products
	err = json.NewDecoder(bytes.NewBuffer(productsData)).Decode(&product)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return err
	}

	return c.JSON(product)
}
