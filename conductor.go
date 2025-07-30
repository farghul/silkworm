package main

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Read updates.txt and take action based on the length of the produced array
func sifter() {
	goals := read(jira.Source)
	updates := strings.Split(string(goals), "\n")
	if len(updates) > 1 {
		for _, s := range updates {
			engine(strings.TrimSpace(s))
		}
	}
}

// Iterate through the updates array and assign plugin and ticket values
func engine(entry string) {
	if len(entry) > 25 {

		/* See if the ticket already exists */
		firstsplit := strings.Split(entry, "/")
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
			post.Fields.Summary = entry
			// body, _ := json.Marshal(post)
			// execute("-v", "curl", "-H", "Authorization: Basic "+token.Jira, "-X", "POST", "--data", string(body), "-H", "Content-Type: application/json", jira.URL+"issue")

			/* Get the new DESSO key and log the ticket creation */
			apiget(firstsplit[1])
			inform("Jira ticket " + sre.Issues[0].Key + " created.")
		} else {
			inform("Jira ticket " + sre.Issues[0].Key + " already exists.")
		}
	}
}

// Grab the ticket information from Jira in order to extract the DESSO-XXXX identifier
func apiget(ticket string) {
	result := execute("-c", "curl", "--request", "GET", "--url", jira.URL+"search?jql="+jira.Criteria+ticket, "--header", "Authorization: Basic "+token.Jira, "--header", "Accept: application/json")
	err := json.Unmarshal(result, &sre)
	inspect(err)
}

// Sort the query based on repository name
func switchboard() {
	switch repo {
	case "premium-plugin":
		premium(label)
	case "freemius":
		substitution(changelog.Spotlight, filter.OPH2+"v"+version+filter.ESP)
	case "wpengine":
		substitution(changelog.WordPress+"advanced-custom-fields/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", ephemeral[0])
	default:
		substitution(changelog.WordPress+label+"/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", ephemeral[0])
	}
}

// Apply special conditions to the premium in-house plugins
func premium(label string) {
	v := bytes.ReplaceAll([]byte(version), []byte(versions[0][0]), []byte(versions[0][1]))
	switch label {
	case "events-calendar-pro":
		substitution(changelog.Calendar+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "event-tickets-plus":
		substitution(changelog.Tickets+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "events-virtual":
		substitution(changelog.Virtual+string(v)+"/", "/"+version+filter.Event)
		eventfilter()
	case "gravityforms":
		substitution(changelog.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		substitution(changelog.Poly, filter.OPH4+version+filter.End)
	case "wp-all-export-pro":
		substitution(changelog.WPExport, "/"+version+filter.CLH4)
		content = execute("-c", "sed", "${/h3./d;}", ephemeral[0])
	}
}

// Find and replace/delete html tags
func substitution(link, filter string) {
	execute("-v", "curl", "-s", link, "-o", ephemeral[1])
	grep := execute("-c", "sed", "-n", filter, ephemeral[1])
	for _, v := range deletions {
		replace := bytes.ReplaceAll(grep, []byte(v), []byte(""))
		grep = replace
	}
	for i := range len(replacements) {
		replace := bytes.ReplaceAll(grep, []byte(replacements[i][0]), []byte(replacements[i][1]))
		grep = replace
	}
	document(ephemeral[0], grep)
	content = execute("-c", "sed", "/^$/d ; s/	//g", ephemeral[0])
	document(ephemeral[0], content)
}

// Special filter to handle the Events Calendar suite of updates
func eventfilter() {
	content = execute("-c", "grep", "-v", "<", ephemeral[0])
	document(ephemeral[0], content)
	content = execute("-c", "sed", "1,3d", ephemeral[0])
	content = append([]byte("h3. "+version+"\n"), content...)
}
