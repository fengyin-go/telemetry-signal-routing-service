package flow16

import (
	"errors"

	"telemetry-signal-routing-service/internal/state16"
)

func Forward(source *state16.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state16.Normalize(source.Next())
		if err == nil {
			return nil
		}
		// 线路拒收回执是不可恢复的失败：立即返回，且保持其类型可识别，
		// 而非继续重试后在步骤耗尽时被吞成正常。
		var rejected *state16.Rejected
		if errors.As(err, &rejected) {
			return err
		}
		last = err
	}
	return last
}
