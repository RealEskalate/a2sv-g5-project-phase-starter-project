package ihash

type Service interface {

	// Hash hashes the provided plain text string and returns the hashed value or
	// An error if the hashing process fails.
	Hash(word string) (string, error)

	// Match hashes the plainWord and compare that with hashedWord and
	// Returns: 1. true if the hashed string matches the plain text string 2.false if they do not match. 3. An error if the comparison process fails.
	Match(hashedWord, plainWord string) (bool, error)
}
