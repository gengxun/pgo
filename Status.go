package pgo

import (
    "fmt"
    "net/http"

    "github.com/pinguo/pgo/Util"
)

// status component, configuration:
// {
//     "useI18n": false,
//     "mapping": {
//         "11002": "Verify Sign Error"
//     }
// }
type Status struct {
    useI18n bool
    mapping map[int]string
}

func (s *Status) Construct() {
    s.useI18n = false
    s.mapping = make(map[int]string)
}

func (s *Status) SetUseI18n(useI18n bool) {
    s.useI18n = useI18n
}

func (s *Status) SetMapping(m map[string]interface{}) {
    for k, v := range m {
        s.mapping[Util.ToInt(k)] = Util.ToString(v)
    }
}

func (s *Status) GetText(status int, ctx *Context, dft ...string) string {
    txt, ok := s.mapping[status]
    if !ok {
        if len(dft) == 0 || len(dft[0]) == 0 {
            if txt = http.StatusText(status); len(txt) == 0 {
                panic(fmt.Sprintf("unknown status code: %d", status))
            }
        } else {
            txt = dft[0]
        }
    }

    if s.useI18n && ctx != nil {
        al := ctx.GetHeader("Accept-Language", "")
        txt = App.GetI18n().Translate(txt, al)
    }

    return txt
}
