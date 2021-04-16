package hermes

// Flat is a theme
type EvouchaSAASQREmail struct{}

// Name returns the name of the flat theme
func (dt *EvouchaSAASQREmail) Name() string {
	return "evoucha_saas_qr_email"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *EvouchaSAASQREmail) HTMLTemplate() string {
	return `<!DOCTYPE html>
	<html lang="en"
	  style="margin: 0; padding: 0; font-size: 14px; font-family: 'Roboto', sans-serif; background-color: #f5f6fa; overflow-x: hidden; overflow-y: auto; box-sizing: border-box;">
	​
	<head>
	  <meta charset="UTF-8">
	  <meta http-equiv="X-UA-Compatible" content="IE=edge">
	  <meta name="viewport" content="width=device-width, initial-scale=1">
	</head>
	​
	<body style="margin: 0; padding: 0">
	  <div class="box" style="position:relative; max-width: 100%; background-color: #f5f6fa; padding: 1.5rem">
		<div class="box-inner" style="position: relative; background-color: #fff; width: 100%; overflow-x: hidden;">
		  <div class="content" style="width: 100%; max-width: 550px; margin: 0 auto; padding: 3rem 1rem; background-color: #fff; box-sizing: border-box;">
			<div class="brand-image" style="display: block; position: relative;">
			  <img src="{{ .Email.Body.EvouchaSAAS.PartnerLogo }}" style="width: 100%; max-width: 150px; height: auto">
			</div>
			<p style="margin-top: 50px; color: {{ .Email.Body.EvouchaSAAS.MainTextColor }}; ">Hi, {{ .Email.Body.EvouchaSAAS.EvouchaSAASMailInfo.CustomerName }}</p>
			<p style="color: {{ .Email.Body.EvouchaSAAS.MainTextColor }}">
				{{ .Email.Body.EvouchaSAAS.Greeting }}
			</p>
			<div class="box-product"
			  style="position: relative; background-color: #f5f5f5; padding: 1.5rem; display: block;">
			  <table style="width:100%">
				<tbody>
				  <tr>
					<td>Product</td>
					<td style="padding-left: 10px; color: {{ .Email.Body.EvouchaSAAS.ImportantTextColor }}">{{ .Email.Body.EvouchaSAAS.EvouchaSAASMailInfo.ProductName }}</td>
				  </tr>
				  <tr>
					<td style="padding-top: 5px;">Denom</td>
					<td style="padding-top: 5px; padding-left: 10px;"><span style="font-weight: 600;">{{ .Email.Body.EvouchaSAAS.EvouchaSAASMailInfo.Denom }}</span></td>
				  </tr>
				</tbody>
			  </table>
 				{{ if .Email.Body.EvouchaSAAS.EvouchaSAASMailInfo.QRCode }}
			 <p style="text-align: center; margin-top: 40px;">Scan QR Code or input the code to redeem</p>
			  <div style="position: relative; display: block; text-align: center;">
				<img
				  src="{{ .Email.Body.EvouchaSAAS.EvouchaSAASMailInfo.QRCode }}"
				  style="width:100%;max-width: 250px;" />
			 	 <p style="text-align: center; font-size: 17px; margin-top: 5px">iahd98hjx</p>
			  </div>
			    {{ else }}
				  <div style="position: relative; display: block; text-align: center;">
					<p style="text-align: center; font-size: 17px; margin-top: 5px; color: {{ .Email.Body.EvouchaSAAS.MainTextColor }};" > {{ .Email.Body.EvouchaSAAS.EvouchaSAASMailInfo.ActivationCode }} </p>
				  </div>
				{{ end }}
			</div>
			<p style="color: {{ .Email.Body.EvouchaSAAS.MainTextColor }};'">if you need help, send us email at, <a
				style="color: {{ .Email.Body.EvouchaSAAS.ImportantTextColor }} " href="mailto:{{ .Email.Body.EvouchaSAAS.EmailAddress }}">{{ .Email.Body.EvouchaSAAS.EmailAddress }}</a></p>
			<p style="margin-top: 35px; color: {{ .Email.Body.EvouchaSAAS.MainTextColor }};">
			  <span>Your Sincerely,</span>
			  <br />
			  <strong>{{ .Email.Body.EvouchaSAAS.PartnerName}} </strong>
			</p>
		  </div>
		</div>
	  </div>
	</body>
	​
	</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *EvouchaSAASQREmail) PlainTextTemplate() string {
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
