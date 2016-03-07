oEmbed
======
[![GoDoc](https://godoc.org/github.com/widnyana/oembed?status.svg)](https://godoc.org/github.com/widnyana/oembed)
[![Build Status](https://travis-ci.org/widnyana/oembed.svg)](https://travis-ci.org/widnyana/oembed)


golang [oEmbed](http://oembed.com/) library

    oEmbed is a format for allowing an embedded representation of a URL on third party sites.
    The simple API allows a website to display embedded content (such as photos or videos)
    when a user posts a link to that resource, without having to parse the resource directly.
    -- http://oembed.com/


Supported Providers
-------------------
- 23HQ
- Alpha App Net
- Animoto
- AudioSnaps
- BlipTV
- Cacoo
- ChartBlocks
- CircuitLab
- Clyp
- CollegeHumor
- Coub
- Crowd Ranking
- Daily Mile
- Dailymotion
- Deviantart.com
- Dipity
- Dotsub
- Embed Articles
- Embedly
- Flickr
- FunnyOrDie
- GMEP
- Geograph Britain and Ireland
- Geograph Channel Islands
- Geograph Germany
- Getty Images
- HuffDuffer
- Hulu
- IFTTT
- Infogram
- Instagram
- Kickstarter
- LearningApps.org
- Meetup
- MixCloud
- Moby Picture
- Official FM
- On Aol
- Ora TV
- Poll Daddy
- Poll Everywhere
- Portfolium
- Quiz.biz
- Quizz.biz
- RapidEngage
- Rdio
- ReleaseWire
- ReverbNation
- Roomshare
- Sapo Videos
- Screenr
- Scribd
- ShortNote
- Shoudio
- Sketchfab
- SlideShare
- SmugMug
- SoundCloud
- SpeakerDeck
- Ted
- They Said So
- Topy
- Twitter
- Ustream
- Viddler
- VideoFork
- VideoJug
- Vimeo
- WordPress.com
- YFrog
- YouTube
- amCharts Live Editor
- chirbit.com
- edocr
- iFixit
- iSnare Articles
- nfb.ca


Instalation
-----------
just like another golang library, after defining your `$GOPATH`, execute this command:

```
go get github.com/widnyana/oembed
```

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
