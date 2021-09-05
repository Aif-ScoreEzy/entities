package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"github.com/AIF-user-system-management/entities/helper"
	protos "github.com/AIF-user-system-management/entities/protos"
	"github.com/golang/protobuf/jsonpb"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//User Model For Gorm ORM
type User struct {
	ID        string         `json:"id"  gorm:"primaryKey,index,column:id"`
	Name      string         `gorm:"type:varchar(255);NOT NULL;" json:"name"`
	Username  string         `gorm:"type:varchar(255);NOT NULL;unique" json:"username"`
	Email     string         `gorm:"type:varchar(255);NOT NULL;unique" json:"email"`
	Password  string         `gorm:"type:varchar(255);NOT NULL" json:"-"`
	APIKey    string         `gorm:"type:varchar(255);NOT NULL;unique" json:"api_key"`
	ParentID  string         `json:"parent_id"`
	Child     []*User        `json:"child" gorm:"foreignkey:ParentID"`
	CreatedAt time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

//BeforeCreate Hook before insert to database
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewV4().String()
	}

	if u.APIKey == "" {
		u.APIKey = helper.GenerateAPIKey()
	}

	if u.Password == "" {
		u.Password = helper.Hash(u.Password)
	}

	err = u.Validate()
	return
}

//Validate object of model
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.APIKey == "" {
		return errors.New("api_key is required")
	}

	return nil
}

//Serialize Convert from protobuf to struct  object
func (u *User) Serialize(data *protos.User) (err error) {
	msh := jsonpb.Marshaler{EnumsAsInts: true}
	result, _ := msh.MarshalToString(data)
	err = json.Unmarshal([]byte(result), u)
	return
}

//Deserialize Convert from struct to protobuf  object
func (u *User) Deserialize(data *protos.User) (err error) {
	byteArray, err := json.Marshal(u)
	if err != nil {
		return err
	}
	r := bytes.NewReader(byteArray)
	msh := jsonpb.Unmarshaler{AllowUnknownFields: true}
	err = msh.Unmarshal(r, data)
	if err != nil {
		return err
	}
	return
}
