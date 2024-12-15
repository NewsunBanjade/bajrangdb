package database
import "bajrangdb/database/model"

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"path/filepath"
	
)

// Database holds the local in-memory and on-disk store
type Database struct {
	data   map[string]model.Entry
	dbFile string
	mutex  sync.RWMutex
}

// NewDatabase initializes a new database instance
func NewDatabase(dbPath string) (*Database, error) {
    // Ensure the parent directory exists
    dir := filepath.Dir(dbPath)
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        return nil, fmt.Errorf("error creating directory: %v", err)
    }

    db := &Database{
        data:   make(map[string]model.Entry),
        dbFile: dbPath,
    }

    if err := db.loadFromDisk(); err != nil {
        return nil, err
    }

    return db, nil
}


func (db *Database) Insert(id, data string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.data[id] = model.Entry{ID: id, Data: data}
	db.saveToDisk()
}

func (db *Database) Get(id string) (model.Entry, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	entry, exists := db.data[id]
	return entry, exists
}

func (db *Database) Delete(id string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	delete(db.data, id)
	db.saveToDisk()
}

func (db *Database) loadFromDisk() error {
	if _, err := os.Stat(db.dbFile); os.IsNotExist(err) {
		return nil
	}
	file, err := os.Open(db.dbFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&db.data)
}

func (db *Database) saveToDisk() {
	file, err := os.Create(db.dbFile)
	if err != nil {
		fmt.Println("Error saving to disk:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(db.data); err != nil {
		fmt.Println("Error encoding data:", err)
	}
}
