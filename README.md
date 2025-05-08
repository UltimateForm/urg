# Urenge

urenge is a CLI tool for generating random tokens

## Root

```
Usage:
  urg [command]

Available Commands:
  help        Help about any command
  scr         scramble a string
  str         generate random string

Flags:ur
  -h, --help   help for urg

Use "urg [command] --help" for more information about a command.
```

## Strings

```
generate random string

Usage:
  urg str [flags]

Examples:
urg str --min=30 --max=50 -lnsu
urg str -un
urg str

Flags:
  -h, --help                  help for str
  -l, --lowers uint8[=255]    include lowercase letters
      --max uint8             max length of generated string (default 16)
      --min uint8             min length of generated string (default 14)
  -n, --numbers uint8[=255]   include numbers
  -s, --special uint8[=255]   include special characters
  -u, --uppers uint8[=255]    include uppercase letters (default 255)
```

## Scrambler

```
scramble a string

Usage:
  urg scr [flags]

Examples:
urg scr lorem ipsum

Flags:
  -h, --help   help for scr
```