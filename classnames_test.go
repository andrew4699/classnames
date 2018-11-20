package classnames_test

import (
	"strings"
	"testing"

	"github.com/andrew4699/classnames"
)

// Ensures that the result doesn't violate any basic invariants
func invariant(t *testing.T, result string) {
	l := len(result)

	if l > 0 {
		if result[0] == ' ' {
			t.Error("Leading whitespace")
		}

		if result[l-1] == ' ' {
			t.Error("Trailing whitespace")
		}
	}
}

// Returns all permutations of "list"
func permutations(list *[]string) [][]string {
	perms := [][]string{}
	permutationsRecursive(list, &[]string{}, &perms)
	return perms
}

func permutationsRecursive(list *[]string, cur *[]string, deposit *[][]string) {
	c := *cur
	l := *list

	if len(*list) == 0 {
		*deposit = append(*deposit, *cur)
	} else {
		for i, s := range *list {
			*cur = append(*cur, s)
			*list = append(l[:i], l[i+1:]...)
			permutationsRecursive(list, cur, deposit)
			*list = append(*list, s)
			*cur = append(c[:len(c)-1], c[len(c):]...)
		}
	}
}

// Returns true if "str" equals any space-delimited permutation of "list"
func equalsPermutation(str string, list *[]string) bool {
	perms := permutations(list)
	//perms := genPerm(list)

	for _, p := range perms {
		if str == strings.Join(p, " ") {
			return true
		}
	}

	return false
}

func TestEmpty(t *testing.T) {
	result := classnames.Classnames()
	invariant(t, result)

	if result != "" {
		t.Fail()
	}
}

func TestEmptyArr(t *testing.T) {
	result := classnames.Classnames([]string{})
	invariant(t, result)

	if result != "" {
		t.Fail()
	}
}

func TestSingle(t *testing.T) {
	c := "class-a"
	result := classnames.Classnames(c)
	invariant(t, result)

	if result != c {
		t.Fail()
	}
}

func TestMany(t *testing.T) {
	a := "class-a"
	b := "class-b"
	c := "class-c"
	d := "class-d"
	e := "class-e"
	f := "class-f"
	result := classnames.Classnames(a, b, c, d, e, f)
	invariant(t, result)

	if result != (a + " " + b + " " + c + " " + d + " " + e + " " + f) {
		t.Fail()
	}
}

func TestSingleArr(t *testing.T) {
	c := "class-a"
	result := classnames.Classnames([]string{c})
	invariant(t, result)

	if result != c {
		t.Fail()
	}
}

func TestManyArr(t *testing.T) {
	a := "class-a"
	b := "class-b"
	c := "class-c"
	d := "class-d"
	e := "class-e"
	f := "class-f"
	result := classnames.Classnames([]string{a, b, c, d, e, f})
	invariant(t, result)

	if result != (a + " " + b + " " + c + " " + d + " " + e + " " + f) {
		t.Fail()
	}
}

func TestSingleArrStrCombo(t *testing.T) {
	b := "class-b"
	c := "class-c"
	result := classnames.Classnames(b, []string{c})
	invariant(t, result)

	if result != (b + " " + c) {
		t.Fail()
	}
}

func TestManyArrStrCombo(t *testing.T) {
	a := "class-a"
	b := "class-b"
	c := "class-c"
	d := "class-d"
	e := "class-e"
	f := "class-f"
	result := classnames.Classnames(a, b, []string{c, d, e}, f)
	invariant(t, result)

	if result != (a + " " + b + " " + c + " " + d + " " + e + " " + f) {
		t.Fail()
	}
}

func TestManyBadValues(t *testing.T) {
	a := "class-a"
	b := false
	c := "class-c"
	d := 12
	var e *string
	f := "class-f"
	var arr []interface{}
	arr = append(arr, c)
	arr = append(arr, e)
	arr = append(arr, d)
	result := classnames.Classnames(a, b, arr, f)
	invariant(t, result)

	if result != (a + " " + c + " " + f) {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	a := "class-a"
	b := "class-b"
	c := "class-c"
	d := "class-d"
	e := "class-e"
	f := "class-f"
	myMap := make(classnames.Map)
	myMap[a] = true
	myMap[b] = false
	myMap[c] = false
	myMap[d] = true
	myMap[e] = true
	myMap[f] = false
	result := classnames.Classnames(myMap)
	invariant(t, result)

	if result != (a+" "+d+" "+e) &&
		result != (a+" "+e+" "+d) &&
		result != (d+" "+a+" "+e) &&
		result != (d+" "+e+" "+a) &&
		result != (e+" "+d+" "+a) &&
		result != (a+" "+a+" "+d) {
		//if equalsPermutation(result, &[]string{a, d, e}) {
		t.Fail()
	}
}
