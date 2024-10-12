# discord-logger-go

really really really simple logging util to send discord web hooks from a service. i always need a way to remotely get info from an app running on a server somewhere. discord makes for an easy, free way to get notified wherever i'm at.

Simply create a logger with your discord webhook:

``` go
dlog := NewDiscordLogger("https://mydiscordwebhookurl", true)
```

then log messages and errors where needed. the logging and error functions use `Printf` style args.
when these functions are called, they will also generate a `log.Printf` entry with the same contents so you don't have to double log.

``` go
dlog.Printf("Check out this variable: %v", myVar)

dlog.Errorf("Uh oh, we have an error: %v", err)
```
