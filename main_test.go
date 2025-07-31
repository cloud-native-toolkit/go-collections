package go_collections

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

type TestVal struct {
	Name  string
	Value string
}

var _ = Describe("Collection functions", func() {
	var collection []string
	matchFn := func(match string) func(v string) bool {
		return func(v string) bool {
			return v == match
		}
	}

	Describe("Any", func() {
		BeforeEach(func() {
			collection = []string{"one", "two", "three"}
		})

		Context("when a string slice contains at least one match", func() {
			It("should return true", func() {
				Expect(Any(collection, matchFn("one"))).To(BeTrue())
			})
		})

		Context("when a string slice contains no matches", func() {
			It("should return false", func() {
				Expect(Any(collection, matchFn("four"))).To(BeFalse())
			})
		})
	})

	Describe("All", func() {
		BeforeEach(func() {
			collection = []string{"one", "one", "one"}
		})

		Context("when a string slice contains all matches", func() {
			It("should return true", func() {
				Expect(All(collection, matchFn("one"))).To(BeTrue())
			})
		})

		Context("when a string slice contains no matches", func() {
			It("should return false", func() {
				Expect(All(collection, matchFn("four"))).To(BeFalse())
			})
		})

		Context("when a string slice contains some matches", func() {
			It("should return false", func() {
				collection = append(collection, "two")

				Expect(All(collection, matchFn("one"))).To(BeFalse())
			})
		})
	})

	Describe("Map", func() {
		var mapCollection []TestVal

		BeforeEach(func() {
			mapCollection = []TestVal{
				{
					Name:  "test1",
					Value: "value1",
				},
				{
					Name:  "test2",
					Value: "value2",
				},
			}
		})

		Context("when provided slice with TestVals", func() {
			It("should returned slice with mapped vales", func() {
				mapper := func(val TestVal) string {
					return fmt.Sprintf("%s:%s", val.Name, val.Value)
				}

				result := Map(mapCollection, mapper)
				Expect(result).To(Equal([]string{
					"test1:value1",
					"test2:value2",
				}))
			})
		})
	})

	Describe("Filter", func() {
		BeforeEach(func() {
			collection = []string{"one", "oneone", "two", "three"}
		})

		Context("when provided slice with strings", func() {
			It("should returned slice with filtered values", func() {
				filterFn := func(val string) bool {
					return strings.HasPrefix(val, "one")
				}

				result := Filter(collection, filterFn)
				Expect(result).To(Equal([]string{"one", "oneone"}))
			})
		})
	})

	Describe("Reduce", func() {
		Context("when provided slice with ints and summing reducer", func() {
			reduceCollection := []int{1, 2, 3, 4, 5}
			reducer := func(result int, current int, collection []int, index int) int {
				return result + current
			}

			It("should returned sum of all values", func() {
				result := Reduce(reduceCollection, reducer, 0)
				Expect(result).To(Equal(15))
			})
		})

		Context("when provided slice with strings and concat reducer", func() {
			reduceCollection := []string{"one", "two", "three"}
			reducer := func(result string, current string, collection []string, index int) string {
				if index == len(collection)-1 {
					return result + " and " + current
				}

				if result == "" {
					return current
				}
				
				return result + ", " + current
			}

			It("should returned concat of all values", func() {
				result := Reduce(reduceCollection, reducer, "")
				Expect(result).To(Equal("one, two and three"))
			})
		})
	})
})
