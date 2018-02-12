1. export XML from JIRA
2. Parse XML and store attachments in story array
3. Fetch info for all files from clubhouse and store in a key-value store `[EXT_ID, CH_ID]` 
5. for each story to be created, loop through all attachments. For each attachment item:
    5.1 check if items key is in the attachmentMap
    5.2 if not, download file from JIRA
    5.3 upload file to clubhouse
    5.4 update CH-file with giving external_id the item key value
    5.5 append the CH-file-ID to the item's file_ids array


Completing stories with files
1. export XML from JIRA
2. Parse XML and store attachments in story array
3. Fetch info for all files from clubhouse and store in a key-value store `[EXT_ID, CH_ID]` 

need 2 maps:
1. map[JiraItemKey string][]CHFile
2. map[CHStoryExternalID string]CHStoryID int64

1 is made while importing from XML jiraKey
2 is made by reading all stories from CH  (https://api.clubhouse.io/api/v2/projects/{project-public-id}/stories)