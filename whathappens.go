// +build what whathappens

package what

import (
	"log"
)

// Happens logs the current function name and whatever message is passed in.
func Happens(fmt string, args ...interface{}) {
	if isPackageEnabled() {
		log.Printf(funcname(2)+": "+fmt+"\n", args...)
	}
}

func If(yes bool, fmt string, args ...interface{}) {
	if yes && isPackageEnabled() {
		log.Printf(funcname(2)+": "+fmt+"\n", args...)
	}
}
