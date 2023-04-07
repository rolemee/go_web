package test

import (
	"fmt"
	"testing"
	"time"
	"github.com/golang-jwt/jwt"
	jwtex "github.com/kataras/jwt"
)
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func TestVerify(t *testing.T){
	pub,_ := jwt.ParseRSAPublicKeyFromPEM([]byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj9jDeOrOpbtkWm6y54z/
dQnGXAg1lZLZlbsmyBmqXj77/qmy5FfQ8qwUV5+h3HZGNKxJuEGfqH9qiYmxd39T
iy6sJrSoLm7FpR63KFEU3KlvnrrieSOI298Eu/rVxiUDri0iXApkjqdtE/6Shc3B
KBL5WUg/Qxs0ZEkTJiL1CvlTqwE5qS3dn0v7H19h+IwiZTHYwjNxsL4raxKxCWSQ
4Cxx5I45b5AFqzY+O55uJgIM8JuXhUNgIOvEvvsQ//S2z0OpG+AD2zk04ustwAnf
P/lYf/1OuwaSXkUSLuZDgk/kH8eg2DjGg7VQEPhaMprD/QuffetBMJsoKakHVMUF
lQIDAQAB
-----END PUBLIC KEY-----`))
	tokenString := (`eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImlhdCI6MTY4MDg0MTQ0MiwiZXhwIjoxNjg2MjQxNDQyfQ.Qxp8CAV96fRh_3HreQUs4q8NWGuvHCh37pkSSMaCisWQqDDcMX3jj37q5wtkhwdubYfQQmxUhxqNtRXoMqIqGylECge8DPb29Dls9HYGaKGPA1r0rZoenfdON53id1_L38JISuzrE1gfEZm4lc-HpKMOZc5OMKnCNGRZvszAE4QyxLqCRSFTg1GAdGiP9BPkKgVem6ms5O8Az6ey4e57tPCLf0X4wpbhyPaRuW8qNsj55VbRWXCN3eqcOlt0KmxQrPj3V03gs5o0Lhh7tGxRA712M-6nSajPzTaWegiPd_pyKX0CXFU0qxoYB_bjsDUWOHfHoEtPiWxusJoMoNZ__w`)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return pub, nil
    })
	if err != nil {
        fmt.Println(err)
    }

    // Check if the token is valid
    if !token.Valid {
        fmt.Println("invalid token")
    }

    // Check the expiration time
    if claims, ok := token.Claims.(jwt.MapClaims); ok && !claims.VerifyExpiresAt(time.Now().Unix(), true) {
        fmt.Println("token expired")
    }
	// token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("AllYourBase"), nil
	// })

	// // toekn,err := jwt.Verify(jwt.RS256,pub,[]byte(""))
	// if(err!=nil){
	// 	fmt.Println(err)
	// }else{
	// 	token1, err := jwt.Parse(toekn, func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("AllYourBase"), nil
	// 	})
	// }

}

func TestSign(t *testing.T){
	pri,err := jwtex.ParsePrivateKeyRSA([]byte(`-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCP2MN46s6lu2Ra
brLnjP91CcZcCDWVktmVuybIGapePvv+qbLkV9DyrBRXn6HcdkY0rEm4QZ+of2qJ
ibF3f1OLLqwmtKgubsWlHrcoURTcqW+euuJ5I4jb3wS7+tXGJQOuLSJcCmSOp20T
/pKFzcEoEvlZSD9DGzRkSRMmIvUK+VOrATmpLd2fS/sfX2H4jCJlMdjCM3Gwvitr
ErEJZJDgLHHkjjlvkAWrNj47nm4mAgzwm5eFQ2Ag68S++xD/9LbPQ6kb4APbOTTi
6y3ACd8/+Vh//U67BpJeRRIu5kOCT+Qfx6DYOMaDtVAQ+FoymsP9C59960Ewmygp
qQdUxQWVAgMBAAECggEAGZnqsEqaHMUNR0sMbEmb7aiftSGA5+4K66sza8mlmMB8
+QYa4KkAdzBJLo8qr6IdPKKtPC3xzHtg6rmm7C2tIMhdu0XL6PeuFYUx8+r6W4uM
YiPAyu6YHQnEPfCWu306+nTt+JOg8gK6VySnz0ifaBJ4bnt9RuRAntiin9b45crU
PxfP2iePfZGFB/J9YEcI3sT2v2Vx9x0l/3WJubOm90jhroBfcguwQxgBUQT/f+Bx
i4dyyBgCOW+WTDJ+YiWkUIWWjGxtU1h3V5X2HJkEbtG2BMxdqtF8h/VACfnShmiy
dnYaEOqPuJl08tniKLKBDPn2Pju09OPcoj0/JfxuwQKBgQDuSiHjuDC2vcetnLRv
EEIhcfxRuR9BbXSRAOVHNYFX9jzQTY2zykql/4BXpoPPCEhJXUmwX9rwD48mDT3X
JR072V34Vx4smwYzywx9Jjp4aPSQLjZW1dBlITb7z+hdqTz51f586Ksbd4SKJgIW
4MWDo/YxRZex/Ceiozj/hr4KNwKBgQCaibG1EG5d5sMavBgLbaKl3uOVsFFC8z26
GBo/tFcXqgb2aPRGu9o3qh/bJaVSGpD/XLEbIPUHY36wWFmdIgpB77XG2Q88Clf+
SGPVEN/iR0HHeWgZcxnsEMTejXr9XvkV2e1MhFigy1/iF5Xu+7HrundVz0j0U9zW
grWH6NAYkwKBgHJIHX7ATKSn2gcam7KUYpL1vuRW23WekAQnIM8JCzJrhvkD6Mzs
XObF5gCjfJ2jRD/jPL8ZCcCORcIjIaB4TU/xzM6YwYD1DIF0hVXUKsCq3Xf0odz9
iIen6V3VdYTN8M+FYRr4328539f9qIzeWfvre9xfmwQigcPPcjuinKv7AoGARD+R
QsSZ6VBzU8hD2jA2B4kS+1Wd3bJszVW9qeqsF+BYKNSbJO3rZm/0l6TdiRAAfzZh
ZPi+HeA74Ad0tCeHXi1OGx5bUwnCaKNxq3RDW7xaeYzNg4fIp++T8jGLYoBiAtt2
qDc+qOyIUOJXre5sWg3EjBn6PyUWP4oEP5x64o8CgYAlqCpZ8Q4uOTDeOZwUZrlu
i5KJjPc+toiO5FhlSAPz3brNLC42jDXlLss5BXBQtSnMR4h3zm2JKwSi3GnQ5zee
HaUhDYNwjAUPvHFmEGDISyrxP3F0mDpUbs0C2HxP7KqmExkBDqAkcmc9ZAhQDPuj
ry5T7pnLdrnZGkFUHX5gjg==
-----END PRIVATE KEY-----`))
	if err != nil{
		fmt.Println(err)
		fmt.Println(123123)
		return
	}
	token,err := jwtex.Sign(jwtex.RS256,pri,jwtex.Map{"user_name":"jzy","id":10},jwtex.MaxAge(1500 * time.Hour))
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println(string(token))
	}

}
type Claims struct {
    UserName string `json:"user_name"`
    jwt.StandardClaims
}

func TestJ(t *testing.T){
	claims := &Claims{}
	tokenstr := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsImV4cCI6MTY4MDY3MzI0NSwidXNlcl9uYW1lIjoianp5IiwianRpIjoiMWM5NDkxZjgtNjVkOC00OGU2LWExYjAtMjBhYmQyNzI4ZGVjIiwiY2xpZW50X2lkIjoic2MtbWFuYWdlciIsInNjb3BlIjpbImFsbCJdfQ.jDPXgH3UsFJaAreAgUa8PdhYDdX1ckdHBPI46vhIA8glOwFFIDg1-UDeKtt_8cOsekud_1hJ6wMLKZwrYaIS814fqxEQE5JBu8K9maSbcOFUChREXXTUXSNDC1SUIm9BwJ6ZNGDtosfUMgd64MB2blxhUyPVTHNqZTiIFknT3pHkBuGXR6Xj2KmbkDmYR45XD2Je3MzaKfwTYPvh9zw_UEhii6GDu6eyhaWipqknIcEoB4_nT5d-F69NI0eK_mpSepoxIgKgdH7tOrs9FH2xYfjMDM9L6-LwznVjgve1s68dVHwaKdgyqtRI9_9FpG1IjqULDMjGhCL9z27wYVa7tw`
	jwtKey,_ := jwt.ParseRSAPublicKeyFromPEM([]byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj9jDeOrOpbtkWm6y54z/
dQnGXAg1lZLZlbsmyBmqXj77/qmy5FfQ8qwUV5+h3HZGNKxJuEGfqH9qiYmxd39T
iy6sJrSoLm7FpR63KFEU3KlvnrrieSOI298Eu/rVxiUDri0iXApkjqdtE/6Shc3B
KBL5WUg/Qxs0ZEkTJiL1CvlTqwE5qS3dn0v7H19h+IwiZTHYwjNxsL4raxKxCWSQ
4Cxx5I45b5AFqzY+O55uJgIM8JuXhUNgIOvEvvsQ//S2z0OpG+AD2zk04ustwAnf
P/lYf/1OuwaSXkUSLuZDgk/kH8eg2DjGg7VQEPhaMprD/QuffetBMJsoKakHVMUF
lQIDAQAB
-----END PUBLIC KEY-----`))
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
	if !token.Valid{
		fmt.Println(err)
	}
	fmt.Println(claims.UserName)
	
}