package models

import (
	"fmt"
)

type User struct {
	Number  string
	Cust_id string
}

var UserInstance User

func (this *User) AllUser(t interface{}, r *interface{}) error {
	db, err := opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Query("select number,cust_id from cust_inf")
	if err != nil {
		return err
	}
	defer rows.Close()
	ret := []*User{}
	for rows.Next() {
		s := &User{}
		err := rows.Scan(&s.Number, &s.Cust_id)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		ret = append(ret, s)
	}
	*r = ret
	return err
}
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
