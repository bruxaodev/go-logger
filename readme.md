# Download package
```sh
go get -v github.com/bruxaodev/go-logger
```

# Usage

## initialize
```js
logger := logger.new("main", logger.Green)
```
##


## Info
### Logs an info message.
```js
logger.Info("Info message")
```
### Logs a formatted info message.
```js
logger.Infof("Info message %s", "formatted")
```

#Complete Example
```js
package main

import (
    "github.com/bruxaodev/go-logger"
)

func main() {
    logger := new("main", logger.Green)
    logger.Debug("Debug message")
    logger.Debugf("Debug message %s", "formatted")
    logger.Info("Info message")
    logger.Infof("Info message %s", "formatted")
    logger.Warn("Warn message")
    logger.Warnf("Warn message %s", "formatted")
    logger.Error("Error message")
    logger.Errorf("Error message %s", "formatted")
}
```