package hermes

// Flat is a theme
type BIVoucherReceipt struct{}

// Name returns the name of the flat theme
func (dt *BIVoucherReceipt) Name() string {
	return "bi_voucher_receipt"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *BIVoucherReceipt) HTMLTemplate() string {
	return `
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<title>Document</title>
	<link rel="preconnect" href="https://fonts.gstatic.com">
	<link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">
	<link rel="preconnect" href="https://fonts.gstatic.com"/>
	<link href="https://fonts.googleapis.com/css2?family=Roboto&amp;display=swap" rel="stylesheet"/>

	<style type="text/css" rel="stylesheet" media="all">
		.data-table {
      width: 100%;
      margin: 0;
			padding : 10px;
    }
    .data-table th {
      text-align: left;
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid #EDEFF2;
    }
    .data-table th p {
      margin: 0;
      color: #9BA2AB;
      font-size: 12px;
    }
    .data-table td {
      padding: 10px 5px;
      color: #74787E;
      font-size: 15px;
      line-height: 18px;
    }

    .button {
      display: inline-block;
      background-color: #3869D4;
      border-radius: 3px;
      color: #ffffff !important;
      font-size: 15px;
      line-height: 45px;
      text-align: center;
      text-decoration: none;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }

		.color-bi {
			color: #c70773;
		}
		.color-success {
			color: #1da122;
		}
		.color-process {
			color: #FF9300;	
		}
		.color-failed {
			color: #FF2A00;
		}
	</style>
</head>

<body style="font-family:Roboto, sans-serif">
	<div class="container-content" style="max-width:910px;margin:auto">
		<div class="container" style="max-width:800px;margin:auto;color:#333333">
            
			<img class="company-image-header" src="{{ .Email.Body.VoucherReceipt.Logo }}" alt="" srcset="" style="margin-top:20px;width:151px;height:auto" width="151"/>
			<h3 style="color:#646464">Dear {{ .Email.Body.VoucherReceipt.Name }}</h1>
			
			<h2 style="color:#3a3434">{{ .Email.Body.VoucherReceipt.SubTitle }}</h2>

			{{ range $row := .Email.Body.VoucherReceipt.Intros }}
			<p style="font-size:14px; color:#333333">{{ $row }}</p>
    		{{ end }}
			
			<div style="background-color:#fcfcfc;width:95%;padding:10px">
				 	<table class="data-table" cellpadding="0" cellspacing="0">
						<tr>
							<td>Merchant</td>
							<td style="color:#c70773;">{{ .Email.Body.VoucherReceipt.Voucher.Merchant }}</td>
						</tr>
						<tr>
							<td>Voucher Code</td>
							<td>{{ .Email.Body.VoucherReceipt.Voucher.Code }}</td>
						</tr>
						<tr>
							<td>Voucher Redeemed</td>
							<td>{{ .Email.Body.VoucherReceipt.Voucher.Name }}</td>
						</tr>
						<tr>
							<td>Value</td>
							<td>{{ .Email.Body.VoucherReceipt.Voucher.Value }} {{ .Email.Body.VoucherReceipt.Voucher.Currency }}</td>
						</tr>
						<tr>
							<td>Redeemed Time</td>
							<td>{{ .Email.Body.VoucherReceipt.Voucher.RedeemTime }}</td>
						</tr>
						<tr>
							<td>Status</td>
							<td style="
									{{ if eq (.Email.Body.VoucherReceipt.Voucher.StatusCode) 0}}
										color:#FF9300;
									{{ else if eq (.Email.Body.VoucherReceipt.Voucher.StatusCode) 1}}
										color:#1da122;
									{{ if eq (.Email.Body.VoucherReceipt.Voucher.StatusCode) 2}}
										color:#FF2A00;
									{{end}}
									{{end}}
							">{{ .Email.Body.VoucherReceipt.Voucher.Status }}</td>
						</tr>
					</table>
			</div>

		 <p style="font-size:14px; color:#333333">{{ .Email.Body.VoucherReceipt.Help }} <a href="mailto:{{ .Email.Body.VoucherReceipt.HelpEmail }}">{{ .Email.Body.VoucherReceipt.HelpEmail }}</a></p>
			

		<!-- Action -->
		{{ with .Email.Body.Actions }}
			{{ if gt (len .) 0 }}
				{{ range $action := . }}
					<p>{{ $action.Instructions }}</p>
					{{ $length := len $action.Button.Text }}
					{{ $width := add (mul $length 9) 20 }}
					{{if (lt $width 200)}}
						{{$width = 200}}
					{{else if (gt $width 570)}}
						{{$width = 570}}
					{{else}}
					{{end}}
						{{safe "<!--[if mso]>" }}
						{{ if $action.Button.Text }}
							<div style="margin: 30px auto;v-text-anchor:middle;text-align:center">
								<v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" 
									xmlns:w="urn:schemas-microsoft-com:office:word" 
									href="{{ $action.Button.Link }}" 
									style="height:45px;v-text-anchor:middle;width:{{$width}}px;background-color:{{ if $action.Button.Color }}{{ $action.Button.Color }}{{ else }}#3869D4{{ end }};"
									arcsize="10%" 
									{{ if $action.Button.Color }}strokecolor="{{ $action.Button.Color }}" fillcolor="{{ $action.Button.Color }}"{{ else }}strokecolor="#3869D4" fillcolor="#3869D4"{{ end }}
									>
									<w:anchorlock/>
									<center style="color: {{ if $action.Button.TextColor }}{{ $action.Button.TextColor }}{{else}}#FFFFFF{{ end }};font-size: 15px;text-align: center;font-family:sans-serif;font-weight:bold;">
										{{ $action.Button.Text }}
									</center>
								</v:roundrect>
							</div>
						{{ end }}
						{{ if $action.InviteCode }}
							<div style="margin-top:30px;margin-bottom:30px">
								<table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
									<tr>
										<td align="center">
											<table align="center" cellpadding="0" cellspacing="0" style="padding:0;text-align:center">
												<tr>
													<td style="display:inline-block;border-radius:3px;font-family:Consolas, monaco, monospace;font-size:28px;text-align:center;letter-spacing:8px;color:#555;background-color:#eee;padding:20px">
														{{ $action.InviteCode }}
													</td>
												</tr>
											</table>
										</td>
									</tr>
								</table>
							</div>
						{{ end }}   
						{{safe "<![endif]-->" }}
						{{safe "<!--[if !mso]><!-- -->"}}
						<table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
							<tr>
								<td align="center">
									<div>
										{{ if $action.Button.Text }}
											<a href="{{ $action.Button.Link }}" class="button" style="{{ with $action.Button.Color }}background-color: {{ . }};{{ end }} {{ with $action.Button.TextColor }}color: {{ . }};{{ end }} width: {{$width}}px;" target="_blank">
												{{ $action.Button.Text }}
											</a>
										{{end}}
										{{ if $action.InviteCode }}
											<span class="invite-code">{{ $action.InviteCode }}</span>
										{{end}}
									</div>
								</td>
							</tr>
						</table>
						{{safe "<![endif]-->" }}
				{{ end }}
			{{ end }}
		{{ end }}

			<div class="signature" style="margin-top:50px;line-height:10px; color:#333333; font-size:14px">
				<p>Best Regards,</p>
				<p>TruRewards</p>
			</div>
		</div>

		<div class="footer" style="margin-top:50px;background-color:#fff;width:100%;min-height:170px;padding:10px;border-top: 1px solid #c8c8c8" height="170">
			<div class="container" style="max-width:800px;color:#333333">
				<div class="row">
          <h4 style="color:#646464">Contact Us</h4>
        </div>
        <div class="row">
					<div class="column" style="max-width:300px;float:left;font-size:12px;color:#ffffff;padding-right:50px">
						<div style="display: inline-flex;align-items: top;justify-content: center;">
							<img src="https://storage.googleapis.com/bi_icon/map-marker-alt-solid.png" 
								style="width: auto; height: 17px;color:#646464;margin-right:5px; padding-top:16px"/>
              <p style="font-size:14px;color:#646464">{{ .Email.Body.VoucherReceipt.Contact.Address }}</p>
						</div>
					</div>
					<div class="column" style="min-width:250px;float:left;font-size:12px;color:#ffffff">
            
            <p class="pi" style="font-size:14px;text-decoration:none;color:#fff"><img src="https://storage.googleapis.com/bi_icon/envelope-solid.png" 
							width="13px" style="font-size:14px;color:#646464; margin-right:4px"/> <a href="mailto:{{ .Email.Body.VoucherReceipt.Contact.Email }}" target="_blank" style="text-decoration: none; color:#646464">{{ .Email.Body.VoucherReceipt.Contact.Email }}</a></p>
						
            <p class="pi" style="font-size:14px;color:#646464"><img src="https://storage.googleapis.com/bi_icon/phone-solid.png" 
							width="13px" style="margin-right:8px"></span>{{ .Email.Body.VoucherReceipt.Contact.PhoneNumber }}</p>
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
func (dt *BIVoucherReceipt) PlainTextTemplate() string {
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
