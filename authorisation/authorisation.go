package authorisation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID
	Name       string
	Email      string
	Role       []string
	Created_at string
	Updated_at string
}

type Result struct {
	Status  string
	Data    User
	Message string
}

type Role struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name       string               `json:"name" validate:"name,required" bson:"name,omitempty"`
	Permission []primitive.ObjectID `json:"permission" bson:"permissions,omitempty"`
}

type Permission struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" validate:"name,required" bson:"name,omitempty"`
}

func CheckAuthorisation(token string) (*User, error) {

	response, err := http.Get("http://user_management:5000/api/v1/user/findByToken?token=" + token)

	if response.StatusCode != 200 {

		defer response.Body.Close()

		return nil, err
	} else {
		var result Result

		//body, err := ioutil.ReadAll(response.Body)
		json.NewDecoder(response.Body).Decode(&result)
		//json.Unmarshal([]byte(body), &result)

		fmt.Print("response --------------")
		fmt.Print(result.Data.Email)

		defer response.Body.Close()

		return &result.Data, nil
	}

}

func VerifyID(id string) (*User, error) {

	response, err := http.Get("http://user_management:5000/api/v1/user/findByID?id=" + id)

	if response.StatusCode == 200 {
		var result Result
		//body, err := ioutil.ReadAll(response.Body)
		//json.Unmarshal([]byte(body), &result)

		json.NewDecoder(response.Body).Decode(&result)

		defer response.Body.Close()
		return &result.Data, nil

	} else {

		defer response.Body.Close()
		return nil, err
	}

}

func GetRole(id string) (*Role, error) {

	response, err := http.Get("http://user_management:5000/api/v1/role?id=" + id)

	if response.StatusCode == 200 {
		var result Role
		//body, err := ioutil.ReadAll(response.Body)
		//json.Unmarshal([]byte(body), &result)

		json.NewDecoder(response.Body).Decode(&result)

		defer response.Body.Close()
		return &result, nil

	} else {

		defer response.Body.Close()
		return nil, err
	}

}
