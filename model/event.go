package model

type Event struct {
	Id             int
	Id_user        int
	Img            string
	JudulEvent     string
	DeskripsiEvent string
	EventType      string
	TotalDonasi    float64
	TanggalAwal    string
	Expire         string
	Status         string
}

type Detail struct {
	Id       int
	Id_event int
	Donatur  string
	Dana     float64
	Metode   string
	Tgl      string
}

func CreateDetail(metode, tgl, donatur string, id_event int, dana float64) (*Detail, error) {
	return &Detail{
		Id_event: id_event,
		Donatur:  donatur,
		Dana:     dana,
		Metode:   metode,
		Tgl:      tgl,
	}, nil
}

func CreateEvent(img, judul, deksripsi, eventType, tanggal, expire, status string, id_user int, totalDonasi float64) (*Event, error) {
	return &Event{
		// id:          id,
		Id_user:        id_user,
		Img:            img,
		JudulEvent:     judul,
		DeskripsiEvent: deksripsi,
		EventType:      eventType,
		Status:         status,
		TanggalAwal:    tanggal,
		Expire:         expire,
		TotalDonasi:    totalDonasi,
	}, nil
}
