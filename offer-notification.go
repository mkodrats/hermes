package hermes

// Flat is a theme
type OfferNotification struct{}

// Name returns the name of the flat theme
func (dt *OfferNotification) Name() string {
	return "register_confirmation"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *OfferNotification) HTMLTemplate() string {
	return `
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
</head>
<style>
	body {
		font-family:  Arial, 'Helvetica Neue', Helvetica, sans-serif;
	}

	.container {
		max-width: 800px;
		margin: auto;
		color: #333333;
	}

	.container-content {
		max-width: 910px;
		margin: auto;
	}

	.company-image-header {
		margin-top: 20px;
		width: 151px;
		height: auto;
	}

	.btn-complete-registration {
		width: 400px;
		padding: 15px;
		border-radius: 8px;
		background-color: #00aeef;
		font-size: 13px;
		font-weight: 700;
		color: #ffffff;
		font-family: 'Roboto';
		text-decoration: none;
	}

	.help {
		margin: 30px 0px 0px 0px;
		font-size: 10px;
	}

	.link {
		text-decoration: none;
		color: #19a29c;
	}

	.signature {
		margin-top: 50px;
		line-height: 10px;
	}

	.footer {
		margin-top: 50px;
		background-color: #19a29c;
		width: 95%;
		height: 170px;
		padding: 10px;
	}

	.container-footer {
		margin: auto;
	}

	.column {
		margin-top: 20px;
		min-width: 200px;
		float: left;
		font-size: 12px;
		color: #ffffff !important;
	}

	.container-footer {
		margin: auto;
	}

	.column {
		margin-top: 20px;
		min-width: 266px;
		float: left;
		font-size: 12px;
		color: #ffffff !important;
	}

	i {
		font-size: 20px;
	}

	.pi {
		font-size: 12px;
	}

	.column h1 {
		font-size: 16px;
	}

	.column a {
		text-decoration: none;
		color: #ffffff;
	}

	.row:after {
		content: "";
		display: flex;
		clear: both;
	}

	.footer-p {
		text-align: center;
		font-size: 10px;
		color: #ffffff;
	}
</style>

<body>
	<div class="container-content">
		<div class="container">
            <!-- Company Header -->
			<img class="company-image-header" src="{{ .Email.Body.Registration.Logo }}" alt=""srcset="">
			<h1>{{ .Email.Body.Registration.Name }}</h1>
 			{{ range $row := .Email.Body.Registration.Intros }}
			<p>{{ $row }}</p>
    		{{ end }}
			<br>
			<div class="signature">
				<p>Best Regards,</p>
				<p>{{ .Email.Body.Registration.Signature }}</p>
			</div>
		</div>

		<div class="footer">
			<div class="container">
				<div class="row">
					<div class="column">
						<h3>Explore Petronas Love Local</h3>
						<p><a href="{{ .Email.Body.Registration.AboutUs }}">About Us</a></p>
						<p><a href="{{ .Email.Body.Registration.ToU }}">Term Of Use</a></p>
					</div>
					<div class="column">
						<h3>Contact us</h3>
						<p class="pi"><i class="fas fa-phone-alt pi"></i> Mesralink Contact Number</p>
						<p class="pi"><i class="far fa-envelope pi"></i> {{ .Email.Body.Registration.Contact.PhoneNumber }}</p>
					</div>
					<div style="margin-top:50px;min-width:200px;float:left;font-size:12px;color:#ffffff">
						<p class="pi" style="text-decoration: none; color:#fff;"><i class="fas fa-phone-alt pi"></i> <a href="mailto:offers@petronaslove.com.my" target="_blank" style="text-decoration: none; color: #fff">offers@petronaslove.com.my</a></p>
						<p class="pi" style="text-decoration: none; color:#fff;"><i class="far fa-envelope pi"></i> <a href="mailto:{{ .Email.Body.Registration.Contact.Email }}" target="_blank" style="text-decoration: none; color: #fff">{{ .Email.Body.Registration.Contact.Email }}</a></p>
					</div>
				</div>
			</div>
			<div class="hhh"></div>
			<div style="margin-top:10px;min-width:100%;text-align:center;font-size:12px;color:#ffffff;display: flow-root;">
				<p class="footer-p">{{ .Email.Body.Registration.Copyright }}</p>
			</div>
		</div>
	</div>
</body>

</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *OfferNotification) PlainTextTemplate() string {
	return `<h2>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }}{{ end }},</h2>
{{ with .Email.Body.Intros }}
  {{ range $line := . }}
    <p>{{ $line }}</p>
  {{ end }}
{{ end }}
{{ if (ne .Email.Body.FreeMarkdown "") }}
  {{ .Email.Body.FreeMarkdown.ToHTML }}
{{ else }}
  {{ with .Email.Body.Dictionary }}
    <ul>
    {{ range $entry := . }}
      <li>{{ $entry.Key }}: {{ $entry.Value }}</li>
    {{ end }}
    </ul>
  {{ end }}
  {{ with .Email.Body.Table }}
    {{ $data := .Data }}
    {{ $columns := .Columns }}
    {{ if gt (len $data) 0 }}
      <table class="data-table" width="100%" cellpadding="0" cellspacing="0">
        <tr>
          {{ $col := index $data 0 }}
          {{ range $entry := $col }}
            <th>{{ $entry.Key }} </th>
          {{ end }}
        </tr>
        {{ range $row := $data }}
          <tr>
            {{ range $cell := $row }}
              <td>
                {{ $cell.Value }}
              </td>
            {{ end }}
          </tr>
        {{ end }}
      </table>
    {{ end }}
  {{ end }}
  {{ with .Email.Body.Actions }} 
    {{ range $action := . }}
      <p>
        {{ $action.Instructions }} 
        {{ if $action.InviteCode }}
          {{ $action.InviteCode }}
        {{ end }}
        {{ if $action.Button.Link }}
          {{ $action.Button.Link }}
        {{ end }}
      </p> 
    {{ end }}
  {{ end }}
{{ end }}
{{ with .Email.Body.Outros }} 
  {{ range $line := . }}
    <p>{{ $line }}<p>
  {{ end }}
{{ end }}
<p>{{.Email.Body.Signature}},<br>{{.Hermes.Product.Name}} - {{.Hermes.Product.Link}}</p>

<p>{{.Hermes.Product.Copyright}}</p>
`
}
