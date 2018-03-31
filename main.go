// Package main is the impleùentation of the API
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/frodonLD/GoLangRESTAPIWithGin/model"
	"github.com/frodonLD/GoLangRESTAPIWithGin/router"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	model.Notifications = append(model.Notifications, model.Notification{ID: "1", Message: "Something is wrong in this world", Level: &model.NotificationLevel{Name: "Warning", Bloquant: false}})
	model.Notifications = append(model.Notifications, model.Notification{ID: "2", Message: "This is the end of the world", Level: &model.NotificationLevel{Name: "Critical", Bloquant: true}})

	app := cli.NewApp()
	app.Name = "Rest API Test With MUX"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{{Name: "frodonLD"}}
	app.Copyright = "frodonLD " + strconv.Itoa(time.Now().Year())

	app.Action = func(c *cli.Context) error {
		router.Start()
		return nil
	}

	// run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
