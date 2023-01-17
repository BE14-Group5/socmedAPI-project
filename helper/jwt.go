package helper

import "github.com/golang-jwt/jwt"

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	userId := -1
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userID"].(type) {
		case float64:
			userId = int(claims["userID"].(float64))
		case int:
			userId = claims["userID"].(int)
		}
		return int(userId)
	}
	return -1
}
