
## In Class Assignment

Make a GIN API that takes a post request `/run`. This request runs python code and sends it back. The request body should contain the following JSON:
```json
{
	"code": "<python code here>"
}
```

and the response will be:

```json
{
	"output": "<output of the python code here>"
}
```

