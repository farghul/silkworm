package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
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
func engine(summary string) {
	if len(summary) > 25 {
		firstsplit := strings.Split(summary, "/")
		repo := firstsplit[0]
		secondsplit := strings.Split(firstsplit[1], ":")
		label := secondsplit[0]
		version := secondsplit[1]

		switchboard(repo, label, version)
		changelog := append([]byte(header), content...)

		adf, err := converter(string(changelog))
		inspect(err)

		description := ADFDoc{adf.Type, adf.Version, adf.Content}

		/* Create a new Jira ticket using Description & Summary */
		post.Fields.Description = description
		post.Fields.Summary = summary
		body, _ := json.Marshal(post)
		spawn(body)
	}
}

// Sort the query based on repository name
func switchboard(repo, label, version string) {
	switch repo {
	case "premium-plugin":
		premium(label, version)
	case "freemius":
		webscrape(changelog.Spotlight, filter.OPH2+"v"+version+filter.ESP)
	case "wpengine":
		webscrape(changelog.WordPress+"advanced-custom-fields/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", ephemeral[0])
	default:
		webscrape(changelog.WordPress+label+"/#developers", "/Changelog"+filter.CLH2)
		content = execute("-c", "sed", "1d", ephemeral[0])
	}
}

func webscrape(link, filter string) {
	execute("-v", "curl", "-s", link, "-o", ephemeral[1])
	grep := execute("-c", "sed", "-n", filter, ephemeral[1])
	document(ephemeral[0], grep)
	content = execute("-c", "sed", "/^$/d ; s/	//g", ephemeral[0])
	document(ephemeral[0], content)
}

// Apply special conditions to the premium in-house plugins
func premium(label, version string) {
	v := bytes.ReplaceAll([]byte(version), []byte(versions[0][0]), []byte(versions[0][1]))
	switch label {
	case "events-calendar-pro":
		webscrape(changelog.Calendar+string(v)+"/", "/"+version+filter.Event)
		eventfilter(version)
	case "event-tickets-plus":
		webscrape(changelog.Tickets+string(v)+"/", "/"+version+filter.Event)
		eventfilter(version)
	case "events-virtual":
		webscrape(changelog.Virtual+string(v)+"/", "/"+version+filter.Event)
		eventfilter(version)
	case "gravityforms":
		webscrape(changelog.Gravity, filter.OPH3+version+filter.End)
	case "polylang-pro":
		webscrape(changelog.Poly, filter.OPH4+version+filter.End)
	case "wp-all-export-pro":
		webscrape(changelog.WPExport, "/"+version+filter.CLH4)
		content = execute("-c", "sed", "${/h3./d;}", ephemeral[0])
	}
}

// Special filter to handle the Events Calendar suite of updates
func eventfilter(version string) {
	content = execute("-c", "grep", "-v", "<", ephemeral[0])
	document(ephemeral[0], content)
	content = execute("-c", "sed", "1,3d", ephemeral[0])
	content = append([]byte("h3. "+version+"\n"), content...)
}

// Send a POST request to the Jira API
func spawn(body []byte) {
	req, _ := http.NewRequest("POST", jira.URL+"issue", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Basic "+jira.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Jira error: %s", body)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		inspect(err)
	} else {
		inform("Jira ticket " + result.Key + " created.")
	}

	defer resp.Body.Close()
}
