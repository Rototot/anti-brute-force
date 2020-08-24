package queries

type ResetAttemptsQuery struct {
	Login string `json:"login" validate:"max:1024"`
	IP    string `json:"ip" validate:"ipv4"`
}
