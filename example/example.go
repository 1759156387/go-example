package example

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func RunClient(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	f := func() {

	}
	for {
		time.Sleep(1 * time.Second)
		f()
	}
}
func RunServer(wg *sync.WaitGroup, addr string, mysql_addr string) error {
	wg.Add(1)
	defer wg.Done()
	s.mysql_addr = mysql_addr
	http.HandleFunc("/", s.ServeHTTP)
	return http.ListenAndServe(addr, nil)
}

type Server struct {
	mysql_addr string
	handlers   map[string]handle
}

func (this *Server) opendb() (*sql.DB, error) {
	return sql.Open("mysql", this.mysql_addr)

}

var (
	s = Server{}
)

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	write_result := true
	var err error
	var data interface{}
	defer func(b *bool, e *error, data *interface{}) {
		if *b {
			ret := ret{}
			ret.Code = 0
			if *e != nil {
				ret.Code = 1 //failed
				ret.ErrMsg = (*e).Error()
			}
			if *data != nil {
				ret.Data = *data
			}

			b, _ := json.Marshal(ret)
			w.Write(b)
		}
	}(&write_result, &err, &data)
	err = r.ParseForm()
	if err != nil {
		return
	}
	action := r.FormValue("a")
	if action == "" {
		err = fmt.Errorf("action is nil string")
		return
	}
	h, ok := this.handlers[action]
	if !ok {
		err = fmt.Errorf("no such handle")
		return
	}
	err = h(w, r, &write_result, &data)
}

func init() {
	s.handlers = make(map[string]handle)
	s.handlers["add"] = s.add
	s.handlers["update"] = s.update
	s.handlers["del"] = s.del
}

func (this *Server) add(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	db, err := this.opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into user values(?,?)", r.FormValue("name"), r.FormValue("age"))
	*ret = "for test"
	return err
}
func (this *Server) update(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	db, err := this.opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("update user set age=? where name=?", r.FormValue("age"), r.FormValue("name"))
	*ret = "for test"
	return err
}
func (this *Server) del(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	db, err := this.opendb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from user where name=?", r.FormValue("name"))
	*ret = "for test"
	return err
}

type handle func(r http.ResponseWriter, w *http.Request, need_json *bool, ret *interface{}) error

type ret struct {
	Code   int
	ErrMsg string
	Data   interface{}
}
