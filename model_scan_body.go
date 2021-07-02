/*
 * Global Search
 *
 * Search for resources with the global and shared resource properties repository integrated in the IBM Cloud platform. The search repository stores and searches cloud resources attributes, which categorize or classify resources. A resource is a physical or logical component that can be created or reserved for an application or service instance and is owned by resource providers, such as Cloud Foundry, IBM Kubernetes Service, or resource controller in IBM Cloud. Resources are uniquely identified by a Cloud Resource Name (CRN)  or by an IMS ID. The properties of a resource include tags and system properties. Both properties are defined in an IBM Cloud billing account, and span across many regions.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request body associated with a search request.
type ScanBody struct {
	// The Lucene-formatted query string. Default to '*' if not set
	Query string `json:"query,omitempty"`
	// The list of the fields returned by the search. Defaults to all. `crn` is always returned.
	Fields []string `json:"fields,omitempty"`
	// An opaque search cursor that is returned on each operation call and that must be set on next call.
	SearchCursor string `json:"search_cursor,omitempty"`
}