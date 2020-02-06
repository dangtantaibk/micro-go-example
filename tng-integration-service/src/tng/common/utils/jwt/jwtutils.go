package jwt

import (
	"fmt"
	beecontext "github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	"tng/common/models/shiper"
	"tng/common/utils/cfgutil"
	"tng/shipper-service/dtos"
)

const (
	ValidJwtToken   int = 1
	InvalidJwtToken int = 2
)

func ParseAuthHeaderToken(authHeader string) string {
	var authToken string
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return ""
	}
	if len(bearerToken) == 2 {
		authToken = bearerToken[1]
	}
	return authToken
}

func TokenVerify(authHeader string) (int, error) {
	bearerToken := strings.Split(authHeader, " ")
	key := []byte(cfgutil.Load("JWT_SECRET_KEY"))

	if len(bearerToken) != 2 {
		return InvalidJwtToken, nil
	}

	if len(bearerToken) == 2 {
		authToken := bearerToken[1]
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil

		})

		if err != nil {
			return InvalidJwtToken, nil
		}
		if !token.Valid {
			return InvalidJwtToken, nil
		}

	}
	return ValidJwtToken, nil
}

func GenerateShipperToken(user *shiper.Shipper, input beecontext.BeegoInput) (string, error) {
	//duration, err := strconv.Atoi(cfgutil.Load("JWT_EXPIRATION"))
	//if err != nil {
	//	return "", err
	//}
	expirationTime := time.Now().Add(15 * 24 * time.Hour)
	//ip := input.IP()
	//userAgent := input.UserAgent()
	//socialId := strconv.FormatInt(user.SocialId, 10)
	//subjectString := socialId + "%%%" + ip + "%%%" + userAgent

	//claims := &jwt.StandardClaims{
	//	ExpiresAt: expirationTime.Unix(),
	//	IssuedAt:  time.Now().Unix(),
	//	Subject:   subjectString,
	//}
	claims := &dtos.Claims{
		PhoneNumber:    user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfgutil.Load("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GenerateToken(user *shiper.User, input beecontext.BeegoInput) (string, error) {
	//duration, err := strconv.Atoi(cfgutil.Load("JWT_EXPIRATION"))
	//if err != nil {
	//	return "", err
	//}
	expirationTime := time.Now().Add(15 * 24 * time.Hour)
	//ip := input.IP()
	//userAgent := input.UserAgent()
	//socialId := strconv.FormatInt(user.SocialId, 10)
	//subjectString := socialId + "%%%" + ip + "%%%" + userAgent

	//claims := &jwt.StandardClaims{
	//	ExpiresAt: expirationTime.Unix(),
	//	IssuedAt:  time.Now().Unix(),
	//	Subject:   subjectString,
	//}
	claims := &dtos.Claims{
		PhoneNumber:    user.PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfgutil.Load("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseWithClaims(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfgutil.Load("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dtos.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}