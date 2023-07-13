package catclient

type CatClient struct {
	count int
}

func NewCatClient() *CatClient {
	return &CatClient{}
}

func (c *CatClient) HTTPGet(name string) (int, string) {
	c.count++
	switch c.count {
	case 1:
		return 200, "Hello " + name
	case 2:
		return 200, "Bye " + name
	}
	return 403, "limit exceeded"
}
