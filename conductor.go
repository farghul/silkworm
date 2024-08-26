package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
)

var (
	flag = os.Args
	// hmdr, _ = os.UserHomeDir()
	// satis   strings.Builder
	// managed strings.Builder
)

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			json.Unmarshal(data, &post)
		case 1:
			json.Unmarshal(data, &filter)
		case 2:
			json.Unmarshal(data, &link)
		case 3:
			json.Unmarshal(data, &jira)
		}
	}
}

// Read updates.txt and take action based on the length of the produced array
func sifter() {
	goals := read(jira.Source)
	updates := strings.Split(string(goals), "\n")
	if len(updates) == 1 {
		engine(0, updates)
	} else {
		updates = strings.Split(string(goals), "\n")
		for i := 0; i < len(updates); i++ {
			engine(i, updates)
		}
	}
}

// Iterate through the updates array and assign plugin and ticket values
func engine(i int, updates []string) {
	if len(updates[i]) > 25 {

		/* See if the ticket already exists */
		firstsplit := strings.Split(updates[i], "/")
		apiget(firstsplit[1])

		if desso.Key == firstsplit[1] {
			repo = firstsplit[0]
			secondsplit := strings.Split(firstsplit[1], ":")
			label = secondsplit[0]
			version = secondsplit[1]

			switchboard()
			changelog := append([]byte(header), content...)

			/* Create Jira ticket using Description & Summary */
			post.Fields.Description = string(changelog)
			post.Fields.Summary = updates[i]
			body, _ := json.Marshal(post)
			document(jira.Path+"temp/data.txt", body)
			// execute("-e", "curl", "-H", "Authorization: Basic "+jira.Token, "-X", "POST", "--data", "@"+jira.Path+"temp/data.txt", "-H", "Content-Type: application/json", jira.Base+"issue")
			apiget(firstsplit[1])
			journal("Jira ticket " + desso.Key + " created at ")
		}
	}
}

// Grab the ticket information from Jira in order to extract the DESSO-XXXX identifier
func apiget(ticket string) {
	/* Actual command for quering the Jira API */
	result := execute("-c", "curl", "--request", "GET", "--url", jira.Base+"search?jql=summary%20~%20"+ticket, "--header", "Authorization: Basic "+jira.Token, "--header", "Accept: application/json")
	json.Unmarshal(result, &desso)
}

// Sort the query based on repository name
func switchboard() {
	if label == "spotlight-social-photo-feeds" {
		repo = "freemius"
	}
	switch repo {
	case "bcgov-plugin":
		premium(label)
	case "freemius":
		substitution(link.Spotlight, filter.OPH2+"v"+version+filter.ESP)
	case "wpengine":
		substitution(link.WordPress+"advanced-custom-fields/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", temp[0])
	default:
		substitution(link.WordPress+label+"/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", temp[0])
	}
}

// Apply special conditions to the premium in-house plugins
func premium(label string) {
	v := bytes.ReplaceAll([]byte(version), []byte(versions[0][0]), []byte(versions[0][1]))
	switch label {
	case "events-calendar-pro":
		substitution(link.Calendar+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "event-tickets-plus":
		substitution(link.Tickets+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "events-virtual":
		substitution(link.Virtual+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "gravityforms":
		substitution(link.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		substitution(link.Poly, filter.OPH4+version+filter.End)
	case "wp-all-export-pro":
		substitution(link.WPExport, "/"+version+filter.CLH4)
		content = execute("-c", "sed", "${/h3./d;}", temp[0])
	}
}

// Find and replace/delete html tags
func substitution(link, filter string) {
	execute("-e", "curl", "-s", link, "-o", temp[1])
	grep := execute("-c", "sed", "-n", filter, temp[1])
	for _, v := range deletions {
		replace := bytes.ReplaceAll(grep, []byte(v), []byte(""))
		grep = replace
	}
	for i := 0; i < len(replacements); i++ {
		replace := bytes.ReplaceAll(grep, []byte(replacements[i][0]), []byte(replacements[i][1]))
		grep = replace
	}
	document(temp[0], grep)
	content = execute("-c", "sed", "/^$/d ; s/	//g", temp[0])
	document(temp[0], content)
}

// Special filter to handle the Events Calendar suite of updates
func eventfilter() {
	content = execute("-c", "grep", "-v", "<", temp[0])
	document(temp[0], content)
	content = execute("-c", "sed", "1,3d", temp[0])
	content = append([]byte("h3. "+version+"\n"), content...)
}
