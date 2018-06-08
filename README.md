# hmi

## API

### LEDs on/off
```POST <beagle-bone-ip-addr>:9000/led/:led```

The beaglebone black has (4) user leds: { "usr0", "usr1", "usr2", "usr3" }

a json doc with state with the state must also be passed.  Possible states are: { "on", "off" }

example:


```
curl -X POST 192.168.0.12:9000/led/usr3 \
        -H "Content-Type: application/json" \
        -d '{"state":"on"}'
```

will turn `usr3` led on

## Build

```$ GOARM=7 GOARCH=arm GOOS=linux go build```