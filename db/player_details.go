package db

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/suryateja1698/golang-cassandra/models"
)

func AddPlayerDetails(ctx context.Context, b *models.PlayerDetails) (string, error) {

	if b.ID == "" {
		id := uuid.New()
		b.ID = id.String()
	}

	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	insertQuery := Session.Query(`INSERT INTO player_details player_details (id, first_name, last_name, age, country,
	                            position, joined_year, joined_month, created_at,
								updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		b.ID, b.FirstName, b.LastName, b.Age, b.Country, b.Position, b.JoinedYear, b.JoinedMonth, b.CreatedAt, b.UpdatedAt)

	if err := insertQuery.WithContext(ctx).Exec(); err != nil {
		log.Println("Error while inserting into database")
		return "", err
	}

	return b.ID, nil
}
