# k8s-apim-operator Scenarios

## Scenario 5

> ##### This scenario demonstrates ratelimiting (throttling) with APIM

- Follow the main README and deploy the apim-operator and configuration files. Make sure to set the analyticsEnabled to "true" and deploy analytics secret with credentials to analytics server and certificate, if you want to check analytics
 
##### Navigate to the scenarios/scenario-5 directory and execute the following command

- Deploy ratelimiting kind <br /> 
    - ***apimcli apply -f ratelimiting_instances.yaml***
    - Note: Deploying ratelimiting custom resource to enforce 5 requests per min

- Create API <br /> 
    - ***apimcli add api -n petstore --from-file=petstore_basic.yaml***

- Update API <br /> 
    - ***apimcli update api -n petstore --from-file=petstore_basic.yaml***
    
- Get available API <br /> 
    - ***apimcli get apis***

- Get service details to invoke the API<br />
    - ***apimcli get services***
    - Note: Get the external IP of the service
 
- Invoking the API <br />
    - ***TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6Ik5UQXhabU14TkRNeVpEZzNNVFUxWkdNME16RXpPREpoWldJNE5ETmxaRFUxT0dGa05qRmlNUT09In0=.eyJhdWQiOiJodHRwOlwvXC9vcmcud3NvMi5hcGltZ3RcL2dhdGV3YXkiLCJzdWIiOiJhZG1pbkBjYXJib24uc3VwZXIiLCJhcHBsaWNhdGlvbiI6eyJvd25lciI6ImFkbWluIiwidGllciI6IlVubGltaXRlZCIsIm5hbWUiOiJKV1RBcHAiLCJpZCI6Mn0sInNjb3BlIjoiYW1fYXBwbGljYXRpb25fc2NvcGUgZGVmYXVsdCIsImlzcyI6Imh0dHBzOlwvXC93c28yYXBpbTo5NDQzXC9vYXV0aDJcL3Rva2VuIiwidGllckluZm8iOnt9LCJrZXl0eXBlIjoiUFJPRFVDVElPTiIsInN1YnNjcmliZWRBUElzIjpbXSwiY29uc3VtZXJLZXkiOiI1QkRSdTNMTG14TlgxUllRaTBfUGpwcExkeEFhIiwiZXhwIjoxNTY4OTA2MzEwLCJpYXQiOjE1Njg5MDI3MTAsImp0aSI6ImU5NzA5MmUyLThlNTktNGYyYy1iYmRiLTAyMmI2ZWI3MzVjMCJ9.ah1h5Gjy5m3CrmMsfaylFiMYjBQW7ndB1GDNHJk08FKzOpxBikEjgaOnoc0yGe43v6xMo7OZZVhvE3hJBfY7S86-seCl-zDIKPHOMcFTM-AYpLY53Yorex-JK1-zd75fi1qwXfar9fI1IE0skSIyR0JJXJC7j228Ho-yaa1NlXVnjC9Y2mVa_zPvXsgYMBVgbsOBCANhYmLMEYOgA5zSdHbEue8ROWw5HU3psW20CU8T4hxvhDlBklLfPcYD961tmBdzymUy4ywk2VxS4jjGWhwI6_bSB2L78QeiM48oS2tE3MRqjRHgEo5vCNx2fj1hrHXkXX8ozpRFB9Y_YTH9sQ==***
   
    - ***curl -X GET "https://\<external IP of LB service>:9095/petstore/v1/pet/55" -H "accept: application/xml" -H "Authorization:Bearer $TOKEN" -k***
    - Note: When you invoke more than 5 times continuously, the requestes will be throttled out after the 5th request
     
- Delete API <br /> 
    - ***apimcli delete api petstore***