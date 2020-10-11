# revregexp

[![GoDoc](https://godoc.org/github.com/pipe01/revregexp?status.svg)](https://godoc.org/github.com/pipe01/revregexp)
[![Go Report Card](https://goreportcard.com/badge/github.com/pipe01/revregexp)](https://goreportcard.com/report/github.com/pipe01/revregexp)

Small library for randomly generating strings from a regex pattern.

## Examples

| Regex pattern             | Sample output |
|---------------------------|---------------|
| `.*`                      | hM7R          |
| `\d+`                     | 415           |
| `[a-d4-8]{5}`             | a8c6c         |
| `[a-f0-9]{4}-[a-f0-9]{6}` | 893a-9034d1   |

## Sample usage

```go
import "github.com/pipe01/revregexp

str := revregexp.Reverse("Hello (World|Everyone)")

fmt.Println(str)
```
