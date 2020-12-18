# geoecho

This is a Go service and API to get geolocation data given IP addresses.  It has one endpoint to get the geolocation in json format, with the IP address being the only input.  There is also a Swagger UI path where you can try the API service in a browser.

# Requirements

Go version 1.14 is required, and Go modules also need to be enabled.


# Building and Running

Build the executable binary

`go build`

Uing the built binary, invoke

`geoecho <apikey> 8080`

where `<apikey>` is your API key for IPGeolocation's API, and `8080` could be replaced by another port.  These arguments are required.

Similarly you can use `go run` as well if you don't want to build the binary

`go run main.go <apikey> 8080`


# Endpoints

`GET /geolocate/<ip_address>` 

IP geolocation that queries an IP address and returns location data. `<ip_address>` is a valid IPv4 or IPv6 address, i.e. `8.8.8.8` or `2001:0db8:85a3:0000:0000:8a2e:0370:7334`

`GET /swagger/index.html`


# Testing

Once the service is running, you can simply hit it with `curl` or `curlie`, e.g.

```
curl -X GET "http://localhost:8080/geolocate/8.8.8.8" -H  "accept: application/json"
```

or

```
curlie localhost:8080/geolocate/2001:0db8:85a3:0000:0000:8a2e:0370:7334
```

# Swagger Docs

The Swagger UI is accessible via `http://<host>:8080/swagger/index.html` once the service is running, where `<host>` is simply `localhost` if it is running on your local machine. The service will still work on other ports, just that unfortunately it has to run on port `8080` for the Swagger UI to work.


# Swagger Code Client

The generated code client is in the `/client` folder


# Development

Go modules are required - set the environment variable `GO111MODULE` to `on`.  You can either run `go mod download` to download the dependencies, or they will also be downloaded when you run via `go run main.go ...`

Generate the Swagger docs UI via  `swag init -g main.go`

Generate the code client via `swagger generate client -f docs/swagger.json -A geoecho`


# Test Cases