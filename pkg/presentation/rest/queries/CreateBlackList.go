package queries

type CreateBlackListQuery struct {
	Subnet string `json:"subnet" validate:"cidrv4"`
}
