
# PayPal Payment API with Go

This is an example of a Go-based API that allows you to create and execute PayPal payments. The API interacts with the PayPal API for payment processing. 

## Prerequisites

- Go installed in your environment.
- PayPal credentials: Replace `"your_client_id_of_paypal_here"` e `"your_client_secret_do_paypal_here"` with your actual PayPal client ID and client secret.
- Familiarity with Go and HTTP requests.

## Configuration

1. Set your PayPal credentials in the code.

2. Run the Go server:

   ```shell
   go run main.go
   ```

The server will be running on port 8080.

## Endpoints

### 1. Create Payment

- **URL:** `/create-payment`
- **Method:** `POST`
- **Request Body (JSON):** 

  ```json
  {
      "amount": 1000
  }
  ```

- **Example Response:**

  ```json
  {
      "approval_url": "https://www.sandbox.paypal.com/checkoutnow?token=EC-1234567890"
  }
  ```

### 2. Execute Payment

- **URL:** `/execute-payment`
- **Method:** `POST`
- **Request Body (JSON):** 

  ```json
  {
      "payment_id": "payment_id_here",
      "payer_id": "payer_id_here"
  }
  ```

- **Example Response:**

  ```json
  {
      "message": "Payment executed successfully"
  }
  ```

## Workflow

1. To create a payment, make a POST request to `/create-payment` with the desired payment amount in the request body.

2. The API interacts with the PayPal API to create a payment and responds with an approval URL. The user should be redirected to this URL for payment approval.

3. After the user approves the payment on the PayPal website, they will be redirected back to your specified return URL.

4. To execute the payment, make a POST request to `/execute-payment` with the payment ID and payer ID received from PayPal.

5. The API interacts with the PayPal API to execute the payment.

Ensure that you follow best security practices when handling sensitive information and payment details.

---

This is a simple example of a PayPal payment API with Go. Customize and enhance this code as needed to meet the specific requirements of your project.

## Developed By:

**Francisco Inoque**