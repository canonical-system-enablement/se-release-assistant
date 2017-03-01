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
	"github.com/bergotorino/go-trello"
)

type SeClient struct {
	trello trello.Client
}

func NewSeClient(secrets TrelloSecrets) (*SeClient, error) {
	trello, err := trello.NewAuthClient(secrets.AppKey, &secrets.Token)
	if err != nil {
		return nil, err
	}

	return &SeClient{trello: *trello}, nil
}
