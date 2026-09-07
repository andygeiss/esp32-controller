# ESP32 Controller

[![License](https://img.shields.io/github/license/andygeiss/esp32-controller)](https://github.com/andygeiss/esp32-controller/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/andygeiss/esp32-controller)](https://github.com/andygeiss/esp32-controller/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/andygeiss/esp32-controller.svg)](https://pkg.go.dev/github.com/andygeiss/esp32-controller)

Write an ESP32 program in Go, and let
[esp32-transpiler](https://github.com/andygeiss/esp32-transpiler) turn it into an Arduino
sketch. It is for anyone who would rather have a program type-checked and tested on a
computer than debug it down a serial cable.

Each package mirrors the Arduino calls a sketch makes, so the Go source and the sketch it
becomes read the same way.

## Install

```sh
go get github.com/andygeiss/esp32-controller
```

## Blink an LED

Implement `Controller`. `Setup` becomes the sketch's `setup()`, and `Loop` becomes
`loop()`, which the board calls over and over.

```go
package main

import (
	"github.com/andygeiss/esp32-controller"
	"github.com/andygeiss/esp32-controller/digital"
	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/timer"
)

const led = 2

type blinker struct{}

func (b *blinker) Setup() error {
	serial.Begin(serial.BaudRate115200)
	digital.PinMode(led, digital.ModeOutput)
	return nil
}

func (b *blinker) Loop() error {
	digital.Write(led, digital.High)
	timer.Delay(500)
	digital.Write(led, digital.Low)
	timer.Delay(500)
	return nil
}

func main() {
	var c esp32_controller.Controller = &blinker{}
	if err := c.Setup(); err != nil {
		return
	}
	for {
		if err := c.Loop(); err != nil {
			return
		}
	}
}
```

## What each package covers

| Package | Go call | Arduino call |
|---|---|---|
| `digital` | `PinMode` | [`pinMode`](https://www.arduino.cc/reference/en/language/functions/digital-io/pinmode/) |
| `digital` | `Write` | [`digitalWrite`](https://www.arduino.cc/reference/en/language/functions/digital-io/digitalwrite/) |
| `random` | `Num`, `NumBetween` | [`random`](https://www.arduino.cc/reference/en/language/functions/random-numbers/random/) |
| `random` | `Seed` | [`randomSeed`](https://www.arduino.cc/reference/en/language/functions/random-numbers/randomseed/) |
| `serial` | `Available` | [`Serial.available`](https://www.arduino.cc/reference/en/language/functions/communication/serial/available/) |
| `serial` | `Begin` | [`Serial.begin`](https://www.arduino.cc/reference/en/language/functions/communication/serial/begin/) |
| `serial` | `Print`, `Println` | [`Serial.print`](https://www.arduino.cc/reference/en/language/functions/communication/serial/print/) |
| `timer` | `Delay` | [`delay`](https://www.arduino.cc/reference/en/language/functions/time/delay/) |
| `wifi` | `Begin`, `BeginEncrypted` | [`WiFi.begin`](https://www.arduino.cc/en/Reference/WiFiBegin) |
| `wifi` | `BSSID`, `RSSI`, `SSID` | [`WiFi.BSSID`](https://www.arduino.cc/en/Reference/WiFiBSSID) and friends |
| `wifi` | `Disconnect` | [`WiFi.disconnect`](https://www.arduino.cc/en/Reference/WiFiDisconnect) |
| `wifi` | `EncryptionType` | [`WiFi.encryptionType`](https://www.arduino.cc/en/Reference/WiFiEncryptionType) |
| `wifi` | `LocalIP`, `SetDNS` | [`WiFi.localIP`](https://www.arduino.cc/en/Reference/WiFiLocalIP) |
| `wifi` | `ScanNetworks` | [`WiFi.scanNetworks`](https://www.arduino.cc/en/Reference/WiFiScanNetworks) |
| `wifi` | `Status` | [`WiFi.status`](https://www.arduino.cc/en/Reference/WiFiStatus) |
| `wifi` | `Client` | [`WiFiClient`](https://www.arduino.cc/en/Reference/WiFiClientConstructor) |

## Build and test

`make` is the only command surface. `make check` runs every gate against the working
tree; `make ci` runs the same gates against the commit, and is what has to be green
before a push.

## Baseline deviations

This module follows the [engineering baseline](https://github.com/andygeiss/baseline).
Four of its library rules collide with the job in [SPEC.md](SPEC.md), because the
exported names have to mirror Arduino's C API for the transpiler to match them. All four
are waived on the record:

- **No package-level mutable state**
  ([project-types/library.md](https://github.com/andygeiss/baseline/blob/main/project-types/library.md))
  — waived 2026-09-06 by Andy. An Arduino sketch keeps pin and network state in globals,
  and the transpiler maps these onto them. Scoped to the exported variables in `digital`
  (`GPIOModes`, `GPIOValues`), `serial` (`AvailableN`, `Baud`) and `wifi` (`Current*`);
  no other package holds state, and none of it is safe to touch from more than one
  goroutine, which the package docs say.

- **`context.Context` first on anything that blocks**
  ([stack/go.md](https://github.com/andygeiss/baseline/blob/main/stack/go.md))
  — waived 2026-09-06 by Andy. Arduino's `delay()` takes no context and a sketch has
  nothing to cancel it from. Scoped to `timer.Delay`, the only call in the module that
  blocks.

- **A library returns, it does not print**
  ([project-types/library.md](https://github.com/andygeiss/baseline/blob/main/project-types/library.md))
  — waived 2026-09-06 by Andy. Printing is what these calls are: they become
  `Serial.print` in the sketch. Scoped to `serial.Print`, `serial.Println`,
  `wifi.Client.Println` and `wifi.Client.Write`, which write to standard output and
  nowhere else. Nothing in the module logs.

- **The consumer declares the interface**
  ([patterns/go-ports-adapters.md](https://github.com/andygeiss/baseline/blob/main/patterns/go-ports-adapters.md))
  — waived 2026-09-06 by Andy. `Controller` is declared here rather than by the caller,
  because the transpiler has to recognise one agreed shape to emit `setup()` and
  `loop()` from. Scoped to that one two-method interface; adding a third method would be
  a breaking change, so it stays at two.
