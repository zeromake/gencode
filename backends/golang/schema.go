package golang

import (
	"fmt"

	"github.com/zeromake/gencode/schema"
)

type Walker struct {
	Needs     []string
	Offset    int
	IAdjusted bool
	Unsafe    bool
}

func (w *Walker) WalkSchema(s *schema.Schema, Package string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`package %s
	
	import (
		"unsafe"
		"io"
		"time"
	)
	`, Package))
	for _, st := range s.Structs {
		p, err := w.WalkStruct(st)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	return
}
