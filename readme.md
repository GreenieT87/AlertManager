# Alert Handler

This is supposed to CRUD a folder based structure for Prometheus Alerts. Including documentation for every Alert, i the from of a `readme.md`.  The `readme.md` will get converted into HTML to post to confluence. The current conflunece page will get stored as `conflunence.json` in each alert folder. 

The folder structre looks like this: 
```
./alerts/groupName/alertName/.meta.yaml
                             rule.yaml  
                             .confluence.json  
                             readme.md  
```

## `.meta.yaml`
Contains meta information about the alert. Do not touch this  
 like: 
 - last published version(should check if a updated version on confluence exists, throw warn dont change confluence)
 - alertname
 - groupname
 - mod time(FUN)
 - creation time

## `rule.yaml`
Typical prometheus rules format
- title
- expr 

## `confluence.json`
Stores the last published confluence page. Do not touch this 

## `readme.md`
Is used to generate the confluence page, this can be __BUT SHOULD NOT__ be left empty. It gets prefilled with the `alertName` and the `rule.yaml`.  
  - Idea:  
    - create a template file 