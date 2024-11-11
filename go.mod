module go-assistant

go 1.22.6

toolchain go1.22.9

require (
	github.com/briandowns/openweathermap v0.19.0
	github.com/huin/goupnp v1.3.0
	github.com/lib-x/edgetts v0.3.7
	gopkg.in/yaml.v3 v3.0.1
)

// replace github.com/lib-x/edgetts => ../edgetts

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
)
