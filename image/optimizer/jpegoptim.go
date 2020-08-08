package optimizer

import (
	"bytes"
	"os/exec"
	"strconv"

	"github.com/juju/errors"
)

const (
	jpegOptimMinQuality = 1
	jpegOptimMaxQuality = 100
)

func (o *Optimizer) jpegOptim(data []byte, quality int) ([]byte, error) {
	if quality < jpegOptimMinQuality || quality > jpegOptimMaxQuality {
		return nil, errors.Errorf("quality has to be between %d and %d", jpegOptimMinQuality, jpegOptimMaxQuality)
	}
	mParam := "-m" + strconv.Itoa(quality)
	cmd := exec.Command("jpegoptim", "-s", mParam, "--stdin", "--stdout")
	cmd.Stdin = bytes.NewReader(data)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, errors.Trace(err)
	}
	return out.Bytes(), nil
}
