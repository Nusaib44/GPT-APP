package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct to represent the JSON request body
type PromptRequest struct {
	Prompt string `json:"prompt"`
}

// Struct to represent the JSON response body
type CodeResponse struct {
	Code string `json:"code"`
}

func main() {
	router := gin.Default()

	// Serve the HTML file
	router.StaticFile("/", "./index.html")

	// Handle POST request to generate code
	router.POST("/generate-code", GenerateCodeHandler)

	router.Run(":8080")
}

func GenerateCodeHandler(c *gin.Context) {
	// Parse JSON request body
	var promptRequest PromptRequest
	if err := c.ShouldBindJSON(&promptRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process prompt and generate code (replace with your implementation)
	generatedCode := GenerateCode(promptRequest.Prompt)

	// Create JSON response
	codeResponse := CodeResponse{
		Code: generatedCode,
	}

	// Send JSON response
	c.JSON(http.StatusOK, codeResponse)
}

// Function to generate code based on the prompt
func GenerateCode(prompt string) string {
	// Replace this with your implementation to interact with GPT-3.5 API
	// Example:
	// - Send the prompt to GPT-3.5 API
	// - Receive generated code from GPT-3.5 API
	// - Return the generated code
	return fmt.Sprintf("Generated code based on the prompt: %s", prompt)
}
