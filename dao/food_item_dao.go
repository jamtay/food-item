package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"food-item-api/models"
)

type FoodItemDao struct {
	Server string
	Database string
}

var foodItemDb *mgo.Database

const (
	FoodItemCollection = "food_item"
)

func (f *FoodItemDao) Connect() {
	session, err := mgo.Dial(f.Server)
	if err != nil {
		log.Fatal(err)
	}
	foodItemDb = session.DB(f.Database)
}

func (f *FoodItemDao) FindAll() ([]models.FoodItem, error) {
	var foodItems []models.FoodItem
	err := foodItemDb.C(FoodItemCollection).Find(bson.M{}).All(&foodItems)
	return foodItems, err
}

func (f *FoodItemDao) FindById(id string) (models.FoodItem, error) {
	var foodItem models.FoodItem
	err := foodItemDb.C(FoodItemCollection).FindId(bson.ObjectIdHex(id)).One(&foodItem)
	return foodItem, err
}

func (f *FoodItemDao) Insert(foodItem models.FoodItem) error {
	err := foodItemDb.C(FoodItemCollection).Insert(&foodItem)
	return err
}
