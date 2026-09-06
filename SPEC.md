# SPEC

**Job:** Write an ESP32 program in Go — pins, serial, WiFi, timing — and let
[esp32-transpiler](https://github.com/andygeiss/esp32-transpiler) turn it into an
Arduino sketch.

**Why:** An Arduino sketch has no `go vet` and no `go test`. Writing the program in Go
means it type-checks and its logic is tested on a computer, before it ever reaches a
board.

**Guardrails:** The transpiler's `internal/transpile/handlers/mapping.go` is the
contract: every name it lists is matched to an Arduino call one for one, so renaming
or removing one breaks it. A name that file does not list is ordinary Go, free to
change. The deviations that shape forces are listed under *Baseline deviations* in
[README.md](README.md).

**Done means:**
[checklists/library.md](https://github.com/andygeiss/baseline/blob/main/checklists/library.md)
walked, and `make ci` green on the commit being pushed.
