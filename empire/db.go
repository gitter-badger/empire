package empire

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
)

func newDB(uri string) (*gorm.DB, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open(u.Scheme, uri)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to connect to postgres: %v", err)
	}

	db, err := gorm.Open(u.Scheme, conn)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
