package utils

func MustNotError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
