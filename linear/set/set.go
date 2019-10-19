//Package set implements the set data structures
package set

//Set represents a set data structure
type Set struct {
	integerMap map[int]bool
}

//NewSet returns a pointer to a Set type
func NewSet() *Set {
	return &Set{
		integerMap: make(map[int]bool),
	}
}

//AddElement adds a new element to the set
func (set *Set) AddElement(element int) {
	if !set.Contains(element) {
		set.integerMap[element] = true
	}
}

//InterSect returns an intersectionset that consists of the intersection of set and another set
func (set *Set) InterSect(anotherSet *Set) *Set {
	is := NewSet()
	for key := range anotherSet.integerMap {
		if set.Contains(key) {
			is.AddElement(key)
		}
	}
	return is
}

//UnionSet returns an Unionset that consists of the union of the set and another set
func (set *Set) UnionSet(anotherSet *Set) *Set {
	us := NewSet()
	for key := range set.integerMap {
		us.AddElement(key)
	}
	for key2 := range anotherSet.integerMap {
		us.AddElement(key2)
	}
	return us
}

//DeleteElement deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

//Contains returns true if element is not found in the set, and false otherwise
func (set *Set) Contains(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}
