# email-service

This service is a frontend to Amazon's managed SES email offering.

There are only two endpoints, `/` and `/health`. The health endpoint is used by the cluster's readiness and liveness probes.

To send an email, you send a POST request with a simple JSON payload structured as follows:

```json
{
    "recipients": [ "firstname.lastname@acme.com", "anothername.lastname@acme.com" ],
    "subject": "Some subject that suits an automated message",
    "markdownBody": "Some *text* with basic **formatting**"
}
```

HTML-capable email clients will display HTML text based on the Markdown source; text-only clients display the Markdown source itself.

The sender is always as configured in the Helm Values file, `chart/values.yaml`. By default, it is `noreply@acme.com`.

Please note that recipients need to accept an invitation to take part before they can receive messages.

## Calling the service
Any system that can send POST requests will do, but if you happen to use `curl`, you might enter:

```
curl -XPOST http://email-service.default.svc.cluster.local:8080/ \
  --header "Content-Type: application/json" \
  -d '{ "subject": "The Red Wheelbarrow", "recipients": ["gerald.schmidt@acme.com"], "markdownBody": "So much depends upon a **red** ~~tractor~~wheelbarrow."}'
```
