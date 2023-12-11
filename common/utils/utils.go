package utils

import "miniproject/common/crypto"

func CheckRole(role crypto.Role, roles []crypto.Role) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}
	return false
}
