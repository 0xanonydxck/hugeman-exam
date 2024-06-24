package errors

import (
	"errors"

	"github.com/dxckboi/hugeman-exam/pkg/util"
	"gorm.io/gorm"
)

func ParsePostgresError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFound("record not found")
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return BadRequest("duplicated key")
	}

	obj := map[string]interface{}{}
	if err := util.Recast(err, &obj); err != nil {
		return err
	}

	if code, ok := obj["Code"].(string); ok {
		if val, ok := postgresErrorMap[code]; ok {
			return val
		}
	}

	return err
}

var postgresErrorMap = map[string]error{
	"02000": NotFound("no data"),
	"55000": NotFound("object not in prerequisite state"),
	"23505": BadRequest("unique violation"),
	"23502": BadRequest("not null violation"),
	"23503": BadRequest("foreign key violation"),
	"23514": BadRequest("check violation"),
	"22001": BadRequest("string data right truncation"),
	"42701": BadRequest("duplicate column"),
	"22003": BadRequest("numeric value out of range"),
	"22007": BadRequest("invalid datetime format"),
	"42P09": BadRequest("ambiguous alias"),
	"42P18": BadRequest("indeterminate datatype"),
	"22023": BadRequest("invalid parameter value"),
	"42P02": BadRequest("undefined parameter"),
	"42703": InternalServerError("undefined column"),
	"42P01": InternalServerError("undefined table"),
}
