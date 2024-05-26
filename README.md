# smtp2telegram
This is an SMTP Server for sending emails into Telegram text messages using Go which accepts authentication with any username and password.

First, you must have a Telegram Chat bot and get a Telegram token from BotFather. Search the internet for more details on how to get it.

After getting the Telegram token, use the configuration in the .env file with several configurations as below:

```
TELEGRAM_TOKEN=xxxxxxxxx:xxxxxxxxxxxxxxxxxxxxxxxxxxx
SMTP_LISTEN_ADDRESS=0.0.0.0:25
EMAIL_DOMAIN_TELEGRAM=ksatele.gram
```

## How to use
Send an email with the destination \<idchat\>@ksatele.gram to this SMTP server.

To get the target IDChat, the target must first get an IDChat, one way is that the target must chat via Telegram first with IDBot, and there you will get a chat ID.

# License
MIT License

Free Software, Yeah!

By Kusumah Sasmita