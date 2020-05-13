package model

// type ArticleStore interface {
// 	All() []Article
// 	Save(*Article) error
// 	Find(int) *Article
// 	Update(*Article) error
// 	Delete(article *Article) error
// }

type EventStore interface {
	All() []Event
	Save(*Event) error
	Find(int) *Event
	Update(*Event) error
	// Delete(article *Event) error
}
