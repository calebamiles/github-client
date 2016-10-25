package pages

import "bytes"

// Join collapses a slice of JSON byte slices into a single slice of JSON
func Join(pages [][]byte) []byte {
	switch len(pages) {
	case 0:
		return []byte{}
	case 1:
		return pages[0]
	case 2:
		firstPage := pages[0]
		firstPage = bytes.TrimSpace(firstPage)
		firstPage = bytes.TrimSuffix(firstPage, []byte("]"))
		firstPage = append(firstPage, ',')

		lastPage := pages[len(pages)-1]
		lastPage = bytes.TrimSpace(lastPage)
		lastPage = bytes.TrimPrefix(lastPage, []byte("["))
		return append(firstPage, lastPage...)
	default:
		firstPage := pages[0]
		firstPage = bytes.TrimSpace(firstPage)
		firstPage = bytes.TrimSuffix(firstPage, []byte("]"))
		firstPage = append(firstPage, ',')

		lastPage := pages[len(pages)-1]
		lastPage = bytes.TrimSpace(lastPage)
		lastPage = bytes.TrimPrefix(lastPage, []byte("["))

		collapsedPages := firstPage
		for i := 1; i < len(pages)-1; i++ {
			middlePage := pages[i]
			middlePage = bytes.TrimSpace(middlePage)
			middlePage = bytes.Trim(middlePage, "[]")
			middlePage = append(middlePage, ',')
			collapsedPages = append(collapsedPages, middlePage...)
		}

		collapsedPages = append(collapsedPages, lastPage...)

		return collapsedPages
	}
}
