# :notebook: We Log
This package provides a singleton of a [uber/zap](https://github.com/uber-go/zap) logger instance to be used on every application without any dependency and using all the appropriate default values.

## Usage

Just import the package, init the logger and start using. 

### Example

> :warning: It is very important that you initialize the logger before anything runs in your main function, so it can be used to log any further errors/info.

```go
package main

import (
	log "github.com/grupo-sbf/we-log"
)

func main() {
	log.InitLog()
	
	log.Log.Info("Test INFO")
	log.Log.Debug("Test DEBUG")
	log.Log.Error("Test ERROR")
}
```

## Environment Variables

- `LOG_LEVEL` - String of the log level ["debug", "info", "warn", "error"].
- `LOG_FORMAT` - String of the log format ["json", "console"].
- `SERVICE_NAME` - Name of the microservice that will be added to all logging events.

## Stackdriver support

We are using [zapdriver](https://github.com/blendle/zapdriver) as the base encoder of our zap instance, which supports the full array of structured logging capabilities of Stackdriver.
