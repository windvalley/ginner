package user

import "use-gin/model"

// Create create user by username and password
func Create(username, password string) error {
	u := &model.User{
		Username: username,
		Password: password,
	}

	if err := u.EncryptPassword(); err != nil {
		return err
	}

	err := u.Create()
	return err
}
