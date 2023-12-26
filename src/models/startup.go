package models

type Startup struct {
	Name             string `json:"name" bson:"name"`
	CEO              string `json:"ceo" bson:"ceo"`
	ValueProposition string `json:"value_proposition" bson:"value_proposition"`
	Industry         string `json:"industry" bson:"industry"`
}

type Founder struct {
	Name       string `json:"name" bson:"name"`
	Profession string `json:"profession" bson:"profession"`
	Age        int    `json:"age" bson:"age"`
}
