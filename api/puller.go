package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jrmycanady/gocronometer"
)

type User struct {
	CRONPULLER_USERNAME string
	CRONPULLER_PASSWORD string
}

func login() *gocronometer.Client {
	var c = gocronometer.NewClient(nil)

	// look for user.json
	var username = ""
	var password = ""
	var userfile = "../user.json"

	if _, err := os.Stat(userfile); err == nil {
		// file exists
		fmt.Println("user.json exists")
		var u User
		b, err_1 := os.ReadFile(userfile)
		if err_1 != nil {
			fmt.Println(err.Error())
			return nil
		}

		err_2 := json.Unmarshal(b, &u)
		if err_2 != nil {
			fmt.Println(err.Error())
			return nil
		}
		username = u.CRONPULLER_USERNAME
		password = u.CRONPULLER_PASSWORD
	} else if os.IsNotExist(err) {
		// file does not exist
		fmt.Println("user.json does not exist")
		username = os.Getenv("CRONPULLER_USERNAME")
		password = os.Getenv("CRONPULLER_PASSWORD")
	}

	var err = c.Login(context.Background(), username, password)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func get_raw_data(c *gocronometer.Client, d int) {
	var start = time.Now().Add(time.Hour * 24 * time.Duration(-d))
	var end = time.Now().Add(time.Hour * 24)
	rawCSV, err := c.ExportDailyNutrition(context.Background(), start, end)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rawCSV)
	var timestamp = time.Now().Format("20060102")
	var filename = fmt.Sprintf("../data/%s.csv", timestamp)
	// check if data directory exists
	if _, err := os.Stat("../data"); os.IsNotExist(err) {
		fmt.Println("data directory does not exist")
		os.Mkdir("../data", 0755)
	}

	err = os.WriteFile(filename, []byte(rawCSV), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer c.Logout(context.Background())

}

func main() {

	var dFlag = flag.Int("d", 1, "Cronpuller will get the previous d days of data, including today.")

	flag.Parse()

	var c = login()
	if c == nil {
		return
	}

	get_raw_data(c, *dFlag)
}
