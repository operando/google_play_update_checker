# TONKOTSU

Google Play Application Update Checker in Go


## Run

```bash
go run googleplay_update_checker.go config.go -c config.toml
```

## Config File

```toml
log = "debug"
package = "com.mercariapp.mercari" # your Android application package name
sleeptime = 1

[slack]
text = "TONKOTSU TEST"
username = "TONKOTSU bot"
icon_emoji = ":pig:"
channel = "#test"

[webhook]
url = "webhook_url" # your Incoming WebHooks URL for Slack

```
