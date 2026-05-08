package main

// Tokens holds the jira API token used
type Tokens struct {
	Jira string `json:"jira"`
}

// Jira builds the Jira API address and update source
type Jira struct {
	URL      string `json:"url"`
	Source   string `json:"source"`
	Summary  string `json:"summary"`
	Criteria string `json:"criteria"`
}

// Changelogs builds a collection of urls to target changelogs
type Changelogs struct {
	ACF       string `json:"acf"`
	Calendar  string `json:"calendar"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	Spotlight string `json:"spotlight"`
	Tickets   string `json:"tickets"`
	Virtual   string `json:"virtual"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
}

// Desso builds the result of the API search
type Desso struct {
	Issues []struct {
		Key    string `json:"key"`
		Fields struct {
			Summary string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

// Filters builds the parameters for sed to execute on the scrape.txt file
type Filters struct {
	OPH2  string `json:"oph2"`
	OPH3  string `json:"oph3"`
	OPH4  string `json:"oph4"`
	CLH2  string `json:"clh2"`
	CLH3  string `json:"clh3"`
	CLH4  string `json:"clh4"`
	End   string `json:"end"`
	ESP   string `json:"esp"`
	Event string `json:"event"`
}

type JiraIssueRequest struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	IssueType   IssueType `json:"issuetype"`
	Labels      []string  `json:"labels,omitempty"`
	Reporter    Reporter  `json:"reporter,omitempty"`
	Project     Project   `json:"project"`
	Description ADFDoc    `json:"description"`
	Summary     string    `json:"summary"`
	Priority    Priority  `json:"priority,omitempty"`
}

type IssueType struct {
	ID string `json:"id"`
}

type Reporter struct {
	AccountID string `json:"accountId"`
}

type Project struct {
	Key string `json:"key"`
}

type Priority struct {
	ID string `json:"id"`
}

type ADFDoc struct {
	Type    string     `json:"type"`    // "doc"
	Version int        `json:"version"` // 1
	Content []ADFBlock `json:"content"`
}

type ADFBlock struct {
	Type    string      `json:"type"` // "paragraph"
	Content []ADFInline `json:"content"`
}

type ADFInline struct {
	Type string `json:"type"` // "text"
	Text string `json:"text"`
}

const (
	bv       string = "1.0.0"
	reset    string = "\033[0m"
	bgred    string = "\033[41m"
	green    string = "\033[32m"
	yellow   string = "\033[33m"
	bgyellow string = "\033[43m"
	halt     string = "program halted"
	header   string = "\nh2. Changelog\n"
	meta     string = "/data/automation/jsons/"
	temp     string = "/data/automation/temp/"
	tokens   string = "/data/automation/tokens/"
)

var (
	sre        Desso
	post       JiraIssueRequest
	label      string
	repo       string
	version    string
	content    []byte
	jira       Jira
	token      Tokens
	filter     Filters
	changelog  Changelogs
	versions   = [1][2]string{{".", "-"}}
	ephemeral  = []string{temp + "grep.txt", temp + "scrape.txt"}
	persistent = []string{meta + "changelogs.json", meta + "filters.json", meta + "jira.json", meta + "ticket.json", tokens + "tokens.json"}
	deletions  = []string{
		"<header>", "</header>",
		"</div>", "<p>", "</p>",
		"</h3>", "</h4>", "</li>",
		"<ul>", "</ul>", "</div>",
		"<br />", "</h1>", "</h2>",
		"<span>", "<entry>", "</entry>",
		"</span>", "<footer>", "</footer>",
	}
	replacements = [15][2]string{
		{"<em>", "*"},
		{"</em>", "*"},
		{"<li>", "- "},
		{"<code>", "*"},
		{"</code>", "*"},
		{"<h1>", "h1. "},
		{"<h2>", "h2. "},
		{"<h3>", "h3. "},
		{"<h4>", "h4. "},
		{"<strong>", "*"},
		{"</strong>", "*"},
		{"&#8211;", " - "},
		{"&#8216;", "'"},
		{"&#8217;", "'"},
		{"<li class=\"free\">", "- "},
	}
)
