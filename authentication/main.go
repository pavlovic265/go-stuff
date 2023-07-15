package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.Handle("/api", ValidateJWT(Home))
	http.HandleFunc("/jwt", GetJWT)

	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}

var SECRET = []byte("super-secret-auth-key")
var API_KEY = "api-key"

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Not authorized"))
				}
				return SECRET, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Println(err.Error())
				w.Write([]byte("Something went wrong" + err.Error()))
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
		}

	})
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -> ", r.Header)
	if r.Header["Api"] != nil {
		if r.Header["Api"][0] == API_KEY {
			token, err := CreateJWT()
			fmt.Println("token -> ", token)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Fprint(w, token)
		}
	}
}
