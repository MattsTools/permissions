package permissions

// An imutable struct indicating permission metadata
type Permissions struct {
	invokingUserID               string
	cachedPermittedSubscriptions []string
}

func (g Permissions) CompareUserID(user string) bool {
	return user == g.invokingUserID
}

func (g Permissions) GetInvokingUserID() string {
	return g.invokingUserID
}

func ConstructPermissions(invokingUserID string, cachedPermittedSubscriptions []string) Permissions {
	return Permissions{
		invokingUserID:               invokingUserID,
		cachedPermittedSubscriptions: cachedPermittedSubscriptions,
	}
}
