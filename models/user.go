package models

type User struct {
	Name string
	Age  string
}

var UserInstance User

func (this *User) Add(name, age interface{}) error {
	db, err := opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into user values(?,?)", name, age)
	return err
}
func (this *User) Update(name, age interface{}) error {
	db, err := opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("update user set age=? where name=?", name)
	return err
}
func (this *User) Del(name, age interface{}) error {
	db, err := opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from user where name=?", name)
	return err
}
