package main

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Read updates.txt and take action based on the length of the produced array
func sifter() {
	goals := read(assets + "updates/updates.txt")
	updates := strings.Split(string(goals), "\n")
	if len(updates) == 1 {
		engine(0, updates)
	} else if len(updates) > 1 {
		for i := 0; i < len(updates); i++ {
			engine(i, updates)
		}
	} else {
		journal("No new tickets to create")
	}
}

// Iterate through the updates array and assign plugin and ticket values
func engine(i int, updates []string) {
	if len(updates[i]) > 25 {

		/* See if the ticket already exists */
		firstsplit := strings.Split(updates[i], "/")
		apiget(firstsplit[1])

		/* If not, create it */
		if len(sre.Issues) == 0 {
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
			execute("-v", "curl", "-H", "Authorization: Basic "+jira.Token, "-X", "POST", "--data", string(body), "-H", "Content-Type: application/json", jira.URL+"issue")

			/* Get the new DESSO key and log the ticket creation */
			apiget(firstsplit[1])
			journal("Jira ticket " + sre.Issues[0].Key + " created.")
		}
	}
}

// Grab the ticket information from Jira in order to extract the DESSO-XXXX identifier
func apiget(ticket string) {
	result := execute("-c", "curl", "--request", "GET", "--url", jira.URL+"search?jql=summary%20~%20"+ticket, "--header", "Authorization: Basic "+jira.Token, "--header", "Accept: application/json")
	err := json.Unmarshal(result, &sre)
	inspect(err)
}

// Sort the query based on repository name
func switchboard() {
	switch repo {
	case "premium-plugin":
		premium(label)
	case "freemius":
		substitution(changes.Spotlight, filter.OPH2+"v"+version+filter.ESP)
	case "wpengine":
		substitution(changes.WordPress+"advanced-custom-fields/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", temp[0])
	default:
		substitution(changes.WordPress+label+"/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", temp[0])
	}
}

// Apply special conditions to the premium in-house plugins
func premium(label string) {
	v := bytes.ReplaceAll([]byte(version), []byte(versions[0][0]), []byte(versions[0][1]))
	switch label {
	case "events-calendar-pro":
		substitution(changes.Calendar+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "event-tickets-plus":
		substitution(changes.Tickets+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "events-virtual":
		substitution(changes.Virtual+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "gravityforms":
		substitution(changes.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		substitution(changes.Poly, filter.OPH4+version+filter.End)
	case "wp-all-export-pro":
		substitution(changes.WPExport, "/"+version+filter.CLH4)
		content = execute("-c", "sed", "${/h3./d;}", temp[0])
	}
}

// Find and replace/delete html tags
func substitution(link, filter string) {
	execute("-v", "curl", "-s", link, "-o", temp[1])
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
