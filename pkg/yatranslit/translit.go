package yatranslit

type Translit struct {
	strategy  Translater
	maxLength int
}

func (t *Translit) SetStrategy(strategy Translater) {
	if strategy == nil {
		strategy = NewDictionaryStrategy()
	}
	t.strategy = strategy
}

func NewTranslit() *Translit {
	return &Translit{
		strategy:  NewDictionaryStrategy(),
		maxLength: -1,
	}
}

func (t *Translit) SetMaxLength(maxLength int) {
	t.maxLength = maxLength
}

func (t *Translit) Transform(text string) string {
	return t.strategy.Translate(text, t.maxLength)
}
