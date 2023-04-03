package helper

import "github.com/morkid/paginate"

var Paginate = paginate.New(&paginate.Config{
	DefaultSize:          5,
	FieldSelectorEnabled: true,
	CustomParamEnabled:   true,
})
