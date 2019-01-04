# GEDCOM 5.x library for Go

The goal of this library is to provide parsing and generation of GEDCOM files
as generated by, among others, the Brothers Keeper generalogical software.

## Usage

For detailed usage instructions, please take a look at [godoc.org][gdo].


```go
import "gitlab.com/zerok/go-gedcom5"
import "os"

func main() {
    var file gedcom5.File
    gedcom5.NewDecoder(os.Stdin).Decode(&file)
}
```

Internally, this library uses [zerolog][] for logging. If you want to inject a
zerolog Logger using a context, you can use the `WithContext(context.Context)`
method of the decoder:

```go
import "gitlab.com/zerok/go-gedcom5"
import "github.com/rs/zerolog"
import "os"

func main() {
    var file gedcom5.File
    logger := zerolog.New(zerlog.ConsoleWriter{Out: os.Stderr})
    dec := gedcom5.NewDecoder(os.Stdin)
    dec = dec.WithContext(logger.WithContext(context.Background()))
    dec.Decode(&file)
}
```


[gdo]: https://godoc.org/gitlab.com/zerok/go-gedcom5
[zerolog]: https://github.com/rs/zerolog
