package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/paypal/rest-api-sdk-go/v1"
)

const (
	payPalClientID     = "seu_client_id_do_paypal_aqui"
	payPalClientSecret = "seu_client_secret_do_paypal_aqui"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create-payment", createPayment).Methods("POST")
	r.HandleFunc("/execute-payment", executePayment).Methods("POST")

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type PaymentRequest struct {
	Amount float64 `json:"amount"`
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := paypalsdk.NewClient(paypalsdk.PayPalConfig{
		ClientID:     payPalClientID,
		ClientSecret: payPalClientSecret,
		Mode:         paypalsdk.APIBaseSandBox, // Mude para paypalsdk.APIBaseLive em produção
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.GetAccessToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payment := paypalsdk.Payment{
		Intent: "sale",
		Payer: &paypalsdk.Payer{
			PaymentMethod: "paypal",
		},
		Transactions: []paypalsdk.Transaction{
			{
				Amount: &paypalsdk.Amount{
					Total:    fmt.Sprintf("%.2f", paymentRequest.Amount),
					Currency: "USD",
				},
			},
		},
		RedirectURLs: &paypalsdk.RedirectURLs{
			ReturnURL: "http://example.com/success",
			CancelURL: "http://example.com/cancel",
		},
	}

	resp, err := client.CreatePayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, link := range resp.Links {
		if link.Rel == "approval_url" {
			response := map[string]interface{}{
				"approval_url": link.Href,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Approval URL not found", http.StatusInternalServerError)
}

func executePayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest struct {
		PaymentID string `json:"payment_id"`
		PayerID   string `json:"payer_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := paypalsdk.NewClient(paypalsdk.PayPalConfig{
		ClientID:     payPalClientID,
		ClientSecret: payPalClientSecret,
		Mode:         paypalsdk.APIBaseSandBox, // Mude para paypalsdk.APIBaseLive em produção
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.GetAccessToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	executePayment := paypalsdk.PaymentExecution{
		PayerID: paymentRequest.PayerID,
	}

	_, err = client.ExecuteApprovedPayment(paymentRequest.PaymentID, executePayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Payment executed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
