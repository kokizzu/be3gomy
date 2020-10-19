package mahasiswa

import (
	"be3gomy/model"
	"database/sql"
	"fmt"
	"time"
)

const table = `mahasiswa`
const dateformat = `2006-01-02 15:04:05`

func SelectAll(db *sql.DB) (mahasiswas []model.Mahasiswa,err error) {
	sql := fmt.Sprintf(`SELECT * FROM %s ORDER BY id DESC`, table)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	mahasiswas = []model.Mahasiswa{}
	for rows.Next() {
		m := model.Mahasiswa{}
		createdAt, updatedAt := ``, ``
		err = rows.Scan(&m.ID,
			&m.NIM,
			&m.Name,
			&m.Semester,
			&createdAt,
			&updatedAt)
		m.CreatedAt, _ = time.Parse(dateformat,createdAt)
		m.UpdatedAt, _ = time.Parse(dateformat,updatedAt)
		if err !=nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, m)
	}
	return
}

func Insert(db *sql.DB, m *model.Mahasiswa) (err error) {
	sql := fmt.Sprintf(`INSERT INTO %v (nim, name, semester, created_at, updated_at) 
VALUES(?,?,?,?,?)`, table)
	now := time.Now()
	res, err := db.Exec(sql,m.NIM,m.Name,m.Semester,now,now)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = int(lastId)
	m.CreatedAt = now
	m.UpdatedAt = now
	return
}
