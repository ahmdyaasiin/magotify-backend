package midtrans

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"os"
)

func CreateToken(idTransaction uuid.UUID, product *entity.Product) (string, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  idTransaction.String(),
			GrossAmt: int64(product.Price),
		},
		Callbacks: &snap.Callbacks{
			Finish: fmt.Sprintf("%s/transaction", "https://example.com"),
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
			snap.PaymentTypeBankTransfer,
		},
		Expiry: &snap.ExpiryDetails{
			Duration: 5,
			Unit:     "minute",
		},
	}

	var client snap.Client
	client.New(serverKey, midtrans.Sandbox)
	client.Options.SetPaymentOverrideNotification(fmt.Sprintf("%s/product/payment/callback", "https://example.com"))

	snapResp, err := client.CreateTransactionToken(req)
	if err != nil {
		return "", err
	}

	return snapResp, err
}

func VerifyPayment(idTransaction uuid.UUID) (*coreapi.TransactionStatusResponse, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

	var client coreapi.Client
	client.New(serverKey, midtrans.Sandbox)

	transactionStatusResp, e := client.CheckTransaction(idTransaction.String())
	if e != nil {
		return nil, e
	} else {
		return transactionStatusResp, nil
	}
}
