# goCurl

A simple tool that allows you to make HTTP requests from the command line providing a textfile with the requests.

## Usage

```bash
goCurl -f <file>
```

## Sample httpFile 

```plaintext
GET http://localhost:8080/getToken
{
    "user_name":"Thomas",
    "password":"password",
    "app_id":2
}
```
