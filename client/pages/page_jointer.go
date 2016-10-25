package pages

// A PageJoinerFunc collapses multiple pages into a single page
type PageJoinerFunc func([][]byte) []byte
