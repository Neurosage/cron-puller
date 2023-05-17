package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jrmycanady/gocronometer"
)

func login() *gocronometer.Client {
	var c = gocronometer.NewClient(nil)
	var username = os.Getenv("CRONPULLER_USERNAME")
	var password = os.Getenv("CRONPULLER_PASSWORD")
	var err = c.Login(context.Background(), username, password)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func get_raw_data(c *gocronometer.Client) {
	rawCSV, err := c.ExportDailyNutrition(context.Background(), time.Now(), time.Now().Add(time.Hour*24))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rawCSV)

}

func main() {
	var c = login()
	if c == nil {
		return
	}

	get_raw_data(c)
}
