package queries

type LoginAttemptQuery struct {
	Login    string `json:"login" validate:"max:1024"`
	Password string `json:"password" validate:"max:1024"`
	IP       string `json:"ip" validate:"ipv4"`
}
