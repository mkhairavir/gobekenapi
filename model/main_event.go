package model

import (
	"database/sql"
	"log"
	"os"
	// "fmt"
)

type MainEventStore struct {
	DB *sql.DB
}

func NewMainEvent() EventStore {
	dsn := os.Getenv("DATABASE_USER") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true&clientFoundRows=true"
	// "root:@tcp(localhost:3306)/db_charty?parseTime=true&clientFoundRows=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return &MainEventStore{DB: db}
}

func (store *MainEventStore) All() []Event {
	events := []Event{}
	rows, err := store.DB.Query(`SELECT * FROM test`)

	if err != nil {

		return events
	}

	event := Event{}
	// event.name = "hallo"
	for rows.Next() {
		rows.Scan(&event.Id, &event.Img, &event.Name, &event.EventType, &event.Id_user, &event.TotalDonasi, &event.Status)
		events = append(events, event)
	}

	return events
}

func (store *MainEventStore) Save(event *Event) error {
	// tahun, bulan, hari := time.Now().Date()
	// tanggal := strconv.Itoa(tahun) + "-" + bulan.String() + "-" + strconv.Itoa(hari)

	result, err := store.DB.Exec(`
		INSERT INTO main_event(img, name, event_type, id_user, total_donasi,tgl, status) VALUES(?,?,?,?,?,?,?)`,
		event.Img,
		event.Name,
		event.EventType,
		event.Id_user,
		event.TotalDonasi,
		event.TanggalAwal,
		event.Status,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	// update article.ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.Id = int(lastID)

	return nil
}

func (store *MainEventStore) Find(id int) *Event {
	event := Event{}

	err := store.DB.
		QueryRow(`SELECT * FROM articles WHERE id=?`, id).
		Scan(
			&event.Id,
			&event.Id_user,
			&event.Img,
			&event.Name,
			&event.EventType,
			&event.TotalDonasi,
			&event.TanggalAwal,
			&event.Expire,
			&event.Status,
		)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &event
}

func (store *MainEventStore) Update(event *Event) error {
	result, err := store.DB.Exec(`
		UPDATE articles SET img = ?, name = ?, event_type = ? WHERE id = ?`,
		event.Img,
		event.Name,
		event.EventType,
		event.Id,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
