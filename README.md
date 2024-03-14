Small demo repo for GORM interacting with a Dolt database. [This blog](https://www.dolthub.com/blog/2024-03-15-gorm-with-dolt) describes how to use this demo.

# Setup

Two dependencies are required:
- [golang](https://go.dev/doc/install)
- [dolt](https://docs.dolthub.com/introduction/installation)
```
sudo bash -c 'curl -L https://github.com/dolthub/dolt/releases/latest/download/install.sh | bash'
```

Clone the [letters repo](https://www.dolthub.com/repositories/max-hoffman/letters):

```
dolt clone evan/letters
```

Start a sql-server:

```
cd letters
dolt sql-server -l trace
```

# Organization

CRUD operations are wrapped in helper functions. Uncomment a function to test it's action against the letters database:

```
func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/letters?interpolateParams=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	readFirstLetter(db)
	//insertLetter(db)
	//updateLetter(db)
	//chainUpdateLetter(db)
	//testTicker(db)
	//generateModels(db)
	//getLastCommit(db)
	//getOneDiff(db)
	//getOneSchemaDiff(db)
}
```

GORM's [dry run mode](https://gorm.io/docs/sql_builder.html#DryRun-Mode) can be used to print SQL queries on the client side. The letters server
log will detail the exact queries sent from the client.
