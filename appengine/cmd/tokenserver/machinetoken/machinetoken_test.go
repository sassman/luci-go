// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package machinetoken

import (
	"crypto/x509"
	"math/big"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/server/auth/signing"

	"github.com/luci/luci-go/common/api/tokenserver"
	"github.com/luci/luci-go/common/api/tokenserver/admin/v1"

	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMintParamsValidation(t *testing.T) {
	Convey("with token params", t, func() {
		params := MintParams{
			FQDN: "host.domain",
			Cert: &x509.Certificate{
				SerialNumber: big.NewInt(12345),
			},
			Config: &admin.CertificateAuthorityConfig{
				KnownDomains: []*admin.DomainConfig{
					{
						Domain:               []string{"domain"},
						Location:             "test-location",
						MachineTokenLifetime: 3600,
					},
				},
			},
			SignerServiceAccount: "token-server@example.com",
		}

		Convey("good params", func() {
			So(params.Validate(), ShouldBeNil)
		})

		Convey("bad FQDN case", func() {
			params.FQDN = "HOST.domain"
			So(params.Validate(), ShouldErrLike, "expecting FQDN in lowercase")
		})

		Convey("bad FQDN", func() {
			params.FQDN = "host"
			So(params.Validate(), ShouldErrLike, "not a valid FQDN")
		})

		Convey("bad char in FQDN", func() {
			params.FQDN = "host@.domain"
			So(params.Validate(), ShouldErrLike, "forbidden character")
		})

		Convey("not whitelisted", func() {
			params.FQDN = "host.blah"
			So(params.Validate(), ShouldErrLike, "not whitelisted in the config")
		})

		Convey("tokens are not allowed", func() {
			params.Config.KnownDomains[0].Location = ""
			So(params.Validate(), ShouldErrLike, "are not allowed")
		})

		Convey("tokens are not allowed 2", func() {
			params.Config.KnownDomains[0].MachineTokenLifetime = 0
			So(params.Validate(), ShouldErrLike, "are not allowed")
		})

		Convey("bad SN", func() {
			params.Cert.SerialNumber = big.NewInt(-1)
			So(params.Validate(), ShouldErrLike, "invalid certificate serial number")
		})
	})
}

func TestMint(t *testing.T) {
	Convey("with mock context", t, func() {
		ctx := context.Background()
		ctx, _ = testclock.UseTime(ctx, time.Date(2015, time.February, 3, 4, 5, 6, 7, time.UTC))

		Convey("works", func() {
			params := MintParams{
				FQDN: "host.domain",
				Cert: &x509.Certificate{
					SerialNumber: big.NewInt(12345),
				},
				Config: &admin.CertificateAuthorityConfig{
					KnownDomains: []*admin.DomainConfig{
						{
							Domain:               []string{"domain"},
							Location:             "test-location",
							MachineTokenLifetime: 3600,
						},
					},
				},
				SignerServiceAccount: "token-server@example.com",
				Signer:               fakeSigner{},
			}
			body, token, err := Mint(ctx, params)
			So(err, ShouldBeNil)
			So(body, ShouldResemble, &tokenserver.MachineTokenBody{
				MachineId: "host@test-location",
				IssuedBy:  "token-server@example.com",
				IssuedAt:  1422936306,
				Lifetime:  3600,
				CaId:      0,
				CertSn:    12345,
			})
			So(token, ShouldEqual, "CjoKEmhvc3RAdGVzdC1sb2NhdGlvbhIYdG9rZW4tc2VydmV"+
				"yQGV4YW1wbGUuY29tGPKRwaYFIJAcMLlgEgZrZXlfaWQaCXNpZ25hdHVyZQ")
		})
	})
}

type fakeSigner struct{}

func (fakeSigner) SignBytes(c context.Context, blob []byte) (keyID string, sig []byte, err error) {
	return "key_id", []byte("signature"), nil
}

func (fakeSigner) Certificates(c context.Context) (*signing.PublicCertificates, error) {
	panic("not implemented yet")
}
