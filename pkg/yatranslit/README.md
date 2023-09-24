# Yandex-standard russian-english transliteration
Golang adaptation of https://github.com/denismitr/translit


## Usage
```go
    translit := translit_yandex.NewTranslit()
    var out string = translit.Transform("Строка для транслитерации, по правилам Яндекс!")
	// out == "stroka-dlya-transliteracii-po-pravilam-yandeksa"
```
To define max length of the output do:
```go
    translit := translit_yandex.NewTranslit()
	translit.SetMaxLenght(10)
    slug := translit.transform("очень длинный текст...");
    // slug == "ochen-dlin"
```

You can provide your own translit implementation as long as it implements the 
`Translater` interface
and inject it into the `Translit` class constructor, this way it will override the default behavior like so:
```go
    CustomTranslit := NewTranslit()
	CustomTranslit.SetStrategy(NewYourTranslitStrategyImpl())
```

