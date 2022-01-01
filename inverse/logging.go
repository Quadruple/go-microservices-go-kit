package inverse

import (
	"go-microservices/matrix"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	InverseService
}

func NewLoggingService(logger log.Logger, invService InverseService) InverseService {
	return &loggingService{logger, invService}
}

func (logService *loggingService) GetInverse(matrixHolder *matrix.MatrixHolder) (_ float64, err error) {
	defer func(begin time.Time) {
		logService.logger.Log(
			"Method", "GetInverse",
			"Received At:", begin.Format(time.ANSIC),
			"Took", time.Since(begin),
			"Error", err,
		)
	}(time.Now())
	return logService.InverseService.GetInverse(matrixHolder)
}
