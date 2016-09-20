package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/756445638/go-example/models"
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
	models.DBSTR = mysql_addr
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

			b, _ := json.MarshalIndent(ret, "", "\t")
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
	s.handlers["addUser"] = s.addUser
	s.handlers["updateUser"] = s.updateUser
	s.handlers["delUser"] = s.delUser
	s.handlers["allUser"] = s.allUser
}

func (this *Server) allUser(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	return models.UserInstance.AllUser(r.FormValue("type"), ret)
}

func (this *Server) addUser(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	*ret = "for test"
	return models.UserInstance.Add(r.FormValue("name"), r.FormValue("age"))
}
func (this *Server) updateUser(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	*ret = "for test"
	return models.UserInstance.Update(r.FormValue("name"), r.FormValue("age"))
}
func (this *Server) delUser(w http.ResponseWriter, r *http.Request, write_result *bool, ret *interface{}) error {
	*ret = "for test"
	return models.UserInstance.Del(r.FormValue("name"), r.FormValue("age"))
}

type handle func(r http.ResponseWriter, w *http.Request, need_json *bool, ret *interface{}) error

type ret struct {
	Code   int
	ErrMsg string
	Data   interface{}
}
