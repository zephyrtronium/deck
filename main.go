package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"text/tabwriter"
)

var usage = `deck {"n"|"k"} N K n/k

N is the number of cards in the deck.
K is the number of lands in the deck.
n is the number of cards drawn.
k is the number of lands drawn.
If the support is n, then n/k is k, the pmf is the probability of drawing the
kth land on the nth draw, and the cdf is the probability of k within n draws.
If the support is k, then n/k is n, the pmf is the probability of drawing
exactly k lands in n draws, and the cdf is the probability of k or fewer.
`

func e(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(0)
}

func main() {
	args := os.Args[1:]
	if len(args) != 4 || args[0] != "n" && args[0] != "k" {
		e(usage)
	}
	var N, K, n, k int64
	var err error
	N, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		e(err)
	}
	K, err = strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		e(err)
	}
	n, err = strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		e(err)
	}

	start := 0
	var pmf []*big.Rat
	if args[0] == "n" {
		start = int(n)
		pmf = make([]*big.Rat, 0, N-n)
		for k = n; n <= N; n++ {
			pmf = append(pmf, neghypergeo(N, K, n, k))
		}
	} else {
		max := n
		if max > K {
			max = K
		}
		pmf = make([]*big.Rat, 0, max)
		for k = 0; k <= max; k++ {
			pmf = append(pmf, hypergeo(N, K, n, k))
		}
	}
	cdf := make([]*big.Rat, len(pmf))
	cdf[0] = pmf[0]
	for i, v := range pmf[1:] {
		cdf[i+1] = new(big.Rat).Add(cdf[i], v)
	}

	tw := tabwriter.NewWriter(os.Stdout, 3, 4, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintf(tw, "%s\tpmf\tcdf\t\n", args[0])
	for i, v := range pmf {
		fmt.Fprintf(tw, "%d\t%s\t%s\t\n", start+i, v.FloatString(20), cdf[i].FloatString(20))
	}
	fmt.Fprintf(tw, "%s\tpmf\tcdf\t\n", args[0])
	tw.Flush()
}
