# Discord Flux Package

This InfluxDB Package can be used to send messages to a Discod channel
via your webhook. Import the discord package.

### Basic Syntax

Here are a few examples of the language to get an idea of the syntax.

    // Several data types are built-in
    // Custom discord function
    // `url` - string - URL of the discord webhook endpoint
    // `username` - string - Username posting the message.
    // `content` - string - The text to display in discord.
    import "discord"
    import "http"
    import "json"

    discord.send = (url, username, content) => {
      data = {
          username: username,
          content: content
        }

      headers = {
          "Content-Type": "application/json"
        }
      encode = json.encode(v:data)
      return http.post(headers: headers, url: url, data: encode)
    }

## Contact

Provide a way for users to get in touch with you if they have questions or need help using your template. What information you give is up to you, but we encourage providing those below.

- Author: Craig Hobbs
- Email: craig@influxdata.com
- Github: [@chobbs](https://github.com/chobbs)
- Influx Slack: [@craig](https://influxdata.com/slack)
