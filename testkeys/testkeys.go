// Package testkeys contains RSA keys and TLS certificates for use in development and testing
package testkeys

// RSA Key
//
// RSA key is generated with following commands:
//     openssl genrsa -out jwt.pem 2048
//     openssl rsa -in jwt.pem -pubout > jwt.pub.pem
//
const (
	Private = "testkeys/jwt.pem"
	Public  = "testkeys/jwt.pub.pem"
)

// TLS certificate
//
// Generated with following command:
//     go run generate_cert.go -ca -duration=87600h0m0s -host="localhost,127.0.0.1"
//
// generate_cert.go is available at http://golang.org/src/pkg/crypto/tls/generate_cert.go
const (
	Key  = "testkeys/key.pem"
	Cert = "testkeys/cert.pem"
)
