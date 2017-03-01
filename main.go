/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	trelloSecretsFile = flag.String("secrets", "trello_secrets.json", "Trello Secrets configuration")
)

func main() {

	flag.Parse()

	trelloSecrets, err := NewTrelloSecrets(*trelloSecretsFile)
	if err != nil {
		log.Fatal(err)
	}

	trello, err := NewSeClient(*trelloSecrets)
	if err != nil {
		log.Fatal(err)
	}

	snapsToRelease, err := trello.SnapsToRelease()
	if err != nil {
		log.Fatal(err)
	}

	mergeProposals, err := snapsToRelease.MergeProposals()
	if err != nil {
		log.Fatal(err)
	}

	for _, mp := range mergeProposals {
		fmt.Println(mp)
	}
}
