package main

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

type User struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Users struct {
	Users []User `json:"users" xml:"user"`
}

func parseJSONFromString(data string) (*Users, error) {
	var users Users
	err := json.Unmarshal([]byte(data), &users)
	return &users, err
}

func parseXMLFromString(data string) (*Users, error) {
	var users Users
	err := xml.Unmarshal([]byte(data), &users)
	return &users, err
}

func TestParseJSON(t *testing.T) {
	mockJSON := `{"users":[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]}`
	users, err := parseJSONFromString(mockJSON)
	if err != nil {
		t.Fatalf("Error parsing JSON: %v", err)
	}
	if len(users.Users) != 2 {
		t.Fatalf("Expected 2 users, got %d", len(users.Users))
	}
}

func TestParseXML(t *testing.T) {
	mockXML := `<users><user><id>1</id><name>Alice</name></user><user><id>2</id><name>Bob</name></user></users>`
	users, err := parseXMLFromString(mockXML)
	if err != nil {
		t.Fatalf("Error parsing XML: %v", err)
	}
	if len(users.Users) != 2 {
		t.Fatalf("Expected 2 users, got %d", len(users.Users))
	}
}
