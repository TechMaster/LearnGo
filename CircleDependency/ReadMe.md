```go
import cycle not allowed
package main
	imports github.com/TechMaster/LearnGo/CircleDependency/apackage
	imports github.com/TechMaster/LearnGo/CircleDependency/bpackage
	imports github.com/TechMaster/LearnGo/CircleDependency/apackage
```

Vấn đề ở đây là gì?

apackage import bpackage. Sau đó bpackage lại import apackage.
```
      +------------------+
      |                  |
 +----------+       +----v-----+
 | apackage |       | bpackage |
 +----^-----+       +----------+
      |                  |
      |                  |
      +------------------+
```