package models

// List of all database models to create tables using GORM migration
var TABLES = []interface{}{
	&User{},
	&Note{},
	&Notebook{},
	&Auth{},
}
