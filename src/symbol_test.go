package Symbol_test

import (
	"sync"

	Symbol "github.com/go-composites/symbol/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Symbol", func() {

	ginkgo.It("constructor interns a symbol by name", func() {
		gomega.Expect(Symbol.New("foo").ToGoString()).To(gomega.Equal("foo"))
	})

	ginkgo.It("exposes its name via Name as an alias of ToGoString", func() {
		s := Symbol.New("alias")
		gomega.Expect(s.Name()).To(gomega.Equal(s.ToGoString()))
	})

	ginkgo.It("returns the SAME underlying instance for the same name", func() {
		first := Symbol.New("interned")
		second := Symbol.New("interned")
		// Pointer identity: like Ruby's :interned.
		gomega.Expect(first).To(gomega.BeIdenticalTo(second))
	})

	ginkgo.It("returns distinct instances for distinct names", func() {
		a := Symbol.New("a")
		b := Symbol.New("b")
		gomega.Expect(a).NotTo(gomega.BeIdenticalTo(b))
	})

	ginkgo.It("is Equal to another symbol of the same name", func() {
		gomega.Expect(Symbol.New("same").Equal(Symbol.New("same"))).To(gomega.BeTrue())
	})

	ginkgo.It("is not Equal to a symbol of a different name", func() {
		gomega.Expect(Symbol.New("left").Equal(Symbol.New("right"))).To(gomega.BeFalse())
	})

	ginkgo.It("is not Equal to the Null-Object symbol", func() {
		gomega.Expect(Symbol.New("real").Equal(Symbol.Null())).To(gomega.BeFalse())
	})

	ginkgo.It("reports that a real symbol is not null", func() {
		gomega.Expect(Symbol.New("real").IsNull()).To(gomega.BeFalse())
	})

	ginkgo.It("inspects a real symbol as :name", func() {
		gomega.Expect(Symbol.New("foo").Inspect()).To(gomega.Equal(":foo"))
	})

	ginkgo.It("interns safely under concurrent access", func() {
		const goroutines = 64
		var wg sync.WaitGroup
		results := make([]Symbol.Interface, goroutines)
		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func(idx int) {
				defer wg.Done()
				results[idx] = Symbol.New("concurrent")
			}(i)
		}
		wg.Wait()
		for i := 1; i < goroutines; i++ {
			gomega.Expect(results[i]).To(gomega.BeIdenticalTo(results[0]))
		}
	})

	ginkgo.Context("the Null-Object Symbol", func() {
		ginkgo.It("reports that it is null", func() {
			gomega.Expect(Symbol.Null().IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("has an empty name", func() {
			gomega.Expect(Symbol.Null().ToGoString()).To(gomega.Equal(""))
			gomega.Expect(Symbol.Null().Name()).To(gomega.Equal(""))
		})
		ginkgo.It("is never equal to a real symbol", func() {
			gomega.Expect(Symbol.Null().Equal(Symbol.New("x"))).To(gomega.BeFalse())
		})
		ginkgo.It("is not even equal to itself", func() {
			gomega.Expect(Symbol.Null().Equal(Symbol.Null())).To(gomega.BeFalse())
		})
		ginkgo.It("inspects as :null", func() {
			gomega.Expect(Symbol.Null().Inspect()).To(gomega.Equal(":null"))
		})
	})
})
