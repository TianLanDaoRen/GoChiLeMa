package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type FoodItem struct {
	Id   int    `json:"id"`
	Date string `json:"date"`
	Food string `json:"food"`
}

type FoodTable struct {
	Items []FoodItem `json:"items"`
	path  string
	mutex sync.Mutex
}

func NewFoodTable(path string) *FoodTable {
	// Check if path of json file exists
	if info, err := os.Stat(path); err != nil || info.IsDir() {
		// Return empty FoodTable
		return &FoodTable{
			path: path,
		}
	}
	// Read json file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Read bytes from file
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal json file to FoodTable
	table := FoodTable{
		path: path,
	}
	err = json.Unmarshal(bytes, &table)
	if err != nil {
		fmt.Println(err)
	}
	// Return FoodTable
	return &table
}

func (table *FoodTable) Save() {
	go (func() {
		// Lock mutex
		table.mutex.Lock()
		defer table.mutex.Unlock()
		// Marshal FoodTable to json bytes
		jsonBytes, err := json.Marshal(table)
		if err != nil {
			fmt.Println(err)
		}
		// Write json bytes to path
		fo, err := os.Create(table.path)
		if err != nil {
			fmt.Println(err)
		}
		defer fo.Close()
		// Write bytes to file
		_, err = fo.Write(jsonBytes)
		if err != nil {
			fmt.Println(err)
		}
	})()
}

func (table *FoodTable) GetNextId() int {
	// Lock mutex
	table.mutex.Lock()
	defer table.mutex.Unlock()
	// Return next id by timestamp
	timestamp := time.Now().Unix()
	return int(timestamp)
}

func (table *FoodTable) AppendItem(item FoodItem) {
	// Lock mutex
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.Items = append(table.Items, item)
}

func (table *FoodTable) RemoveItemById(id int) {
	// Lock mutex
	table.mutex.Lock()
	defer table.mutex.Unlock()
	for i, item := range table.Items {
		if item.Id == id {
			table.Items = append(table.Items[:i], table.Items[i+1:]...)
			break
		}
	}
}

func (table *FoodTable) RemoveItemsByIds(id []int) {
	for _, i := range id {
		table.RemoveItemById(i)
	}
}
