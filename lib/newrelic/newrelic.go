package main

import (
	"context"
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

var schemaDrop = `DROP TABLE IF EXISTS person`

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
)`

// Person is a person in the database
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func createApp() *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("SQLx"),
		newrelic.ConfigLicense("---newrelic-license-key---"),
		//newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if nil != err {
		log.Fatalln(err)
	}
	if err := app.WaitForConnection(5 * time.Second); nil != err {
		log.Fatalln(err)
	}
	return app
}

func main() {
	// Create application
	app := createApp()
	defer app.Shutdown(10 * time.Second)
	// Start a transaction
	txn := app.StartTransaction("main")
	defer txn.End()
	// Add transaction to context
	ctx := newrelic.NewContext(context.Background(), txn)

	db, err := sqlx.Connect("nrmysql", "isucari:isucari@tcp(127.0.0.1:3306)/isucari?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.ExecContext(ctx, schemaDrop)
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}
	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	// Add people to the database
	// When the context is passed, DatastoreSegments will be created
	tx := db.MustBegin()
	tx.MustExecContext(ctx, "INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExecContext(ctx, "INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.Commit()

	// Read from the database
	// When the context is passed, DatastoreSegments will be created
	people := []Person{}
	err = db.SelectContext(ctx, &people, "SELECT * FROM person ORDER BY `first_name` ASC")
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	jason := Person{}
	err = db.GetContext(ctx, &jason, "SELECT * FROM person WHERE `first_name` = ?", "Jason")
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}
	err =db.GetContext(ctx, &jason, "SELECT * FROM person WHERE `first_name` = ?", "Jason")
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}
	err = db.GetContext(ctx, &jason, "SELECT * FROM person WHERE first_name=?", "Jason")
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}
}