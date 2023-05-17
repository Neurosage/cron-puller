Cron puller will look for a user.json file in the installation directory with the following format:


    {
        "CRONPULLER_USERNAME": <username>,
        "CRONPULLER_PASSWORD": <password>
    }

If user.json does not exist, it will look for environment variables `CRONPULLER_USERNAME` and `CRONPULLER_PASSWORD`.

You need to follow the guide [here](https://docs.gspread.org/en/latest/oauth2.html) to authenticate with Google Sheets. Also note that the cells the code updates currently is just based on my personal spreadsheet so you will need to change it accordingly.