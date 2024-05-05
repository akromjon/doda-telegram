**DODA TELEGRAM CLI** - Send Message and Files Telegram Using Telegram Bot on Servers, VMs, local.

**USAGE**: <br>
`doda-cli message <message>`- this sends message.<br>
`doda-cli file <file path>` - send local, files on servers.

**Installation**<br>
```
curl -sSL https://raw.githubusercontent.com/akromjon/doda-telegram/main/install.sh | bash
```

`.env` - the file needs to be created and moved `/usr/local/bin` directory with the following variables:<br>
```
BOT_TOKEN="bot_token"
CHAT_ID="chat_id"
```
