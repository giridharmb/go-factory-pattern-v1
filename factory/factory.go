package factory

import "log"

type mongoDB struct {
	database map[string]string
}

type sqliteDB struct {
	database map[string]string
}

// -----------------------------------------------------

type IDatabase interface {
	GetData(string) string
	PutData(string, string)
}

// -----------------------------------------------------

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

// -----------------------------------------------------

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sqlDB sqliteDB) PutData(query string, data string) {
	sqlDB.database[query] = data
}

// -----------------------------------------------------

func NewDatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return sqliteDB{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

// -----------------------------------------------------

type file struct {
	name    string
	content string
}

type ntfs struct {
	files map[string]file
}

type ext4 struct {
	files map[string]file
}

// -----------------------------------------------------

type IFileSystem interface {
	CreateFile(string)
	FileFile(string) file
}

// -----------------------------------------------------

func (n ntfs) CreateFile(s string) {
	myFile := file{
		name:    "ntfs file",
		content: s,
	}
	n.files[s] = myFile
	log.Printf("ntfs...")
}

func (n ntfs) FileFile(s string) file {
	if _, ok := n.files[s]; !ok {
		return file{}
	}
	return n.files[s]
}

// -----------------------------------------------------

func (e ext4) CreateFile(s string) {
	myFile := file{
		name:    "ext4 file",
		content: s,
	}
	e.files[s] = myFile
	log.Printf("ext4...")
}

func (e ext4) FileFile(s string) file {
	if _, ok := e.files[s]; !ok {
		return file{}
	}
	return e.files[s]
}

// -----------------------------------------------------

func NewFilesystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

// -----------------------------------------------------

type Factory func(string) interface{}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return NewDatabaseFactory
	case "filesystem":
		return NewFilesystemFactory
	default:
		return nil
	}
}
