1. export XML from JIRA
2. Parse XML and store attachments in story array
3. Fetch info for all files from clubhouse and store in a key-value store `[EXT_ID, CH_ID]` 
5. for each story to be created, loop through all attachments. For each attachment item:
    5.1 check if items key is in the attachmentMap
    5.2 if not, download file from JIRA
    5.3 upload file to clubhouse
    5.4 update CH-file with giving external_id the item key value
    5.5 append the CH-file-ID to the item's file_ids array
