Cron puller will look for a user.json file in the installation directory with the following format:

`
{
    "CRONPULLER_USERNAME": <username>
    "CRONPULLER_PASSWORD": <password>
}
`

If user.json does not exist, it will look for environment variables `CRONPULLER_USERNAME` and `CRONPULLER_PASSWORD`.