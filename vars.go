package main

// Links builds a collection of urls to target changlogs
type Links struct {
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

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	Team    string `json:"team"`
	Base    string `json:"base"`
	Path    string `json:"path"`
	Token   string `json:"token"`
	Issue   string `json:"issue"`
	Source  string `json:"source"`
	Project string `json:"project"`
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

type Desso struct {
	Issues []struct {
		Key    string `json:"key"`
		Fields struct {
			Summary string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

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
	bv     string = "1.0"
	halt   string = "program halted"
	header string = "\nh2. Changelog\n"
)

var (
	sre       Desso
	post      Posts
	link      Links
	label     string
	repo      string
	version   string
	content   []byte
	filter    Filters
	jira      Atlassian
	versions  = [1][2]string{{".", "-"}}
	temp      = []string{jira.Path + "temp/grep.txt", jira.Path + "temp/scrape.txt"}
	jsons     = []string{jira.Path + "jsons/ticket.json", jira.Path + "jsons/filters.json", jira.Path + "jsons/links.json", jira.Path + "jsons/jira.json"}
	deletions = []string{
		"<header>", "</header>",
		"</div>", "<p>", "</p>",
		"</h3>", "</h4>", "</li>",
		"<ul>", "</ul>", "</div>",
		"<br />", "</h1>", "</h2>",
		"<span>", "<entry>", "</entry>",
		"</span>", "<footer>", "</footer>",
	}
	replacements = [12][2]string{
		{"<em>", "*"},
		{"</em>", "*"},
		{"<li>", "- "},
		{"<code>", "*"},
		{"</code>", "*"},
		{"<h1>", "h1. "},
		{"<h2>", "h2. "},
		{"<h3>", "h3. "},
		{"<h4>", "h3. "},
		{"<strong>", "*"},
		{"</strong>", "*"},
		{"<li class=\"free\">", "- "},
	}
)
