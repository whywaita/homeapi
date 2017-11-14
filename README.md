# Home API

wrapper API for [whywaita](https://github.com/whywaita)'s house.

## Install

There are configuration(authentication, infrared light) for [IRKit](http://getirkit.com/) in `vars` package.

```
package vars

const (
        JsonHomeLight = `{"format":"raw","freq":38,"data":[...]}` # toggle home light signal
        JsonAirConOn  = `{"format":"raw","freq":38,"data":[...]}` # Power On air‐conditioner
        JsonAirConOff = `{"format":"raw","freq":38,"data":[...]}` # Power Off air‐conditioner
)

const (
        ClientKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" # your irkit client key
        DeviceID  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" # your irkit device id
)
```

client key and device id for IRKit can get [here](http://getirkit.com/) `IRKit Internet HTTP API`.

### Build from source code

```
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
$ go build .
```

## Usage

yayoi have two mode (server & CLI). default mode is server.

### Server Mode

```
$ ./yayoi -m server
$ ./yayoi # either possible
```

### CLI Mode

```
$ ./yayoi -m cli
```

if HomeLight down, you can type it.

```
$ ./yayoi -m cli
>>> irkit light off
HomeLight Off...
```
