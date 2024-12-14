package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type User struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Users struct {
	Users []User `json:"users" xml:"users>user"`
}

func parseJSON(filename string) (*Users, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users Users
	err = json.Unmarshal(data, &users)
	return &users, err
}

func parseXML(filename string) (*Users, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users Users
	err = xml.Unmarshal(data, &users)
	return &users, err
}

func TestParseJSON(t *testing.T) {
	_, err := parseJSON("data.json")
	if err != nil {
		t.Fatalf("Error parsing JSON: %v", err)
	}
}

func TestParseXML(t *testing.T) {
	_, err := parseXML("data.xml")
	if err != nil {
		t.Fatalf("Error parsing XML: %v", err)
	}
}
