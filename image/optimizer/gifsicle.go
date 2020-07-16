package optimizer

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/juju/errors"
)

const (
	gifsicleMinLossy = 1
	gifsicleMaxLossy = 100
)

func (o *Optimizer) gifsicle(data []byte, optimizationLevel int, lossy int) ([]byte, error) {
	if lossy < gifsicleMinLossy || lossy > gifsicleMaxLossy {
		return nil, errors.Errorf("lossy has to be between %d and %d", gifsicleMinLossy, gifsicleMaxLossy)
	}
	tmpFile := os.TempDir() + "gifsicle.tmp.gif"
	if err := ioutil.WriteFile(tmpFile, data, 0600); err != nil {
		return nil, errors.Trace(err)
	}
	cmd := exec.Command("gifsicle", "-O"+strconv.Itoa(optimizationLevel), "--lossy="+strconv.Itoa(lossy),
		"--use-col=web", tmpFile)
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
