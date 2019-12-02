package database

func AddUser(user *User) (err error ){
	d := Passport.Create(user)
	if d.Error != nil {
		err = d.Error
		return err
	}
	return nil
}

func ModifiedUserPwd(username, pwd string) (err error) {
	d := Passport.Model(&User{}).Where("username = ?", username).Update("pwd", pwd)
	if d.Error != nil {
		err = d.Error
		return err
	}
	return nil
}

func DeleteUserByID(ID uint64) (err error) {
	d := Passport.Where("ID = ?", ID).Delete(User{})
	if  d.Error != nil {
		err = d.Error
		return err
	}
	return nil
}

func DeleteUserByName(name string) (err error) {
	d := Passport.Where("username = ?", name).Delete(User{})
	if d.Error != nil {
		err = d.Error
		return err
	}
	return nil
}

func QueryPwdByName(name string) (string, error) {
	user := User{}
	d := Passport.Select("pwd").Where("username = ?", name).First(&user)
	if d.Error != nil {
		return "", d.Error
	}
	if d.RecordNotFound() {
		return "RecordNotFound", nil
	}
	return user.Pwd, nil
}
