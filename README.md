# Colin's Transfer Tool

This is a really basic tool to transfer firmware files to my retro computer systems over a serial port. This removes the need for flashing EEPROM chips to test every software change. Instead, I run a small bootloader on the EEPROM and load new firmware into RAM with this tool.

## Protocol

The "protocol" is trivially simple. The target system starts listening for serial port data and stores all incoming bytes into RAM until it sees three bytes in a row with the value `0xDE`. This tool then loads a file from disk and sends it to the listening target over the serial port.

There is no error checking or acknowledgments of any kind, but I've found this to be very reliable for transferring firmware.

There is one flaw: if the data being sent ever actually contains three bytes with value `0xDE` in a row, the transfer will stop prematurely. This has not yet been an issue, but it will require an update to the protocol if it ever comes up.

In short: this is a hack to speed up firmware development, not a robust file transfer solution.

## Usage

```
ctt -p <port> -f <file> -b [baudrate]
```

`ctt` requires a port and a filename be specified and optionally the baudrate (defaults to 115200)
