// Package temperrcatcher provides a TempErrCatcher object,
// which implements simple error-retrying functionality.
// It is meant to be used with things like net.Lister.Accept:
//
//   import (
//     tec "github.com/jbenet/go-temp-err-catcher"
//   )
//
//   func listen(listener net.Listener) {
//     var c tec.TempErrCatcher
//
//     for {
//       conn, err := listener.Accept()
//       if err != nil && c.IsTemporary(c) {
//         continue
//       }
//       return conn, err
//     }
//   }
//
// You can make your errors implement `Temporary`:
//
//   type errTemp struct {
//     e error
//   }
//
//   func (e errTemp) Temporary() bool {
//     return true
//   }
//
//   func (e errTemp) Error() string {
//     return e.e.Error()
//   }
//
//   err := errors.New("beep boop")
//   var c tec.TempErrCatcher
//   c.IsTemporary(err)              // false
//   c.IsTemporary(errTemp{err}) // true
//
// Or just use `ErrTemp`:
//
//   err := errors.New("beep boop")
//   var c tec.TempErrCatcher
//   c.IsTemporary(err)              // false
//   c.IsTemporary(tec.ErrTemp{err}) // true
//
//
// You can also define an `IsTemp` function to classify errors:
//
//   var ErrSkip = errors.New("this should be skipped")
//   var ErrNotSkip = errors.New("this should not be skipped")
//
//   var c tec.TempErrCatcher
//   c.IsTemp = func(e error) bool {
//     return e == ErrSkip
//   }
//
//   c.IsTemporary(ErrSkip) // true
//   c.IsTemporary(ErrNotSkip) // false
//   c.IsTemporary(ErrTemp) // false! no longer accepts Temporary()
//
package temperrcatcher
