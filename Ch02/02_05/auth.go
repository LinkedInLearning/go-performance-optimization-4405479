package auth

type User struct {
	Login string
	Token string
}

type Store []User

// Token returns auth token for login, empty string if not found.
func (s Store) Token(login string) string {
	i := s.userIndex(login)
	if i == -1 {
		return ""
	}

	return s[i].Token
}

// userIndex return index where user it, -1 if not found.
func (s Store) userIndex(login string) int {
	for i, u := range s {
		if u.Login == login {
			return i
		}
	}

	return -1
}
