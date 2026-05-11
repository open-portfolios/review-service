package infra

import (
	"context"
	"time"

	"github.com/open-portfolios/review/internal/biz"
	"github.com/open-portfolios/review/internal/conf"
	"github.com/sony/sonyflake/v2"
)

type snowflakeRepo struct {
	flake *sonyflake.Sonyflake
}

func NewSnowflakeRepo(s *conf.Snowflake) (biz.SnowflakeRepo, error) {
	startTime, err := time.Parse(time.DateOnly, s.StartTime)
	if err != nil {
		return nil, err
	}

	var machineID func() (int, error)
	if s.MachineId != 0 {
		machineID = func() (int, error) { return int(s.MachineId), nil }
	}

	flake, err := sonyflake.New(sonyflake.Settings{
		StartTime: startTime,
		MachineID: machineID,
	})
	if err != nil {
		return nil, err
	}
	return &snowflakeRepo{flake}, nil
}

func (r *snowflakeRepo) Generate(ctx context.Context) (int64, error) {
	resultCh := make(chan int64, 1)
	errCh := make(chan error, 1)

	go func() {
		defer close(resultCh)
		defer close(errCh)
		result, err := r.flake.NextID()
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case <-ctx.Done():
		return 0, context.Canceled
	case err := <-errCh:
		return 0, err
	case v := <-resultCh:
		return v, nil
	}
}
