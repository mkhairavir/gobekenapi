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
	FindEvent(id, id_user int) *Event
	Update(*Event) error
	SaveDet(*Detail) error
	EventDet(id int) []Detail
	History() []Event
	UserEvent(id int) []Event
	// SingleEvent(id int) *Event
	AllDet() []Detail
	DeleteEvent(event *Event) error
}
