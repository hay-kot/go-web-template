// Code generated by "/pkgs/automapper"; DO NOT EDIT.
package mapper

import (
	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/internal/types"
)

func UserOutFromModel(from ent.User) types.UserOut {
	return types.UserOut{
		ID:          from.ID,
		Name:        from.Name,
		Email:       from.Email,
		Password:    from.Password,
		IsSuperuser: from.IsSuperuser,
	}
}

func UserOutToModel(from types.UserOut) ent.User {
	return ent.User{
		ID:          from.ID,
		Name:        from.Name,
		Email:       from.Email,
		Password:    from.Password,
		IsSuperuser: from.IsSuperuser,
	}
}
