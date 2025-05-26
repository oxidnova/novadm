package schema

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginToken struct {
	AccessToken string `json:"accessToken"`
}

type UserInfo struct {
	Username string   `json:"username"`
	RealName string   `json:"realName"`
	Roles    []string `json:"roles"`
	Password string   `json:"-"`
}

type ArrayItem struct {
	Items  any    `json:"items"`
	Total  int    `json:"total"`
	LastId string `json:"lastId,omitempty"`
}
