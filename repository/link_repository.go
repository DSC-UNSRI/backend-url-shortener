package repository

import (
	"database/sql"

	"github.com/DSC-UNSRI/backend-url-shortener/model"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) GetAllLinks() ([]model.Link, error) {
	rows, err := r.db.Query("SELECT * FROM links")
	if err != nil {
		return nil, err
	}
	var links []model.Link

	for rows.Next() {
		var link model.Link
		err := rows.Scan(&link.Id, &link.OriginLink, &link.ShortenedLink, &link.CreatedAt)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}
	return links, nil
}

func (r *LinkRepository) GetLinkByShortenedLink(shortenedLink string) (model.Link, error) {
	var link model.Link
	row := r.db.QueryRow("SELECT * FROM links WHERE shortened_link = ?", shortenedLink)
	err := row.Scan(&link.Id, &link.OriginLink, &link.ShortenedLink, &link.CreatedAt)
	if err != nil {
		return link, err
	}

	return link, nil
}
