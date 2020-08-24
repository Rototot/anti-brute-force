package queries

type DeleteBlackListQuery struct {
	Subnet string `json:"ip" validate:"cidrv4"`
}
