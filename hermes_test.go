package hermes

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testedThemes = []Theme{
	// Insert your new theme here
	new(Default),
	new(Flat),
	new(WithQRCode),
}

/////////////////////////////////////////////////////
// Every theme should display the same information //
// Find below the tests to check that              //
/////////////////////////////////////////////////////

// Implement this interface when creating a new example checking a common feature of all themes
type Example interface {
	// Create the hermes example with data
	// Represents the "Given" step in Given/When/Then Workflow
	getExample() (h Hermes, email Email)
	// Checks the content of the generated HTML email by asserting content presence or not
	assertHTMLContent(t *testing.T, s string)
	// Checks the content of the generated Plaintext email by asserting content presence or not
	assertPlainTextContent(t *testing.T, s string)
}

// Scenario
type SimpleExample struct {
	theme Theme
}

func (ed *SimpleExample) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name:      "HermesName",
			Link:      "http://hermes-link.com",
			Copyright: "Copyright © Hermes-Test",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
		TextDirection:      TDLeftToRight,
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Dictionary: []Entry{
				{"Firstname", "Jon"},
				{"Lastname", "Snow"},
				{"Birthday", "01/01/283"},
			},
			Table: Table{
				Data: [][]Entry{
					{
						{Key: "Item", Value: "Golang"},
						{Key: "Description", Value: "Open source programming language that makes it easy to build simple, reliable, and efficient software"},
						{Key: "Price", Value: "$10.99"},
					},
					{
						{Key: "Item", Value: "Hermes"},
						{Key: "Description", Value: "Programmatically create beautiful e-mails using Golang."},
						{Key: "Price", Value: "$1.99"},
					},
				},
				Columns: Columns{
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignment: map[string]string{
						"Price": "right",
					},
				},
			},
			Actions: []Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	return h, email
}

func (ed *SimpleExample) assertHTMLContent(t *testing.T, r string) {

	// Assert on product
	assert.Contains(t, r, "HermesName", "Product: Should find the name of the product in email")
	assert.Contains(t, r, "http://hermes-link.com", "Product: Should find the link of the product in email")
	assert.Contains(t, r, "Copyright © Hermes-Test", "Product: Should find the Copyright of the product in email")
	assert.Contains(t, r, "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png", "Product: Should find the logo of the product in email")
	assert.Contains(t, r, "If you’re having trouble with the button &#39;Confirm your account&#39;, copy and paste the URL below into your web browser.", "Product: Should find the trouble text in email")
	// Assert on email body
	assert.Contains(t, r, "Hi Jon Snow", "Name: Should find the name of the person")
	assert.Contains(t, r, "Welcome to Hermes", "Intro: Should have intro")
	assert.Contains(t, r, "Birthday", "Dictionary: Should have dictionary")
	assert.Contains(t, r, "Open source programming language", "Table: Should have table with first row and first column")
	assert.Contains(t, r, "Programmatically create beautiful e-mails using Golang", "Table: Should have table with second row and first column")
	assert.Contains(t, r, "$10.99", "Table: Should have table with first row and second column")
	assert.Contains(t, r, "$1.99", "Table: Should have table with second row and second column")
	assert.Contains(t, r, "started with Hermes", "Action: Should have instruction")
	assert.Contains(t, r, "Confirm your account", "Action: Should have button of action")
	assert.Contains(t, r, "#22BC66", "Action: Button should have given color")
	assert.Contains(t, r, "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010", "Action: Button should have link")
	assert.Contains(t, r, "Need help, or have questions", "Outro: Should have outro")
}

