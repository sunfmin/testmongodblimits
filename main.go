package main

import (
	"fmt"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Email string
	Name  string
}

func (u *User) MakeId() interface{} {
	return bson.NewObjectId()
}

func main() {
	i := 1
	for {
		dbname := fmt.Sprintf("testmongodblimits_%d", i)
		db := mgodb.NewDatabase("localhost", dbname)
		err := db.Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})
		if err != nil {
			fmt.Println(err)
		}
		i = i + 1
		fmt.Printf("database created: %s\n", dbname)
	}
}
