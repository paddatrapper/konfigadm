package phases

import (
	"github.com/moshloop/configadm/pkg/os"
	"github.com/moshloop/configadm/pkg/types"
)

func GetTags(os os.OS) []types.Flag {
	return []types.Flag{*types.GetTag(os.GetTag())}
}