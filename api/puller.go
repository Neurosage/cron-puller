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
			fmt.Println(err_2.Error())
			return nil
		}
		username = u.CRONPULLER_USERNAME
		password = u.CRONPULLER_PASSWORD
	} else if os.IsNotExist(err) {
		// file does not exist
		fmt.Println("user.json does not exist")
		fmt.Println("Should it be created? (y/n)")
		var response string
		fmt.Scanln(&response)
		if response == "y" {
			fmt.Println("Please enter your Cronometer username/email: ")
			fmt.Scanln(&username)
			fmt.Println("Please enter your Cronometer password: ")
			fmt.Scanln(&password)

			var u = User{
				CRONPULLER_USERNAME: username,
				CRONPULLER_PASSWORD: password,
			}

			b, err_1 := json.Marshal(u)
			if err_1 != nil {
				fmt.Println(err_1.Error())
				return nil
			}

			err_2 := os.WriteFile(userfile, b, 0644)

			if err_2 != nil {
				fmt.Println(err_2.Error())
				return nil
			}
		} else {
			username = os.Getenv("CRONPULLER_USERNAME")
			password = os.Getenv("CRONPULLER_PASSWORD")

			if username == "" || password == "" {
				fmt.Println("No user.json and no environment variables set")
				fmt.Println("Please configure one of these options")
				return nil
			}
		}

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
	}

	rawBiometricsCSV, err := c.ExportBiometrics(context.Background(), start, end)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var filename = "../data/nutrition.csv"
	// check if data directory exists
	if _, err := os.Stat("../data"); os.IsNotExist(err) {
		fmt.Println("data directory does not exist, creating it")
		os.Mkdir("../data", 0755)
	}

	// list all files in data directory
	files, err := os.ReadDir("../data")
	if err != nil {
		fmt.Println(err.Error())
	}

	// delete any existing file
	for _, file := range files {
		os.Remove(fmt.Sprintf("../data/%s", file.Name()))
	}

	err = os.WriteFile(filename, []byte(rawCSV), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("../data/biometrics.csv", []byte(rawBiometricsCSV), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer c.Logout(context.Background())

}

func main() {

	var dFlag = flag.Int("d", 0, "Cronpuller will get the previous d days of data, including today.")

	flag.Parse()

	var c = login()
	if c == nil {
		return
	}

	get_raw_data(c, *dFlag)
}
