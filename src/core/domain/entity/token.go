package entity

import "time"

type Token struct {
	AccessToken string
	ExpiresIn   time.Time
}
