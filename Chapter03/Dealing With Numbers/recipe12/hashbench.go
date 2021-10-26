package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/fnv"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type (
	bps        int
	hashFunc   func() hash.Hash
	hashFunc32 func() hash.Hash32
	hashFunc64 func() hash.Hash64
	hashData   struct {
		name string
		f    hashFunc
		f32  hashFunc32
		f64  hashFunc64
		bs   int
		bps  bps
		err  error
		d    time.Duration
		sum  []byte
	}
)

var (
	data  []byte
	bench = []*hashData{
		// Hashes defined by/known to crypto package
		{name: "MD4", f: md4.New},
		{name: "MD5", f: md5.New},
		{name: "SHA1", f: sha1.New},
		{name: "SHA224", f: sha256.New224},
		{name: "SHA256", f: sha256.New},
		{name: "SHA384", f: sha512.New384},
		{name: "SHA512", f: sha512.New},
		{name: "SHA512_224", f: sha512.New512_224},
		{name: "SHA512_256", f: sha512.New512_256},
		{name: "RIPEMD160", f: ripemd160.New},
		{name: "SHA3_224", f: sha3.New224},
		{name: "SHA3_256", f: sha3.New256},
		{name: "SHA3_384", f: sha3.New384},
		{name: "SHA3_512", f: sha3.New512},
		{name: "BLAKE2s_256", f: blake(blake2s.New256)},
		{name: "BLAKE2b_256", f: blake(blake2b.New256)},
		{name: "BLAKE2b_384", f: blake(blake2b.New384)},
		{name: "BLAKE2b_512", f: blake(blake2b.New512)},

		// Other hash methods from standard lib
		{name: "ADLER32", f32: adler32.New},
		{name: "FNV-1_128", f: fnv.New128},
		{name: "FNV-1a_128", f: fnv.New128a},
		{name: "FNV-1_32", f32: fnv.New32},
		{name: "FNV-1a_32", f32: fnv.New32a},
		{name: "FNV-1_64", f64: fnv.New64},
		{name: "FNV-1a_64", f64: fnv.New64a},
	}
)

func (b bps) String() string {
	var units = []string{"B/s", "kB/s", "MB/s", "GB/s", "TB/s", "PB/s"}
	var i, u = float64(b), 0
	for i > 1000 && u < 5 {
		i = i / 1024
		u++
	}
	return fmt.Sprintf("%.3f %s", i, units[u])
}

func blake(blakeHasher func(key []byte) (hash.Hash, error)) hashFunc {
	return func() hash.Hash {
		h, err := blakeHasher(nil)
		if err != nil {
			return nil
		}
		return h
	}
}

func benchmark(h *hashData) {
	var hasher hash.Hash
	switch {
	case h.f != nil:
		hasher = h.f()
	case h.f32 != nil:
		hasher = h.f32()
	case h.f64 != nil:
		hasher = h.f64()
	default:
		return
	}
	h.bs = hasher.BlockSize()
	start := time.Now()
	_, err := hasher.Write(data)
	h.sum = hasher.Sum(nil)
	h.d = time.Since(start)
	h.err = err
	bbps := int(float64(len(data)) / h.d.Seconds())
	h.bps = bps(bbps)
}

func generateData() {
	start := time.Now()
	data = make([]byte, 250*1024*1024)
	rand.Read(data)
	log.Printf("Test data generated in %s", time.Since(start))
}

func main() {
	generateData()

	var wg sync.WaitGroup
	bhChan := make(chan *hashData)
	for i := 1; i < runtime.NumCPU(); i++ {
		wg.Add(1)

		go func(i int) {
			var bc int
			for b := range bhChan {
				benchmark(b)
				bc++
			}
			log.Printf("Stopping bench GO-routine #%d (%d benchmarks done)", i, bc)
			wg.Done()
		}(i)
	}

	for _, b := range bench {
		bhChan <- b
	}
	close(bhChan)

	wg.Wait()

	log.Print("        Hash  Block size        Time           BPS  Sum")
	for _, b := range bench {
		log.Printf("%12s  % 10d  %10s  %12s  %x", b.name, b.bs, b.d, b.bps, b.sum)
	}

}
