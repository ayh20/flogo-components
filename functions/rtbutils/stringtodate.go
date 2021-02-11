package rtbutils

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

const (
	stdLongMonth      = "January"
	stdMonth          = "Jan"
	stdNumMonth       = "1"
	stdZeroMonth      = "01"
	stdLongWeekDay    = "Monday"
	stdWeekDay        = "Mon"
	stdDay            = "2"
	stdUnderDay       = "_2"
	stdZeroDay        = "02"
	stdHour           = "15"
	stdHour12         = "3"
	stdZeroHour12     = "03"
	stdMinute         = "4"
	stdZeroMinute     = "04"
	stdSecond         = "5"
	stdZeroSecond     = "05"
	stdLongYear       = "2006"
	stdYear           = "06"
	stdPM             = "PM"
	stdpm             = "pm"
	stdTZ             = "MST"
	stdISO8601TZ      = "Z0700"  // prints Z for UTC
	stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
	stdNumTZ          = "-0700"  // always numeric
	stdNumShortTZ     = "-07"    // always numeric
	stdNumColonTZ     = "-07:00" // always numeric
)

func init() {
	function.Register(&fnstringtodate{})
}

type fnstringtodate struct {
}

func (fnstringtodate) Name() string {
	return "stringtodate"
}

func (fnstringtodate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnstringtodate) Eval(params ...interface{}) (interface{}, error) {

	date, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format date first argument must be string")
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Format date second argument must be string")
	}

	t, err := time.Parse(format, date)

	if err != nil {
		return nil, fmt.Errorf("Error Occured: %w", err)
	}

	return t, nil
}
