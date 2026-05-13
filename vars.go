package main

// Jira builds the Jira API address and update source
type Jira struct {
	Source string `json:"source"`
	Token  string `json:"token"`
	URL    string `json:"url"`
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

type Request struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	IssueType   IssueType `json:"issuetype"`
	Labels      []string  `json:"labels,omitempty"`
	Reporter    Reporter  `json:"reporter"`
	Project     Project   `json:"project"`
	Description ADFDoc    `json:"description"`
	Summary     string    `json:"summary"`
	Priority    Priority  `json:"priority"`
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
	Type    string     `json:"type"`
	Version int        `json:"version"`
	Content []ADFBlock `json:"content"`
}

type ADFBlock struct {
	Type    string         `json:"type"`
	Attrs   map[string]int `json:"attrs,omitempty"`
	Content []ADFInline    `json:"content,omitempty"`
}

type ADFInline struct {
	Type    string      `json:"type"`
	Text    string      `json:"text,omitempty"`
	Content []ADFInline `json:"content,omitempty"`
}

type Response struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

type Color string

const (
	Reset           = "\033[0m"
	Black    Color  = "\033[30m"
	Red      Color  = "\033[31m"
	Green    Color  = "\033[32m"
	Yellow   Color  = "\033[33m"
	Blue     Color  = "\033[34m"
	Magenta  Color  = "\033[35m"
	Cyan     Color  = "\033[36m"
	White    Color  = "\033[37m"
	BGRed    Color  = "\033[41m"
	BGYellow Color  = "\033[43m"
	bv       string = "1.1.0"
	halt     string = "program halted"
	header   string = "<H2>Changelog</H2>\n"
	meta     string = "/data/automation/jsons/"
	temp     string = "/data/automation/temp/"
)

var (
	content    []byte
	jira       Jira
	filter     Filters
	post       Request
	changelog  Changelogs
	versions   = [1][2]string{{".", "-"}}
	ephemeral  = []string{temp + "grep.txt", temp + "scrape.txt"}
	persistent = []string{meta + "changelogs.json", meta + "filters.json", meta + "jira.json", meta + "ticket.json"}
)
