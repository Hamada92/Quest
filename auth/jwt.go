package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secretKey []byte
}

var _ TokenGenerator = (*JWT)(nil)

func NewJWT(secretKey []byte) (*JWT, error) {
	// validate key length in the future

	return &JWT{secretKey: secretKey}, nil
}

func (j *JWT) CreateToken(claims *Payload, duration time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secretKey)
}

func (j *JWT) VerifyToken(tokenString string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(j.secretKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyfunc)

	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(*Payload)

	if !ok {
		log.Fatal("Uknown claims type")
	}

	return claims, nil
}

// func main() {
// 	j := JWT{
// 		secretKey: []byte("ASDff35ysf"),
// 	}
// 	claims := &Payload{
// 		"ahmad",
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
// 		},
// 	}

// 	token, err := j.CreateToken(claims, time.Minute)

// 	if err != nil {
// 		log.Fatal("error creating token: ", err)
// 	}

// 	fmt.Println(token)

// 	payload, err := j.VerifyToken(token)

// 	if err != nil {
// 		log.Fatal("error verifying token: ", err)
// 	}

// 	fmt.Println("user: ", payload.UserID)
// }
