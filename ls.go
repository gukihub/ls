package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var prog_name string = path.Base(os.Args[0])
var version string = "1.0"

func main() {

	var dir string
	var err error
	var fi_buf os.FileInfo
	var fi_tab []os.FileInfo
	var f *os.File

	var v_flag *bool

	v_flag = flag.Bool("version", false, "display program version")

	flag.Parse()

	if *v_flag == true {
		fmt.Printf("%s version %s\n", prog_name, version)
		return
	}

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s directory\n", prog_name)
		os.Exit(1)
	}

	dir = os.Args[1]

	f, err = os.Open(dir)

	if err != nil {
		fmt.Printf("%s: failed to open %s (%v)\n", prog_name, dir, err)
		os.Exit(1)
	}

	fi_buf, err = f.Stat()

	if err != nil {
		fmt.Printf("%s: failed to stat %s (%v)\n", prog_name, dir, err)
		os.Exit(1)
	}

	if fi_buf.IsDir() == false {
		fmt.Printf("%s: %s is not a directory\n", prog_name, dir)
		os.Exit(1)
	}

	fi_tab, err = f.Readdir(0)

	if err != nil {
		fmt.Printf("%s: failed to read %s (%v)\n", prog_name, dir, err)
		os.Exit(1)
	}

	for _, g := range fi_tab {
		fmt.Printf("%-11s %5d %s %s\n",
			g.Mode(),
			g.Size(),
			g.ModTime().Format("2006-01-02 15:04:05"),
			g.Name())
	}

}
