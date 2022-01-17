/*
 * Copyright (c) 2022, GolangDevs
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package goemail

import (
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"path/filepath"
	"strings"
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

func (m *Message) Attach(file string) error {
	return m.attach(file, false)
}

func (m *Message) Inline(file string) error {
	return m.attach(file, true)
}

func (m *Message) addHeader(key string, value string) Header {
	newHeader := Header{Key: key, Value: value}
	m.Headers = append(m.Headers, newHeader)
	return newHeader
}

func newMessage(subject string, body string, bodyContentType string) *Message {
	m := &Message{Subject: subject, Body: body, BodyContentType: bodyContentType}
	m.Attachments = make(map[string]*Attachment)

	return m
}

func NewMessage(subject string, body string) *Message {
	return newMessage(subject, body, "text/plain")
}

func NewHTMLMessage(subject string, body string) *Message {
	return newMessage(subject, body, "text/html")
}

func (m *Message) Tolist() []string {
	rcptList := []string{}

	toList, _ := mail.ParseAddressList(strings.Join(m.To, ","))
	for _, to := range toList {
		rcptList = append(rcptList, to.Address)
	}

	ccList, _ := mail.ParseAddressList(strings.Join(m.CC, ","))
	for _, cc := range ccList {
		rcptList = append(rcptList, cc.Address)
	}

	bccList, _ := mail.ParseAddressList(strings.Join(m.Bcc, ","))
	for _, bcc := range bccList {
		rcptList = append(rcptList, bcc.Address)
	}

	return rcptList
}

func Send(addr string, auth smtp.Auth, m *Message) error {
	return smtp.SendMail(addr, auth, m.From.Address, m.Tolist())
}
