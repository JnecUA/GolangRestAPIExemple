package platform

//Team type to work with db table
type Team struct {
	UsersIds  []string `json:"usersIds"`
	Name      string   `json:"name"`
	CreatorID string   `json:"creatorId"`
}
