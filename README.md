# Discord-Webhook-Queue 
###### _by [Prizzle#4655](https://twitter.com/bypassedpx)_

**Never loose a (Checkout-)Webhook again!**

This is a simple queue system for Discord webhooks. It is designed to be used with sneaker bots, but can be used for anything.

The program will queue your webhooks and sends them when the rate limit is over. 

## How to setup

1. Download the latest release for your plattform from [here](https://github.com/Prizzledizle/discord-webhook-queue/releases/)
2. Extract the zip file
3. Fill out the settings.json file
    
    this is an example:
   ```json
     {
          "webhooks": [
               {
                    "webhook": "https://discord.com/api/webhooks/123456789/abc",
                    "alias": "MyWebhook"
               },
               {
                    "webhook": "https://discord.com/api/webhooks/123456789/abc",
                    "alias": "MyWebhook2"
               }
          ],
          "port": 8080
     }
   ```
4. Run the exectuable
5. Put in the `http://localhost:PORT/ALIAS` (printed in the CLI) as webhook url in your bot
6. Make sure to test once! :heart:

### Settings.json explanation

| Key | Explanation | Type |
| --- | --- | --- |
| webhook | The discord webhook you want to use | String |
| port | The port the server will run on | Integer |
| alias | The alias of the webhook, make sure not use spaces | String |

## Got any questions?
Feel free to DM me on Discord (Prizzle#4655), on [Twitter](https://twitter.com/bypassedpx) or open an issue on this repo.

Also thanks to [Jonah](https://twitter.com/jonahxyz) for helping me fixing a little error in the code.