package models

// {
// 	"access_token": "6qrZcUqja7812RVdnEKjpzOL4CvHBFG",
// 	"token_type": "Bearer",
// 	"expires_in": 604800,
// 	"refresh_token": "D43f5y0ahjqew82jZ4NViEr2YafMKhue",
// 	"scope": "identify"
// }

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    uint64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
