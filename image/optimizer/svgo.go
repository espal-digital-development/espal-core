package optimizer

import (
	"bytes"
	"os/exec"

	"github.com/juju/errors"
)

func (o *Optimizer) svgo(data []byte) ([]byte, error) {
	cmd := exec.Command("svgo", "--multipass", "-i", "-", "-o", "-")
	cmd.Stdin = bytes.NewReader(data)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, errors.Trace(err)
	}
	return out.Bytes(), nil
}
