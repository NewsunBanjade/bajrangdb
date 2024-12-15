package main

import (
	"bajrangdb/database"
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

// Function to create an entry
func create(db *database.Database, id, data string) {
	db.Insert(id, data)
	fmt.Println("Data inserted successfully.")
}

// Function to update an entry
func update(db *database.Database, id, data string) {
	if _, exists := db.Get(id); exists {
		db.Insert(id, data)
		fmt.Println("Data updated successfully.")
	} else {
		fmt.Println("Entry does not exist.")
	}
}

// Function to delete an entry
func delete(db *database.Database, id string) {
	db.Delete(id)
	fmt.Println("Data deleted successfully.")
}

// Function to get an entry
func get(db *database.Database, id string) {
	entry, exists := db.Get(id)
	if exists {
		fmt.Printf("Data for %s: %s\n", id, entry.Data)
	} else {
		fmt.Println("Entry not found.")
	}
}

func main() {
	// Initialize the database
	dbPath := "./db/data.json"
	db, err := database.NewDatabase(dbPath)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "bajrang",
		Short: "Bajrang is a CLI tool to interact with the database",
		Long:  "Bajrang is a command-line tool that allows you to create, update, delete, and get entries in the local database.",
	}

	// Create command
	var createCmd = &cobra.Command{
		Use:   "create [id] [data]",
		Short: "Create a new entry in the database",
		Args:  cobra.ExactArgs(2), // Require exactly 2 arguments (id and data)
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			data := args[1]
			create(db, id, data)
		},
	}

	// Update command
	var updateCmd = &cobra.Command{
		Use:   "update [id] [data]",
		Short: "Update an existing entry in the database",
		Args:  cobra.ExactArgs(2), // Require exactly 2 arguments (id and data)
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			data := args[1]
			update(db, id, data)
		},
	}

	// Delete command
	var deleteCmd = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete an entry from the database",
		Args:  cobra.ExactArgs(1), // Require exactly 1 argument (id)
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			delete(db, id)
		},
	}

	// Get command
	var getCmd = &cobra.Command{
		Use:   "get [id]",
		Short: "Get an entry from the database",
		Args:  cobra.ExactArgs(1), // Require exactly 1 argument (id)
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			get(db, id)
		},
	}

	// Add commands to the root command
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(getCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
