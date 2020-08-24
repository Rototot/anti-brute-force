package queries

type CreateBlackListQuery struct {
	Subnet string `json:"ip" validate:"cidrv4"`
}
