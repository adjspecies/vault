// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package db_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/adjspecies/vault/db"
)

var connStrTests = []struct {
	about          string
	options        *db.Options
	expectedString string
	expectedError  string
}{
	{
		about: "using all options",
		options: &db.Options{
			User:                    "rose",
			Pass:                    "thedoctor",
			DBName:                  "master",
			Host:                    "no-daleks.com",
			Port:                    8042,
			SSLMode:                 "disable",
			FallbackApplicationName: "notpq",
			ConnectionTimeout:       300,
			SSLCert:                 "/path/to/cert",
			SSLKey:                  "/path/to/key",
			SSLRootCert:             "/path/to/ca",
		},
		expectedString: "dbname=master user=rose password=thedoctor host=no-daleks.com port=8042 sslmode=disable fallback_application_name=notpq connection_timeout=300 sslcert=/path/to/cert sslkey=/path/to/key sslrootcert=/path/to/ca",
	},
	{
		about: "using some optoins",
		options: &db.Options{
			User:   "rose",
			Pass:   "thedoctor",
			DBName: "master",
		},
		expectedString: "dbname=master user=rose password=thedoctor",
	},
	{
		about: "escaping strings",
		options: &db.Options{
			User:   "rose tyler",
			Pass:   "the doctor's favorite",
			DBName: "master",
		},
		expectedString: "dbname=master user='rose tyler' password='the doctor\\'s favorite'",
	},
	{
		about:         "dbname is required",
		options:       &db.Options{},
		expectedError: "dbname is required",
	},
	{
		about: "invalid sslmode",
		options: &db.Options{
			DBName:  "master",
			SSLMode: "bad-wolf",
		},
		expectedError: "invalid sslmode bad-wolf, expected one of [disable require verify-ca verify-full]",
	},
}

func TestOptions(t *testing.T) {
	Convey("it should be able to generate a connection string", t, func() {
		for _, test := range connStrTests {
			Convey(fmt.Sprintf("...%s", test.about), func() {
				connStr, err := test.options.ConnectionString()
				if test.expectedError != "" {
					So(err, ShouldBeError, test.expectedError)
				} else {
					So(connStr, ShouldEqual, test.expectedString)
				}
			})
		}
	})
}
