/*
Package goutil implements many useful little tools for Go
*/
package goutil

// PanicOnError panic upon receiving an non-nil error
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
