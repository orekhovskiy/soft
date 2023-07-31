package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	router := gin.Default()
	router.GET("/calculate", authenticateMiddleware, calculateHandler)
	router.Run(":8080")
}

// Authentication middleware
func authenticateMiddleware(c *gin.Context) {
	userAccess := c.GetHeader("User-Access")
	if userAccess != "superuser" {
		fmt.Println("Access denied. Wrong header 'User-Access'")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		c.Abort()
		return
	}
	c.Next()
}

// Handler for performing arithmetic operations
func calculateHandler(c *gin.Context) {
	var req struct {
		Expression string `json:"expression"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	expression := req.Expression
	if expression == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expression"})
		return
	}

	result, err := calculate(expression)
	fmt.Println(result, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

// Function to perform arithmetic operations
func calculate(expression string) (float64, error) {
	// Split string for operands
	operations := strings.Split(expression, "+")
	// Execute sum
	var result float64
	for _, op := range operations {
		// Checking for subtract
		if strings.Contains(op, "-") {
			subOp := strings.Split(op, "-")
			// First one is always positive
			num, err := strconv.ParseFloat(subOp[0], 64)
			if err != nil {
				return 0, fmt.Errorf("invalid expression: %s", expression)
			}
			result += num
			// The rest are negative
			for i := 1; i < len(subOp); i++ {
				num, err := strconv.ParseFloat(subOp[i], 64)
				if err != nil {
					return 0, fmt.Errorf("invalid expression: %s", expression)
				}
				result -= num
			}
		} else {
			num, err := strconv.ParseFloat(op, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid expression: %s", expression)
			}
			result += num
		}
	}
	return result, nil
}
