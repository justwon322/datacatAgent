package main

import (
	"context"
	"datacatAgent/ent"
	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Create a new MySQL database connection.
	dsn := "admin:oper-2018@tcp(138.2.116.22:3306)/datacat?parseTime=true"
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Fetch the Program by ID
	//ctx := context.Background()
	//
	//// Fetch the User by ID
	//fetchedUser, err := client.User.
	//	Query().
	//	Where(user.ID(1)).
	//	Only(ctx)
	//if err != nil {
	//	log.Fatalf("failed fetching user: %v", err)
	//}
	//
	//fmt.Printf("User fetched: %v\n", fetchedUser)
}
