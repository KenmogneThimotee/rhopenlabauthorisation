### rhopenlabauthorisation


##There are three functions:

**1**
CheckAuthorisation(token string) (*User, error)
To check if user is authenticated

type User struct {
	ID         primitive.ObjectID
	Name       string
	Email      string
	Role       []primitive.ObjectID
	Created_at string
	Updated_at string
}

**2**
VerifyID(id string) (*User, error)
To know if the user exist it return user struct or nil 

type User struct {
	ID         primitive.ObjectID
	Name       string
	Email      string
	Role       []primitive.ObjectID
	Created_at string
	Updated_at string
}


**3**
GetRole(id string) (*Role, error)
To get role based on his name. It return role struct or nil

type Role struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name       string               `json:"name" validate:"name,required" bson:"name,omitempty"`
	Permission []primitive.ObjectID `json:"permission" bson:"permissions,omitempty"`
}
