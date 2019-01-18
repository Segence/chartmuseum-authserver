package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	cmAuth "github.com/chartmuseum/auth"
	"github.com/gin-gonic/gin"
)

var (
	tokenGenerator    *cmAuth.TokenGenerator
	tokenExpiry       = flag.Duration("token-expiry", time.Minute*5, "The duration that the generated token is valid for")
	requiredGrantType = flag.String("required-grant-type", "client_credentials", "The grant type name to request a token from the server")
	masterAccessKey   = flag.String("master-access-key", "MASTERKEY", "The key used to request a token from the server")
	privateKeyPath    = flag.String("private-key-path", "../config/server.key", "The file path to the private key file")
	servicePort       = flag.Int("service-port", 8080, "The HTTP port to use")
)

func oauthTokenHandler(c *gin.Context) {
	authHeader := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	if authHeader != *masterAccessKey {
		c.JSON(401, gin.H{"error": fmt.Sprintf(authHeader)})
		return
	}

	grantType := c.Query("grant_type")
	if grantType != *requiredGrantType {
		c.JSON(400, gin.H{"error": fmt.Sprintf("grant_type must equal %s", *requiredGrantType)})
		return
	}

	scope := c.Query("scope")
	parts := strings.Split(scope, ":")
	if len(parts) != 3 || parts[0] != cmAuth.AccessEntryType {
		c.JSON(400, gin.H{"error": fmt.Sprintf("scope is missing or invalid")})
		return
	}

	access := []cmAuth.AccessEntry{
		{
			Name:    parts[1],
			Type:    cmAuth.AccessEntryType,
			Actions: strings.Split(parts[2], ","),
		},
	}
	accessToken, err := tokenGenerator.GenerateToken(access, *tokenExpiry)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"access_token": accessToken})
}

func main() {

	flag.Parse()

	fmt.Println("Configuration:")
	fmt.Printf("  Token expiry:        %v\n", *tokenExpiry)
	fmt.Printf("  Required grant type: %s\n", *requiredGrantType)
	fmt.Printf("  Master access key:   %s\n", *masterAccessKey)
	fmt.Printf("  Private key path:    %s\n", *privateKeyPath)
	fmt.Printf("  Service port:        %d\n", *servicePort)
	fmt.Println("")

	var err error
	tokenGenerator, err = cmAuth.NewTokenGenerator(&cmAuth.TokenGeneratorOptions{
		PrivateKeyPath: *privateKeyPath,
	})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/oauth/token", oauthTokenHandler)
	r.Run(fmt.Sprintf("%s:%d", "", *servicePort))
}
