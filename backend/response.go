package main

import (
	"fmt"
	"time"
)

type JSONTime time.Time

const JSONTimeLayout = "2006-01-02 15:04:05"

func (j JSONTime) MarshalJSON() ([]byte, error) {
	stime := fmt.Sprintf("\"%s\"", time.Time(j).Format(JSONTimeLayout))
	return []byte(stime), nil
}

type RecoutResponse struct {
	Message   string   `json:"message"`
	CreatedAt JSONTime `json:"created_at"`
}
