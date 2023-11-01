package main

import "time"

type Event struct {
	EventName string    `json:"event_name"`
	UserId    string    `json:"user_id"`
	Data      string    `json:"data"`
	Time      time.Time `json:"time"`
}
