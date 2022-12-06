## Ports
The ports will be placed in the directory ./internal/core/ports. It contains the interfaces definition used to communicate with actors.

For Example:
```
type GamesRepository interface {
	  Get(id string) (domain.Game, error)
	  Save(domain.Game) error
}

type GamesService interface {
	  Get(id string) (domain.Game, error)
	  Create(name string, size uint, bombs uint) (domain.Game, error)
	  Reveal(id string, row uint, col uint) (domain.Game, error)
}
```