package types

const (
	AuthUrl = "authinfo"
)

type AuthInfo struct {
	BaseUser
	Permissions []string `json:"permissions"`
}

// CompareAuthInfo
//
// Compares 2 authinfo structs, used for testing
func CompareAuthInfo(a, b *AuthInfo) bool {
	if a.Username != b.Username || a.Id != b.Id {
		return false
	}

	if len(a.Permissions) != len(b.Permissions) {
		return false
	}

	for i := range a.Permissions {
		if a.Permissions[i] != b.Permissions[i] {
			return false
		}
	}

	return true
}
