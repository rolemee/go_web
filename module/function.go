package module

import (
	"sort"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt"
)

func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}
type Claims struct {
    UserName string `json:"user_name"`
	UserId int `json:"id"`
    jwt.StandardClaims
}
var pub,_ = jwt.ParseRSAPublicKeyFromPEM([]byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj9jDeOrOpbtkWm6y54z/
dQnGXAg1lZLZlbsmyBmqXj77/qmy5FfQ8qwUV5+h3HZGNKxJuEGfqH9qiYmxd39T
iy6sJrSoLm7FpR63KFEU3KlvnrrieSOI298Eu/rVxiUDri0iXApkjqdtE/6Shc3B
KBL5WUg/Qxs0ZEkTJiL1CvlTqwE5qS3dn0v7H19h+IwiZTHYwjNxsL4raxKxCWSQ
4Cxx5I45b5AFqzY+O55uJgIM8JuXhUNgIOvEvvsQ//S2z0OpG+AD2zk04ustwAnf
P/lYf/1OuwaSXkUSLuZDgk/kH8eg2DjGg7VQEPhaMprD/QuffetBMJsoKakHVMUF
lQIDAQAB
-----END PUBLIC KEY-----`))
func Checkjwt(tokens string) bool{
		token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return pub, nil
		})
		if err != nil {
			return false
		}
		if !token.Valid {
			return false
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return false
		}
		return true
}

func ParseJwt(tokenstr string) *Claims{
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
        return pub, nil
    })
	if !token.Valid{
		fmt.Println(err)
	}
	return claims
}