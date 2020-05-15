package discord

import "http"
import "json"

// `url` - string - URL of the discord webhook endpoint
// `username` - string - Username posting the message.
// `content` - string - The text to display in discord.

send = (url, username, content) => {
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
