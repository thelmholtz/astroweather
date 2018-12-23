//Package except contains a json exportable error type, with error code and description fields.
//It extends on the default error which accepts nil values; to allow Go's idiomatic error checking
/*Usage:
```encoder := json.NewEncoder(outputWriter)
if err := failableMethod; err != nil {
	encoder.Encode(err)
}
... success code```
*/
package except

import "fmt"

//E is an interface implementing a json marshallable error type.
type E interface {
	error
	Code() string
	Description() string
}

//Except implements except
type except struct {
	C string `json:"code"`
	D string `json:"description"`
}

//Error returns a json message
func (e except) Error() string {
	return fmt.Sprintf("Runtime ERROR: {\n\t\"code\":%s,\n\t\"description\":%s\n}\n", e.C, e.D) //There should be a way to use the json encoder to build this messagecd
}

//New returns a new error with this code and description
func New(c, d string) E {
	return except{C: c, D: d}
}

//Code returns the exception error code
func (e except) Code() string {
	return e.C
}

//Description returns the exception description
func (e except) Description() string {
	return e.D
}
