# Action Stats Server

This is an example usage of the github.com/gscherer/action_stats library. 
It includes a web server which allows you to create and read time values for various actions (e.g. jump, walk, run)

# Installation

Install using the following commands:
```sh
go get github.com/gscherer/action_stats_server
go install github.com/gscherer/action_stats_server
```
This should fetch the code and install the binary in the `bin` folder configured for your $GOPATH. 

The default is `~/go/bin/action_stats_server`.

# Usage

Assuming the binary was installed in `~/go/bin/action_stats_server` you can start the server by running it, with an optional `-port` option.
This will start a web server on the specified port - the default is 8080.

## Examples

### Start the server

```sh
~/go/bin/action_stats_server -port 9000
```

### Creating Time Entries

As an example we can start by creating a "jump" action with a time of 80
```curl
curl -i -H "Content-Type: application/json" \
  -X POST -d '{"action": "jump", "time": 80}' \
  http://localhost:9000/action-stats
```
The server should respond with an array of all current actions and their time averages
```sh
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Date: Wed, 06 Jan 2021 17:04:16 GMT
Content-Length: 28

[{"action":"jump","avg":80}]
```
Creating additional entries will adjust the averages for each action.
```sh
curl -i -H "Content-Type: application/json" \
  -X POST -d '{"action": "jump", "time": 90}' \
  http://localhost:9000/action-stats
```
```sh
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Date: Wed, 06 Jan 2021 17:04:25 GMT
Content-Length: 28

[{"action":"jump","avg":85}]
```

### Get Stats

You can also read the current averages for each action by sending a GET request
```sh
curl -i http://localhost:9000/action-stats
```sh
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 06 Jan 2021 17:10:56 GMT
Content-Length: 28

[{"action":"jump","avg":85}]
```

### Errors

Sending malformed JSON or invalid/empty action and time entries will result in responses with a HTTP status of 400 and a JSON encoded error body.

```sh
curl -i -H "Content-Type: application/json" \
  -X POST -d '{"action": "jump", "time": ""}' \
  http://localhost:9000/action-stats
```
```sh
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Wed, 06 Jan 2021 17:13:36 GMT
Content-Length: 58

{"error":"Bad Request","message":"Invalid action or time"}
```
