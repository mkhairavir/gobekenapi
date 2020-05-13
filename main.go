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
	// curl http://localhost:8080/articles
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

	e.POST("/events", func(c echo.Context) error {
		layoutUS := "January 2, 2006"
		tanggalan := time.Now()
		// tanggal := strconv.Itoa(tahun) + "-" + bulan.String() + "-" + strconv.Itoa(hari)

		img := c.FormValue("img")
		name := c.FormValue("name")
		eventType := c.FormValue("eventType")
		status := c.FormValue("status")
		idUser, _ := strconv.Atoi(c.FormValue("id_user"))
		totalDonasi, _ := strconv.ParseFloat(c.FormValue("totaldonasi"), 64)
		tanggal := tanggalan.Format(layoutUS)
		expire := tanggalan.AddDate(0, 1, 0).Format(layoutUS)

		event, _ := model.CreateEvent(img, name, eventType, tanggal, expire, status, idUser, totalDonasi)

		store.Save(event)

		return c.JSON(http.StatusOK, event)
	})

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
