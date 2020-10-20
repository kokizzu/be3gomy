package main

import (
	"be3gomy/config"
	"be3gomy/mahasiswa"
	"be3gomy/model"
	"be3gomy/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
const addr = `:7000`
type Server struct {
	db      *sql.DB
	ViewDir string
} 
type handler func(w http.ResponseWriter, r *http.Request)
func InitServer() *Server {
	db, err := config.Mysql()
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		db:      db,
		ViewDir: `views/`,
	}
}
func (s *Server) Listen() {
	log.Println(`listen at `+addr)
	http.HandleFunc(`/mahasiswa`, s.Mahasiswa())
	http.HandleFunc(`/mahasiswa/create`,s.MahasiswaCreate())
	// update: id, namaBaru, nimBaru, semesterBaru; {jumlahBerhasilDiupdate:0}
	// delete: id; response: {berhasil:bool,record:{}}
	http.HandleFunc(`/mahasiswa/delete`,s.MahasiswaDelete())
	err := http.ListenAndServe(addr,nil)
	if err != nil {
		fmt.Println(err)
	}
}
func (s *Server) Mahasiswa() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		mahasiswas, err := mahasiswa.SelectAll(s.db)
		if utils.IsError(w,err) {
			return
		}
		utils.ResponseJson(w, mahasiswas)
	}
}
func (s *Server) MahasiswaCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == `GET` {
			http.ServeFile(w,r,s.ViewDir+`mahasiswa_create.html`)
			return
		}
		m := model.Mahasiswa{}
		err := json.NewDecoder(r.Body).Decode(&m)
		if utils.IsError(w,err) {
			return
		}
		err = mahasiswa.Insert(s.db,&m)
		if utils.IsError(w,err) {
			return
		}
		utils.ResponseJson(w, m)
	}
}

func (s *Server) MahasiswaDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == `GET` {
			http.ServeFile(w,r,s.ViewDir+`mahasiswa_delete.html`)
			return
		}
		m := model.Mahasiswa{}
		err := json.NewDecoder(r.Body).Decode(&m)
		if utils.IsError(w,err) {
			return
		}
		ok, err := mahasiswa.Delete(s.db,&m)
		if utils.IsError(w,err) {
			return
		}
		res := map[string]interface{}{}
		res[`deletedRecord`] = m
		res[`berhasil`] = ok
		utils.ResponseJson(w, res)
	}
}

func main() {
	server := InitServer()
	server.Listen()
}
