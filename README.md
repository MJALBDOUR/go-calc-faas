# go-calc-faas
Learning GCP FaaS by implementing a simple calculator logic ( Add | Sub | Mul | Div )

## Available Functions
* Add2
* Sub2
* Mul2
* Div2

## How funcs are deployed
using gcloud sdk, I enter the command:
gcloud functions deploy <*function_name*> --runtime=go113 --trigger-http --allow-unauthenticated

However, in the future, using *cloudbuild.yaml* + build *triggers* would be way more *efficient*
*example*
steps:
- name: 'gcr.io/google.com/cloudsdktool/cloudsdk'
  args:
  - functions
  - deploy
  - <*function_name*>
  - --runtime=go113
  - --region=<*function_region*>
  - --source=<*function_source_code*>
  - --trigger-http

## Func Call Example
### CURL
curl -X POST <*function_url*> -H "Content-Type:application/json" -d '{"a": <*float64*>, "b": <*float64*>}'

### POSTMAN
* METHOD: POST
* URL: <function_url>
* HEADERS: KEY:Content-type VALUE: application/json
* BODY: raw --> {"a": <*float64*>, "b": <*float64*>}
* SEND

### EXPECTED RESPONSES -- no *json* for now
* A + B = %v | Invalid input + Bad Request
* A - B = %v | Invalid input + Bad Request
* A * B = %v | Invalid input + Bad Request
* A / B = %v | Invalid input + Bad Request | Division by zero
