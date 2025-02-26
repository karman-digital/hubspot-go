package blogs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (b *BlogService) GetAllBlogPosts(opts hubspotmodels.BlogFilterOptions) (hubspotmodels.BlogPostsResponse, error) {
	req, err := retryablehttp.NewRequest(http.MethodGet, "https://api.hubapi.com/cms/v3/blogs/posts", nil)
	if err != nil {
		return hubspotmodels.BlogPostsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.AccessToken()))
	queryParams := url.Values{}
	if opts.Filters != nil {
		for key, value := range opts.Filters {
			queryParams.Add(key, value)
		}
	}
	if opts.Sort != "" {
		queryParams.Add("sort", opts.Sort)
	}
	if opts.Limit != 0 {
		queryParams.Add("limit", strconv.Itoa(opts.Limit))
	}
	if opts.After != "" {
		queryParams.Add("after", opts.After)
	}
	if opts.State != "" {
		queryParams.Add("state", opts.State)
	}
	req.URL.RawQuery = queryParams.Encode()
	fmt.Println(req.URL.String())
	fmt.Println(req)
	resp, err := b.Client().Do(req)
	if err != nil {
		return hubspotmodels.BlogPostsResponse{}, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return hubspotmodels.BlogPostsResponse{}, fmt.Errorf("error reading response body: %s", err)
	}
	fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		return hubspotmodels.BlogPostsResponse{}, fmt.Errorf("error getting blog posts: %s", resp.Status)
	}
	var blogPostsResponse hubspotmodels.BlogPostsResponse
	err = json.Unmarshal(body, &blogPostsResponse)
	if err != nil {
		return hubspotmodels.BlogPostsResponse{}, fmt.Errorf("error unmarshalling response body: %s", err)
	}
	return blogPostsResponse, nil
}
