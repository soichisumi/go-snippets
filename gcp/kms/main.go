package main

import (
	cloudkms "cloud.google.com/go/kms/apiv1"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"log"
	"time"
)

// validateToken ...
func validateToken(token string, pubKey *rsa.PublicKey) (*jwt.Token, error) {
	// Parse function parse and validate token.
	// keyFunc receive the parsed token and should return the key for validating
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("Unexpected signing method: %v", t.Header["alg"])
			return nil, fmt.Errorf("invalid token")
		}
		return pubKey, nil
	})

	if vErr, ok := err.(*jwt.ValidationError); ok {
		switch {
		case vErr.Errors&jwt.ValidationErrorMalformed != 0:
			return nil, errors.New("malformed token")

		case vErr.Errors&jwt.ValidationErrorExpired != 0:
			return nil, errors.New("token is expired")

		case vErr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
			return nil, errors.New("signature is invalid")

		default:
			return nil, errors.New("token is invalid")
		}
	} else if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	usedAfterIssued := claims.VerifyIssuedAt(time.Now().Unix(), true)
	if !usedAfterIssued {
		return nil, errors.New("token is used before issued")
	}
	validIssuer := claims.VerifyIssuer("issuer", true)
	if !validIssuer {
		return nil, errors.New("issuer is invalid")
	}
	return jwtToken, nil
}

func generateAndVerifyJwt(){
	fmt.Printf("yo")
	now := time.Now()
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "issuer"
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour).Unix()

	tokenString, err := token.SigningString()
	if err != nil {
		log.Fatalf("err: %+v", err)
	}
	fmt.Printf("token: %+v\n", tokenString)

	hasher := crypto.SHA256
	digest := hasher.New()
	digest.Write([]byte(tokenString))

	projectID := ""
	keyRingID := ""
	keyID := ""
	keyName := fmt.Sprintf("projects/%s/locations/global/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%d", projectID, keyRingID, keyID, 1)

	req := &kmspb.AsymmetricSignRequest{
		Name: keyName,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: digest.Sum(nil),
			},
		},
	}

	// sign
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	// Call the API.
	signResponse, err := client.AsymmetricSign(ctx, req)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	encodedSignature := jwt.EncodeSegment(signResponse.Signature)

	signedToken := tokenString + "." + encodedSignature// validate token
	fmt.Printf("signedToken: %s\n", signedToken)

	// Retrieve the public key from KMS. response, err := client.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: name})
	pubKeyResponse, err := client.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: keyName})
	if err != nil {
		fmt.Printf("GetPublicKey: %v\n", err)
		return
	}
	// Parse the key.
	block, _ := pem.Decode([]byte(pubKeyResponse.Pem))
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKIXPublicKey: %v\n", err)
		return
	}
	//ecKey, ok := parsedKey.(*ecdsa.PublicKey)
	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		fmt.Printf("key '%s' is not EC\n", keyName)
		return
	}

	token, err = validateToken(signedToken, pubKey)
	if err != nil{
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("token: %+v", token)
}

func main(){
	 generateAndVerifyJwt()
}