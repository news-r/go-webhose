package webhose

import (
	"log"

	"github.com/google/go-querystring/query"
)

// NewsClient creates a client to collect news and blog articles
type NewsClient struct {
	Token     string `url:"token"`
	Query     string `url:"q"`
	Sort      string `url:"sort,omitempty"`
	Order     string `url:"order,omitempty"`
	From      string `url:"from,omitempty"`
	TimeStamp string `url:"ts,omitempty"`
	Format    string `url:"format,omitempty"`
	Size      int    `url:"size,omitempty"`
	Accuracy  string `url:"accuracy_confidence,omitempty"`
	Highlight bool   `url:"highlight,omitempty"`
	Latest    bool   `url:"latest,omitempty"`
}

// getParameters returns encoded parameters
func getParameters(client *NewsClient) string {
	values, err := query.Values(client)
	if err != nil {
		log.Fatal(err)
	}
	return values.Encode()
}
