package smoothengine

import (
	smoothgorm "github.com/focardoso/smooth-engine/gorm"
	"github.com/focardoso/smooth-orm"
)

func Gorm(config smoothgorm.Config) smooth.Engine {
	return smoothgorm.Open(config)
}
