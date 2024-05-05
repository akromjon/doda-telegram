**DODA TELEGRAM CLI** - Send Message and Files Telegram Using Telegram Bot on Servers, VMs, local.

**USAGE**:
`doda-cli message <message>`- this sends message.
`doda-cli file <file path>` - send local, files on servers.

**Installation**
`curl -sSL https://raw.githubusercontent.com/akromjon/doda-telegram/main/install.sh | bash`

`.env` - the file needs to be created and moved `/usr/local/bin` directory with following variables:
`BOT_TOKEN="bot_token"`
`CHAT_ID="chat_id"`