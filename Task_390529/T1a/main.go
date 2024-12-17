package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Page struct {
	Total int    `json:"total"`
	Items []Item `json:"items"`
}

func main() {
	http.HandleFunc("/items", paginatedItems)
	http.ListenAndServe(":8080", nil)
}

func paginatedItems(w http.ResponseWriter, r *http.Request) {
	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")

	var items []Item
	for i := 0; i < 100; i++ {
		items = append(items, Item{ID: i + 1, Name: fmt.Sprintf("Item %d", i+1)})
	}

	totalItems := len(items)
	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)
	if offsetInt < 0 {
		offsetInt = 0
	}
	if limitInt <= 0 {
		limitInt = 10
	}
	if offsetInt+limitInt > totalItems {
		limitInt = totalItems - offsetInt
	}

	pagedItems := items[offsetInt : offsetInt+limitInt]
	page := Page{Total: totalItems, Items: pagedItems}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(page)
}
