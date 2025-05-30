package security

import "time"

type Builder interface {
	CreateToken(user string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
