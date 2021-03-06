/*
 * Global Search
 *
 * Search for resources with the global and shared resource properties repository integrated in the IBM Cloud platform. The search repository stores and searches cloud resources attributes, which categorize or classify resources. A resource is a physical or logical component that can be created or reserved for an application or service instance and is owned by resource providers, such as Cloud Foundry, IBM Kubernetes Service, or resource controller in IBM Cloud. Resources are uniquely identified by a Cloud Resource Name (CRN)  or by an IMS ID. The properties of a resource include tags and system properties. Both properties are defined in an IBM Cloud billing account, and span across many regions.
 *
 * API version: 2.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SearchApiService SearchApi service
type SearchApiService service

type ApiSearchRequest struct {
	ctx _context.Context
	ApiService *SearchApiService
	body *ScanBody
	transactionId *string
	accountId *string
	limit *int32
	timeout *int32
	sort *[]string
}

func (r ApiSearchRequest) Body(body ScanBody) ApiSearchRequest {
	r.body = &body
	return r
}
func (r ApiSearchRequest) TransactionId(transactionId string) ApiSearchRequest {
	r.transactionId = &transactionId
	return r
}
func (r ApiSearchRequest) AccountId(accountId string) ApiSearchRequest {
	r.accountId = &accountId
	return r
}
func (r ApiSearchRequest) Limit(limit int32) ApiSearchRequest {
	r.limit = &limit
	return r
}
func (r ApiSearchRequest) Timeout(timeout int32) ApiSearchRequest {
	r.timeout = &timeout
	return r
}
func (r ApiSearchRequest) Sort(sort []string) ApiSearchRequest {
	r.sort = &sort
	return r
}

func (r ApiSearchRequest) Execute() (ScanResult, *_nethttp.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
 * Search Find instances of resources (v3)
 * Find Cloud Foundry resources, IAM-enabled resources, or 
storage and network resources running on classic infrastructure in a 
specific account ID. You can apply query strings if necessary. 

To filter results, you can insert a string using the Lucene syntax and the 
query string is parsed into a series of terms and operators. A term can be 
a single word or a phrase, in which case the search is performed for all 
the words, in the same order. To filter for a specific value regardless of 
the property that contains it, type the search term without specifying a 
field. Only resources that belong to the account ID and that are accessible 
by the client are returned. 

You must use `/v3/resources/search` when you need to fetch more than `10000` 
resource items. The `/v2/resources/search` prohibits paginating through such 
a big number. On the first call, the operation returns a live cursor on the 
data that you must use on all the subsequent calls to get the next batch of 
results until you get the empty result set. By default, the fields returned 
for every resource are "crn", "name", "family", "type", and "account_id". You 
can specify the subset of the fields you want in your request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchRequest
 */
func (a *SearchApiService) Search(ctx _context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ScanResult
 */
func (a *SearchApiService) SearchExecute(r ApiSearchRequest) (ScanResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScanResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.Search")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/resources/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.accountId != nil {
		localVarQueryParams.Add("account_id", parameterToString(*r.accountId, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.timeout != nil {
		localVarQueryParams.Add("timeout", parameterToString(*r.timeout, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.transactionId != nil {
		localVarHeaderParams["transaction-id"] = parameterToString(*r.transactionId, "")
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["IAM"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV2SearchRequest struct {
	ctx _context.Context
	ApiService *SearchApiService
	body *SearchBody
	transactionId *string
	accountId *string
	limit *int32
	timeout *int32
	offset *int32
	sort *string
	provider *string
}

func (r ApiV2SearchRequest) Body(body SearchBody) ApiV2SearchRequest {
	r.body = &body
	return r
}
func (r ApiV2SearchRequest) TransactionId(transactionId string) ApiV2SearchRequest {
	r.transactionId = &transactionId
	return r
}
func (r ApiV2SearchRequest) AccountId(accountId string) ApiV2SearchRequest {
	r.accountId = &accountId
	return r
}
func (r ApiV2SearchRequest) Limit(limit int32) ApiV2SearchRequest {
	r.limit = &limit
	return r
}
func (r ApiV2SearchRequest) Timeout(timeout int32) ApiV2SearchRequest {
	r.timeout = &timeout
	return r
}
func (r ApiV2SearchRequest) Offset(offset int32) ApiV2SearchRequest {
	r.offset = &offset
	return r
}
func (r ApiV2SearchRequest) Sort(sort string) ApiV2SearchRequest {
	r.sort = &sort
	return r
}
func (r ApiV2SearchRequest) Provider(provider string) ApiV2SearchRequest {
	r.provider = &provider
	return r
}

func (r ApiV2SearchRequest) Execute() (SearchResults, *_nethttp.Response, error) {
	return r.ApiService.V2SearchExecute(r)
}

/*
 * V2Search Find instances of resources (v2)
 * Find resources in the internal search platform for a specific account
ID. You can optionally apply query strings. To filter results, you can insert
a string using the Lucene syntax and the query string is parsed into a series
of terms and operators. A term can be a single word or a phrase, in which
case the search is performed for all the words, in the same order. To filter
for a specific value regardless of the property that contains it, use an asterisk
as the key name. Results are filtered by the account ID and IAM subsystem.
Only resources that belong to the account ID and that are accessible by the
client are returned.
The usage of this API is deprecated for cases when the provider is not specified or 
`provider=ghost`. 
`/v3/resources/search` replaces this API for the deprecated cases.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiV2SearchRequest
 */
func (a *SearchApiService) V2Search(ctx _context.Context) ApiV2SearchRequest {
	return ApiV2SearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchResults
 */
func (a *SearchApiService) V2SearchExecute(r ApiV2SearchRequest) (SearchResults, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.V2Search")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/resources/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.accountId != nil {
		localVarQueryParams.Add("account_id", parameterToString(*r.accountId, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.timeout != nil {
		localVarQueryParams.Add("timeout", parameterToString(*r.timeout, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.transactionId != nil {
		localVarHeaderParams["transaction-id"] = parameterToString(*r.transactionId, "")
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["IAM"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
