package main

import (
	"context"
	"crypto/tls"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func main() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"https://192.168.64.5:32100"},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("test", "test"),
	})
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	activeUser := true
	options := &driver.UserOptions{"Secret123", &activeUser, ""}
	createUser, err := c.CreateUser(ctx, "ingest", options)
	if err != nil {
		fmt.Println(err)
	}
	createUser, err = c.CreateUser(ctx, "ax-admin", options)
	if err != nil {
		fmt.Println(err)
	}

	createdb, err := c.CreateDatabase(ctx, "axdb", nil)
	if err != nil {
		fmt.Println(err)
	}

	addUserToDb, err := c.User(ctx, "ingest")
	axdb, err := c.Database(ctx, "axdb")
	addDbUser := addUserToDb.GrantReadWriteAccess(ctx, axdb)
	addUserToDb, err = c.User(ctx, "ax-admin")
	addDbUser = addUserToDb.GrantReadWriteAccess(ctx, axdb)

	fmt.Println(createUser)
	fmt.Println(createdb)
	fmt.Println(addUserToDb)
	fmt.Println(addDbUser)
}
