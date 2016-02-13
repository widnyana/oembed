oEmbed
======
golang [oEmbed](http://oembed.com/) library

    oEmbed is a format for allowing an embedded representation of a URL on third party sites. 
    The simple API allows a website to display embedded content (such as photos or videos) 
    when a user posts a link to that resource, without having to parse the resource directly.
    -- http://oembed.com/

Instalation
-----------
just like another golang library, after defining your `$GOPATH`, execute this command:

```go get github.com/widnyana/oembed```

Usage
-----

oembed came with providers rule called `providers.json`, this module rely on that list of provider, 
and you can add more providers as you need. kudos to Vitaly Dyatlov for this great design!

this is a basic example on how to use this library:

```
import (
    "bytes"
    "io/ioutil"
    "github.com/widnyana/oembed"
)

func main() {
    rule, err := ioutil.ReadFile("/path/to/provider.json")

    if err != nil {
        panic(err)
    }

    svc := oembed.NewOembed()
    svc.ParseProviders(bytes.NewReader(rule))

    item := svc.FindItem(url)
    if item != nil {
        info, err := item.FetchOembed(url, nil)
        if err != nil {
            fmt.Printf("An error occured: %s\n", err.Error()
        } else {
            if info.Status >= 300 {
                fmt.Printf("Response status code is: %d\n", info.Status)
            } else {
                fmt.Printf("Oembed info:\n%s\n", info)
            }
        }
    } else {
        fmt.Printf("oEmbed for URL: %s Not found", url)
    }
}
```


Credits
-------
this piece of code is just a small improvement of [dyatlov/go-oembed](https://github.com/dyatlov/go-oembed).

License
-------
[MIT License](http://widnyana.mit-license.org)