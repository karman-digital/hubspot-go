package hubspotmodels

type BlogPost struct {
	ArchivedAt                    int64                    `json:"archivedAt"`
	ArchivedInDashboard           bool                     `json:"archivedInDashboard"`
	AttachedStylesheets           []interface{}            `json:"attachedStylesheets"`
	AuthorName                    string                   `json:"authorName"`
	BlogAuthorId                  string                   `json:"blogAuthorId"`
	CategoryId                    int                      `json:"categoryId"`
	ContentGroupId                string                   `json:"contentGroupId"`
	ContentTypeCategory           int                      `json:"contentTypeCategory"`
	Created                       string                   `json:"created"`
	CreatedById                   string                   `json:"createdById"`
	CurrentState                  string                   `json:"currentState"`
	CurrentlyPublished            bool                     `json:"currentlyPublished"`
	Domain                        string                   `json:"domain"`
	EnableGoogleAmpOutputOverride bool                     `json:"enableGoogleAmpOutputOverride"`
	FeaturedImage                 string                   `json:"featuredImage"`
	FeaturedImageAltText          string                   `json:"featuredImageAltText"`
	HtmlTitle                     string                   `json:"htmlTitle"`
	ID                            string                   `json:"id"`
	Language                      string                   `json:"language"`
	LayoutSections                map[string]interface{}   `json:"layoutSections"`
	MetaDescription               string                   `json:"metaDescription"`
	Name                          string                   `json:"name"`
	PostBody                      string                   `json:"postBody"`
	PostSummary                   string                   `json:"postSummary"`
	PublicAccessRules             []map[string]interface{} `json:"publicAccessRules"`
	PublicAccessRulesEnabled      bool                     `json:"publicAccessRulesEnabled"`
	PublishDate                   string                   `json:"publishDate"`
	PublishImmediately            bool                     `json:"publishImmediately"`
	RssBody                       string                   `json:"rssBody"`
	RssSummary                    string                   `json:"rssSummary"`
	Slug                          string                   `json:"slug"`
	State                         string                   `json:"state"`
	TagIds                        []int64                  `json:"tagIds"`
	Translations                  map[string]interface{}   `json:"translations"`
	Updated                       string                   `json:"updated"`
	UpdatedById                   string                   `json:"updatedById"`
	URL                           string                   `json:"url"`
	UseFeaturedImage              bool                     `json:"useFeaturedImage"`
	WidgetContainers              map[string]interface{}   `json:"widgetContainers"`
	Widgets                       map[string]interface{}   `json:"widgets"`
}

type BlogPostsResponse struct {
	Total   int        `json:"total"`
	Results []BlogPost `json:"results"`
}

type BlogFilterOptions struct {
	CreatedAt     string            `url:"createdAt,omitempty"`     // Return blog posts created at this exact time
	CreatedAfter  string            `url:"createdAfter,omitempty"`  // Return blog posts created after this time
	CreatedBefore string            `url:"createdBefore,omitempty"` // Return blog posts created before this time
	UpdatedAt     string            `url:"updatedAt,omitempty"`     // Return blog posts updated at this exact time
	UpdatedAfter  string            `url:"updatedAfter,omitempty"`  // Return blog posts updated after this time
	UpdatedBefore string            `url:"updatedBefore,omitempty"` // Return blog posts updated before this time
	Sort          string            `url:"sort,omitempty"`          // Fields to sort by (createdAt, name, updatedAt, createdBy, updatedBy)
	After         string            `url:"after,omitempty"`         // Cursor token for next page of results
	Limit         int               `url:"limit,omitempty"`         // Maximum number of results to return (default: 20)
	Filters       map[string]string `url:"filters,omitempty"`       // Filters to apply to the results
}
