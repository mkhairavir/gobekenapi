package main

import (
	// "fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// untuk tampil spesifik event
	e.GET("/events/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		event := store.Find(id)

		// Response
		return c.JSON(http.StatusOK, event)
	})

	e.GET("/detail/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		detail := store.FindDet(id)

		// Response
		return c.JSON(http.StatusOK, detail)
	})

	// untuk tampil event detail
	e.GET("/event/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		details := store.EventDet(id)
		return c.JSON(http.StatusOK, details)
	})

	//untuk tampil semua detail
	e.GET("/detail", func(c echo.Context) error {
		// id, _ := strconv.Atoi(c.Param("id"))

		details := store.AllDet()
		return c.JSON(http.StatusOK, details)
	})

	// untuk tampil event buatan user
	e.GET("/user/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		events := store.UserEvent(id)
		return c.JSON(http.StatusOK, events)
	})

	//untuk tampil history
	e.GET("/histories", func(c echo.Context) error {
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
		deskripsi := c.FormValue("deskripsi")
		eventType := c.FormValue("event_type")
		status := c.FormValue("status")
		idUser, _ := strconv.Atoi(c.FormValue("id_user"))
		targetDonasi, _ := strconv.ParseFloat(c.FormValue("target_donasi"), 64)
		totalDonasi := float64(0)
		tanggal := tanggalan.Format(layoutISO)
		expire := tanggalan.AddDate(0, 1, 0).Format(layoutISO)

		event, _ := model.CreateEvent(img, name, deskripsi, eventType, tanggal, expire, status, idUser, targetDonasi, totalDonasi)

		store.Save(event)

		return c.JSON(http.StatusOK, event)
	})

	// untuk pos donasi
	e.POST("/donasi", func(c echo.Context) error {
		layoutISO := "2006-01-02"
		tanggalan := time.Now()

		id_event, _ := strconv.Atoi(c.FormValue("id_event"))
		donatur := c.FormValue("donatur")
		dana_donasi, _ := strconv.ParseFloat(c.FormValue("dana_donasi"), 64)
		metode_donasi := c.FormValue("metode_donasi")
		tgl_donasi := tanggalan.Format(layoutISO)
		status := "Hutang"

		detail, _ := model.CreateDetail(metode_donasi, tgl_donasi, donatur, status, id_event, dana_donasi)

		store.SaveDet(detail)

		return c.JSON(http.StatusOK, detail)
	})

	// untuk edit event berdasar id
	e.PUT("/user/:id_user/event/:id", func(c echo.Context) error {
		// e.PUT("/events/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		id_user, _ := strconv.Atoi(c.Param("id_user"))

		event := store.FindEvent(id, id_user)
		event.Img = c.FormValue("img")
		event.JudulEvent = c.FormValue("name")
		event.DeskripsiEvent = c.FormValue("deskripsi")
		event.EventType = c.FormValue("event_type")

		store.Update(event)

		return c.JSON(http.StatusOK, event)
	})

	e.PUT("/detail/:id", func(c echo.Context) error {
		// e.PUT("/events/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		detail := store.FindDet(id)

		detail.Status = c.FormValue("status")

		store.UpdateDet(detail)

		return c.JSON(http.StatusOK, detail)
	})

	// untuk delete event
	e.DELETE("/event/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		event := store.Find(id)

		store.DeleteEvent(event)

		return c.JSON(http.StatusOK, event)
	})

}

func main() {

	godotenv.Load()
	// init data store
	store := model.NewMainEvent()

	// Create new instance echo framework
	e := echo.New()
	e.Use(middleware.CORS())
	// our apps
	app(e, store)

	// e.Logger.Fatal(e.Start(":8880"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
