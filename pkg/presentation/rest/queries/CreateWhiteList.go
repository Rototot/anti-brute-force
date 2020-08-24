package queries

type CreateWhiteListQuery struct {
	Subnet string `json:"ip" validate:"cidrv4"`
}
