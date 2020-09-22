package platform

type Team struct {
	usersIds  []string `json:"usersIds"`
	name      string   `json:"name"`
	creatorId string   `json:"creatorId"`
}
