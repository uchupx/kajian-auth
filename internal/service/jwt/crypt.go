package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/uchupx/kajian-auth/config"
)

type CryptService interface {
	CreateSignPSS(value string) (signatureStr string, err error)
	Verify(value string, signature string) (resp bool, err error)
	CreateJWTToken(expDuration time.Duration, content interface{}) (token *string, err error)
	VerifyJWTToken(token string) (result interface{}, err error)
}

type cryptService struct {
	conf *config.Config
}

type Params struct {
	Conf *config.Config
}

func NewCryptService(params Params) CryptService {
	return &cryptService{
		conf: params.Conf,
	}
}

func (s *cryptService) loadRsaPrivateKey() (rsaKey *rsa.PrivateKey, err error) {
	key := strings.ReplaceAll(s.conf.RSA.Private, "\\n", "\n") // remove double slash, because it will affected when convert to byte
	rsaKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		// s.logger.Errorf("[loadRsaPrivateKey] failed parse private key, err: %+v", err)
		return
	}

	return
}

func (s *cryptService) loadRsaPublicKey() (rsaPub *rsa.PublicKey, err error) {
	key := strings.ReplaceAll(s.conf.RSA.Public, "\\n", "\n") // remove double slash, because it will affected when convert to byte
	rsaPub, err = jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	if err != nil {
		// s.logger.Errorf("[loadRsaPublicKey] failed parse public key, err: %+v", err)
		return
	}

	return
}

func (s *cryptService) CreateSignPSS(value string) (signatureStr string, err error) {
	msg := []byte(value)
	msgHash := sha256.New()

	_, err = msgHash.Write(msg)
	if err != nil {
		return
	}

	msgHashSum := msgHash.Sum(nil)

	privateKey, err := s.loadRsaPrivateKey()
	if err != nil {
		return
	}

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		// s.logger.Errorf("[CreateSignPSS] failed to create signature, err: %+v", err)
		return
	}

	signatureStr = base64.URLEncoding.EncodeToString(signature)
	return
}

func (s *cryptService) Verify(value string, signature string) (resp bool, err error) {
	resp = false

	signatureByte, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return
	}

	msg := []byte(value)
	msgHash := sha256.New()

	_, err = msgHash.Write(msg)
	if err != nil {
		return
	}

	msgHashSum := msgHash.Sum(nil)

	publicKey, err := s.loadRsaPublicKey()
	if err != nil {
		return
	}

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signatureByte, nil)
	if err != nil {
		return
	}

	resp = true

	return
}

func (s *cryptService) CreateJWTToken(expDuration time.Duration, content interface{}) (token *string, err error) {
	privateKey, err := s.loadRsaPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("[CreateJWTToken] failed load rsa private key, err: %+v", err)
	}

	publicKey, err := s.loadRsaPublicKey()
	if err != nil {
		return nil, fmt.Errorf("[CreateJWTToken] failed load rsa public key, err: %+v", err)
	}

	jwtServicecryptService := NewJWT(privateKey, publicKey)

	return jwtServicecryptService.Create(expDuration, content)
}

func (s *cryptService) VerifyJWTToken(token string) (result interface{}, err error) {
	privateKey, err := s.loadRsaPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("[VerifyJWTToken] failed load rsa private key, err: %+v", err)
	}

	publicKey, err := s.loadRsaPublicKey()
	if err != nil {
		return nil, fmt.Errorf("[VerifyJWTToken] failed load rsa public key, err: %+v", err)
	}

	jwtServicecryptService := NewJWT(privateKey, publicKey)

	return jwtServicecryptService.Validate(token)
}
