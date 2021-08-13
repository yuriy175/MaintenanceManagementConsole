package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

func CheckQueryParameter(queryString url.Values, paramName string, w http.ResponseWriter) string {
	params, ok := queryString[paramName]
	if !ok || len(params[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	return params[0]
}

func CheckOptionalQueryParameter(queryString url.Values, paramName string, w http.ResponseWriter) string {
	params, ok := queryString[paramName]
	if !ok || len(params[0]) < 1 {
		return ""
	}
	return params[0]
}

func CheckUserAuthorization(authService interfaces.IAuthService, w http.ResponseWriter, r *http.Request) *models.UserClaims {
	claims := checkAuthorization(authService, w, r)
	return claims
}

func CheckAdminAuthorization(authService interfaces.IAuthService, w http.ResponseWriter, r *http.Request) *models.UserClaims {
	claims := checkAuthorization(authService, w, r)
	if claims != nil && claims.Role != models.AdminRoleName {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		return nil
	}

	return claims
}

func checkAuthorization(authService interfaces.IAuthService, w http.ResponseWriter, r *http.Request) *models.UserClaims {
	if checkOptionsAndSetCORSMethod(w, r) {
		return nil
	}

	tokenString := checkRequestAuthorization(w, r)
	if tokenString == "" {
		return nil
	}

	claims := authService.VerifyToken(tokenString)
	if claims == nil {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
	}

	return claims
}

// CheckAuthorization check authorization token presence
func checkRequestAuthorization(w http.ResponseWriter, r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization") //Получение токена

	if tokenHeader == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		return ""
	}

	splitted := strings.Split(tokenHeader, " ") //Токен обычно поставляется в формате `Bearer {token-body}`, мы проверяем, соответствует ли полученный токен этому требованию
	if len(splitted) != 2 {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		return ""
	}

	tokenPart := splitted[1]

	return tokenPart
}

// CheckOptionsAndSetCORSMethod check if method OPTIONS and sets CORS
func checkOptionsAndSetCORSMethod(w http.ResponseWriter, r *http.Request) bool {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return true
	}

	return false
}
