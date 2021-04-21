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

type Relation struct {
	pkg.BaseDomain
	RelationType core_v1.Relation
	RelationID   string `gorm:"type:varchar(100)"`
	UserID       string `gorm:"type:varchar(100)"`
}

func (r *Relation) ToBytes() (*bytes.Buffer, error) {
	var rBytes bytes.Buffer
	enc := gob.NewEncoder(&rBytes)
	err := enc.Encode(*r)
	return &rBytes, err
}

func (r *Relation) ToJson() (string, error) {
	rBytes, err := json.Marshal(*r)
	if err != nil {
		return "{}", err
	}
	return string(rBytes), nil
}

func (r *Relation) String() string {
	return fmt.Sprintf("{\"external_id\": %v, \"relation_id\": %v, \"user_id\": %v, \"relation_type\": %v}", r.ExternalId, r.RelationID, r.UserID, r.RelationType)
}

func (r *Relation) GetName() pkg.DomainName {
	return "relations"
}

func (r *Relation) ToDto() interface{} {
	return pikachu_v1.RelationDto{
		// TODO add relation id
		UserId: r.UserID,
		Relation: r.RelationType,
	}
}

func (r *Relation) FillProperties(dto interface{}) pkg.Base {
	rDto := dto.(*pikachu_v1.RelationDto)
	r.UserID = rDto.UserId
	r.RelationType = rDto.Relation
	return r
}

func (r *Relation) Merge(other interface{}) {
	oDomain := other.(*Relation)
	if oDomain.RelationType != 0 {
		r.RelationType = oDomain.RelationType
	}
}

func (r *Relation) FromSqlRow(rows *sql.Rows) (pkg.Base, error) {
	for rows.Next(){
		err := rows.Scan(&r.ExternalId, &r.RelationID, &r.UserID, &r.RelationType)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (r *Relation) SetExternalId(externalId string) {
	panic("implement me")
}
