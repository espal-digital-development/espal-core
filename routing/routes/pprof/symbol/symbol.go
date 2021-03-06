package symbol

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/system/units"
	"github.com/juju/errors"
)

// Route processor.
type Route struct{}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	context.SetContentType("text/plain")
	responseBuffer := bytes.NewBuffer(nil)
	responseBuffer.WriteString("num_symbols: 1\n")
	var bodyReader *bufio.Reader
	if context.Method() == http.MethodPost {
		bodyReader = bufio.NewReader(context.RequestBody())
	} else {
		bodyReader = bufio.NewReader(bytes.NewReader([]byte(context.QueryString())))
	}
	for {
		word, err := bodyReader.ReadSlice('+')
		if err == nil {
			word = word[0 : len(word)-1]
		}
		pc, err := strconv.ParseUint(string(word), units.Base10, units.BitWidth64Bit)
		if err != nil {
			pc = 0
		}
		if pc != 0 {
			function := runtime.FuncForPC(uintptr(pc))
			if function != nil {
				fmt.Fprintf(responseBuffer, "%#x %s\n", pc, function.Name())
			}
		}
		if err != nil {
			if io.EOF != err {
				fmt.Fprintf(responseBuffer, "reqding request: %v\n", err)
			}
			break
		}
	}
	if _, err := responseBuffer.WriteTo(context); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
