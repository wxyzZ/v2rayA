package controller

import (
	"fmt"
)

var (
	updating bool

	processingErr = fmt.Errorf("the last request is being processed")
)
