[![Test](https://github.com/alekseinovikov/gredis/actions/workflows/test.yml/badge.svg)](https://github.com/alekseinovikov/gredis/actions/workflows/test.yml)

## Build

To build result binary executable application use the following commands:

```shell
#!/bin/sh

go get fyne.io/fyne/v2/cmd/fyne

go build
fyne package -icon logo.png

# result is a platform specific package
# for the current operating system.
```