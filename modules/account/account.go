package account

import (
  "database/sql"
)

var apiDB *sql.DB
var realmdDB *sql.DB

func Init(adb, rdb *sql.DB) {
  apiDB = adb
  realmdDB = rdb
}
