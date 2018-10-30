## Instaget

Have you ever wanted to download an image you see on Instagram and,
lost your way in the page source? Now, you don't have to.

### Installation

```
$ go get -u github.com/cakturk/instaget
```

### Usage

Download a single image or a video

```sh
$ instaget https://www.instagram.com/p/BoaV78VjZGv/?taken-by=instagram
```

Download all images from a sidecar image a.k.a. carousel

```sh
$ instaget https://www.instagram.com/p/BoHk1haB5tM/?taken-by=instagram
```

Download posts made within specified time range. Note that the
time interval must be in reverse chronological order.


```sh
$ instaget -from="2014-09-24 3:23" --to="2013-12-09 3:23" https://instagram.com/instagram
```
