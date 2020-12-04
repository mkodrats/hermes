package hermes

// Flat is a theme
type BITemporaryPassword struct{}

// Name returns the name of the flat theme
func (dt *BITemporaryPassword) Name() string {
	return "bi_temporary_password"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *BITemporaryPassword) HTMLTemplate() string {
	return `
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
	<link rel="preconnect" href="https://fonts.gstatic.com">
	<link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">
	<link rel="preconnect" href="https://fonts.gstatic.com"/>
	<link href="https://fonts.googleapis.com/css2?family=Roboto&amp;display=swap" rel="stylesheet"/>
	<link href="//netdna.bootstrapcdn.com/twitter-bootstrap/2.3.2/css/bootstrap-combined.no-icons.min.css" rel="stylesheet"/>
	<link href="//netdna.bootstrapcdn.com/font-awesome/3.2.1/css/font-awesome.css" rel="stylesheet"/>
</head>

<body style="font-family:Roboto, sans-serif">
	<div class="container-content" style="max-width:910px;margin:auto">
		<div class="container" style="max-width:800px;margin:auto;color:#333333">
            
			<img class="company-image-header" src="{{ .Email.Body.Registration.Logo }}" alt="" srcset="" style="margin-top:20px;width:151px;height:auto" width="151"/>
			<h3 style="color:#646464">Dear {{ .Email.Body.Registration.Name }}</h1>
 			
			{{ range $row := .Email.Body.Registration.Intros }}
			<p style="font-size:14px; color:#333333">{{ $row }}</p>
    		{{ end }}
    		
			<br/>
			
			<div style="background-color:#FBEEF2;width:95%;padding:10px">
				 <p style="color:#333333">Username : {{.Email.Body.Registration.Username}}</p>
				 <p style="color:#333333">Password : {{.Email.Body.Registration.TemporaryPassword}}</p>
			</div>

		 <p style="font-size:14px; color:#333333">{{ .Email.Body.Registration.Help }}</p>

			
			<div class="signature" style="margin-top:50px;line-height:10px; color:#333333; font-size:14px">
				<p>Best Regards,</p>
				<p>TruRewards</p>
			</div>
		</div>

		<div class="footer" style="margin-top:50px;background-color:#fff;width:95%;min-height:170px;padding:10px;border-top: 1px solid #c8c8c8" height="170">
			<div class="container" style="max-width:800px;color:#333333">
				<div class="row">
          <h4 style="color:#646464">Contact Us</h4>
        </div>
        <div class="row">
					<div class="column" style="max-width:300px;float:left;font-size:12px;color:#ffffff;padding-right:50px">
						<div style="display: inline-flex;align-items: top;justify-content: center;">
							<img src="https://storage.googleapis.com/bi_icon/map-marker-alt-solid.png" 
								style="width: auto; height: 17px;color:#646464;margin-right:5px; padding-top:16px"/>
              <p style="font-size:14px;color:#646464">{{ .Email.Body.Registration.Contact.Address }}</p>
						</div>
					</div>
					<div class="column" style="min-width:250px;float:left;font-size:12px;color:#ffffff">
            
            <p class="pi" style="font-size:14px;text-decoration:none;color:#fff"><img src="https://storage.googleapis.com/bi_icon/envelope-solid.png" 
							width="13px" style="font-size:14px;color:#646464; margin-right:4px"/> <a href="mailto:{{ .Email.Body.Registration.Contact.Email }}" target="_blank" style="text-decoration: none; color:#646464">{{ .Email.Body.Registration.Contact.Email }}</a></p>
						
            <p class="pi" style="font-size:14px;color:#646464"><img src="https://storage.googleapis.com/bi_icon/phone-solid.png" 
							width="13px" style="margin-right:8px"></span>{{ .Email.Body.Registration.Contact.PhoneNumber }}</p>
					</div>
				</div>
			</div>
			<div class="hhh"></div>
			<div style="margin-top:10px;min-width:100%;text-align:center;font-size:12px;color:#ffffff;display: flow-root;">
			</div>
		</div>
	</div>

</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *BITemporaryPassword) PlainTextTemplate() string {
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
