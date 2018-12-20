## Astroweather

There's a solar system with some weird coupling between it's planets relative location in space and the meteorological effects of their atmospheres.
This is a REST service to tell's them how the weather is gonna be a given day.
Except for rain, we haven't cracked that yet; so you might want to take an umbrella with you every time you leave the house.

Run:

```
prompt$ cd github.com/
prompt$ go get github.com/gorilla/mux github.com/thelholtz/except
prompt$ go install
prompt$ $GOPATH/bin/astroweather
```

To start serving on :8080.
Happy forecasting!