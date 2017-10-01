// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package db

import (
	"database/sql"
	"fmt"
	"strings"

	// Blank import used for managing PostGres within this package only
	_ "github.com/lib/pq"
	errgo "gopkg.in/errgo.v1"
)

// New creates a new DB object and connects to the database specified by the
// options.
func New(options *Options) (*DB, error) {
	db := &DB{
		options: options,
	}
	connStr, err := db.options.ConnectionString()
	if err != nil {
		return nil, errgo.Mask(err)
	}
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.conn = conn
	return db, nil
}

// Options holds all the options allowing connections to a database.
type Options struct {
	User                    string
	Pass                    string
	DBName                  string
	Host                    string
	Port                    int
	SSLMode                 string
	FallbackApplicationName string
	ConnectionTimeout       int
	SSLCert                 string
	SSLKey                  string
	SSLRootCert             string
}

func escape(s string) string {
	// Escape spaces and single quotes.
	if strings.ContainsAny(s, " '") {
		return fmt.Sprintf("'%s'", strings.Replace(s, "'", "\\'", -1))
	}
	return s
}

// ConnectionString builds the connection string required to connect to a
// postgres database.
func (o *Options) ConnectionString() (string, error) {
	var opts []string
	if o.DBName != "" {
		opts = append(opts, fmt.Sprintf("dbname=%s", escape(o.DBName)))
	} else {
		return "", errgo.New("dbname is required")
	}
	if o.User != "" {
		opts = append(opts, fmt.Sprintf("user=%s", escape(o.User)))
	}
	if o.Pass != "" {
		opts = append(opts, fmt.Sprintf("password=%s", escape(o.Pass)))
	}
	if o.Host != "" {
		opts = append(opts, fmt.Sprintf("host=%s", escape(o.Host)))
	}
	if o.Port != 0 {
		opts = append(opts, fmt.Sprintf("port=%d", o.Port))
	}
	if o.SSLMode != "" {
		validOptions := []string{"disable", "require", "verify-ca", "verify-full"}
		ok := false
		for _, key := range validOptions {
			if o.SSLMode == key {
				ok = true
			}
		}
		if !ok {
			return "", errgo.New(fmt.Sprintf("invalid sslmode %s, expected one of %v", o.SSLMode, validOptions))
		}
		opts = append(opts, fmt.Sprintf("sslmode=%s", o.SSLMode))
	}
	if o.FallbackApplicationName != "" {
		opts = append(opts, fmt.Sprintf("fallback_application_name=%s", escape(o.FallbackApplicationName)))
	}
	if o.ConnectionTimeout != 0 {
		opts = append(opts, fmt.Sprintf("connection_timeout=%d", o.ConnectionTimeout))
	}
	if o.SSLCert != "" {
		opts = append(opts, fmt.Sprintf("sslcert=%s", escape(o.SSLCert)))
	}
	if o.SSLKey != "" {
		opts = append(opts, fmt.Sprintf("sslkey=%s", escape(o.SSLKey)))
	}
	if o.SSLRootCert != "" {
		opts = append(opts, fmt.Sprintf("sslrootcert=%s", escape(o.SSLRootCert)))
	}
	return strings.Join(opts, " "), nil
}

// DB holds connection information and allows querying of the database
type DB struct {
	conn    *sql.DB
	options *Options
}
