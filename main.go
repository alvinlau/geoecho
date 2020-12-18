// GET /geolocate/<ip_address> IP geolocation that queries an IP address and returns location data, using a free service such as ipgeolocation.io
// GET /swagger/index.html an endpoint that displays Swagger/OpenAPI specification of the endpoint you created above. We recommend using echo-swagger for this.
// Finally, once you've generated your OpenAPI specification you will use Swagger codegen to generate a Golang SDK so that other Go applications can use and access your API. Add the client code in a folder of your submission named /client.

package main

import (
  "net/http"
  "io/ioutil"
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
  e.GET("/geolocate", geolocate)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func geolocate(c echo.Context) error {
  resp, err := http.Get("https://httpbin.org/get")
  if err != nil {
    // handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    // handle error
  }
  return c.String(http.StatusOK, string(body))
}