package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	// Get database URL from environment or command line
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		if len(os.Args) > 1 {
			dsn = os.Args[1]
		} else {
			fmt.Println("Usage: go run backend/cmd/db/migrate.go <postgres_connection_string>")
			fmt.Println("Or set DATABASE_URL environment variable")
			os.Exit(1)
		}
	}

	// Connect to database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to database")
	return db
}

func getMigrationFiles() []string {
	// Get the executable's directory to find migrations
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	migrationsDir := filepath.Join(exeDir, "migrations")

	// Read directory contents
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v", err)
	}

	// Filter and collect SQL files
	var migrationFiles []string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".sql") {
			migrationFiles = append(migrationFiles, f.Name())
		}
	}

	// Sort files by name (they should be prefixed with numbers for ordering)
	sort.Strings(migrationFiles)

	// Print found migrations for debugging
	fmt.Printf("Found %d migration files:\n", len(migrationFiles))
	for _, f := range migrationFiles {
		fmt.Printf("  - %s\n", f)
	}

	return migrationFiles
}

func getAppliedMigrations(db *sql.DB) map[string]bool {
	ctx := context.Background()
	applied := make(map[string]bool)

	// Query applied migrations
	rows, err := db.QueryContext(ctx, "SELECT name FROM migrations")
	if err != nil {
		// If the table doesn't exist yet, that's fine - no migrations have been applied
		if strings.Contains(err.Error(), "relation \"migrations\" does not exist") {
			fmt.Println("No migrations have been applied yet")
			return applied
		}
		log.Fatalf("Failed to query migrations table: %v", err)
	}
	defer rows.Close()

	// Read all applied migrations
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalf("Failed to scan migration name: %v", err)
		}
		applied[name] = true
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error reading migrations: %v", err)
	}

	// Print applied migrations for debugging
	fmt.Printf("Found %d applied migrations:\n", len(applied))
	for name := range applied {
		fmt.Printf("  - %s\n", name)
	}

	return applied
}

func applyMigrations(db *sql.DB, migrations []string, applied map[string]bool) {
	ctx := context.Background()

	// Get the executable's directory for migration files
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	migrationsDir := filepath.Join(exeDir, "migrations")

	for _, fname := range migrations {
		if applied[fname] {
			fmt.Printf("Skipping already applied migration: %s\n", fname)
			continue
		}

		// Read migration file
		path := filepath.Join(migrationsDir, fname)
		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", fname, err)
		}

		// Start transaction
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to start transaction for %s: %v", fname, err)
		}

		// Apply migration
		fmt.Printf("Applying migration: %s\n", fname)
		if _, err := tx.ExecContext(ctx, string(sqlBytes)); err != nil {
			tx.Rollback()
			log.Fatalf("Failed to apply migration %s: %v", fname, err)
		}

		// Record migration
		if _, err := tx.ExecContext(ctx,
			"INSERT INTO migrations (name) VALUES ($1)",
			fname,
		); err != nil {
			tx.Rollback()
			log.Fatalf("Failed to record migration %s: %v", fname, err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			log.Fatalf("Failed to commit migration %s: %v", fname, err)
		}

		fmt.Printf("Successfully applied migration: %s\n", fname)
	}
}

func main() {
	db := connectDB()
	defer db.Close()

	migrations := getMigrationFiles()
	applied := getAppliedMigrations(db)
	applyMigrations(db, migrations, applied)
}
