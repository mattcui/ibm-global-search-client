package main

import (
	"context"
	"fmt"

	igsc "github.com/mattcui/ibm-global-search-client"
)

func main() {
	token := "eyJraWQiOiIyMDIxMDYxOTE4MzciLCJhbGciOiJSUzI1NiJ9.eyJpYW1faWQiOiJJQk1pZC0wNjAwMDA1OTBEIiwiaWQiOiJJQk1pZC0wNjAwMDA1OTBEIiwicmVhbG1pZCI6IklCTWlkIiwic2Vzc2lvbl9pZCI6IkMtYjYzNWY2MDYtYTJkMy00NTk1LTg5NGItY2UxMmIyMDhmMzIzIiwianRpIjoiNTM4NzE4OWMtNDE3Yy00NDBmLWJkNWEtNWFkZTg4OTQ3ZGQ4IiwiaWRlbnRpZmllciI6IjA2MDAwMDU5MEQiLCJnaXZlbl9uYW1lIjoiWFVFIFhJQU5HIiwiZmFtaWx5X25hbWUiOiJDVUkiLCJuYW1lIjoiWFVFIFhJQU5HIENVSSIsImVtYWlsIjoiY3VpeHVleEBjbi5pYm0uY29tIiwic3ViIjoiY3VpeHVleEBjbi5pYm0uY29tIiwiYXV0aG4iOnsic3ViIjoiY3VpeHVleEBjbi5pYm0uY29tIiwiaWFtX2lkIjoiSUJNaWQtMDYwMDAwNTkwRCIsIm5hbWUiOiJYVUUgWElBTkcgQ1VJIiwiZ2l2ZW5fbmFtZSI6IlhVRSBYSUFORyIsImZhbWlseV9uYW1lIjoiQ1VJIiwiZW1haWwiOiJjdWl4dWV4QGNuLmlibS5jb20ifSwiYWNjb3VudCI6eyJib3VuZGFyeSI6Imdsb2JhbCIsInZhbGlkIjp0cnVlLCJic3MiOiJjMjA1NDQ2YTY0ODdjMDEwNGVlOTU2MzE0YjY5MTg3MyIsImltc191c2VyX2lkIjoiNzMzMzExMSIsImZyb3plbiI6dHJ1ZSwiaW1zIjoiMTc1OTIxMyJ9LCJpYXQiOjE2MjUyODI5MjEsImV4cCI6MTYyNTI4NDEyMSwiaXNzIjoiaHR0cHM6Ly9pYW0uY2xvdWQuaWJtLmNvbS9pZGVudGl0eSIsImdyYW50X3R5cGUiOiJ1cm46aWJtOnBhcmFtczpvYXV0aDpncmFudC10eXBlOmFwaWtleSIsInNjb3BlIjoiaWJtIG9wZW5pZCIsImNsaWVudF9pZCI6ImJ4IiwiYWNyIjoxLCJhbXIiOlsicHdkIl19.CXDU7x4qzznUtfbofVL6e9gnPdi-3TyIWfgQ5kZknQz20RHW0um7p8eTjfbK4myRAn8pB2KlGr5ebdlxEYifOITfqs1dadEXhqZAzgXD-D86Jnyq9P3qgCTNF_xqFzn3xciuUivbpBgYQs8n5t24gVT5Viw1SmzSYpOLLta47FngjEglRS0KQOoYc_awmNuG4zRi3KptX7dP2UsImfurcXQqqhNrNSga9NTqjwJ34Q9QuRbF3mZ5ZbGxj9_CMWc58_YAcy1Stkks5U1RWn4b4u3kyaiEFl7LYs97Am4A6mtAUve7dZyFuyOEpf39yDNoO4E8s1-dN3ha4aYkwe4PGA"
	ctx := context.WithValue(context.Background(), "accesstoken", token)

	cfg := igsc.NewConfiguration()
	client := igsc.NewAPIClient(cfg)

	searchBody := igsc.NewSearchBodyWithDefaults()

	searchApiV2SearchOpts := igsc.SearchApiV2SearchOpts{
		
	}

	searchResult, res, _ := client.SearchApi.V2Search(ctx, searchBody, searchApiV2SearchOpts)
	fmt.Printf("result=%v response=%v", searchResult, res)

}
