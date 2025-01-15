# \VersionAPI

All URIs are relative to *https://ossindex.sonatype.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](VersionAPI.md#Get) | **Get** /api/v3/version | Get service version information



## Get

> Get(ctx).Execute()

Get service version information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	ossindex "github.com/sonatype-nexus-community/ossindex-api-client-go"
)

func main() {

	configuration := ossindex.NewConfiguration()
	apiClient := ossindex.NewAPIClient(configuration)
	r, err := apiClient.VersionAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

