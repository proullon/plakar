/*
 * Copyright (c) 2021 Gilles Chehade <gilles@poolp.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/poolpOrg/plakar/snapshot"
)

func cmd_rm(ctx Plakar, args []string) int {
	flags := flag.NewFlagSet("rm", flag.ExitOnError)
	flags.Parse(args)

	if flags.NArg() == 0 {
		log.Fatalf("%s: need at least one snapshot ID to rm", flag.CommandLine.Name())
	}

	snapshots, err := getSnapshots(ctx.Store(), flags.Args())
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	for _, snap := range snapshots {
		wg.Add(1)
		go func(snap *snapshot.Snapshot) {
			err := ctx.Store().Purge(snap.Uuid)
			if err == nil {
				fmt.Fprintf(os.Stdout, "%s: OK\n", snap.Uuid)
			} else {
				fmt.Fprintf(os.Stdout, "%s: KO\n", snap.Uuid)
			}
			wg.Done()
		}(snap)
	}
	wg.Wait()

	return 0
}
