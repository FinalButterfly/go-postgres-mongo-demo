package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// Конструктор объекта хранилища.
func New(connectionString string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			author_id,
			title,
			content,
			created_at
		FROM posts
		ORDER BY id;
	`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.Id,
			&t.AuthorID,
			&t.Title,
			&t.Content,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)

	}
	return posts, rows.Err()
}

func (s *Store) AddPost(post storage.Post) (interface{}, error) {
	var id int
	var err error
	err = s.db.QueryRow(context.Background(), `
		INSERT into posts(
			author_id,
			title,
			content,
			created_at,
			published_at
		)
		values(
			$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id;
	`,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
		post.PublishedAt,
	).Scan(&id)
	return id, err
}

func (s *Store) UpdatePost(post storage.Post) error {
	var err error
	_, err = s.db.Exec(context.Background(), `
		UPDATE posts
		set (
			author_id,
			title,
			content,
			created_at
		)
		= 
		(
			$1,
			$2,
			$3,
			$4
		)
		where id = $5
	`,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
		post.Id,
	)
	return err
}
func (s *Store) DeletePost(id int) error {
	var err error
	_, err = s.db.Exec(context.Background(),
		`
		DELETE FROM posts
		where id = $1
		`,
		id,
	)
	return err
}
