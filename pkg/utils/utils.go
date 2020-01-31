package utils

/**
Index returns the first index of the target string t, or -1 if no match is found.
*/
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

/**
Include returns true if the target string t is in the slice.
*/
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

/**
Any returns true if one of the strings in the slice satisfies the predicate f.
*/
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

/**
All returns true if all of the strings in the slice satisfy the predicate f.
*/
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

/**
Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
*/
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

/**
Map returns a new slice containing the results of applying the function f to each string in the original slice.
*/
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

/*
Difference return a new slice that in target slice and not in base slice.
*/
func Difference(target []string, base []string) []string {
	var diff []string

	// Loop one times, first to find target strings not in base,
	for _, s1 := range target {
		found := false
		for _, s2 := range base {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}
	return diff
}
