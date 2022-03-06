package scrapePkg

// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

type Monitor struct {
	Address string `json:"address"`
	Path    string `json:"path"`
	Count   uint64 `json:"count"`
	Size    uint64 `json:"size"`
}

func NewMonitor() Monitor {
	monitor := new(Monitor)
	return *monitor
}

// func (mon *Monitor) Freshen(opts *ScrapeOptions) (uint64, error) {
// 	i := 0
// 	fmt.Println("I am here", i)
// 	i++
// 	var expOpts exportPkg.ExportOptions
// 	expOpts.Addrs = append(expOpts.Addrs, mon.Address)
// 	expOpts.Globals = opts.Globals
// 	expOpts.Freshen = true
// 	fmt.Println("I am here", i)
// 	i++
// 	opts.Globals.PassItOn("acctExport", expOpts.ToCmdLine())
// 	fmt.Println("I am here", i, mon.Path)
// 	i++
// 	in, err := os.Stat(mon.Path)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	mon.Size = uint64(in.Size())
// 	mon.Count = mon.Size / 8
// 	return mon.Count, nil
// }
