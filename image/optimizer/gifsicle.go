package optimizer

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/espal-digital-development/system/permissions"
	"github.com/juju/errors"
)

const (
	gifsicleDefaultOptimizationLevel = 3
	gifsicleMinLossy                 = 1
	gifsicleMaxLossy                 = 100
)

func (o *Optimizer) gifsicle(data []byte, optimizationLevel int, lossy int) ([]byte, error) {
	if lossy < gifsicleMinLossy || lossy > gifsicleMaxLossy {
		return nil, errors.Errorf("lossy has to be between %d and %d", gifsicleMinLossy, gifsicleMaxLossy)
	}
	tmpFile := strings.TrimSuffix(os.TempDir(), "/") + "/gifsicle.tmp.gif"
	if err := ioutil.WriteFile(tmpFile, data, permissions.UserReadWrite); err != nil {
		return nil, errors.Trace(err)
	}
	oParam := "-O" + strconv.Itoa(optimizationLevel)
	cmd := exec.Command("gifsicle", oParam, "--lossy="+strconv.Itoa(lossy), "--use-col=web", tmpFile)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, errors.Trace(err)
	}
	if err := os.Remove(tmpFile); err != nil {
		return nil, errors.Trace(err)
	}
	return out.Bytes(), nil
}
