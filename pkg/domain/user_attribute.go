package domain

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type UserAttribute struct {
	pkg.BaseDomain
	AttributeKey   string
	AttributeValue string
	UserID         string `gorm:"type:varchar(100)"`
}

func (i *UserAttribute) ToBytes() (*bytes.Buffer, error) {
	var rBytes bytes.Buffer
	enc := gob.NewEncoder(&rBytes)
	err := enc.Encode(*i)
	return &rBytes, err
}

func (i *UserAttribute) ToJson() (string, error) {
	rBytes, err := json.Marshal(*i)
	if err != nil {
		return "{}", err
	}
	return string(rBytes), nil
}

func (i *UserAttribute) String() string {
	return fmt.Sprintf("{\"external_id\": %v, \"attribute_key\": %v, \"attribute_value\": %v}", i.ExternalId, i.AttributeKey, i.AttributeValue)
}

func (i *UserAttribute) GetName() pkg.DomainName {
	return "user_attributes"
}

func (i *UserAttribute) ToDto() interface{} {
	return pikachu_v1.UserAttributeDto{
		AttributeKey:   i.AttributeKey,
		AttributeValue: i.AttributeValue,
		ExternalId:     i.ExternalId,
	}
}

func (i *UserAttribute) FillProperties(dto interface{}) pkg.Base {
	userAttributeDto := dto.(pikachu_v1.UserAttributeDto)
	i.AttributeKey = userAttributeDto.AttributeKey
	i.AttributeValue = userAttributeDto.AttributeValue
	return i
}

func (i *UserAttribute) Merge(other interface{}) {
	userAttributeDto := other.(*UserAttribute)
	if userAttributeDto.AttributeValue != "" {
		i.AttributeValue = userAttributeDto.AttributeValue
	}
}

func (i *UserAttribute) FromSqlRow(rows *sql.Rows) (pkg.Base, error) {
	err := rows.Scan(&i.Id, &i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Status, &i.AttributeKey, &i.AttributeValue)
	return i, err
}

func (i *UserAttribute) SetExternalId(externalId string) {
	i.ExternalId = externalId
}
