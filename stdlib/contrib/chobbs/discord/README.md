# Discord Package


Use this Flux Package to send a single message to a Discord channel using a webhook.

## Parameters

  - `botuser:` **_Data type:_** String - overrides the current username of the webhook.
  - `content:` **_Data type:_** String - simple message, the message contains (up to 2000 characters)
  - `webhook:` **_Data type:_** String - url generated on discord to post messages to a channel

## Basic Example

Here's an example definition for the `discord.send()` function.

    import "discord"
    import "influxdata/influxdb/secrets"

    //this value can be stored in the secret-store()
    hook = secrets.get(key: "DISCORD_HOOK")

    lastReported =
      from(bucket: "example-bucket")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "statuses")
        |> last()
        |> tableFind(fn: (key) => true)
        |> getRecord(idx: 0)

    discord.send(
      url:hook,
      botuser:"chobbs",
      content: "Great Scott!- Disk usage is: \"${lastReported.status}\"."
      )


## Contact

Provide a way for users to get in touch with you if they have questions or need help using your package. What information you give is up to you, but we encourage providing those below.

- Author: Craig Hobbs
- Email: craig@influxdata.com
- Github: [@chobbs](https://github.com/chobbs)
- Influx Slack: [@craig](https://influxdata.com/slack)
