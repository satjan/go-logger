## Initialize

logger.Init(fmt.Sprintf("%s/api-gateway.log", os.Getenv("LOG_PATH")))

## How to use

```
logger.Info("info message")
logger.Error("error message")
```