// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "encoding/json"
    "testing"

    "github.com/algolia/algoliasearch-client-go/algolia/opt"
    "github.com/stretchr/testify/require"
)

func TestOptionalWords(t *testing.T) {
    for _, c := range []struct {
        opts     []interface{}
        expected opt.OptionalWordsOption
    }{
        {
            opts:     []interface{}{nil},
            expected: opt.OptionalWords(nil),
        },
        {
            opts:     []interface{}{opt.OptionalWords("value1")},
            expected: opt.OptionalWords("value1"),
        },
        {
            opts:     []interface{}{opt.OptionalWords("value1", "value2", "value3")},
            expected: opt.OptionalWords("value1", "value2", "value3"),
        },
    } {
        var (
            in  = ExtractOptionalWords(c.opts...)
            out opt.OptionalWordsOption
        )
        data, err := json.Marshal(&in)
        require.NoError(t, err)
        err = json.Unmarshal(data, &out)
        require.NoError(t, err)
        require.ElementsMatch(t, c.expected.Get(), out.Get())
    }
}