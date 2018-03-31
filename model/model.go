package model

import "time"

// EHR : Electronic Health Record
type EHR struct {
	ID                string        `json:"id"`
	Firstname         string        `json:"firstname"`
	Lastname          string        `json:"lastname"`
	SocialSecurityNum string        `json:"socialSecurityNum"`
	Birthday          time.Time     `json:"birthday"`
	Appointments      []Appointment `json:"visits"`
}

// Appointment public for access outside the CC
type Appointment struct {
	DrID    string    `json:"drId"`
	Date    time.Time `json:"date"`
	Comment string    `json:"comment"`
}
