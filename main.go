package main

import (
	// "fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/mkhairavir/gobekenapi/model"
)

func app(e *echo.Echo, store model.EventStore) {
	// untuk tampil event
	e.GET("/events", func(c echo.Context) error {
		// Process
		events := store.All()

		// Response
		return c.JSON(http.StatusOK, events)
	})

	// curl http://localhost:8080/articles/1
	e.GET("/events/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// untuk tampil event detail
	e.GET("/event/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		details := store.AllDet(id)
		return c.JSON(http.StatusOK, details)
	})

	// untuk tampil event buatan user
	e.GET("/user/events/", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		details := store.AllDet(id)
		return c.JSON(http.StatusOK, details)
	})

	e.GET("/history", func(c echo.Context) error {
		history := store.History()

		return c.JSON(http.StatusOK, history)
	})

	// untuk post event
	e.POST("/events", func(c echo.Context) error {
		layoutISO := "2006-01-02"
		tanggalan := time.Now()
		// tanggal := strconv.Itoa(tahun) + "-" + bulan.String() + "-" + strconv.Itoa(hari)

		img := c.FormValue("img")
		name := c.FormValue("name")
		eventType := c.FormValue("eventType")
		status := c.FormValue("status")
		idUser, _ := strconv.Atoi(c.FormValue("id_user"))
		totalDonasi, _ := strconv.ParseFloat(c.FormValue("totaldonasi"), 64)
		tanggal := tanggalan.Format(layoutISO)
		expire := tanggalan.AddDate(0, 1, 0).Format(layoutISO)

		event, _ := model.CreateEvent(img, name, eventType, tanggal, expire, status, idUser, totalDonasi)

		store.Save(event)

		return c.JSON(http.StatusOK, event)
	})

	// untuk pos donasi
	e.POST("/user/eventdet", func(c echo.Context) error {
		layoutISO := "2006-01-02"
		tanggalan := time.Now()

		id_event, _ := strconv.Atoi(c.FormValue("id_event"))
		donatur := c.FormValue("donatur")
		dana_donasi, _ := strconv.ParseFloat(c.FormValue("dana_donasi"), 64)
		metode_donasi := c.FormValue("metode_donasi")
		tgl_donasi := tanggalan.Format(layoutISO)

		detail, _ := model.CreateDetail(metode_donasi, tgl_donasi, donatur, id_event, dana_donasi)

		store.SaveDet(detail)

		return c.JSON(http.StatusOK, detail)
	})

	// untuk edit event berdasar id
	e.PUT("/events/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		event := store.Find(id)
		event.Img = c.FormValue("img")
		event.Name = c.FormValue("name")
		event.EventType = c.FormValue("event_type")

		store.Update(event)

		return c.JSON(http.StatusOK, event)
	})

}

func main() {
	// init data store
	store := model.NewMainEvent()

	// Create new instance echo framework
	e := echo.New()

	// our apps
	app(e, store)

	e.Logger.Fatal(e.Start(":8880"))
}
