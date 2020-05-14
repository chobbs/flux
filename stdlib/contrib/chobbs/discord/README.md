# Discord Package

Use this Flux Package to send an automated message to a Discord channel using a webhook.

### Basic Syntax

Here's the usage definition for the `discord.send()` function.

    // Custom discord function
    // `url` - string - URL of the discord webhook endpoint
    // `username` - string - Username posting the message.
    // `content` - string - The text to display in discord.

    import "discord"
    import "http"
    import "json"
    import "influxdata/influxdb/secrets"

    //username - overrides the current username of the webhook.
    //content - simple message, the message contains (up to 2000 characters)
    //webhook - generated on discord to post messages to a channel

    //this value can be stored in the secret-store()
    hook = secrets.get(key: "DISCORD_HOOK")
    
    lastReported =
      from(bucket: "example-bucket")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "statuses")
        |> last()
        |> tableFind(fn: (key) => exists key._level)

   discord.send = (
      url:hook,
      username:"chobbs",
      content:"Great Scott! -  Disk usages is at \"${lastReported.status}\"."."
      )


## Contact

Provide a way for users to get in touch with you if they have questions or need help using your template. What information you give is up to you, but we encourage providing those below.

- Author: Craig Hobbs
- Email: craig@influxdata.com
- Github: [@chobbs](https://github.com/chobbs)
- Influx Slack: [@craig](https://influxdata.com/slack)
