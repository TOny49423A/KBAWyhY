// 代码生成时间: 2025-08-26 18:26:42
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "encoding/json"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID          uint   `json:"id"`
    Name        string `json:"name"`
    Price       float64 `json:"price"`
    Quantity    uint   `json:"quantity"`
}

// ShoppingCart holds a list of cart items
type ShoppingCart struct {
    Items []CartItem `json:"items"`
}

// AddItemToCart adds an item to the shopping cart
func AddItemToCart(cart *ShoppingCart, item CartItem) *ShoppingCart {
    cart.Items = append(cart.Items, item)
    return cart
}

// CalculateTotal calculates the total price of the shopping cart
func CalculateTotal(cart *ShoppingCart) float64 {
    var total float64
    for _, item := range cart.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger(), middleware.Recover())

    // Define routes
    e.POST("/cart", addCartItem)
    e.GET("/cart", getCart)
    e.DELETE("/cart/:id", removeCartItem)
    e.GET("/cart/total", calculateTotal)

    // Start the server
    e.Start(":8080")
}

// addCartItem adds a new item to the shopping cart
func addCartItem(c echo.Context) error {
    var item CartItem
    if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid item data"})
    }
    cart := ShoppingCart{Items: []CartItem{}}
    AddItemToCart(&cart, item)
    return c.JSON(http.StatusOK, cart)
}

// getCart retrieves the current state of the shopping cart
func getCart(c echo.Context) error {
    cart := ShoppingCart{Items: []CartItem{}}
    // In a real-world scenario, you would retrieve the cart from a database or session
    return c.JSON(http.StatusOK, cart)
}

// removeCartItem removes an item from the shopping cart by ID
func removeCartItem(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    cart := ShoppingCart{Items: []CartItem{}}
    // In a real-world scenario, you would update the cart in a database or session
    return c.JSON(http.StatusOK, map[string]string{"message": "Item removed"})
}

// calculateTotal calculates the total price of the shopping cart
func calculateTotal(c echo.Context) error {
    cart := ShoppingCart{Items: []CartItem{}}
    // In a real-world scenario, you would retrieve the cart from a database or session
    total := CalculateTotal(&cart)
    return c.JSON(http.StatusOK, map[string]float64{"total": total})
}