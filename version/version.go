package version

import "fmt"

const Version = "0.1.0"

func Display() {
	fmt.Println(`{"version" : "`, Version, `"}`)
}
