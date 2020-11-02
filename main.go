package main
import (
  "os"
  "fmt"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)
func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  port := os.Getenv("PORT")
  fmt.Println("PORT:", port)
  e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
