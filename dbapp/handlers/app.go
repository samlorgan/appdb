package handlers

import (
	"dbapp/db"
	"dbapp/ent"
)

type App struct {
	DbClient *ent.Client
	Database db.IDatabase
}
