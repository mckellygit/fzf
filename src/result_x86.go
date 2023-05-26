//go:build 386 || amd64

package fzf

import "unsafe"
import "strings"
// import "fmt"

func compareRanks(irank Result, jrank Result, tac bool) bool {
	left := *(*uint64)(unsafe.Pointer(&irank.points[0]))
	right := *(*uint64)(unsafe.Pointer(&jrank.points[0]))
	if left < right {
		return true
	} else if left > right {
		return false
    }

    // fmt.Print("i,j = ", irank.item.Index(), " ", irank.item.text.ToString(), " ", jrank.item.Index(), " ", jrank.item.text.ToString())
    // fmt.Println()

    var cval = strings.Compare(irank.item.text.ToString(), jrank.item.text.ToString())

    if (!tac) {

        if cval <= 0 {
            return true
        } else {
            return false
        }

    } else {

        if cval <= 0 {
            return false
        } else {
            return true
        }

    }

	// return (irank.item.Index() <= jrank.item.Index()) != tac
}
