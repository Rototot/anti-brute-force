package queries

type DeleteBlackListQuery struct {
	Subnet string `json:"subnet" validate:"cidrv4"`
}
