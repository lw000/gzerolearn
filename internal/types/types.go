// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"` // parameters are auto validated
}

type Response struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Token string `json:"toekn"`
}