func (ed *SimpleExample) assertPlainTextContent(t *testing.T, r string) {

	// Assert on product
	assert.Contains(t, r, "HermesName", "Product: Should find the name of the product in email")
	assert.Contains(t, r, "http://hermes-link.com", "Product: Should find the link of the product in email")
	assert.Contains(t, r, "Copyright © Hermes-Test", "Product: Should find the Copyright of the product in email")
	assert.NotContains(t, r, "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png", "Product: Should not find any logo in plain text")

	// Assert on email body
	assert.Contains(t, r, "Hi Jon Snow", "Name: Should find the name of the person")
	assert.Contains(t, r, "Welcome to Hermes", "Intro: Should have intro")
	assert.Contains(t, r, "Birthday", "Dictionary: Should have dictionary")
	assert.Contains(t, r, "Open source", "Table: Should have table content")
	assert.Contains(t, r, `+--------+--------------------------------+--------+
|  ITEM  |          DESCRIPTION           | PRICE  |
+--------+--------------------------------+--------+
| Golang | Open source programming        | $10.99 |
|        | language that makes it easy    |        |
|        | to build simple, reliable, and |        |
|        | efficient software             |        |
| Hermes | Programmatically create        | $1.99  |
|        | beautiful e-mails using        |        |
|        | Golang.                        |        |
+--------+--------------------------------+--------`, "Table: Should have pretty table content")
	assert.Contains(t, r, "started with Hermes", "Action: Should have instruction")
	assert.NotContains(t, r, "Confirm your account", "Action: Should not have button of action in plain text")
	assert.NotContains(t, r, "#22BC66", "Action: Button should not have color in plain text")
	assert.Contains(t, r, "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010", "Action: Even if button is not possible in plain text, it should have the link")
	assert.Contains(t, r, "Need help, or have questions", "Outro: Should have outro")
}

type WithTitleInsteadOfNameExample struct {
	theme Theme
}

func (ed *WithTitleInsteadOfNameExample) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name:  "Jon Snow",
			Title: "A new e-mail",
		},
	}
	return h, email
}

func (ed *WithTitleInsteadOfNameExample) assertHTMLContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Name: should not find greetings from Jon Snow because title should be used")
	assert.Contains(t, r, "A new e-mail", "Title should be used instead of name")
}

func (ed *WithTitleInsteadOfNameExample) assertPlainTextContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Name: should not find greetings from Jon Snow because title should be used")
	assert.Contains(t, r, "A new e-mail", "Title shoud be used instead of name")
}

type WithGreetingDifferentThanDefault struct {
	theme Theme
}

func (ed *WithGreetingDifferentThanDefault) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Greeting: "Dear",
			Name:     "Jon Snow",
		},
	}
	return h, email
}

func (ed *WithGreetingDifferentThanDefault) assertHTMLContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Should not find greetings with 'Hi' which is default")
	assert.Contains(t, r, "Dear Jon Snow", "Should have greeting with Dear")
}

func (ed *WithGreetingDifferentThanDefault) assertPlainTextContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Should not find greetings with 'Hi' which is default")
	assert.Contains(t, r, "Dear Jon Snow", "Should have greeting with Dear")
}

type WithSignatureDifferentThanDefault struct {
	theme Theme
}

func (ed *WithSignatureDifferentThanDefault) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name:      "Jon Snow",
			Signature: "Best regards",
		},
	}
	return h, email
}

func (ed *WithSignatureDifferentThanDefault) assertHTMLContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Yours truly", "Should not find signature with 'Yours truly' which is default")
	assert.Contains(t, r, "Best regards", "Should have greeting with Dear")
}

func (ed *WithSignatureDifferentThanDefault) assertPlainTextContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Yours truly", "Should not find signature with 'Yours truly' which is default")
	assert.Contains(t, r, "Best regards", "Should have greeting with Dear")
}

type WithInviteCode struct {
	theme Theme
}

func (ed *WithInviteCode) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			Actions: []Action{
				{
					Instructions: "Here is your invite code:",
					InviteCode:   "123456",
				},
			},
		},
	}
	return h, email
}

func (ed *WithInviteCode) assertHTMLContent(t *testing.T, r string) {
	assert.Contains(t, r, "Here is your invite code", "Should contains the instruction")
	assert.Contains(t, r, "123456", "Should contain the short code")
}

func (ed *WithInviteCode) assertPlainTextContent(t *testing.T, r string) {
	assert.Contains(t, r, "Here is your invite code", "Should contains the instruction")
	assert.Contains(t, r, "123456", "Should contain the short code")
}

