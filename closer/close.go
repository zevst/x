package closer

import (
	"fmt"
	"github.com/zevst/x/log"
	"go.uber.org/zap"
	"io"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Error("Closing", zap.String("struct", fmt.Sprintf("%T", c)), zap.Error(err))
	}
}
