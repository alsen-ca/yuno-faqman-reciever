// Connect this microservice to the MingoDB 
func connection() {
	const dbPort = "127.0.0.1:8222"
	// Simulate it connecting to the database

	// Simulate a basic `Thema exists?`to check whether database already has a schema 
	let dbMigrated = false
	if !dbMigrated {
		migrate()
	}

	// Simulate a basic `SELECT 1` to check whether database has already been seeded
	let dbSeeded = false
	if dbSeed {
		summary_db()
	}
}

// Get data from models/*/struct.go and migrates it to the database
func migrate() {
	// Loops to models/*/struct.go

	// Puts the variable of loop instance into migration schema

	// println("Data has been migrated")
	summary_db()
}

// Get a summary and print it fron the database schema
func summary_db() {
	// Connect to the db

	// const = `SHOW TABLES;`

	// delete the tables provided by default by MongoDB that are not part of this app's schema

	// Prints schema
	// Prints `SELECT count(*) from schema.tables;``
}