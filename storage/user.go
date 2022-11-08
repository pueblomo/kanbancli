package storage

import "os/user"

var dir string

func Init() {
	usr, _ := user.Current()
	dir = usr.HomeDir
}
