// 代码生成时间: 2025-08-18 01:40:36
Author: [Your Name]
Date: [Today's Date]
*/

package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"

    "github.com/labstack/echo"
)

// InventoryItem 表示库存项
type InventoryItem struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Quantity int    `json:"quantity"`
}

// Inventory 管理库存项
type Inventory struct {
    items []InventoryItem
    mutex sync.Mutex
}

// NewInventory 创建新的库存实例
func NewInventory() *Inventory {
    return &Inventory{items: []InventoryItem{}}
}

// AddItem 向库存中添加一个项
func (i *Inventory) AddItem(item InventoryItem) error {
    i.mutex.Lock()
    defer i.mutex.Unlock()
    i.items = append(i.items, item)
    return nil
}

// UpdateItem 更新库存中的一个项
func (i *Inventory) UpdateItem(id int, quantity int) error {
    i.mutex.Lock()
    defer i.mutex.Unlock()
    for index, item := range i.items {
        if item.ID == id {
            i.items[index].Quantity = quantity
            return nil
        }
    }
    return fmt.Errorf("item with id %d not found", id)
}

// GetItemByID 通过ID获取库存项
func (i *Inventory) GetItemByID(id int) (*InventoryItem, error) {
    i.mutex.Lock()
    defer i.mutex.Unlock()
    for _, item := range i.items {
        if item.ID == id {
            return &item, nil
        }
    }
    return nil, fmt.Errorf("item with id %d not found", id)
}

// GetAllItems 获取所有库存项
func (i *Inventory) GetAllItems() []InventoryItem {
    i.mutex.Lock()
    defer i.mutex.Unlock()
    return i.items
}

func main() {
    e := echo.New()

    // 实例化库存
    inventory := NewInventory()

    // 初始库存项
    initialItems := []InventoryItem{{ID: 1, Name: "Laptop", Quantity: 10}, {ID: 2, Name: "Mouse", Quantity: 15}}
    for _, item := range initialItems {
        _ = inventory.AddItem(item) // 这里忽略了错误处理以简化示例
    }

    // 添加路径和处理函数
    e.GET("/items", func(c echo.Context) error {
        items := inventory.GetAllItems()
        return c.JSON(http.StatusOK, items)
    })

    e.POST("/items", func(c echo.Context) error {
        item := InventoryItem{}
        if err := c.Bind(&item); err != nil {
            return err
        }
        if err := inventory.AddItem(item); err != nil {
            return err
        }
        return c.JSON(http.StatusCreated, item)
    })

    e.PUT("/items/:id", func(c echo.Context) error {
        id := c.Param("id")
        itemID, _ := strconv.Atoi(id)
        quantity := c.QueryParam("quantity")
        qty, _ := strconv.Atoi(quantity)
        if err := inventory.UpdateItem(itemID, qty); err != nil {
            return err
        }
        return c.NoContent(http.StatusNoContent)
    })

    e.GET("/items/:id", func(c echo.Context) error {
        id := c.Param("id")
        itemID, _ := strconv.Atoi(id)
        item, err := inventory.GetItemByID(itemID)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, item)
    })

    // 启动服务器
    log.Fatal(e.Start(":8080"))
}