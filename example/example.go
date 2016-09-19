package example

import (
	"database/sql"
	"encoding/json"
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
}

func init() {
	s.handlers = make(map[string]handle)
	s.handlers["add"] = s.add
	s.handlers["update"] = s.update
	s.handlers["del"] = s.del
}

func (this *Server) add(r http.ResponseWriter, w *http.Request, need_json *bool) error {
	return nil
}
func (this *Server) update(r http.ResponseWriter, w *http.Request, need_json *bool) error {
	return nil
}
func (this *Server) del(r http.ResponseWriter, w *http.Request, need_json *bool) error {
	return nil
}

type handle func(r http.ResponseWriter, w *http.Request, need_json *bool) error

type ret struct {
	Code   int
	ErrMsg string
	Data   interface{}
}
