package main

import (
	log "github.com/sirupsen/logrus"
)

// ProcessData uses a channel of items from fetch.go and
// saves items to the database.
func ProcessData(cp chan CollectionPagination) {
	for r := range cp {
		go func(i CollectionPagination) {
			for _, item := range i.Results {
				err := item.Save()
				if err != nil {
					log.WithFields(log.Fields{
						"item": item,
						"err":  err,
					}).Error("Error saving item: ", err)
				}
			}
		}(r)
	}
}
