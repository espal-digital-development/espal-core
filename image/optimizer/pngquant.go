package optimizer

import (
	"bytes"
	"os/exec"
	"strconv"

	"github.com/juju/errors"
)

const (
	pngQuantMinSpeed = 1
	pngQuantMaxSpeed = 10
)

func (o *Optimizer) pngQuant(data []byte, speed int) ([]byte, error) {
	if speed < pngQuantMinSpeed || speed > pngQuantMaxSpeed {
		return nil, errors.Errorf("speed has to be between %d and %d", pngQuantMinSpeed, pngQuantMaxSpeed)
	}
	speedParam := strconv.Itoa(speed)
	cmd := exec.Command("pngquant", "-", "--speed", speedParam)
	cmd.Stdin = bytes.NewReader(data)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, errors.Trace(err)
	}
	return out.Bytes(), nil
}
