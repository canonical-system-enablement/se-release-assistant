SE Release Assistant
====================

This utility helps the SE team make the releases of the snaps.

It currently print the list of merge proposals related to each snap that
is being released. The MPs are fetched from "MPs" checklist. This checklist
is searched on:

 a) the release story
 b) the work story that is related to the release story

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

```
konrad at annapurna in se-release-assistant (master) % ./se-release-assistant 
https://code.launchpad.net/~morphis/snappy-hwe-snaps/+git/modem-manager/+merge/318102
https://code.launchpad.net/~kzapalowicz/snappy-hwe-snaps/+git/bluez/+merge/315108
https://code.launchpad.net/~kzapalowicz/snappy-hwe-snaps/+git/engineering-tests/+merge/317756
konrad at annapurna in se-release-assistant (master) %
```
