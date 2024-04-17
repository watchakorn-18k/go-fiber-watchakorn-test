package entities

type NewUserBody struct {
	Name string `json:"name" bson:"name,omitempty"`
	Age  int    `json:"age" bson:"age"`
	Text string `json:"text" bson:"text"`
}
type UserDataFormat struct {
	Name string `json:"name" bson:"name,omitempty"`
	Age  int    `json:"age" bson:"age"`
	Text string `json:"text" bson:"text"`
}
