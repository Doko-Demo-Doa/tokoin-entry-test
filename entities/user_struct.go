package entities

//easyjson:json
type UserEntityList []UserEntity

//easyjson:json
type UserEntity struct {
	Id         int    `json:"_id"`
	Url        string `json:"url"`
	ExternalId string `json:"external_id"`
	Name       string `json:"name"`
}
