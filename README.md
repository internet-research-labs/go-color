# Cloud API for Sharing Random Colors :art:

Server that:
- Displays a message about wtf color api is
- Allows HTTP requests to request the current color
- Allows websocket connections which sends a JSON message

## API

* `GET` `http://localhost:8080/` Returns an HTML page displaying instructions.
* `GET` `http://localhost:8080/color` Returns a JSEND respond announcing the current color.
* `WS` `ws://localhost:8080/socket` Requests a socket connection to receive push-notifications when the color updates
