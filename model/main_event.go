package model

import (
	"database/sql"
	"log"
	// "fmt"
)

type MainEventStore struct {
	DB *sql.DB
}

func NewMainEvent() EventStore {
	// dsn := os.Getenv("DATABASE_USER") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true&clientFoundRows=true"
	dsn := "sql3339915:QIU6tupy3K@tcp(sql3.freemysqlhosting.net)/sql3339915?parseTime=true&clientFoundRows=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return &MainEventStore{DB: db}
}

func (store *MainEventStore) All() []Event {
	events := []Event{}
	rows, err := store.DB.Query(`SELECT * FROM main_event`)

	if err != nil {

		return events
	}

	event := Event{}
	// event.name = "hallo"
	for rows.Next() {
		rows.Scan(&event.Id, &event.Id_user, &event.Img, &event.Name, &event.EventType, &event.TanggalAwal, &event.Expire, &event.TotalDonasi, &event.Status)
		events = append(events, event)
	}

	return events
}

func (store *MainEventStore) AllDet(id int) []Detail {
	details := []Detail{}

	rows, err := store.DB.Query(`SELECT * FROM main_event_detail WHERE id_event = ?`, id)

	if err != nil {
		return details
	}

	detail := Detail{}

	for rows.Next() {
		rows.Scan(&detail.Id, &detail.Id_event, &detail.Donatur, &detail.Dana, &detail.Metode, &detail.Tgl)
		details = append(details, detail)
	}

	return details

}

func (store *MainEventStore) SaveDet(detail *Detail) error {

	result, err := store.DB.Exec(`
		INSERT INTO main_event_detail(id_event, donatur, dana_donasi, metode_donasi, tgl_donasi) VALUES(?,?,?,?,?)`,
		detail.Id_event,
		detail.Donatur,
		detail.Dana,
		detail.Metode,
		detail.Tgl,
	)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	detail.Id = int(lastID)

	return nil

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

func (store *MainEventStore) History() []Event {
	histories := []Event{}

	rows, err := store.DB.Query(`SELECT * from main_event where expire <= CURRENT_DATE()`)

	if err != nil {
		return histories
	}

	history := Event{}
	for rows.Next() {
		rows.Scan(&history.Id, &history.Id_user, &history.Img, &history.Name, &history.EventType, &history.TanggalAwal, &history.Expire, &history.TotalDonasi, &history.Status)
		histories = append(histories, history)
	}
	return histories
}
