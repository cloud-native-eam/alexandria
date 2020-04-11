package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// yes I know, good code looks differently...

func main() {
	//new connection
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"https://alexandria-db:8529"},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	// new client to connect to ADB
	rootUser := os.Getenv("ROOT-USER")
	rootUserPW := os.Getenv("ROOT-PW")
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(rootUser, rootUserPW),
	})
	if err != nil {
		fmt.Println(err)
	}
	// read env
	ingestPw := os.Getenv("INGEST-PW")
	adminPw := os.Getenv("ADMIN-PW")
	fmt.Println(ingestPw)
	fmt.Println(adminPw)

	ctx := context.Background()
	activeUser := true
	// set password options for user
	ingestOptions := &driver.UserOptions{ingestPw, &activeUser, ""}
	// create user ingest
	createUser, err := c.CreateUser(ctx, "ingest", ingestOptions)
	if err != nil {
		fmt.Println(err)
	}
	// create ax DB admin user
	adminOptions := &driver.UserOptions{adminPw, &activeUser, ""}
	createUser, err = c.CreateUser(ctx, "ax-admin", adminOptions)
	if err != nil {
		fmt.Println(err)
	}
	// create axdb Database
	createdb, err := c.CreateDatabase(ctx, "axdb", nil)
	if err != nil {
		fmt.Println(err)
	}
	// add ingest user to axdb
	addUserToDb, err := c.User(ctx, "ingest")
	axdb, err := c.Database(ctx, "axdb")
	addDbUser := addUserToDb.GrantReadWriteAccess(ctx, axdb)
	// add admin user to axdb
	addUserToDb, err = c.User(ctx, "ax-admin")
	addDbUser = addUserToDb.GrantReadWriteAccess(ctx, axdb)

	os.Clearenv()
	fmt.Println(createUser)
	fmt.Println(createdb)
	fmt.Println(addUserToDb)
	fmt.Println(addDbUser)
}
