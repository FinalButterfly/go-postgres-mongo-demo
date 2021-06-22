package storage

// Post - публикация.
type Post struct {
	Id          interface{} `json:"_id" bson:"_id"`
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error)                 // получение всех публикаций
	AddPost(post Post) (interface{}, error) // создание новой публикации
	UpdatePost(Post) error                  // обновление публикации
	DeletePost(id interface{}) error        // удаление публикации по ID
}
