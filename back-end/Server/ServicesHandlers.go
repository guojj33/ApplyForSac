package Server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/request"

	"github.com/dgrijalva/jwt-go"

	"../Models"
	"../Services"
	"github.com/unrolled/render"
)

var tokens []Token

const SecretKey = "guojj33"
const Issuer = "guojj"

const (
	Subject_User  = "USER"
	Subject_Admin = "ADMIN"
)

type Token struct {
	SAC_TOKEN string `json:"SAC_TOKEN"`
}

//-> tokenStr
func CreateToken(secretKey []byte, issuer string, audience string, subject string) (Token, error) {
	claims := &jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Issuer:   issuer,
		Audience: audience,
		Subject:  subject,
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return Token{}, err
	}
	token := Token{
		SAC_TOKEN: tokenStr,
	}
	return token, nil
}

//登录 发送 token
func LoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		newLoginReq := struct {
			Id       string
			Password string
		}{}
		err := json.Unmarshal(payload, &newLoginReq)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		var accountType Models.AccountType
		accountType, err = Services.LogIn(newLoginReq.Id, newLoginReq.Password)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to login.\n"+err.Error()+"\n")
			return
		} else {
			//登录信息验证成功，签发 token
			var subject string
			switch accountType {
			case 0:
				subject = Subject_User
			case 1:
				subject = Subject_Admin
			}
			token, err := CreateToken([]byte(SecretKey), Issuer, newLoginReq.Id, subject)
			if err != nil {
				formatter.Text(w, http.StatusInternalServerError, "Failed to create token.\n"+err.Error()+"\n")
				return
			}
			//创建 token 成功并保存到服务器
			tokens = append(tokens, token)
			info := struct {
				Token          Token
				CurAccountType int
			}{
				Token:          token,
				CurAccountType: int(accountType),
			}
			formatter.JSON(w, http.StatusOK, info)
		}
	}
}

//验证 token 有效性的中间件
func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//println("reqURI:" + req.RequestURI)
		//不需要 token 就可以请求的 api
		if req.RequestURI == "/api/login" || req.RequestURI == "/api/register" || req.RequestURI == "/api/comments" {
			next.ServeHTTP(w, req)
			return
		}
		token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKey), nil
			})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Failed to parse token.\n" + err.Error() + "\n"))
			return
		}
		if token.Valid {
			exist := false
			var tokenStr string
			for k, v := range req.Header {
				if k == "Authorization" {
					tokenStr = v[0][7:] //直接从 header 中获取 token， 格式为 "Bearer XXXX.XXXX.XXXX"
					break
				}
			}
			//println("tokenStr:" + tokenStr)
			//服务器端对应 token 判断
			for _, curToken := range tokens {
				if curToken.SAC_TOKEN == tokenStr {
					exist = true
					break
				}
			}
			if !exist {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid token.\n"))
				return
			}
			//检查是否有权限请求相应的 api
			//未完成
			next.ServeHTTP(w, req)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token.\n" + err.Error() + "\n"))
			return
		}
	})
}

func RegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		newRegisterReq := struct {
			UserId   string
			Password string
			Email    string
		}{}
		err := json.Unmarshal(payload, &newRegisterReq)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		_, err = Services.RegisterNewUser(newRegisterReq.UserId, newRegisterReq.Password, newRegisterReq.Email)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to register.\n"+err.Error()+"\n")
			return
		} else {
			formatter.Text(w, http.StatusOK, "Register succeed.\n")
		}
	}
}

func LogoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var tokenStr string
		for k, v := range req.Header {
			if k == "Authorization" {
				tokenStr = v[0][7:] //直接从 header 中获取 token， 格式为 "Bearer XXXX.XXXX.XXXX"
				break
			}
		}
		//println("tokenStr:" + tokenStr)
		//服务器端对应 token 判断
		exist := false
		for i, curToken := range tokens {
			if curToken.SAC_TOKEN == tokenStr {
				tokens = append(tokens[:i], tokens[i+1:]...)
				exist = true
				break
			}
		}
		if exist {
			formatter.Text(w, http.StatusOK, "Logout secceed.\n")
		} else {
			formatter.Text(w, http.StatusBadRequest, "Logout failed.\n")
		}
	}
}
