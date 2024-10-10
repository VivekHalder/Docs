package services

import (
	"log"

	"github.com/VivekHalder/TryingDocs/database"
	"github.com/VivekHalder/TryingDocs/models"
)

func SaveDocument(doc models.Document) error {
	_, err := database.DB.Exec("INSERT INTO documents (id, content) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET content=$2", doc.ID, doc.Content)

	if err != nil {
		log.Println("Error saving the document: ", err)
		return err
	}

	return nil
}
