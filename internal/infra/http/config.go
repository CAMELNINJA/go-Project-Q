package http

type Config struct {
	Address             string `short:"a" long:"address" env:"ADDRESS" description:"Service address" `
	JWTPrivateKey       string `long:"jwt-private-key" env:"JWT_PRIVATE_KEY" description:"Path to JWT private key" `
	PushServiceEndpoint string `long:"push-service-endpoint" env:"PUSH_SERVICE_ENDPOINT" description:"push service endpoint" `
}
