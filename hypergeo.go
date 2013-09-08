package main

import "math/big"

// N: number of cards in deck
// K: number of lands in deck
// n: number of draws
// k: number of drawn lands
// return: probability that k is the number of drawn lands after n draws
func hypergeo(N, K, n, k int64) *big.Rat {
	num := new(big.Int).Binomial(n, k)
	num.Mul(num, new(big.Int).Binomial(N-n, K-k))
	return new(big.Rat).SetFrac(num, new(big.Int).Binomial(N, K))
}

// N: number of cards in deck
// K: number of lands in deck
// n: number of draws
// k: number of drawn lands
// return: probability that n is the number of draws to get k lands
func neghypergeo(N, K, n, k int64) *big.Rat {
	num := new(big.Int).Binomial(K, k-1)
	num.Mul(num, new(big.Int).Binomial(N-K, n-k))
	p := new(big.Rat).SetFrac(num, new(big.Int).Binomial(N, n-1))
	return p.Mul(p, big.NewRat(K-k+1, N-n+1))
}
