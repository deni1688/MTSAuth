package main

import (
	"github.com/deni1688/motusauth/api"
	"github.com/deni1688/motusauth/db"
	"github.com/deni1688/motusauth/models"
)

func main() {
	db.Init([]interface{}{models.User{}})
	api.Init()
}
