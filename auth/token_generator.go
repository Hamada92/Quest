package auth

import "time"

type TokenGenerator interface {
	CreateToken(payload *Payload, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
