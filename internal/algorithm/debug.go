package algorithm

import "fmt"

var DebugEnabled = false

func Debug(format string, a ...any) {
	if DebugEnabled {
		fmt.Printf("[DEBUG] "+format+"\n", a...)
	}
}
