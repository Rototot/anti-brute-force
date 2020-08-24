package queries

type DeleteWhiteListQuery struct {
	Subnet string `json:"ip" validate:"cidrv4"`
}
