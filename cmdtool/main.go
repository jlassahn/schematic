
package main

import (
	"fmt"
	"os"
	"github.com/jlassahn/schematic"
)

func main() {
	
	schem, err := schematic.LoadSchematic(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	schem.Save(os.Args[2])
	schem.ExportSVG("out.svg", 1)
}

