# Geoecho

This is a Go service and API to get geolocation data given IP addresses.  It has one endpoint to get the geolocation in json format, with the IP address being the only input.  There is also a Swagger UI path where you can try the API service in a browser.

# Requirements

Go version 1.14 is required, and Go modules also need to be enabled.


# Building and Running

Build the executable binary

`go build` which will produce the binary `geoecho`

Using the built binary `geoecho`, invoke

`geoecho <apikey> 8080`

where `<apikey>` is your API key for IPGeolocation's API. You can sign up for an account at https://ipgeolocation.io and get an API key.  Also, `8080` could be replaced by another port.  These arguments are required at the moment.

Similarly you can use `go run` as well if you don't want to build the binary

`go run main.go <apikey> 8080`


# Endpoints

`GET /geolocate/<ip_address>` 

IP geolocation that queries an IP address and returns location data. `<ip_address>` is a valid IPv4 or IPv6 address, i.e. `8.8.8.8` or `2001:0db8:85a3:0000:0000:8a2e:0370:7334`

`GET /swagger/index.html`

Swagger UI meant to be used in browser.  Allows user to test the `geolocate` endpoint in browser.


# Testing

Once the service is running, you can simply hit it with `curl` or [`curlie`](https://curlie.io/) , e.g.

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

The generated code client is in the `/client` folder.  The main file to start from is `geoecho_client.go`.


# Development

Go modules are required - set the environment variable `GO111MODULE` to `on`.  You can either run `go mod download` to download the dependencies, or they will also be downloaded when you run via `go run main.go <apikey> <port>`

Generate the Swagger docs UI via  `swag init -g main.go`

Generate the code client via `swagger generate client -f docs/swagger.json -A geoecho`


# Edges Cases Tested

Valid and invalid IPv4/v6 addresses are handled, where the invalid IP addesses give a graceful error.

```
validIPV4  "10.40.210.253"
invalidIPV4  "1000.40.210.253"
valiIPV6  "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
invalidIPV6  "2001:0db8:85a3:0000:0000:8a2e:0370:7334:3445"
```

White spaces around the IP addesses are handled, e.g. `_8.8.8.8_`

Escaped surrounding white spaces are handled and discarded, e.g. `"http://localhost:8080/geolocate/%208.8.8.8%20"`

Escaped colon characters `%3A` are handled for IPv6 addesses and converted to actual colons `:`, e.g. `http://localhost:8080/geolocate/2001%3A0db8%3A85a3%3A0000%3A0000%3A8a2e%3A0370%3A7334`

Generally invalid IP address arguments with malformated strings are handled with graceful error response

Missing IP addess arguments are also detected and handled.