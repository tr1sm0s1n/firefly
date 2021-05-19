// Copyright © 2021 Kaleido, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiserver

import (
	"net/http"

	"github.com/kaleido-io/firefly/internal/config"
	"github.com/kaleido-io/firefly/pkg/database"
	"github.com/kaleido-io/firefly/pkg/fftypes"
	"github.com/kaleido-io/firefly/internal/i18n"
	"github.com/kaleido-io/firefly/internal/oapispec"
)

var getTxnById = &oapispec.Route{
	Name:   "getTxnById",
	Path:   "namespaces/{ns}/transactions/{txnid}",
	Method: http.MethodGet,
	PathParams: []oapispec.PathParam{
		{Name: "ns", ExampleFromConf: config.NamespacesDefault, Description: i18n.MsgTBD},
		{Name: "txnid", Description: i18n.MsgTBD},
	},
	QueryParams:     nil,
	FilterFactory:   database.TransactionQueryFactory,
	Description:     i18n.MsgTBD,
	JSONInputValue:  func() interface{} { return nil },
	JSONOutputValue: func() interface{} { return &fftypes.Transaction{} },
	JSONOutputCode:  http.StatusOK,
	JSONHandler: func(r oapispec.APIRequest) (output interface{}, err error) {
		output, err = r.Or.GetTransactionById(r.Ctx, r.PP["ns"], r.PP["txnid"])
		return output, err
	},
}