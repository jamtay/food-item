package models

import "gopkg.in/mgo.v2/bson"

type FoodItem struct {
	ID	bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name,omitempty"`
	Description string `bson:"description" json:"description,omitempty"`
	Cost float32 `bson:"cost" json:"cost,omitempty"`
	Calories int16 `bson:"calories" json:"calories,omitempty"`
	Protein int16 `bson:"protein" json:"protein,omitempty"`
	IsVegan bool `bson:"is_vegan" json:"is_vegan,omitempty"`
}