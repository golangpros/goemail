/*
 * Copyright (c) 2022, GolangDevs
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package goemail

import (
	"net/mail"
)

/* Attachment: attach files */
type Attachment struct {
	Filename string
	Data string
	Inline bool
}

/* Header: an additional email header */
type Header struct {
	Key string
	Value string
}

type Message struct {
	From mail.Address
	To []string
	CC []string
	Bcc []string
}