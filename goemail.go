/*
 * Copyright (c) 2022, GolangDevs
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package goemail

import (
	"io/ioutil"
	"net/mail"
	"path/filepath"
)

/* Attachment: attach files */
type Attachment struct {
	Filename string
	Data     []byte
	Inline   bool
}

/* Header: an additional email header */
type Header struct {
	Key   string
	Value string
}

/* Message: From, To, ReplyTo, Subject, Body, Attachments... */
type Message struct {
	From            mail.Address
	To              []string
	CC              []string
	Bcc             []string
	ReplyTo         string
	Subject         string
	Body            string
	BodyContentType string
	Headers         []Header
	Attachments     map[string]*Attachment
}

func (m *Message) attach(file string, inline bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(file)

	m.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

func (m *Message) addTo(address mail.Address) []string {
	m.To = append(m.To, address.String())
	return m.To
}

func (m *Message) addCC(address mail.Address) []string {
	m.CC = append(m.CC, address.String())
	return m.CC
}

func (m *Message) addBcc(address mail.Address) []string {
	m.Bcc = append(m.Bcc, address.String())
	return m.Bcc
}
