package nekolog

type Formatter interface {
	Format(e *entry) error
}
