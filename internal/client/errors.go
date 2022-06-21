package client

import "fmt"

var (
	errorUnobtainedAccess = fmt.Errorf("can't udate, no access token obtained")
	errorConnection       = fmt.Errorf("response has invalid status code")
)
