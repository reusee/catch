package catch

import "fmt"

type Err struct {
	Pkg, Info string
	Err       error
}

func (e Err) Error() string {
	return fmt.Sprintf("%s: %s\n%v", e.Pkg, e.Info, e.Err)
}

func PkgChecker(pkg string) func(error, string) {
	return func(err error, info string) {
		if err != nil {
			panic(Err{
				Pkg:  pkg,
				Info: info,
				Err:  err,
			})
		}
	}
}

func Catch(err *error) {
	if p := recover(); p != nil {
		*err = p.(error)
	}
}
