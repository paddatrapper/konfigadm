package phases

import (
	"github.com/moshloop/configadm/pkg/utils"
)

var Environment Phase = environment{}

type environment struct{}

func (p environment) ApplyPhase(sys *SystemConfig, ctx *SystemContext) ([]Command, Filesystem, error) {
	var commands []Command
	files := Filesystem{}

	if len(sys.Environment) > 0 {
		files["/etc/environment"] = File{Content: utils.MapToIni(sys.Environment)}
	}
	return commands, files, nil
}