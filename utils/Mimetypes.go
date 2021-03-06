package utils

var allowedMimetypes = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
}

// IsMimetypeAllowed returns whether this mimetype is allowed to be uploaded
func IsMimetypeAllowed(mime string) bool {
	for s := range allowedMimetypes {
		if allowedMimetypes[s] == mime {
			return true
		}
	}
	return false
}