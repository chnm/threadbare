package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func (cr CooperItem) String() string {
	return cr.Objects[0].Title
}

func (cr CooperItem) Save() error {
	log.Info("Writing items to database...")

	query := `
	INSERT INTO connthreads.connthreads_items(id, title, date, description, type, medium, url, country)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT DO NOTHING;
	`

	api, err := json.Marshal(cr)
	if err != nil {
		log.Debug("Error marshalling API: ", err)
	}

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		return err
	}

	description := ""
	if len(cr.Objects[0].Description) > 0 {
		description = cr.Objects[0].Description
	}

	// set a timestamp for the item
	timestamp := "NOW()"

	_, err = stmt.Exec(cr.Objects[0].ID, cr.Objects[0].Title, description, cr.Objects[0].URL, timestamp, cr.Objects[0].Medium, cr.Objects[0].Date, cr.Objects[0].URL, cr.Objects[0].Country, cr.Objects[0].Type, api)
	if err != nil {
		return err
	}

	return nil

}
