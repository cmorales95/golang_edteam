// name of the directory
package greet

// The variables scope is a global package level
var emoji = ":D"

// Export funcionts and Not export funcionts (Applicable for var/cons)
// First Letter Uppercase: Is exported, It could be consume by other packages
// First letter lowercase: Is not exported, It is only used into the package scope
func English() string {
	return "Hi " + emoji
}

func Italian() string {
	return "Ciao " + emoji
}