// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2014 Christian Schramm. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "github.com/tmthrgd/shoco"

// Emails is a model optimized for compressing email addresses.
//
// It was trained against a sample of 2,000 email addresses using the
// --optimize-encoding flag. It achieves good compression on short strings
// consisting of just email addresses.
func Emails() *shoco.Model {
	check(emailsModel)
	return emailsModel
}

var emailsModel = &shoco.Model{
	ChrsByChrID: []byte{'.', 'e', 'o', 'a', 'n', 's', 'r', 'c', 'i', 'l', '@', 't', 'h', 'u', 'm', 'd', 'g', 'p', 'y', 'b', 'w', 'k', 'f', 'v', 'j', '-', '1', '0', 'x', 'z', '2', '8'},
	ChrIdsByChr: [256]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 25, 0, -1, 27, 26, 30, -1, -1, -1, -1, -1, 31, -1, -1, -1, -1, -1, -1, -1, 10, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 3, 19, 7, 15, 1, 22, 16, 12, 8, 24, 21, 9, 14, 4, 2, 17, -1, 6, 5, 11, 13, 23, 20, 28, 18, 29, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	SuccessorIDsByChrIDAndChrID: [][]int8{
		{-1, -1, 3, 11, 2, 5, -1, 0, -1, 15, -1, 14, 8, 1, 7, 12, 4, 9, -1, 6, 10, -1, 13, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{3, 13, -1, 9, 4, 2, 0, 10, 14, 5, 6, 1, -1, -1, 12, 7, -1, -1, 8, -1, 11, -1, -1, 15, -1, -1, -1, -1, -1, -1, -1, -1},
		{1, -1, 4, -1, 2, 10, 3, 14, -1, 5, 6, 7, 13, 9, 0, 12, -1, 15, -1, 11, -1, -1, -1, 8, -1, -1, -1, -1, -1, -1, -1, -1},
		{4, -1, -1, -1, 2, 7, 1, 8, 3, 0, 13, 5, 9, 14, 6, 10, 15, -1, 11, -1, -1, -1, -1, 12, -1, -1, -1, -1, -1, -1, -1, -1},
		{2, 1, 13, 7, 10, 6, -1, 12, 9, -1, 3, 4, 0, -1, 15, 5, 8, -1, 14, -1, -1, -1, 11, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{0, 5, 4, 3, -1, 10, -1, 13, 7, 12, 1, 2, 6, 9, 8, -1, -1, 14, 15, -1, 11, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{9, 0, 4, 2, 10, 6, 11, -1, 1, 15, 8, 3, -1, 13, -1, 7, 5, -1, 12, -1, -1, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{7, 4, 0, 1, -1, 11, 6, 10, 9, 8, 14, 3, 2, 12, -1, -1, 15, -1, 13, -1, -1, 5, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{5, 6, 9, 8, 0, 3, 7, 2, -1, 1, 15, 4, -1, -1, 11, 10, 12, -1, -1, -1, -1, -1, 14, 13, -1, -1, -1, -1, -1, -1, -1, -1},
		{1, 0, 5, 3, -1, 6, -1, 14, 2, 4, 7, 8, -1, 13, 15, 9, -1, 12, 10, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 11},
		{-1, 13, -1, 8, 0, 3, 15, 2, -1, 11, -1, 6, 5, -1, 9, 7, 1, 12, 10, 4, 14, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{1, 0, 4, 6, -1, 11, 5, 13, 3, 12, 7, 8, 2, -1, 9, -1, 14, -1, 15, -1, 10, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{5, 3, 2, 1, 9, 0, 10, 12, 4, 13, 6, 8, -1, 7, 14, -1, -1, 15, -1, -1, -1, -1, -1, -1, -1, 11, -1, -1, -1, -1, -1, -1},
		{9, 10, -1, -1, 3, 2, 1, 11, 14, 4, -1, 5, -1, -1, 8, 13, 6, 7, -1, 15, -1, 0, 12, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{3, 1, 4, 0, -1, 9, 15, 5, 2, -1, 6, -1, 7, 12, 10, 13, -1, 8, 14, 11, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{3, 0, 5, 2, -1, 7, 6, 14, 1, 15, 4, -1, 11, 9, 13, 10, -1, 12, 8, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{1, 2, 3, 4, 13, 9, 6, -1, 8, 12, 11, -1, 7, 15, 0, -1, 14, 5, -1, -1, -1, -1, -1, -1, -1, 10, -1, -1, -1, -1, -1, -1},
		{7, 1, 6, 0, -1, 11, 4, 3, 10, 8, 14, 13, 2, 12, 15, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, 5, -1, -1, -1, -1, -1, -1},
		{0, 11, 7, 2, 4, 3, 14, 10, -1, 6, 1, 9, -1, -1, 5, 12, 15, 8, -1, 13, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{9, 0, 3, 1, -1, 8, 2, 13, 4, 7, 14, 5, 11, 6, -1, 15, -1, -1, 12, 10, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{0, 2, 4, 1, 8, 11, 14, 10, 3, 5, 12, -1, 7, -1, -1, 13, -1, -1, 9, -1, -1, 15, -1, 6, -1, -1, -1, -1, -1, -1, -1, -1},
		{1, 0, 6, 3, -1, 5, 13, 9, 4, 8, 2, 11, 7, -1, 15, -1, -1, -1, 10, 12, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{8, 5, 0, 6, -1, 10, 3, -1, 1, 7, 9, 2, -1, 11, 12, -1, -1, 15, -1, -1, -1, -1, 4, -1, -1, 14, -1, -1, -1, -1, -1, 13},
		{2, 0, 4, 3, -1, -1, 6, 7, 1, 15, 5, -1, 9, -1, 11, -1, -1, 12, 13, -1, 8, -1, 14, -1, 10, -1, -1, -1, -1, -1, -1, -1},
		{5, 3, 0, 1, -1, -1, 15, 9, 4, 14, 7, -1, -1, 2, 8, 12, 10, 11, -1, 6, -1, -1, 13, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, 11, -1, 10, 12, 4, -1, 5, -1, 0, -1, 1, 14, 8, 2, 6, 13, 3, -1, 15, 9, -1, 7, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{6, -1, -1, 10, -1, -1, -1, -1, -1, -1, 0, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 3, 1, -1, -1, 4, 12},
		{2, -1, -1, -1, -1, 10, -1, -1, -1, -1, 3, -1, -1, -1, 11, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0, 1, -1, -1, 6, 9},
		{0, 7, 2, 4, -1, 15, -1, 9, 6, -1, 1, 3, -1, -1, 10, -1, -1, 8, 13, -1, 12, -1, 5, -1, -1, 14, -1, -1, 11, -1, -1, -1},
		{4, 3, 0, 1, -1, -1, -1, -1, 2, 14, 5, -1, 7, 8, -1, 10, -1, -1, 9, -1, -1, -1, 11, -1, 12, 13, 15, -1, -1, 6, -1, -1},
		{1, 11, -1, -1, -1, -1, -1, -1, -1, -1, 0, 7, -1, -1, 15, -1, -1, 12, -1, -1, -1, -1, -1, -1, -1, -1, 3, 4, 10, -1, 6, -1},
		{8, -1, -1, -1, -1, 9, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0, 7, -1, -1, -1, 5},
	},
	ChrsByChrAndSuccessorID: [][]byte{
		{'l', 't', 'm', 'p', 's', 'c', 'd', 'f', 'u', 'w', 'a', 'e', 'n', 'g', 'h', 'b'},
		{'c', 'u', 'n', 'o', 'g', 's', 'b', 'm', 'h', 'p', 'w', 'a', 'd', 'f', 't', 'l'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'1', '0', '.', '@', '3', '4', '2', '6', '5', '8', 's', 'm', '\x00', '\x00', '\x00', '\x00'},
		{'@', '0', '6', '1', '2', '3', '.', '7', '4', '9', 'a', '5', '8', '\x00', '\x00', '\x00'},
		{'@', '.', '5', '1', '0', '4', '2', 't', '7', '6', 'x', 'e', 'p', '3', '9', 'm'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'1', '@', '4', '5', '6', '8', '3', '0', '.', 's', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'n', 'g', 'c', 's', 'b', 'h', 't', 'd', 'a', 'm', 'y', 'l', 'p', 'e', 'w', 'r'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'l', 'r', 'n', 'i', '.', 't', 'm', 's', 'c', 'h', 'd', 'y', 'v', '@', 'u', 'g'},
		{'e', 'a', 'r', 'o', 'i', 't', 'u', 'l', 's', '.', 'b', 'h', 'y', 'c', '@', 'd'},
		{'o', 'a', 'h', 't', 'e', 'k', 'r', '.', 'l', 'i', 'c', 's', 'u', 'y', '@', 'g'},
		{'e', 'i', 'a', '.', '@', 'o', 'r', 's', 'y', 'u', 'd', 'h', 'p', 'm', 'c', 'l'},
		{'r', 't', 's', '.', 'n', 'l', '@', 'd', 'y', 'a', 'c', 'w', 'm', 'e', 'i', 'v'},
		{'o', 'i', 't', 'r', 'f', 'e', 'a', 'l', '.', '@', 's', 'u', 'm', '8', '-', 'p'},
		{'m', '.', 'e', 'o', 'a', 'p', 'r', 'h', 'i', 's', '-', '@', 'l', 'n', 'g', 'u'},
		{'s', 'a', 'o', 'e', 'i', '.', '@', 'u', 't', 'n', 'r', '-', 'c', 'l', 'm', 'p'},
		{'n', 'l', 'c', 's', 't', '.', 'e', 'r', 'a', 'o', 'd', 'm', 'g', 'v', 'f', '@'},
		{'o', 'a', 'u', 'e', 'i', '.', 'b', '@', 'm', 'c', 'g', 'p', 'd', 'f', 'l', 'r'},
		{'e', '.', '@', 'a', 'i', 's', 'o', 'h', 'l', 'c', 'y', 't', 'b', 'r', '_', 'm'},
		{'e', '.', 'i', 'a', 'l', 'o', 's', '@', 't', 'd', 'y', '8', 'p', 'u', 'c', 'm'},
		{'a', 'e', 'i', '.', 'o', 'c', '@', 'h', 'p', 's', 'm', 'b', 'u', 'd', 'y', 'r'},
		{'h', 'e', '.', '@', 't', 'd', 's', 'a', 'g', 'i', 'n', 'f', 'c', 'o', 'y', 'm'},
		{'m', '.', 'n', 'r', 'o', 'l', '@', 't', 'v', 'u', 's', 'b', 'd', 'h', 'c', 'p'},
		{'a', 'e', 'h', 'c', 'r', '-', 'o', '.', 'l', 'p', 'i', 's', 'u', 't', '@', 'm'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'e', 'i', 'a', 't', 'o', 'g', 's', 'd', '@', '.', 'n', 'r', 'y', 'u', 'k', 'l'},
		{'.', '@', 't', 'a', 'o', 'e', 'h', 'i', 'm', 'u', 's', 'w', 'l', 'c', 'p', 'y'},
		{'e', '.', 'h', 'i', 'o', 'r', 'a', '@', 't', 'm', 'w', 's', 'l', 'c', 'g', 'y'},
		{'k', 'r', 's', 'n', 'l', 't', 'g', 'p', 'm', '.', 'e', 'c', 'f', 'd', 'i', 'b'},
		{'e', 'i', '.', 'a', 'o', '@', 'r', 'c', 'w', 'h', 'j', 'm', 'p', 'y', 'f', 'l'},
		{'.', 'a', 'e', 'i', 'o', 'l', 'v', 'h', 'n', 'y', 'c', 's', '@', 'd', 'r', 'k'},
		{'.', '@', 'o', 't', 'a', 'f', 'i', 'e', 'p', 'c', 'm', 'x', 'w', 'y', '-', 's'},
		{'.', '@', 'a', 's', 'n', 'm', 'l', 'o', 'p', 't', 'c', 'e', 'd', 'b', 'r', 'g'},
		{'o', 'a', 'i', 'e', '.', '@', 'z', 'h', 'u', 'y', 'd', 'f', 'j', '-', 'l', '1'},
	},
	Packs: []shoco.Pack{
		{0x80000000, 1, 2, [8]uint{26, 24, 24, 24, 24, 24, 24}, [8]int16{15, 3, 0, 0, 0, 0, 0}},
		{0xc0000000, 2, 4, [8]uint{25, 21, 18, 16, 16, 16, 16}, [8]int16{15, 15, 7, 3, 0, 0, 0}},
		{0xe0000000, 4, 8, [8]uint{24, 20, 16, 12, 9, 6, 3, 0}, [8]int16{15, 15, 15, 15, 7, 7, 7, 7}},
	},
	MinChr:        45,
	MaxSuccessorN: 7,
}
