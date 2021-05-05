package domain

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/core_v1"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
	"time"
)

type User struct {
	pkg.BaseDomain
	FirstName      string
	LastName       string
	Gender         core_v1.Gender
	DateOfBirth    time.Time
	Identities     []Identity      `gorm:"foreignKey:UserID;references:ExternalId"`
	Addresses      []Address       `gorm:"foreignKey:UserID;references:ExternalId"`
	Relations      []Relation      `gorm:"foreignKey:UserID;references:ExternalId;foreignKey:RelationID;references:ExternalId"`
	UserAttributes []UserAttribute `gorm:"foreignKey:UserID;references:ExternalId"`
	Age            int64
	Height         float64
	Weight         float64
	Password       string
}

func (u *User) MarshalBinary() ([]byte, error) {
	dto := u.ToDto().(pikachu_v1.UserDto)
	return proto.Marshal(&dto)
}

func (u *User) UnmarshalBinary(buffer []byte) error {
	dto := pikachu_v1.UserDto{}
	err := proto.Unmarshal(buffer, &dto)
	if err != nil {
		return err
	}
	u.FillProperties(dto)
	return nil
}
func (u *User) ToBytes() (*bytes.Buffer, error) {
	var rBytes bytes.Buffer
	enc := gob.NewEncoder(&rBytes)
	err := enc.Encode(*u)
	return &rBytes, err
}

func (u *User) ToJson() (string, error) {
	rBytes, err := json.Marshal(*u)
	if err != nil {
		return "{}", err
	}
	return string(rBytes), nil
}

func (u *User) String() string {
	return fmt.Sprintf("{\"first_name\": %v, \"last_name\": %v, \"age\": %v, \"gender\": %v, \"date_of_birth\": %v, \"status\": %v, \"height\": %v}", u.FirstName, u.LastName, u.Age, u.Gender, u.DateOfBirth, u.Status, u.Height)
}

func (u *User) GetName() pkg.DomainName {
	return "users"
}

func (u *User) ToDto() interface{} {
	dobProto, _ := ptypes.TimestampProto(u.DateOfBirth)
	return &pikachu_v1.UserDto{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Gender:      u.Gender,
		DateOfBirth: dobProto,
		Height:      u.Height,
		Weight:      u.Weight,
		Age:         u.Age,
		ExternalId:  u.ExternalId,
	}
}

func (u *User) FillProperties(dto interface{}) pkg.Base {
	userDto := dto.(pikachu_v1.UserDto)
	u.FirstName = userDto.FirstName
	u.LastName = userDto.LastName
	u.Gender = userDto.Gender
	u.DateOfBirth = userDto.DateOfBirth.AsTime()
	u.Age = userDto.Age
	u.Height = userDto.Height
	u.Weight = userDto.Weight
	u.DeletedAt = nil
	u.Password = userDto.Password
	return u
}

func (u *User) Merge(other interface{}) {
	updatableUser := other.(*User)
	if updatableUser.Age != 0 {
		u.Age = updatableUser.Age
	}
	if updatableUser.Weight != 0 {
		u.Weight = updatableUser.Weight
	}
	if updatableUser.Height != 0 {
		u.Height = updatableUser.Height
	}
	if updatableUser.FirstName != "" {
		u.FirstName = updatableUser.FirstName
	}
	if updatableUser.LastName != "" {
		u.LastName = updatableUser.LastName
	}
	if &updatableUser.DateOfBirth != nil {
		u.DateOfBirth = updatableUser.DateOfBirth
	}
}

func (u *User) FromSqlRow(rows *sql.Rows) (pkg.Base, error) {
	err := rows.Scan(&u.ExternalId, &u.Id, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Status, &u.FirstName,
		&u.LastName, &u.Gender, &u.DateOfBirth, &u.Age, &u.Height, &u.Weight)
	return u, err
}

func (u *User) AddIdentity(identity Identity) error {
	for _, existingIdentity := range u.Identities {
		if existingIdentity.IdentityType == identity.IdentityType && existingIdentity.IdentityValue == identity.IdentityValue {
			return errors.New(fmt.Sprintf("identity with value %v and type %v already exists", identity.IdentityType, identity.IdentityValue))
		}
	}
	u.Identities = append(u.Identities, identity)
	return nil
}

func (u *User) SetExternalId(externalId string) {
	u.ExternalId = externalId
}
