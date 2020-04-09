package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wednesday-solution/go-boiler/pkg/utl/postgres"

	"github.com/go-pg/migrations/v7"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {

	flag.Usage = usage
	flag.Parse()
	db := postgres.MigrationConnect()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
func DropTable(db migrations.DB, tableName string) error {
	fmt.Println(fmt.Sprintf("dropping table %s ...", tableName))
	_, err := db.Exec(fmt.Sprintf("DROP TABLE %s", tableName))
	return err
}

func CreateTableAndAddTrigger(db migrations.DB, createTableQuery string, tableName string) error {
	fmt.Print(fmt.Sprintf("\nCreating %s\n", tableName))
	_, err := db.Exec(createTableQuery)

	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf(`CREATE TRIGGER update_user_modtime BEFORE UPDATE ON %s FOR EACH ROW EXECUTE PROCEDURE  update_updated_at_column();`, tableName))

	if err != nil {
		return err
	}
	return err
}

// HandleMigrations - doesn't work now since migrations.MustRegister requires
// the name of the file that executes it to be in the form of 1_blah_blah.go
func HandleMigrations(tableName string, createTableQuery string) {
	migrations.MustRegister(func(db migrations.DB) error {
		err := CreateTriggerForUpdatedAt(db)
		if err != nil {
			return err
		}
		err = CreateTableAndAddTrigger(db, createTableQuery, tableName)
		if err != nil {
			return err
		}
		return err
	}, func(db migrations.DB) error {
		return DropTable(db, tableName)
	})
}

func CreateTriggerForUpdatedAt(db migrations.DB) error {
	_, err := db.Exec(`CREATE OR REPLACE FUNCTION update_updated_at_column()
			RETURNS TRIGGER AS $$
			BEGIN
			NEW.updated_at = now();
			RETURN NEW;
			END;
			$$ language 'plpgsql';`)
	return err
}
