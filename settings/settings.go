package settings

import (
    "path/filepath"
)


// One directory back of current directory (Project Root)
var BASE_DIR, err = filepath.Abs("./../")

var DB_TYPE = "postgres"
var DATABASE = map[string]string {
        "dbname": "gocrud",
        "user": "postgres",
        "host": "localhost",
        "port": "5432",
        "password": "",
    }
