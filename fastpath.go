package main

import (
	"strings"
)

func calculateDistance1(pairs [][]string) (src, dst string) {
	for i := 0; i < len(pairs); i++ {
		src += pairs[i][0]
		dst += pairs[i][1]
	}

	for i := 0; i < len(pairs); i++ {
		src = strings.Replace(src, pairs[i][1], "", -1)
		dst = strings.Replace(dst, pairs[i][0], "", -1)
	}

	return
}

func calculateDistance2(pairs [][]string) (src, dst string) {
	for i := 0; i < len(pairs); i++ {
		srcI := pairs[i][0]
		dstI := pairs[i][1]

		isSrc, isDst := true, true
		for j := 0; j < len(pairs); j++ {
			if srcI == pairs[j][1] {
				isSrc = false
			}

			if dstI == pairs[j][0] {
				isDst = false
			}
		}

		if isSrc {
			src = srcI
		}

		if isDst {
			dst = dstI
		}
	}

	return
}