type WithFreeMarkdownContent struct {
	theme Theme
}

func (ed *WithFreeMarkdownContent) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			FreeMarkdown: `
> _Hermes_ service will shutdown the **1st August 2017** for maintenance operations. 

Services will be unavailable based on the following schedule:

| Services | Downtime |
| :------:| :-----------: |
| Service A | 2AM to 3AM |
| Service B | 4AM to 5AM |
| Service C | 5AM to 6AM |

---

Feel free to contact us for any question regarding this matter at [support@hermes-example.com](mailto:support@hermes-example.com) or in our [Gitter](https://gitter.im/)

`,
			Intros: []string{
				"An intro that should be kept even with FreeMarkdown",
			},
			Dictionary: []Entry{
				{"Dictionary that should not be displayed", "Because of FreeMarkdown"},
			},
			Table: Table{
				Data: [][]Entry{
					{
						{Key: "Item", Value: "Golang"},
					},
					{
						{Key: "Item", Value: "Hermes"},
					},
				},
			},
			Actions: []Action{
				{
					Instructions: "Action that should not be displayed, because of FreeMarkdown:",
					Button: Button{
						Color: "#22BC66",
						Text:  "Button",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"An outro that should be kept even with FreeMarkdown",
			},
		},
	}
	return h, email
}

func (ed *WithFreeMarkdownContent) assertHTMLContent(t *testing.T, r string) {
	assert.Contains(t, r, "Yours truly", "Should find signature with 'Yours truly' which is default")
	assert.Contains(t, r, "Jon Snow", "Should find title with 'Jon Snow'")
	assert.Contains(t, r, "<em>Hermes</em> service will shutdown", "Should find quote as HTML formatted content")
	assert.Contains(t, r, "<td align=\"center\">2AM to 3AM</td>", "Should find cell content as HTML formatted content")
	assert.Contains(t, r, "<a href=\"mailto:support@hermes-example.com\">support@hermes-example.com</a>", "Should find link of mailto as HTML formatted content")
	assert.Contains(t, r, "An intro that should be kept even with FreeMarkdown", "Should find intro even with FreeMarkdown")
	assert.Contains(t, r, "An outro that should be kept even with FreeMarkdown", "Should find outro even with FreeMarkdown")
	assert.NotContains(t, r, "should not be displayed", "Should find any other content that the one from FreeMarkdown object")
}

func (ed *WithFreeMarkdownContent) assertPlainTextContent(t *testing.T, r string) {
	assert.Contains(t, r, "Yours truly", "Should find signature with 'Yours truly' which is default")
	assert.Contains(t, r, "Jon Snow", "Should find title with 'Jon Snow'")
	assert.Contains(t, r, "> Hermes service will shutdown", "Should find quote as plain text with quote emphaze on sentence")
	assert.Contains(t, r, "2AM to 3AM", "Should find cell content as plain text")
	assert.Contains(t, r, `+-----------+------------+
| SERVICES  |  DOWNTIME  |
+-----------+------------+
| Service A | 2AM to 3AM |
| Service B | 4AM to 5AM |
| Service C | 5AM to 6AM |
+-----------+------------+`, "Should find pretty table as plain text")
	assert.Contains(t, r, "support@hermes-example.com", "Should find link of mailto as plain text")
	assert.NotContains(t, r, "<table>", "Should not find html table tags")
	assert.NotContains(t, r, "<tr>", "Should not find html tr tags")
	assert.NotContains(t, r, "<a>", "Should not find html link tags")
	assert.NotContains(t, r, "should not be displayed", "Should find any other content that the one from FreeMarkdown object")

}

// Test all the themes for the features

func TestThemeSimple(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &SimpleExample{theme})
	}
}

func TestThemeWithTitleInsteadOfName(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithTitleInsteadOfNameExample{theme})
	}
}

func TestThemeWithGreetingDifferentThanDefault(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithGreetingDifferentThanDefault{theme})
	}
}

func TestThemeWithGreetingDiffrentThanDefault(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithSignatureDifferentThanDefault{theme})
	}
}

