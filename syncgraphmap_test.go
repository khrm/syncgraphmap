package syncgraphmap

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"math/rand"
)

func data(n int) []string {
	str := make([]string, n)
	for i := range str {
		str[i] = strconv.Itoa(rand.Int())
	}
	return str
}

var wg sync.WaitGroup

func BenchmarkGraphSingleMutex(b *testing.B) {
	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphSingleMutex()
	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}
	wg.Wait()
}

func BenchmarkGraphDoubleMutex(b *testing.B) {
	str1 := data(b.N)
	str2 := data(b.N)

	m := NewSiteMapDoubleMutex()
	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}
	wg.Wait()
}

func BenchmarkGraphSTD(b *testing.B) {
	str1 := data(b.N)
	str2 := data(b.N)

	m := NewSiteMapDoubleMutex()
	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}
	wg.Wait()
}

func BenchmarkGraphHybrid(b *testing.B) {
	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphHybrid()

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}
	wg.Wait()
}

func BenchmarkGraphSingleMutex90(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphSingleMutex()

	l := float64(b.N) * float64(90) / float64(100)

	b.ResetTimer()
	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkGraphDoubleMutex90(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewSiteMapDoubleMutex()

	l := float64(b.N) * float64(90) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkGraphSTD90(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphSTD()

	l := float64(b.N) * float64(90) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkGraphHybrid90(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphHybrid()

	l := float64(b.N) * float64(90) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkGraphSingleMutex50(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphSingleMutex()

	l := float64(b.N) * float64(50) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkGraphDoubleMutex50(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewSiteMapDoubleMutex()

	l := float64(b.N) * float64(50) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkGraphSTD50(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphSTD()

	l := float64(b.N) * float64(50) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}

func BenchmarkHybrid50(b *testing.B) {

	index := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		index[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})

	str1 := data(b.N)
	str2 := data(b.N)

	m := NewGraphHybrid()

	l := float64(b.N) * float64(50) / float64(100)

	b.ResetTimer()

	add := func(u string) {
		m.AddNode(u)
		wg.Done()
	}
	conn := func(u, v string) {
		go m.AddDirection(u, v)
		wg.Done()
	}

	for i := 0; i < int(l); i++ {
		wg.Add(2)
		go add(str1[i])
		go conn(str1[i], str2[i])
		i++
		if i < b.N {
			wg.Add(2)
			go add(str1[i])
			go conn(str1[i], str2[i])

		}
	}

	for _, v := range index {
		wg.Add(2)
		go add(str1[v])
		go conn(str1[v], str2[v])
	}
	wg.Wait()
}
