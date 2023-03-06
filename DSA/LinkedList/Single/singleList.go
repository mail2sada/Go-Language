package single

import "reflect"



type list struct {
	item any
	next *list
}

func CreateList() *list {
	list := list{next: nil}
	return &list
}

func (lst *list) AddItem(item any) {

	for lst.next != nil {
		lst = lst.next
	}
	lst.item = item
	lst.next = &list{next: nil}
}

func (lst *list) GetList() []any {
	elements := []any{}
	for lst.next != nil {
		elements = append(elements, lst.item)
		lst = lst.next
	}
	return elements
}

func (lst *list) Delete(item any) bool {



	for lst.next != nil {
		if reflect.DeepEqual(lst.item, item) {

		}

	}
	return true
}
