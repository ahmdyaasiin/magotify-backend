package scheduler

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/jasonlvhit/gocron"
)

func Run(tu usecase.InterfaceTransactionUseCase) {
	UpdateExpiredTransaction(tu)
	UpdateExpiredOrder(tu)

	go func() {
		<-gocron.Start()
	}()
}

func UpdateExpiredTransaction(tu usecase.InterfaceTransactionUseCase) {
	gocron.Every(5).Minutes().Do(tu.UpdateExpiredTransaction)
}

func UpdateExpiredOrder(tu usecase.InterfaceTransactionUseCase) {
	gocron.Every(5).Minutes().Do(tu.UpdateExpiredOrder)
}
