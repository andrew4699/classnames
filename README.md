# classnames
## Simple utility for conditionally joining HTML class names together (based on [JedWatson/classnames](https://github.com/JedWatson/classnames))

# Install
```
go get github.com/andrew4699/classnames
```

# Usage
Note that __classnames.B__ is shorthand for __classnames.Build__

```
...

import (
    "github.com/andrew4699/classnames"
)

classnames.B("foo", "bar") // "foo bar"
classnames.B([]string{"foo", "bar"}, "duck") // "foo bar duck"
classnames.B(nil, "foo", []string{"bar"}, nil, nil, nil) // "foo bar"

myMap := classnames.Map{
    "foo": true,
    "bar": false,
    "duck": true,
}

classnames.B(myMap) // "foo duck"

...
```