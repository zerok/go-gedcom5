# GEDCOM 5.x library for Go

**Note:** This is highly experimental and pre-alpha. Use at your own risk!

The goal of this library is to provide parsing and generation of GEDCOM files
as generated by, among others, the [Brothers Keeper][bk] generalogical
software.

My primary motivation for starting this project was that my family has a family
tree stored inside Brothers Keeper and I somehow wanted to work with the data
stored in it. While the ultimate goal here is to provide complete support for
GEDCOM 5.5, getting it to work with an export generated by Brothers Keeper is,
for now, the focus.

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

## Serialization

There also exist an encoder that can be used for serializing a `*gedcom5.File`
into a collection of lines using the `LineWriter` interface. The implementation
there has only a low priority, though, for now.


[gdo]: https://godoc.org/gitlab.com/zerok/go-gedcom5
[zerolog]: https://github.com/rs/zerolog
[bk]: https://bkwin.org/
