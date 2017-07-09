scorodesk
==

Scorodesk exposes a webhook for zendesk to hit with ticket information. This information is parled into something scoro can understand.

The MVP of this software will:

 1. Accept a 'new ticket' with some data
 1. Create a scoro project with this information


building
--

```bash
$ make
go get -v
github.com/jspc/scorodesk
go build -ldflags "-X 'main.buildNumber=1' \
                           -X 'main.buildMachine=grootdeth.local' \
                           -X 'main.builtBy=jspc' \
                           -X 'main.builtWhen=1499613704' \
                           -X 'main.compiler=go version go1.8.3 darwin/amd64' \
                           -X 'main.sha='" \
                           -o 'scorodesk'
```

This API is se4 compatibile.


running
--

```bash
$ ./scorodesk -help
Usage of ./scorodesk:
  -c string
        config file (default "testdata/config.json")
  -httptest.serve string
        if non-empty, httptest.NewServer serves on this address and blocks
```


configuration
--

scorodesk takes a json file holdign configuration:

```json
{
  "listen_address": "localhost:8080"
}
```
