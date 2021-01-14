/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2017 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package encoding

import (
	"context"
	"encoding/base64"

	"github.com/loadimpact/k6/js/common"
	"github.com/loadimpact/k6/js/modules"
)

func init() {
	modules.Register("k6/x/encoding", New())
}

type Encoding struct{}

func New() *Encoding {
	return &Encoding{}
}

func (e *Encoding) Stringify(ctx context.Context, input []byte) string {
	return string(input)
}

func (e *Encoding) B64encode(ctx context.Context, input []byte, encoding string) []byte {
	var enc *base64.Encoding
	switch encoding {
	case "rawstd":
		enc = base64.StdEncoding.WithPadding(base64.NoPadding)
	case "rawurl":
		enc = base64.URLEncoding.WithPadding(base64.NoPadding)
	case "url":
		enc = base64.URLEncoding
	default:
		enc = base64.StdEncoding
	}
	out := make([]byte, enc.EncodedLen(len(input)))
	enc.Encode(out, input)
	return out
}

func (e *Encoding) B64decode(ctx context.Context, input []byte, encoding string) []byte {
	var enc *base64.Encoding
	switch encoding {
	case "rawstd":
		enc = base64.StdEncoding.WithPadding(base64.NoPadding)
	case "rawurl":
		enc = base64.URLEncoding.WithPadding(base64.NoPadding)
	case "url":
		enc = base64.URLEncoding
	default:
		enc = base64.StdEncoding
	}
	output := make([]byte, enc.DecodedLen(len(input)))
	n, err := enc.Decode(output, input)
	if err != nil {
		common.Throw(common.GetRuntime(ctx), err)
	}

	return output[:n]
}
