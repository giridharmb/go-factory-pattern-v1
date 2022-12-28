package main

import (
	"fmt"
	"go-factory-pattern/factory"
	"log"
	"reflect"
)

func SetupConstructors(env string) (factory.IDatabase, factory.IFileSystem) {
	fs := factory.AbstractFactory("filesystem")
	db := factory.AbstractFactory("database")

	return db(env).(factory.IDatabase), fs(env).(factory.IFileSystem)
}

func main() {
	env1 := "production"
	env2 := "development"

	db1, fs1 := SetupConstructors(env1)
	db2, fs2 := SetupConstructors(env2)

	// -----------------------------------------------------

	db1.PutData("test1", "value1")
	log.Printf(db1.GetData("test1"))

	db2.PutData("test2", "value2")
	log.Printf(db2.GetData("test2"))

	log.Printf(reflect.TypeOf(db1).Name())
	log.Printf(reflect.TypeOf(db2).Name())

	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	// -----------------------------------------------------

	fs1.CreateFile("/tmp/fs1.txt")
	fmt.Println(fs1.FileFile("/tmp/fs1.txt"))

	fs2.CreateFile("/tmp/fs2.txt")
	fmt.Println(fs2.FileFile("/tmp/fs2.txt"))

	log.Printf(reflect.TypeOf(fs1).Name())
	log.Printf(reflect.TypeOf(fs2).Name())

	fmt.Println(reflect.TypeOf(&fs1).Elem())
	fmt.Println(reflect.TypeOf(&fs2).Elem())

	// -----------------------------------------------------

}

/*
Output

2022/12/27 21:03:21 MongoDB
2022/12/27 21:03:21 value1
2022/12/27 21:03:21 SqliteDB
2022/12/27 21:03:21 value2
2022/12/27 21:03:21 mongoDB
2022/12/27 21:03:21 sqliteDB
factory.IDatabase
factory.IDatabase
2022/12/27 21:03:21 ntfs...
{ntfs file /tmp/fs1.txt}
2022/12/27 21:03:21 ext4...
{ext4 file /tmp/fs2.txt}
2022/12/27 21:03:21 ntfs
2022/12/27 21:03:21 ext4
factory.IFileSystem
factory.IFileSystem
*/
