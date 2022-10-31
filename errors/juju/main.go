package main

import (
	"fmt"
	"github.com/juju/errors"
)

func main() {
	err := errors.Errorf("original")
	err = errors.Annotatef(err, "context")
	err = errors.Annotatef(err, "more context")
	fmt.Println(errors.ErrorStack(err))
}
