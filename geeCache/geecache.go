package geecache

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

type Getter interface {
	Get(key string) ([]byte, error)
}
