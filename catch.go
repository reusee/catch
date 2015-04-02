package catch

import "fmt"

func Check(info interface{}, err error) {
	if err != nil {
		panic(fmt.Errorf("%v error - %v", info, err))
	}
}

func Catch(err *error) {
	if p := recover(); p != nil {
		*err = p.(error)
	}
}
