Swagger
---

The Swagger Codegen for golang doesn't create perfect output.
There will be a few different things that you will have to modify after code-gen has been completed.

# optional

For some reason, the `import` for the `optional` packge doesn't seem to be included in certain
files. You will need to add these manually.

```golang
import (
	"github.com/antihax/optional"
 )
```

# Delete duplicate definition

When `swagger-codegen` generates the API, it creates a duplicate of `AddConnectionOpts` so you will
need to comment it out in the generated code. They are both in the same namespace so as long as
there is a single definition, we should be OK.
