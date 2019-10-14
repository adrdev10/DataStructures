//Package set implements the set data structures
package set

//Set represents a set data structure
type Set struct {
	integerMap map[int]bool
}

//NewSet returns a pointer to a Set type
func NewSet(im map[int]bool) *Set {
	return &Set{
		integerMap: im,
	}
}

//AddElement adds a new element to the set
func (set *Set) AddElement(element int) {
	if !set.Contains(element) {
		set.integerMap[element] = true
	}
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
