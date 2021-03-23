package hermes

// Flat is a theme
type SengkuniTemporaryPassword struct{}

// Name returns the name of the flat theme
func (dt *SengkuniTemporaryPassword) Name() string {
	return "sengkuni_temporary_password"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *SengkuniTemporaryPassword) HTMLTemplate() string {
	return `
<!DOCTYPE html><html lang="en"><head>
    <link href='https://fonts.googleapis.com/css?family=Muli' rel='stylesheet'>
    <link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'>
    <link href='https://fonts.googleapis.com/css?family=Montserrat' rel='stylesheet'>

	<meta charset="UTF-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	<title>Document</title>
</head>



<body>
	<div style="max-width:910px;margin:auto">
		<div style="max-width:800px;margin:auto;color:#333333;font-family:'Muli';">
            
			<img  src="{{ .Email.Body.Registration.Logo }}" style="margin-top:20px;width:100px;height:auto"/>
            <h1>Hi {{ .Email.Body.Registration.Name }},</h1>
       
 			
			<p>Congratulations, your account has been created. You may now Sign In using following credentials:</p>
    		
            
            <div style="background-color:#eef3fb;width:95%;padding:10px">
                <p style="color:#333333;text-indent: 10px;"> Email : {{.Email.Body.Registration.EmailAddress}}</p>
								<p style="color:#333333;text-indent: 10px;"> Password : {{.Email.Body.Registration.TemporaryPassword}}</p>
           </div>
		
			<p>For security reason, please update your password in the account menu after Sign In.</p>
			<div style="margin-top:50px;line-height:10px">
				<p>Best Regards,</p>
				<p>Authscure</p>
			</div>
        </div>
        
        <br>

		<div style="margin: auto;background-color:#00b0f0;width:100%;height:150px;padding:10px" height="170">
            <div style="font-family:Muli;color: #ffffff; text-align: center;font-size: 14px;">
                <p>This message is automatically generated by Authscure. You received this message because this email address is associated with your Authscure account. </p>
            </div>

            <br>

            <div style="width: 910px;height: 0.5px;margin: 6.5px 0;opacity: 0.5;background-color: #ffffff;"></div>

            <br>

            <div style="font-family: Montserrat;text-align: center;color: #ffffff;font-size: 14px;">Authscure ©2021. All Rights Reserved</div>
		</div>
	</div>



</body></html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *SengkuniTemporaryPassword) PlainTextTemplate() string {
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
