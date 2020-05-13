package model

type Event struct {
	Id          int
	Id_user     int
	Img         string
	Name        string
	EventType   string
	TotalDonasi float64
	TanggalAwal string
	Expire      string
	Status      string
}

func CreateEvent(img, name, eventType, tanggal, expire, status string, id_user int, totalDonasi float64) (*Event, error) {
	return &Event{
		// id:          id,
		Id_user:     id_user,
		Img:         img,
		Name:        name,
		EventType:   eventType,
		Status:      status,
		TanggalAwal: tanggal,
		Expire:      expire,
		TotalDonasi: totalDonasi,
	}, nil
}
