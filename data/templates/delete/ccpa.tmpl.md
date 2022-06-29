Dear Privacy Compliance Officer,

My name is {{.User.FirstName}} {{.User.LastName}}. I reside in California and am exercising my right to delete my personal information under
the California Consumer Privacy Act. I request that {{.Org.OrganizationName}} deletes all of the information
it has collected about me, whether directly from me, through a third party, or through a service provider.

I use the following email addresses:{{range .User.EmailAddresses}}
- {{.}}
{{end}}

My phone numbers are:{{range .User.PhoneNumbers}}
- {{.}}
{{end}}

If you need any more information from me, please let me know as soon as possible. If you cannot comply
with my request–either in whole or in part–please state the reason why you cannot comply. If part of my information
is subject to an exception, please delete all information that is not subject to an exception.
If my request is incomplete, please provide me with specific instructions on how to complete my request.

Sincerely,

{{.User.FirstName}} {{.User.LastName}}

{{.Date}}