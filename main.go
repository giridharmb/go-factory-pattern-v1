package main

import (
	"fmt"
	"go-factory-pattern/factory"
	"log"
	"reflect"
)

func main() {
	env1 := "production"
	env2 := "development"

	db1 := factory.NewDatabaseFactory(env1)
	db2 := factory.NewDatabaseFactory(env2)

	db1.PutData("test1", "value1")
	log.Printf(db1.GetData("test1"))

	db2.PutData("test2", "value2")
	log.Printf(db2.GetData("test2"))

	log.Printf(reflect.TypeOf(db1).Name())
	log.Printf(reflect.TypeOf(db2).Name())

	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(&db2).Elem())

}

/*
Output

2022/12/27 16:49:23 MongoDB
2022/12/27 16:49:23 value1
2022/12/27 16:49:23 SqliteDB
2022/12/27 16:49:23 value2
2022/12/27 16:49:23 mongoDB
2022/12/27 16:49:23 sqliteDB
factory.IDatabase
factory.IDatabase
*/
