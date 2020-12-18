// GET /geolocate/<ip_address> IP geolocation that queries an IP address and returns location data, using a free service such as ipgeolocation.io
// GET /swagger/index.html an endpoint that displays Swagger/OpenAPI specification of the endpoint you created above. We recommend using echo-swagger for this.
// Finally, once you've generated your OpenAPI specification you will use Swagger codegen to generate a Golang SDK so that other Go applications can use and access your API. 
// Add the client code in a folder of your submission named /client.

package main

import (
  // "encoding/json"
  // "fmt"
  "os"
  "net"
  "strings"
  "net/http"
  "net/url"
  "io/ioutil"
  "github.com/k0kubun/pp"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/swaggo/echo-swagger"
  _ "github.com/alvinlau/geoecho/docs"
)


// @title Geolocation API
// @version 1.0
// @description API for getting geolocation info

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
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

// geolocate godoc
// @Summary returns geolocation info given ip address
// @Description inquires IP geolocation's API to get details in json format
// @Tags geolocate
// @Accept plain
// @Produce json
// @Param ip path string true "valid ipv4/v6 address"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "invalid or missing ip address"
// @Failure 500 {string} string "unable to parse response from geolocation API"
// @Router /geolocate/{ip} [get]
func geolocate(c echo.Context) error {
  // parse the IP address
  ip := c.Param("ip") 
  pp.Println(ip)

  unspaced_ip, err := url.QueryUnescape(ip)
  if err != nil {
    return c.String(http.StatusBadRequest, "invalid ip format")
  }

  trimmed_ip := strings.TrimSpace(unspaced_ip)
  if trimmed_ip == "" || len(trimmed_ip) == 0 {
    return c.String(http.StatusBadRequest, "missing ip address input")
  }
  
  escaped_ip, err := url.PathUnescape(trimmed_ip)
  if err != nil {
    return c.String(http.StatusBadRequest, "invalid ip format")
  }  

  pp.Println(escaped_ip)
  if net.ParseIP(escaped_ip) == nil {
    // example valid and invalid ipv4/v6 addresses
    // https://golangbyexample.com/validate-an-ip-address-in-go/
    // validIPV4 := "10.40.210.253"
    // invalidIPV4 := "1000.40.210.253"
    // valiIPV6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
    // invalidIPV6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334:3445"
    msg := "invalid ip address, make sure it is valid ip v4 or v6 address"
    return c.String(http.StatusBadRequest, msg)
  }

  // get API key
  apikey := os.Args[1]
  pp.Println(apikey)
  if strings.TrimSpace(apikey) == "" {
    return c.String(http.StatusInternalServerError, "API key not set")
  }

  // call ipgeolocation API
  resp, err := http.Get("https://api.ipgeolocation.io/ipgeo?apiKey=" + apikey + "&ip=" + escaped_ip)
  
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
