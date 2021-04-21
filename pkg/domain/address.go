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

type Address struct {
	pkg.BaseDomain
	Line1   string
	Line2   string
	ZipCode string
	State   string
	Country string
	UserID  string `gorm:"type:varchar(100)"`
}

func (a *Address) ToBytes() (*bytes.Buffer, error) {
	// this thing escapes to the heap
	var aBytes bytes.Buffer
	enc := gob.NewEncoder(&aBytes)
	err := enc.Encode(*a)
	return &aBytes, err
}

func (a *Address) ToJson() (string, error) {
	jsonBytes, err := json.Marshal(*a)
	if err != nil {
		return "{}", err
	}
	return string(jsonBytes), nil
}

func (a *Address) String() string {
	return fmt.Sprintf("{\"external_id\":%v, \"line1\": %v, \"line2\": %v, \"zip_code\": %v, \"country\": %v, \"user_id\":%v, \"state\": %v}",
		a.ExternalId, a.Line1, a.Line2, a.ZipCode, a.Country, a.UserID, a.State)
}

func (a *Address) GetName() pkg.DomainName {
	return "addresses"
}

func (a *Address) ToDto() interface{} {
	return pikachu_v1.AddressDto{
		Country: a.Country,
		Line1:   a.Line1,
		Line2:   a.Line2,
		State:   a.State,
		ZipCode: a.ZipCode,
	}
}

func (a *Address) FillProperties(dto interface{}) pkg.Base {
	addressDto := dto.(pikachu_v1.AddressDto)
	a.Country = addressDto.Country
	a.Line2 = addressDto.Line2
	a.Line1 = addressDto.Line1
	a.State = addressDto.State
	a.ZipCode = addressDto.State
	return a
}

func (a *Address) Merge(other interface{}) {
	address := other.(pikachu_v1.AddressDto)
	if address.Line1 != "" {
		a.Line1 = address.Line1
	}
	if address.Line2 != "" {
		a.Line2 = address.Line2
	}
	if address.Country != "" {
		a.Country = address.Country
	}
	if address.State != "" {
		a.State = address.State
	}
	if address.ZipCode != "" {
		a.ZipCode = address.ZipCode
	}
}

func (a *Address) FromSqlRow(rows *sql.Rows) (pkg.Base, error) {
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.ZipCode, &a.Line1, &a.Line2, &a.Country, &a.ExternalId, &a.State)
		if err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (a *Address) SetExternalId(externalId string) {
	a.ExternalId = externalId
}
