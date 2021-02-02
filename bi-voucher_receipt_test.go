package hermes

import (
	"fmt"
	"log"
	"testing"
)

func TestBIVoucherReceipt_HTMLTemplate(t *testing.T) {
	h := Hermes{
		Theme: new(BIVoucherReceipt),
		Product: Product{
			Name:      "TruRewards",
			Link:      "",
			Logo:      "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
			Copyright: "TruRewards ©2020. All Rights Reserved",
		},
	}

	email := Email{
		Body: Body{
			Name: "Voucher Redemption",
			VoucherReceipt: VoucherReceipt{
				Logo:     "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
				Name:     "Customer",
				SubTitle: "Congratulations, Your Voucher Successfully Redeemed",
				Intros: []string{
					"Below are details of the redemption for your reference :",
				},
				Voucher: Voucher{
					Merchant:   "Jaya Electronics Store",
					Code:       "ABC1234567890",
					Name:       "Jaya Electronics Store Voucher \nExtra Discount",
					Value:      10,
					Currency:   "MYR",
					RedeemTime: "Wednesday, 16 December 2020, 09:24\nAM",
					Status:     "Success",
					StatusCode: 1,
				},
				Signature: "TruRewards",
				Help:      "If you need help, send us email at ",
				HelpEmail: "support@bankislam.com",
				Contact: Contact{
					Email:       "customer.support@bankislam.com",
					PhoneNumber: "(+60) 19 296 9465",
					Address:     "51G, Jalan Desa 9/6, Bandar Country Homes, 48000 Rawang, Selangor, Malaysia",
				},
				EmailAddress: "oni@authscure.com.my",
			},
			Actions: []Action{
				{
					Button: Button{
						Color: "#c70773",
						Text:  "Visit Website",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
		}}

	res, err := h.GenerateHTML(email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
