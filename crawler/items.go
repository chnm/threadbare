package main

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

// CooperCollection describes the collected items from the Cooper Hewitt API.
type CooperItems struct {
	ID              int           `json:"id"`
	Title           string        `json:"title"`
	Description     []string      `json:"description"`
	URL             string        `json:"url"`
	Timestamp       time.Time     `json:"timestamp"`
	Medium          []string      `json:"medium"`
	Date            string        `json:"date"`
	AccessionNumber string        `json:"accession_number"`
	DepartmentID    string        `json:"department_id"`
	ImageURL        []string      `json:"image_url"`
	Country         string        `json:"country"`
	Type            string        `json:"type"`
	ItemsURL        string        `json:"items_url"`
	TitleRaw        string        `json:"title_raw"`
	Results         []CooperItems `json:"results"`
}

type ItemList struct {
	Results []CooperItems `json:"results"`
}

func (ci CooperItems) String() string {
	return ci.Title
}

func (ci CooperItems) Save() error {
	log.Info("Writing items to database.")

	query := `
	INSERT INTO conthreads_items(id, title, description, url, timestamp, medium, date, accession_number, department_id, image_url, country, type, items_url, title_raw)
	VALUES($1, $2, $3, $4, NOW(), $6, $7, $8, $9, $10, $11, $12, $13, $14)
	ON CONFLICT DO NOTHING;
	`

	api, err := json.Marshal(ci)
	if err != nil {
		log.Debug("Error marshalling API: ", err)
	}

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		return err
	}

	description := ""
	if len(ci.Description) > 0 {
		description = ci.Description[0]
	}

	_, err = stmt.Exec(ci.ID, ci.Title, description, ci.URL, ci.Timestamp, ci.Medium, ci.Date, ci.AccessionNumber, ci.DepartmentID, ci.ImageURL, ci.Country, ci.Type, ci.ItemsURL, ci.TitleRaw, api)
	if err != nil {
		return err
	}

	return nil

}
