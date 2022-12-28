package factory

import "log"

type mongoDB struct {
	database map[string]string
}

type sqliteDB struct {
	database map[string]string
}

type IDatabase interface {
	GetData(string) string
	PutData(string, string)
}

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}
	log.Printf("MongoDB")
	return mdb.database[query]
}

func (sqlDB sqliteDB) GetData(query string) string {
	if _, ok := sqlDB.database[query]; !ok {
		return ""
	}
	log.Printf("SqliteDB")
	return sqlDB.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sqlDB sqliteDB) PutData(query string, data string) {
	sqlDB.database[query] = data
}

func NewDatabaseFactory(env string) IDatabase {
	switch env {
	case "production":
		return mongoDB{database: make(map[string]string)}
	case "development":
		return sqliteDB{database: make(map[string]string)}
	default:
		return nil
	}
}
