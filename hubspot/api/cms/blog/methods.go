package blogs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	blogmodels "github.com/karman-digital/hubspot/hubspot/api/models/cms/blogs"
)

func (b *BlogService) GetAllBlogPosts(opts blogmodels.BlogFilterOptions) (blogmodels.BlogPostsResponse, error) {
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
	path := "/cms/v3/blogs/posts"
	if encoded := queryParams.Encode(); encoded != "" {
		path = fmt.Sprintf("%s?%s", path, encoded)
	}
	resp, err := b.SendRequest(http.MethodGet, path, nil)
	if err != nil {
		return blogmodels.BlogPostsResponse{}, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return blogmodels.BlogPostsResponse{}, fmt.Errorf("error reading response body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return blogmodels.BlogPostsResponse{}, fmt.Errorf("error getting blog posts: %s", resp.Status)
	}
	var blogPostsResponse blogmodels.BlogPostsResponse
	err = json.Unmarshal(body, &blogPostsResponse)
	if err != nil {
		return blogmodels.BlogPostsResponse{}, fmt.Errorf("error unmarshalling response body: %s", err)
	}
	return blogPostsResponse, nil
}
