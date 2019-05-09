package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gamorvi/restapi2/app/models"
	u "github.com/gamorvi/restapi2/utils"
	"github.com/gorilla/context"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

type JwtToken struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	username := req.FormValue("Username")
	password := req.FormValue("Password")
	result := models.GetUsername(username)
	if result == nil {
		u.Respond(w, u.Message(false, "Your credentials do not match our records u"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Your credentials do not match our records p"))
		return
	}
	// jwt ttl
	ttl := 15 * time.Minute
	accessTokenExpire := os.Getenv("access_token_expire")
	min, err := strconv.Atoi(accessTokenExpire)
	if err != nil {
		log.Println(err)
	}
	if accessTokenExpire != "" {
		ttl = time.Duration(min) * time.Minute
	}

	CreateToken(w, username, password, ttl)
}

func CreateToken(w http.ResponseWriter, username string, password string, ttl time.Duration) {
	jwt_secret := os.Getenv("jwt_secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      ttl,
	})

	tokenString, error := token.SignedString([]byte(jwt_secret))
	if error != nil {
		fmt.Println(error)
	}
	resp := u.Message(true, "success")
	resp["data"] = JwtToken{Token: tokenString}
	u.Respond(w, resp)
	return
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		jwt_secret := os.Getenv("jwt_secret")
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(jwt_secret), nil
				})
				if error != nil {
					u.Respond(w, u.Message(false, error.Error()))
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					u.Respond(w, u.Message(false, "Invalid authorization token"))
					return
				}
			}
		} else {
			u.Respond(w, u.Message(false, "An authorization header is required"))
			return
		}
	})
}
