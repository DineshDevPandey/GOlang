// package twofer implements sharing functions 
package twofer

// ShareWith returns sharing with whome.
// it takes an argument name, if no name is passed it will treat it as "you"
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
