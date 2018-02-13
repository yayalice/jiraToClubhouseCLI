# jiraToClubhouseCLI

[Clubhouse.io](http://clubhouse.io) is an excellent modern project management application. This tool helps make the transition from JIRA by importing your JIRA tickets into Clubhouse.

Updated by T3db0t to use an external JSON file to map user IDs and project IDs and assorted other improvements.  See this blog post for more information and a full tutorial: http://log.liminastudio.com/programming/how-and-why-i-switched-from-jira-to-clubhouse

Updated again by marc-maillard to handle default user, multiple projects, estimates, adding jira sprints as tags on Clubhouse-stories and correctly mapping stories with epics.

The tool has two modes: `export` and `import`.  Use `export` to generate a JSON file and `import` to actually send all the data to Clubhouse using your API token.  Generate an API token from `https://app.clubhouse.io/[your_organization]/settings/account/api-tokens` and make it an environment variable like so: `export CLUBHOUSE_API_TOKEN=asdflkjaseflkjdf`

Also included are a few utility scripts in Python: `createTestStory.py`, `deleteArchivedStories.py` and `deleteEmptyEpics.py`.  Use these like so:

`python deleteArchivedStories.py $CLUBHOUSE_API_TOKEN`

## Notes / Known Issues

- The importer will add a "JIRA" label to every ticket it imports.  If something gets messed up and you need to redo an import, this makes it easy to select all imported tickets, archive them and use the `deleteArchivedStories.py` script to delete them.
- Clubhouse Epics will be created from JIRA Epics every time the script is run. If you need to redo an import, after deleting the imported tickets, you can run the `deleteEmptyEpics.py` script as above.
- JIRA Subtasks will be added as Clubhouse tasks for the parent ticket. If the parent is JIRA Epic (which you shouldn't do in the first place), those subtasks won't get imported.
- Workflow mapping needs to be added to the external JSON file, sorry! :)

## User Map

A JSON file with this structure:

```
[
	{
		"jiraUsername": "userA",
		"chID": "476257c9-ac5a-46bc-67d6-4bc8bbfde7be",
		"default: true
	},
	{
		"jiraUsername": "userB",
		"chID": "476257c9-ac5a-46bc-67d6-4bc8bbfde7be"
	},
	{
		"jiraUsername": "userC",
		"chID": "476257c9-ac5a-46bc-67d6-4bc8bbfde7be"
	}
]
```

## Project Map

A JSON file with this structure:

```
[
	{
		"jiraProjectKey": "ProjA",
		"chProjectID": 6
	},
	{
		"jiraProjectKey": "ProjB",
		"chProjectID": 820
	},
	{
		"jiraProjectKey": "ProjC",
		"chProjectID": 1444
	}
]
```		

You'll need to get the right JIRA usernames and JIRA projectkeys and Clubhouse Project ID and user IDs and fill them in with these structures. If you want to replace all unknown JIRA users with one default user, add a default=true to one of the users in userMap.json. Detail on how to do this in [this blog post](http://log.liminastudio.com/programming/how-and-why-i-switched-from-jira-to-clubhouse).

## Workflow Map

Right now you'll have to map workflow states in the script itself (sorry!) in `jiraStructs.go`:

```
switch item.Status {
    case "Ready for Test":
        // ready for test
        state = 500000010
    case "Task In Progress":
        // in progress
        state = 500000015
    case "Selected for Review/Development":
    	// selected
    	state = 500000011
    case "Task backlog":
    	// backlog
        state = 500000014
    case "Done":
    	// Completed
    	state = 500000012
    case "Verified":
    	// Completed
    	state = 500000012
    case "Closed":
    	state = 500000021
    default:
    	// backlog
        state = 500000014
}
```

To get your Clubhouse workflow state IDs, you can look at `curl -X GET \
  -H "Content-Type: application/json" \
  -L "https://api.clubhouse.io/api/v1/workflows?token=$CLUBHOUSE_API_TOKEN"`

## Export
To see what the data will look like (as JSON).

```bash
go run *.go export --in SearchRequest.xml --map userMap.json --project projectMap.json --out file.json
```


### Params
 * `--in` The xml file you want to read from
 * `--map` The user maps
 * `--project` The project maps
 * `--out` The file you want to export the JSOn to.

## Import

To actually import to Clubhouse. Use `--test` to parse input files and 'preview' what the effects will be without uploading to Clubhouse.  Good for checking that everything works before pulling the trigger.

```bash
go run *.go import --in SearchRequest.xml --map userMap.json --project projectMap.json --token $CLUBHOUSE_API_TOKEN
```

### Params
 * `--in` The xml file you want to read from
 * `--map` The user maps
 * `--project` The project maps
 * `--token` The api token for your Clubhouse instance
 * `--test` Test mode: run the program, but do not upload to Clubhouse.

 ## ImportFiles

 ```bash
 go run *.go importFiles --in SearchRequest.xml --map userMap.json --project projectMap.json --token $CLUBHOUSE_API_TOKEN
```

### Params
 * `--in` The xml file you want to read from
 * `--map` The user maps
 * `--project` The project maps
 * `--token` The api token for your Clubhouse instance
 * `--test` Test mode: run the program, but do not upload to Clubhouse.