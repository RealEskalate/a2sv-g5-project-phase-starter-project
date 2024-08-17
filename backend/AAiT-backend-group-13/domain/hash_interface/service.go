package ihash

type Service interface {
	Hash(word string) (string, error)
	Match(hashedWord, plainWord string) (bool, error)
}
