package helper

import (
	"time"

	"my-gram/entity"

	"github.com/kataras/jwt"
)

var SharedKey = []byte("r@h@s14il4h1thathavesecure32char")

type Claims struct {
	AccessClaims  entity.AccessClaim
	DefaultClaims entity.DefaultClaim
}

func GenerateDefaultClaims(username string) entity.DefaultClaim {
	timenow := time.Now()
	return entity.DefaultClaim{
		Expired:   int(timenow.Add(24 * time.Hour).UnixMilli()),
		NotBefore: int(timenow.UnixMilli()),
		IssuedAt:  int(timenow.UnixMilli()),
		Issuer:    "my-gram",
		Audience:  "my-gram",
		JTI:       username,
		Typ:       "",
	}
}

func GenerateToken(userInfo entity.User) (tokenOut entity.Tokens, err error) {
	defaultClaim := GenerateDefaultClaims(userInfo.Username)
	defaultClaim.Typ = "id_token"

	accessClaim := entity.AccessClaim{
		ID:       int(userInfo.ID),
		Username: userInfo.Username,
	}

	userClaims := Claims{
		AccessClaims:  accessClaim,
		DefaultClaims: defaultClaim,
	}

	// Generate JWT
	IDToken, err := jwt.Sign(jwt.HS256, SharedKey, userClaims)
	tokenOut.AccessToken = string(IDToken)
	if err != nil {
		return entity.Tokens{}, err
	}

	return
}

func VerifyToken(token string, Claims any) (err error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, SharedKey, []byte(token))
	if err != nil {
		return err
	}

	err = verifiedToken.Claims(&Claims)
	if err != nil {
		return err
	}

	return
}
