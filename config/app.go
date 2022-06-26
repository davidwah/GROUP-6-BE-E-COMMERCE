package config

import "os"

var SECRET_JWT string = os.Getenv("secret")

