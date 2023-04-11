package user

type User struct {
	Login string
	// TODO: More fields
}

type SliceCache []User

func (c SliceCache) Find(login string) (User, bool) {
	for _, u := range c {
		if u.Login == login {
			return u, true
		}
	}

	return User{}, false
}

type MapCache map[string]User // login -> User

func (c MapCache) Find(login string) (User, bool) {
	user, ok := c[login]
	return user, ok
}
