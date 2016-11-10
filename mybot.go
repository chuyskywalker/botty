/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"bytes"
	"math/rand"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	ws, id := slackConnect(os.Args[1])
	fmt.Printf("mybot ready, id = %s, ^C exits\n", id)

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		// see if we're mentioned
		if m.Type == "message" {
    		fmt.Printf("message: %+v\n", m)
			if strings.Index(m.Text, "<@"+id+">") != -1 && strings.Index(m.Text, " sf ") != -1 {
				// if so try to parse if
				
				cmd := exec.Command("/bin/ls", "-alh")
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				m.Text = fmt.Sprintf("File Listing: ```\n%s\n```\n", out.String())
				// m.Text = fmt.Sprintf("Ok!\n")
				postMessage(ws, m)
			}
			if strings.Index(m.Text, "<@"+id+">") != -1 && strings.HasPrefix(m.Text, "thank") {
				reasons := []string{
				    "No problem",
				    "I live to serve",
				    "T'was my pleasure",
				    "ACKNOWLEDGED",
				    "But of course",
				    "It was nothin'",
				    "Any time",
				}
				n := rand.Int() % len(reasons)
				m.Text = fmt.Sprintf("%s, <@%s>!\n", reasons[n], m.User)
				postMessage(ws, m)
			}
		}
	}
}
