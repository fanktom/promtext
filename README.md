# Promtext

A prometheus collector that collects all textfiles from a loal directory and serves them via HTTP

You can pass the _port_ and the _dir_ to serve:
```
promtext -port 9100 -dir aFolderFullOfTextfiles
```

If the folder contains 20 files with prometheus values, a call to the endpoint with any path will concat all values and offer them for scraping.

## Docker
Promtext uses the `go:onbuild` image. Use it with `system8/promtext`.
