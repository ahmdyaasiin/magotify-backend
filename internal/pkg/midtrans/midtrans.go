package midtrans

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"os"
	"strings"
)

func CreateToken(idTransaction string, totalAmount int64) (string, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  idTransaction,
			GrossAmt: totalAmount,
		},
		Callbacks:       nil,
		EnabledPayments: snap.AllSnapPaymentType,
		Expiry: &snap.ExpiryDetails{
			Duration: 30,
			Unit:     "seconds",
		},
	}

	var client snap.Client
	client.New(serverKey, midtrans.Sandbox)

	var callback string

	if strings.Contains(idTransaction, "SHP") {
		callback = "shop"
	} else {
		callback = "pick_up"
	}

	client.Options.SetPaymentOverrideNotification(fmt.Sprintf("%s/payment/%s/validate", "https://starfish-neutral-coyote.ngrok-free.app/v1", callback))

	snapResp, err := client.CreateTransactionToken(req)
	if err != nil {
		return "", err
	}

	return snapResp, nil
}

func VerifyPayment(idTransaction string) (*coreapi.TransactionStatusResponse, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

	var client coreapi.Client
	client.New(serverKey, midtrans.Sandbox)

	transactionStatusResp, e := client.CheckTransaction(idTransaction)
	if e != nil {
		return nil, e
	} else {
		return transactionStatusResp, nil
	}
}
