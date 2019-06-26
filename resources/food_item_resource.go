package resources

import (
	"encoding/json"
	"food-item-api/config"
	"food-item-api/dao"
	"food-item-api/models"
	"food-item-api/response_handler"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var foodItemDao = dao.FoodItemDao{}
var foodItemConfig = config.Config{}

func init() {
	foodItemConfig.ReadFoodItemConfig()
	foodItemDao.Server = foodItemConfig.Server
	foodItemDao.Database = foodItemConfig.Database
	foodItemDao.Connect()
}

func GetFoodItems(w http.ResponseWriter, r *http.Request) {
	foodItems, err := foodItemDao.FindAll()
	if err != nil {
		response_handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response_handler.RespondWithJson(w, http.StatusOK, foodItems)
}

func GetFoodItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if !bson.IsObjectIdHex(id) {
		response_handler.RespondWithError(w, http.StatusNotFound, id + " is not a valid id")
		return
	}

	foodItem, err := foodItemDao.FindById(id)
	if err != nil {
		response_handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response_handler.RespondWithJson(w, http.StatusOK, foodItem)
}

func CreateFoodItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var foodItem models.FoodItem
	if err := json.NewDecoder(r.Body).Decode(&foodItem); err != nil {
		response_handler.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	foodItem.ID = bson.NewObjectId()
	if err := foodItemDao.Insert(foodItem); err != nil {
		response_handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response_handler.RespondWithJson(w, http.StatusCreated, foodItem)
}

