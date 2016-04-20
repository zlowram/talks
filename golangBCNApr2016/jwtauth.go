// Token creation (iat and exp claims included)
claims := make(map[string]string)
claims["userid"] = "user1"
claims["role"] = "admin"
token, err := s.auth.NewToken(claims)

// Verification performed automagically by the middleware

// Claims passed to next handler via gorilla/context
user := context.Get(r, "user").(map[string]string)
if user["role"] != "admin" {
	fmt.Println("You're not admin duuuude!")
}
