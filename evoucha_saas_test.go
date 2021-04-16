package hermes

import (
	"fmt"
	"log"
	"testing"
)

func TestEvouchaSAASQREmail_HTMLTemplate(t *testing.T) {
	h := Hermes{
		Theme: new(EvouchaSAASQREmail),
		Product: Product{
			Name:      "TruRewards",
			Link:      "",
			Logo:      "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
			Copyright: "TruRewards ©2020. All Rights Reserved",
		},
	}

	email := Email{
		Body: Body{
			EvouchaSAAS: EvouchaSAASConfig{
				PartnerName:        "Bank ABC",
				PartnerLogo:        "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
				MainTextColor:      "black",
				ImportantTextColor: "red",
				Greeting:           "Your voucher is successfully purchased. Here is transaction details with QR Code for redemption at the store.",
				PhoneNumber:        "6281371237",
				EmailAddress:       "support@mail.com",
				Address:            "Address Test",
				EvouchaSAASMailInfo: EvouchaSAASMailInfo{
					CustomerName:   "Koko",
					QRCode:         "",
					ActivationCode: "Activation",
					ProductName:    "Product Name Test",
					Denom:          "100 RM",
				},
			},
		}}

	res, err := h.GenerateHTML(email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
