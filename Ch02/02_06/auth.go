package auth

type User struct {
	Login string
	Token string
}

type Store map[string]User // login -> User

// Token returns auth token for login, empty string if not found.
func (s Store) Token(login string) string {
	u, ok := s[login]
	if !ok {
		return ""
	}

	return u.Token
}
