package service_url

type ReadUrlParameters struct {
	GetUrlRepository func(string) (ReadUrlData, error)
}
type ReadUrlData struct {
	Original string
	Id       int
}

func ReadUrl(p ReadUrlParameters) func(string) (string, error) {
	return func(token string) (string, error) {
		url, err := p.GetUrlRepository(token)
		if err != nil {
			return "", err
		}
		return url.Original, nil
	}
}
