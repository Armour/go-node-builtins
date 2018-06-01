# go-node-builtins

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Template from jarvis](https://img.shields.io/badge/Hi-Jarvis-ff69b4.svg)](https://github.com/Armour/Jarvis)

## Install

```bash
go get github.com/Armour/go-node-builtins
```

## Example

Get list of core modules for specific Node.js version:

```go
package main

import (
	"fmt"

	"github.com/Armour/go-node-builtins/builtins"
)

func main() {
	b, err := builtins.GetVersion("6.0.0")
	if err != nil {
		// handle error
	}
	fmt.Printf("%v", b)
}
```

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

[MIT License](https://github.com/Armour/go-node-builtins/blob/master/LICENSE)
