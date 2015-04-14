package catch

func Check(err error, wrapper func(error) error) {
	if err != nil {
		panic(wrapper(err))
	}
}

func Id(err error) error {
	return err
}

func Catch(err *error) {
	if p := recover(); p != nil {
		*err = p.(error)
	}
}
