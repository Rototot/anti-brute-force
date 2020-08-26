package queries

type DeleteWhiteListQuery struct {
	Subnet string `json:"subnet" validate:"cidrv4"`
}
