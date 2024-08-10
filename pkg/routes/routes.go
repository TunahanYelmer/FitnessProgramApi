package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and defines the routes
func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Define routes
    router.POST("/programs/user", createProgram)
    router.PUT("/programs/:id", updateProgram)
    router.DELETE("/programs/:id", deleteProgram)

    return router
}

// Handler to create a new program
func createProgram(c *gin.Context) {
    user := c.Param("user")
    fmt.Println("CreateProgram called")
    c.JSON(http.StatusOK, gin.H{"message": "Program created successfully"})
}

// Handler to update an existing program
func updateProgram(c *gin.Context) {
    id := c.Param("id")
    // Placeholder for updating a program
    fmt.Printf("UpdateProgram called for ID: %s\n", id)
    c.JSON(http.StatusOK, gin.H{"message": "Program updated successfully"})
}

// Handler to delete a program
func deleteProgram(c *gin.Context) {
    id := c.Param("id")
    // Placeholder for deleting a program
    fmt.Printf("DeleteProgram called for ID: %s\n", id)
    c.JSON(http.StatusOK, gin.H{"message": "Program deleted successfully"})
}