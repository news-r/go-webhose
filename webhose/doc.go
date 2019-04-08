/*
Package webhose provides an interface to the webhose.io API.
Install with
	go get github.com/JohnCoene/go-webhose
Get a free API key on webhose's website then create a client and get articles.
	client := &webhose.NewsClient{Token: "xxXXXX-XXxxXX-xXXxxxXX-xXXXx", Query: "go programming language"}
	news := webhose.GetArticles(client, 2) // client and number of pages of result
*/
package webhose
