package security

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoBuilder struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoBuilder(symmetricKey string) (Builder, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("tamaño de la llave inválido:%d caracteres", chacha20poly1305.KeySize)
	}
	builder := &PasetoBuilder{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return builder, nil
}

func (builder *PasetoBuilder) CreateToken(user string, role string, duration time.Duration) (string, error) {
	payload, err := NewPayload(user, role, duration)
	if err != nil {
		return "", err
	}
	return builder.paseto.Encrypt(builder.symmetricKey, payload, nil)
}

func (builder *PasetoBuilder) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := builder.paseto.Decrypt(token, builder.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrorInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
