package jwt

import (
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"sblog/config"
	"strconv"
	"time"
)

func Encode(des interface{}) (ret string, err error) {
	var iss string
	dest, ok := des.(string)
	if ok {
		iss = dest
	}

	destm, ok := des.(map[string]string)
	if ok {
		isss, _ := (json.Marshal(destm))
		iss = string(isss)
	}

	destmi, ok := des.(map[string]interface{})
	if ok {
		isss, _ := (json.Marshal(destmi))
		iss = string(isss)
	}

	mySigningKey := Get(config.JWT_KEY)
	// Create the Claims
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    iss,
		Subject:   "test",
	}
	//mapClaims := &jwt.MapClaims{
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ret, err = token.SignedString(mySigningKey)
	return
}

func Decode(token string) (ret map[string]string, err error) {
	ret = make(map[string]string)
	mySigningKey := config.JWT_KEY
	retToken, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return Get(mySigningKey), nil
	})
	for index, val := range retToken.Claims.(jwt.MapClaims) {
		vall, ok := val.(string)
		if ok {
			ret[index] = vall
		}

		valint64, ok := val.(int64)
		if ok {
			ret[index] = strconv.Itoa(int(valint64))
		}

		valint, ok := val.(int)
		if ok {
			ret[index] = strconv.Itoa((valint))
		}

	}
	return
}

func Get(jk interface{}) []byte {
	var jwtk []byte
	jwtk, ok := jk.([]byte)
	if ok == true {
		return jwtk
	}
	jwtks, ok := jk.(string)
	if ok {
		jwtk = []byte(jwtks)
	}
	return jwtk
}
