package goeasyjwt

import "fmt"

func main() {
	claims := map[string]any{
		"id":   "nn",
		"name": "joe",
	}
	key := []byte("sdihb skdjbl dskjbliks")
	NewToken, _ := GenerateToken(claims, key, 4)
	fmt.Println("the token is: ", NewToken)
	Neclams, _ := VerifyToken(NewToken, key)
	fmt.Println("claims: ", Neclams)
}
