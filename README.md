# nekolog
相对于标准库的log，支持日志等级，支持json，plain text输出，支持颜色输出。

### Example
```
import log "github.com/neee333ko/nekolog"


func main(){
    logger := log.New(WithOutput(os.Stdout))

    logger.Info("This is an Info")
    logger.Warnf("WARN: %s","This is a Warn")
}
```
