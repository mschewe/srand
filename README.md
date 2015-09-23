# srand

srand contains utility and convenience functions for working with secure randoms.

[![BuildStatus](https://travis-ci.org/toashd/srand.svg)](https://travis-ci.org/toashd/srand)

### Description
srand helps generating secure random numbers, non-consecutive and non-predictable keys (e.g. session keys, CSRF, and HMACS) using crypto/rand and cryptographically secure pseudo-random number generator (CSPRNG).

### API
``` go
func GenerateRandomBytes(n int) ([]byte, error)
```

Generates raw bytes for another cryptographic function, such as HMAC keys.

``` go
func GenerateRandomString(s int) (string, error)
```

Generates keys for session IDs, CSRF tokens etc. It is base64 URL encoded, providing secure strings that can be used in file-names, templates, HTTP headers and to minimize encoding overhead compared to hex encoding.

``` go
func GenerateRandomInt(n int64) (int64, error)
```

Generates a securely generated random number output between 0 and n.

``` go
func GenerateRandomIntRange(m, n int64) (int64, error)
```

Generates a securely generated random number output in range [m, n).

``` go
func GenerateRandomFloat() (float64, error)
```

Generates a securely generated random floating point number output in range [0, 1).

``` go
func GenerateRandomFloatRange(m, n float64) (float64, error)
```

Generates a securely generated random floating point number output in range [m, n).

``` go
func GenerateUUID(n int64) (*big.Int, error)
```

Generates a securely generated UUID compliant to RFC 4122.

### Getting Started

1: Download the package

```bash
go get github.com/toashd/srand
```

2: Import srand to your Go project

```go
import "github.com/toashd/srand"
```

### Example

``` go
package main

import (
  "fmt"
  "github.com/toashd/srand"
)

func main() {
	// Example: generate a 44 byte, base64 encoded, url-safe token.
	token, err := srand.GenerateRandomString(32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", token)

	// Example: generate a random number output between [0, 100).
	rand, err := GenerateRandomInt(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", rand)

	// Example: generate a random floating point number output between [0, 1).
	rand, err := GenerateRandomFloat()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", rand)

	// Example: generate a UUID according to RFC 4122.
	uuid, err := GenerateUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("%v\n", uuid)
}
```

## Contribution
Let me know your utils or feel free to suggest any kind of improvements, refactorings. Just file an
issue, fork and submit a pull request.

## Get in touch

Tobias Schmid, toashd@gmail.com, [@toashd](http://twitter.com/toashd), [toashd.com](http://toashd.com)

### License
srand is available under the MIT license. See the LICENSE file for more info.
