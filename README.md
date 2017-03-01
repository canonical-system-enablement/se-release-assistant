SE Release Assistant
====================

This utility helps the SE team make the releases of the snaps.

Building
========

It uses the:
 - "github.com/bergotorino/go-trello" Trello library for Go.

You can install them with:

```
go get github.com/bergotorino/go-trello
```

Requirements
============

trello_secrets.json file that contains the Application Key and Token.

Follow the instructions here https://trello.com/app-key to obtain an API
key and token.


```
konrad at annapurna in snap-release-email-generator (master) % cat trello_secrets.json 
{"app_id":"APP ID","token":"TOKEN"}

```

Usage
=====

./se-release-assistant

Output
======

List of Merge Proposals related to snaps being currently released
