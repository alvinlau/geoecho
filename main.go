// GET /geolocate/<ip_address> IP geolocation that queries an IP address and returns location data, using a free service such as ipgeolocation.io
// GET /swagger/index.html an endpoint that displays Swagger/OpenAPI specification of the endpoint you created above. We recommend using echo-swagger for this.
// Finally, once you've generated your OpenAPI specification you will use Swagger codegen to generate a Golang SDK so that other Go applications can use and access your API. 
// Add the client code in a folder of your submission named /client.

package main

import (
  // "encoding/json"
  // "fmt"
  "os"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/k0kubun/pp"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/swaggo/echo-swagger"
  _ "github.com/alvinlau/geoecho/swagger"
  //  echoSwagger "github.com/swaggo/echo-swagger"
)


// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/geolocate/:ip", geolocate)
  e.GET("/swagger/*", echoSwagger.WrapHandler)

  // Start server
  e.Logger.Fatal(e.Start(":" + os.Args[2]))
}

// Handlers

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func geolocate(c echo.Context) error {
  // parse the IP address
  ip := c.Param("ip") 
  pp.Println(ip)

  if strings.TrimSpace(ip) == "" || len(ip) == 0 {
    return c.String(http.StatusBadRequest, "missing ip address input")
  }

  // get API key
  apikey := os.Args[1]
  pp.Println(apikey)
  if strings.TrimSpace(apikey) == "" {
    return c.String(http.StatusInternalServerError, "API key not set")
  }

  // call ipgeolocation API
  resp, err := http.Get("https://api.ipgeolocation.io/ipgeo?apiKey=" + apikey + "&ip=" + ip)
  
  if err != nil {
    return c.String(http.StatusBadRequest, "invalid request to geolocation API")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return c.String(http.StatusInternalServerError, "unable to parse response from geolocation API")
  }
  return c.String(http.StatusOK, string(body))
}

func swagger(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}