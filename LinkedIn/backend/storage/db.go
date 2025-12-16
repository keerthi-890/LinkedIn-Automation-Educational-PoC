package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // ✅ Pure Go SQLite (NO CGO)
)



func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS profiles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		title TEXT,
		location TEXT,
		profile_id TEXT UNIQUE,
		status TEXT DEFAULT 'New',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS activity_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		action TEXT,
		details TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(createTables)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func LogActivity(db *sql.DB, action, details string) {
	_, err := db.Exec(
		"INSERT INTO activity_logs (action, details) VALUES (?, ?)",
		action, details,
	)
	if err != nil {
		log.Printf("Error logging activity: %v", err)
	}
}

func GetStats(db *sql.DB) (Stats, error) {
	var s Stats // uses Stats from models.go

	row := db.QueryRow("SELECT COUNT(*) FROM profiles")
	row.Scan(&s.ProfilesFound)

	row = db.QueryRow("SELECT COUNT(*) FROM profiles WHERE status = 'Connected'")
	row.Scan(&s.AcceptedConnections)

	s.RequestsSent = s.AcceptedConnections // PoC simplification
	s.MessagesSent = 0

	return s, nil
}

