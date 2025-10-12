package valueObjects

import (
	"github.com/go-playground/validator/v10"
)

type RoleType struct {
	Value string
}

var (
	RoleTypeAdmin    = RoleType{Value: "ADMIN"}
	RoleTypeCustomer = RoleType{Value: "CUSTOMER"}
)

var validRoleTypes = map[string]struct{}{
	"ADMIN":    {},
	"CUSTOMER": {},
}

func RoleTypeValidation(sl validator.StructLevel) {
	roleType := sl.Current().Interface().(RoleType)

	if roleType.Value == "" {
		return
	}

	_, exists := validRoleTypes[roleType.Value]
	if !exists {
		sl.ReportError(roleType, "RoleType", "RoleType", "invalid", roleType.Value)
	}
}
