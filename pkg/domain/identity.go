package domain

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/core_v1"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type Identity struct {
	pkg.BaseDomain
	IdentityType  core_v1.IdentityType
	IdentityValue string
	UserID        string `gorm:"type:varchar(100)"`
}

func (i *Identity) ToBytes() (*bytes.Buffer, error) {
	var iBytes bytes.Buffer
	enc := gob.NewEncoder(&iBytes)
	err := enc.Encode(*i)
	return &iBytes, err
}

func (i *Identity) ToJson() (string, error) {
	iBytes, err := json.Marshal(*i)
	if err != nil {
		return "{}", err
	}
	return string(iBytes), nil
}

func (i *Identity) String() string {
	return fmt.Sprintf("{\"external_id\": %v, \"identity_type\": %v, \"identity_value\": %v}", i.ExternalId, i.IdentityType, i.IdentityValue)
}

func (i *Identity) GetName() pkg.DomainName {
	return "identities"
}

func (i *Identity) ToDto() interface{} {
	return pikachu_v1.IdentityDto{
		IdentityType:  i.IdentityType,
		IdentityValue: i.IdentityValue,
		ExternalId:    i.ExternalId,
	}
}

func (i *Identity) FillProperties(dto interface{}) pkg.Base {
	identityDto := dto.(pikachu_v1.IdentityDto)
	i.IdentityType = identityDto.IdentityType
	i.IdentityValue = identityDto.IdentityValue
	return i
}

func (i *Identity) Merge(other interface{}) {
	identityDto := other.(*Identity)
	if identityDto.IdentityType != 0 {
		i.IdentityType = identityDto.IdentityType
	}
	if identityDto.IdentityValue != "" {
		i.IdentityValue = identityDto.IdentityValue
	}
}

func (i *Identity) FromSqlRow(rows *sql.Rows) (pkg.Base, error) {
	err := rows.Scan(&i.Id, &i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Status, &i.IdentityType, &i.IdentityValue)
	return i, err
}

func (i *Identity) SetExternalId(externalId string) {
	i.ExternalId = externalId
}
