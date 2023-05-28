A script to pull kcal consumption from [Cronometer](https://cronometer.com/) and sync it to a Google Sheets spreadsheet

Installation:

`git clone https://github.com/Neurosage/cron-puller.git`

`cd cron-puller`

`pip install -r requirements.txt`

Requirements: [golang](https://go.dev/dl/)

Usage: `python main.py`

Arguments: `-d`, `--days`: The number of previous days, starting from today (0), to pull data for (default 0)

Cron puller will look for a user.json file in the installation directory with the following format:


    {
        "CRONPULLER_USERNAME": <username>,
        "CRONPULLER_PASSWORD": <password>
    }
    
Where `<username>` and `<password>` are your Cronometer username and password.

If user.json does not exist, it will look for environment variables `CRONPULLER_USERNAME` and `CRONPULLER_PASSWORD`.

You need to follow the guide [here](https://docs.gspread.org/en/latest/oauth2.html) to authenticate with Google Sheets. Also note that the cells the code updates currently is just based on my personal spreadsheet so you will need to change it accordingly.
