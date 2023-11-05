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

// Saved cart data
var userCarts = make(map[string]structs.UserCart)

func CreateCart(c *fiber.Ctx) error {
	mobileNumber := c.Params("mobile_number")

	// get user's cart or create a new one if it doesn't exist
	cart, exists := userCarts[mobileNumber]
	if !exists {
		cart = structs.UserCart{}
	}

	// Parse request body into a CartItem
	var item structs.CartItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"Error":   err,
		})
	}

	if item.Quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"Error":   "Quantity should not be 0",
		})
	}

	// Add item to the user's cart
	cart.Items = append(cart.Items, item)

	// Update or add the cart
	userCarts[mobileNumber] = cart

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Variant added to cart successfully",
	})
}

func RemoveItemCart(c *fiber.Ctx) error {
	mobileNumber := c.Params("mobile_number")

	// Retrieve user's cart
	cart, exists := userCarts[mobileNumber]
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User's cart not found",
		})
	}

	// Parse request body into CartItem
	var itemToRemove structs.CartItem
	if err := c.BodyParser(&itemToRemove); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	// Find and remove specified item from user's cart
	var updatedItems []structs.CartItem
	var found bool
	for _, item := range cart.Items {
		if item.VariantID == itemToRemove.VariantID {
			found = true
			if item.Quantity > 0 && item.Quantity > itemToRemove.Quantity {
				item.Quantity -= itemToRemove.Quantity
				updatedItems = append(updatedItems, item)
			}
		} else {
			updatedItems = append(updatedItems, item)
		}
	}

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Item not found in the cart",
		})
	}

	// Update user's cart
	cart.Items = updatedItems
	userCarts[mobileNumber] = cart

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product removed from cart successfully",
	})
}

func GetUserCart(c *fiber.Ctx) error {
	mobileNumber := c.Params("mobile_number")

	cart, exists := userCarts[mobileNumber]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User's cart not found",
		})
	}

	// Prepare object to store cart items
	var cartItems []structs.ViewCartItem

	// Convert cart items to desired format
	for _, item := range cart.Items {
		variant := GetVariantFromID(item.VariantID)
		cartItems = append(cartItems, structs.ViewCartItem{
			VariantID:    item.VariantID,
			VariantTitle: variant.Title,
		})
	}

	return c.Status(fiber.StatusOK).JSON(cartItems)
}

func GetVariantFromID(variantID int64) structs.Variant {

	resp, err := util.HttpRequest("GET", consts.VARIANT+fmt.Sprintf("%d", variantID)+".json", nil)
	fmt.Println(err, "get variant err")

	var response structs.VariantResponse

	// Parse the JSON response
	decoder := json.NewDecoder(bytes.NewBuffer(resp))
	if err := decoder.Decode(&response); err != nil {
		return structs.Variant{}
	}
	return response.Variant
}

func CreateOrder(c *fiber.Ctx) error {
	mobileNumber := c.Params("mobile_number")

	// Retrieve the user's cart
	cart, exists := userCarts[mobileNumber]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User's cart not found",
		})
	}

	requestBody := map[string]interface{}{
		"order": map[string]interface{}{
			"line_items": cart.Items,
		},
	}
	shopifyOrderID, err := util.HttpRequest("POST", consts.CREATE_ORDER+consts.ORDER_UTM_PARAMS, requestBody)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return err
	}

	// Parse the Shopify order response to get the order ID
	var shopifyOrderResponse structs.ShopifyOrder

	err = json.NewDecoder(bytes.NewBuffer(shopifyOrderID)).Decode(&shopifyOrderResponse)
	if err != nil {
		return err
	}
	fmt.Println(shopifyOrderResponse, "shopifyOrderResponse")

	return c.Status(fiber.StatusCreated).JSON(shopifyOrderResponse)
}

func GetOrderFromUtmSource(c *fiber.Ctx) error {
	paramValue := c.Query("utm_source")
	fmt.Println(consts.CREATE_ORDER+consts.GET_ORDER_BY_UTM+paramValue, "consts.CREATE_ORDER+consts.GET_ORDER_BY_UTM+paramValue")
	resp, err := util.HttpRequest("GET", consts.CREATE_ORDER+consts.GET_ORDER_BY_UTM+paramValue, nil)
	fmt.Println(err, "get variant err")

	var response structs.ShopifyOrderResponseByUtm

	err = json.NewDecoder(bytes.NewBuffer(resp)).Decode(&response)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return err
	}

	return c.JSON(response)
}
