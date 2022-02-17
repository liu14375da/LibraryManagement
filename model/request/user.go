package request

type AccountRequest struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Code  string `json:"code"`
	Id    string `json:"id"`
}
