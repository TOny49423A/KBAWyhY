// 代码生成时间: 2025-09-01 05:01:35
package main

import (
    "net/http"
    "fmt"
    "github.com/labstack/echo"
    "github.com/go-playground/validator/v10"
)

// Form represents the structure of the form with input fields
type Form struct {
    Name  string `json:"name" validate:"required,min=2,max=100"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required,gte=18"`
}

// ValidationResult contains the result of the validation
type ValidationResult struct {
    IsValid bool   `json:"isValid"`
    Errors   []string `json:"errors"`
}

// validateForm validates the form and returns a ValidationResult
func validateForm(c echo.Context, form Form) (ValidationResult, error) {
    // Create a validator instance
    v := validator.New()
    
    // Validate the form
    if err := v.Struct(form); err != nil {
        // Extract validation errors
        errors := err.(validator.ValidationErrors)
        var validationErrors []string
        for _, e := range errors {
            validationErrors = append(validationErrors, e.Translate(e))
        }
        return ValidationResult{IsValid: false, Errors: validationErrors}, nil
    }
    
    return ValidationResult{IsValid: true, Errors: []string{}}, nil
}

func main() {
    e := echo.New()
    
    // Define a route for POST request that accepts a JSON body
    e.POST("/form", func(c echo.Context) error {
        // Decode the JSON body into a Form struct
        var form Form
        if err := c.Bind(&form); err != nil {
            return err
        }
        
        // Validate the form
        result, err := validateForm(c, form)
        if err != nil {
            return err
        }
        
        // Check if the form is valid
        if result.IsValid {
            return c.JSON(http.StatusOK, map[string]interface{}{"message": "Form is valid"})
        } else {
            return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": result.Errors})
        }
    })
    
    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
