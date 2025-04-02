package main

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	URL    string `json:"url"`
	Token  string `json:"token"`
	Source string `json:"source"`
}

// Changelogs builds a collection of urls to target changlogs
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

// Filters builds the parameters for sed to execute on the scrapped.txt file
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

// Posts builds the ticket information to send to Jira
type Posts struct {
	Fields struct {
		Issuetype struct {
			Self string `json:"self"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"issuetype"`
		Labels   []string `json:"labels"`
		Reporter struct {
			Self         string `json:"self"`
			AccountID    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
		} `json:"reporter"`
		Project struct {
			Self           string `json:"self"`
			ID             string `json:"id"`
			Key            string `json:"key"`
			Name           string `json:"name"`
			ProjectTypeKey string `json:"projectTypeKey"`
		} `json:"project"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Priority    struct {
			Self string `json:"self"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"priority"`
	} `json:"fields"`
}

const (
	bv        string = "1.0.0"
	reset     string = "\033[0m"
	bgred     string = "\033[41m"
	green     string = "\033[32m"
	yellow    string = "\033[33m"
	bgyellow  string = "\033[43m"
	halt      string = "program halted"
	header    string = "\nh2. Changelog\n"
	temp      string = "/data/automation/temp/"
	resources string = "/data/automation/resources/"
)

var (
	sre        Desso
	post       Posts
	label      string
	repo       string
	version    string
	content    []byte
	filter     Filters
	jira       Atlassian
	changes    Changelogs
	versions   = [1][2]string{{".", "-"}}
	ephemeral  = []string{temp + "grep.txt", temp + "scrape.txt"}
	persistent = []string{resources + "changelogs.json", resources + "filters.json", resources + "jira.json", resources + "new.json"}
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