func TestThemeWithFreeMarkdownContent(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithFreeMarkdownContent{theme})
	}
}

func TestThemeWithInviteCode(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithInviteCode{theme})
	}
}

func checkExample(t *testing.T, ex Example) {
	// Given an example
	h, email := ex.getExample()

	// When generating HTML template
	r, err := h.GenerateHTML(email)
	t.Log(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, r)

	// Then asserting HTML is OK
	ex.assertHTMLContent(t, r)

	// When generating plain text template
	r, err = h.GeneratePlainText(email)
	t.Log(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, r)

	// Then asserting plain text is OK
	ex.assertPlainTextContent(t, r)
}

////////////////////////////////////////////
// Tests on default values for all themes //
// It does not check email content        //
////////////////////////////////////////////

func TestHermes_TextDirectionAsDefault(t *testing.T) {
	h := Hermes{
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		TextDirection:      "not-existing", // Wrong text-direction
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	_, err := h.GenerateHTML(email)
	assert.Nil(t, err)
	assert.Equal(t, h.TextDirection, TDLeftToRight)
	assert.Equal(t, h.Theme.Name(), "default")
}

func TestHermes_WithQR(t *testing.T) {
	h := Hermes{
		Theme: new(WithQRCode),
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
			Logo: "https://cdn.freelogovectors.net/wp-content/uploads/2020/02/petronas-logo.png",
		},
		TextDirection:      "not-existing", // Wrong text-direction
		DisableCSSInlining: true,
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			QRCode: `data:image/jpeg;base64,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VldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9sAQwACAgICAgIDAgIDBQMDAwUGBQUFBQYIBgYGBgYICggICAgICAoKCgoKCgoKDAwMDAwMDg4ODg4PDw8PDw8PDw8P/9sAQwECAgIEBAQHBAQHEAsJCxAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQ/90ABABE/9oADAMBAAIRAxEAPwD9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoooo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17/RRQAUUUUAFFFFABRRRQAUUUUAFfAvjD/gmp+yp458Wa1411/RtQfVNfvbi/u2j1GdEae6kaWQqoOFBZicDgV99UUAfNnwA/ZO+DX7M82t3HwpsLmzk8QLbrdm4upLncLYuYwu8/LjzGzjrX0nRRQAV82/H/8AZP8Ag1+0xLok/wAV7C5vH8PrcLaG3upLbaLkoZA2w/Nny1xnpX0lRQB8CeEf+CaX7KfgjxXo3jPQdG1BNT0G8t7+1aTUZ3RZ7aRZYyyk4YBlBweDX33RRQAUUUUAFFFFABRRRQAVieJvD2m+LfDeq+FNZV30/WbSeyuBG5jcw3MZjkCuvKnaxwRyOorbooA/OP8A4dVfsff9ATU//BnP/jX3B8Mfht4S+EHgPSPhv4Ftms9D0ONoraJ3aVgHdpGLO2SxZ3ZiT3Nd5RQAUUUUAFfB3jr/AIJu/ssfETxlrXjzxHot+2reILua+u2h1CaONri4YvIwQHC7mJOBxk1940UAfMPwD/ZC+Cv7NWq6trPwqsryzuNahjt7n7ReS3CMkTF1wrnAIJPPXk+pr6eoooAKKKKACiiigAooooAKKKKACvkT43/sPfs//tC+M08ffEvTLy61hLWKzD297Lbp5MLMyDYhxkFzzX13RQB8W/CH9gP9m/4H/EDTfib4A0q+t9d0kTi3knvppkX7RC8EmUY4OUkYDPTOetfaVFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAf/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAztX1GLR9JvdXnUvHYwSTsq/eKxKWIGe5Ar8fP8Ah8x8Jf8AoQtb/wC/lt/8cr9a/HP/ACJPiD/sH3f/AKJav4dKAP7If2Wv2lfDv7U/w6vPiL4Z0m60a2stSm0x4bsoZDJDFDMWBjLDaRMB65Br6Tr8pf8Agj1/ybBr3/Y2Xv8A6RWNfq1QAUUUUAFFFFABRRRQAV4P+0j8etD/AGbPhZe/FPxFptxqtnZ3Fvbm3tSglZrhwgILkDA6nmveK/OP/gqr/wAmfa3/ANhPTP8A0eKAPB/+HzHwl/6ELW/+/lt/8co/4fMfCX/oQtb/AO/lt/8AHK/nfr2vTP2bf2htZ0201jSPhl4lvbC/ijuLe4h0m7kimhlUOkiOsZDKykFSDgg5FAH7a/8AD5j4S/8AQha3/wB/Lb/45R/w+Y+Ev/Qha3/38tv/AI5X4r/8MtftL/8ARKfFP/gmvP8A41R/wy1+0v8A9Ep8U/8AgmvP/jVAH9I/7Kf7fHgf9q3xvqvgbwz4b1HRbrS9ObUmlu2haN40mjhKjy2J3ZlB6YwDX3rX4E/8EoPgz8Xfh18c/FOteP8AwXrPhqwn8OS28c+pWE9pG8zXlq4RWlRQWKoxwOcA1++1ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAf//X/fyiiigAooooAKKKKAOW8c/8iT4g/wCwfd/+iWr+HSv7i/HP/Ik+IP8AsH3f/olq/h0oA/pa/wCCPX/JsGvf9jZe/wDpFY1+rVflL/wR6/5Ng17/ALGy9/8ASKxr9WqACiiigAooooAKKKKACvzj/wCCqv8AyZ9rf/YT0z/0eK/Ryvzj/wCCqv8AyZ9rf/YT0z/0eKAP5aB1r+0j9mr/AJNz+Ff/AGKmh/8ApBDX8W461/aR+zV/ybn8K/8AsVND/wDSCGgD2uiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/0P38ooooAKKKKACiiigDlvHP/Ik+IP8AsH3f/olq/h0r+4vxz/yJPiD/ALB93/6Jav4dKAP6Wv8Agj1/ybBr3/Y2Xv8A6RWNfq1X5S/8Eev+TYNe/wCxsvf/AEisa/VqgAooooAKKKKACiiigAr84/8Agqr/AMmfa3/2E9M/9Hiv0cr84/8Agqr/AMmfa3/2E9M/9HigD+Wgda/tI/Zq/wCTc/hX/wBipof/AKQQ1/FuOtf2kfs1f8m5/Cv/ALFTQ/8A0ghoA9rooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA/9H9/KKKKACiiigAooooA5bxz/yJPiD/ALB93/6Jav4dK/uL8c/8iT4g/wCwfd/+iWr+HSgD+lr/AII9f8mwa9/2Nl7/AOkVjX6tV+Uv/BHr/k2DXv8AsbL3/wBIrGv1aoAKKKKACiiigAooooAK/OP/AIKq/wDJn2t/9hPTP/R4r9HK/OP/AIKq/wDJn2t/9hPTP/R4oA/lnr7X8M/8FDP2sfCHhvSfCWgeMVttL0S0gsbSL7Fat5dvbRrFGu5oyThVAyTk96+KK/Wj4f8A/BJP4nfEDwH4b8e2PjnR7W28SabZ6nFDJDcF40vIUmVGIXBKh8HHGaAPCv8Ah5d+2P8A9Duv/gBaf/Gq7n4Yf8FF/wBrfX/iX4S0HVvGS3FjqWr2FtcRmxtVDxTXCI65WMEZUkZBB9K9r/4cy/Fj/ooOif8Afi5/+JrrPAX/AASF+J/hLxz4d8VXnjzR54NG1KzvZI44Lje6W8yyMq5GMkLgZ4oA/fGiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP/0v38ooooAKKKKACiiigDlvHP/Ik+IP8AsH3f/olq/h0r+4vxz/yJPiD/ALB93/6Jav4dKAP6Wv8Agj1/ybBr3/Y2Xv8A6RWNfq1X5S/8Eev+TYNe/wCxsvf/AEisa/VqgAooooAKKKKACiiigAr84/8Agqr/AMmfa3/2E9M/9Hiv0cr84/8Agqr/AMmfa3/2E9M/9HigD+Wgda/tI/Zq/wCTc/hX/wBipof/AKQQ1/FuOtf2kfs1f8m5/Cv/ALFTQ/8A0ghoA9rooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA/9P9/KKKKACiiigAooooA5bxz/yJPiD/ALB93/6Jav4dK/uL8c/8iT4g/wCwfd/+iWr+HSgD+lr/AII9f8mwa9/2Nl7/AOkVjX6tV+Uv/BHr/k2DXv8AsbL3/wBIrGv1aoAKKKKACiiigAooooAK/OP/AIKq/wDJn2t/9hPTP/R4r9HK/OP/AIKq/wDJn2t/9hPTP/R4oA/loHWv7SP2av8Ak3P4V/8AYqaH/wCkENfxbjrX9pH7NX/Jufwr/wCxU0P/ANIIaAPa6KKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP/U/fyiiigAooooAKKKKAOd8X21xe+E9bs7SMyzz2NzHGijJZ2iYKB7kmv49v8AhkX9qX/ok3ij/wAFN1/8br+yuigD82P+CWfw68e/DL9nbWNB+Inh++8NalP4lvLmO21C3e2maBrS0RZAkgDbSyMAcckGv0noooAKKKKACiiigAooooAK+Df+Ckfgbxl8RP2WdY8NeA9Eu/EGrSahp0i2ljC9xOyRzgsyxoCxCjk4HA5r7yooA/jUH7Iv7Uv/AESbxR/4Kbr/AON1/Wj8AtI1Pw/8CfhxoOt2sljqOm+G9Htrm3lUpJDNDZxJJG6nkMrAgg9CK9aooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9f9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9D9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9H9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9L9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9P9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9T9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9X9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9b9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD/9k=`,
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	_, err := h.GenerateHTML(email)
	assert.Nil(t, err)
	assert.Equal(t, h.TextDirection, TDLeftToRight)
	assert.Equal(t, h.Theme.Name(), "with_qr_code")
}

func TestHermes_Register(t *testing.T) {
	h := Hermes{
		Theme: new(RegisterConfirmation),
		Product: Product{
			Name:      "Petronas Love Local",
			Link:      "",
			Logo:      "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
			Copyright: "Petronas Love Local ©2020. All Rights Reserved",
		},
	}

	email := Email{
		Body: Body{
			Name:         "Lucifer",
			Intros:       nil,
			Dictionary:   nil,
			Table:        Table{},
			QRCode:       "",
			Actions:      nil,
			Outros:       nil,
			Greeting:     "",
			Signature:    "",
			Title:        "",
			FreeMarkdown: "",
			Registration: Registration{
				Logo: "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
				Name: "Lucifer",
				Intros: []string{
					"We have received your account registration request.",
					"To complete the registration, please click the button or manually copy & paste the provided link in your browser url address bar to verify your email account :",
				},
				ActionButton: "COMPLETE REGISTRATION",
				ActionURL:    "https://www.petronaslovelocal.com/activate/token/1234567890xyzpqr0987",
				Expiration:   "1 hour",
				Signature:    "Petronas Love Local",
				Help:         "If the button above is not clickable, copy and paste the link below to activate your account.",
				Copyright:    "Petronas Love Local ©2020. All Rights Reserved",
				SocialMedia: SocialMedia{
					Facebook:  "facebookURL",
					Instagram: "instagramURL",
					Twitter:   "twitterURL",
					Youtube:   "youtubeURL",
				},
				Contact: Contact{
					Email:       " (+60) 19 200 300",
					PhoneNumber: " offers@petronaslove.com.my",
				},
				AboutUs: "aboutUsURL",
				ToU:     "ToUURL",
			},
		},
	}
	res, err := h.GenerateHTML(email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func TestHermes_ManualInspection(t *testing.T) {
	h := Hermes{
		Theme: new(ManualInspection),
		Product: Product{
			Name:      "Petronas Love Local",
			Link:      "",
			Logo:      "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
			Copyright: "Petronas Love Local ©2020. All Rights Reserved",
		},
	}

	email := Email{
		Body: Body{
			Name:       "Manual Upload",
			Intros:     nil,
			Dictionary: nil,

			Table:        Table{},
			QRCode:       "",
			Actions:      nil,
			Outros:       nil,
			Greeting:     "",
			Signature:    "",
			Title:        "",
			FreeMarkdown: "",
			Registration: Registration{
				Logo:         "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
				Name:         "Manual Upload",
				EmailAddress: "mohdjamilafiq@gmail.com",
				Intros: []string{
					"Petronas Love Local's user has added a new receipt through manual upload.",
				},
				Signature: "Petronas Love Local",
				Copyright: "Petronas Love Local ©2020. All Rights Reserved",
				SocialMedia: SocialMedia{
					Facebook:  "facebookURL",
					Instagram: "instagramURL",
					Twitter:   "twitterURL",
					Youtube:   "youtubeURL",
				},
				Contact: Contact{
					Email:       " (+60) 19 200 300",
					PhoneNumber: " offers@petronaslove.com.my",
				},
				AboutUs: "aboutUsURL",
				ToU:     "ToUURL",
			},
		},
	}
	res, err := h.GenerateHTML(email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func TestHermes_OfferNotification(t *testing.T) {
	h := Hermes{
		Theme: new(OfferNotification),
		Product: Product{
			Name:      "Petronas Love Local",
			Link:      "",
			Logo:      "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
			Copyright: "Petronas Love Local ©2020. All Rights Reserved",
		},
	}

	email := Email{
		Body: Body{
			Name:       "Manual Upload",
			Intros:     nil,
			Dictionary: nil,

			Table:        Table{},
			QRCode:       "",
			Actions:      nil,
			Outros:       nil,
			Greeting:     "",
			Signature:    "",
			Title:        "",
			FreeMarkdown: "",
			Registration: Registration{
				Logo:         "https://storage.googleapis.com/offer-img-stg/logo-1%403x.png",
				Name:         "Manual Upload",
				EmailAddress: "mohdjamilafiq@gmail.com",
				Intros: []string{
					"Petronas Love Local's user has added a new receipt through manual upload.",
				},
				Signature: "Petronas Love Local",
				Copyright: "Petronas Love Local ©2020. All Rights Reserved",
				SocialMedia: SocialMedia{
					Facebook:  "facebookURL",
					Instagram: "instagramURL",
					Twitter:   "twitterURL",
					Youtube:   "youtubeURL",
				},
				Contact: Contact{
					Email:       " (+60) 19 200 300",
					PhoneNumber: " offers@petronaslove.com.my",
				},
				AboutUs: "aboutUsURL",
				ToU:     "ToUURL",
			},
		},
	}
	res, err := h.GenerateHTML(email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func TestHermes_Default(t *testing.T) {
	h := Hermes{}
	setDefaultHermesValues(&h)
	email := Email{}
	setDefaultEmailValues(&email)

	assert.Equal(t, h.TextDirection, TDLeftToRight)
	assert.Equal(t, h.Theme, new(Default))
	assert.Equal(t, h.Product.Name, "Hermes")
	assert.Equal(t, h.Product.Copyright, "Copyright © 2020 Hermes. All rights reserved.")

	assert.Empty(t, email.Body.Actions)
	assert.Empty(t, email.Body.Dictionary)
	assert.Empty(t, email.Body.Intros)
	assert.Empty(t, email.Body.Outros)
	assert.Empty(t, email.Body.Table.Data)
	assert.Empty(t, email.Body.Table.Columns.CustomWidth)
	assert.Empty(t, email.Body.Table.Columns.CustomAlignment)
	assert.Empty(t, string(email.Body.FreeMarkdown))

	assert.Equal(t, email.Body.Greeting, "Hi")
	assert.Equal(t, email.Body.Signature, "Yours truly")
	assert.Empty(t, email.Body.Title)
}
