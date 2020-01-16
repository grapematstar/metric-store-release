package rulesclient_test

import (
	"crypto/tls"
	"net/http"

	. "github.com/cloudfoundry/metric-store-release/src/pkg/rulesclient"

	"github.com/cloudfoundry/metric-store-release/src/internal/testing"
	sharedtls "github.com/cloudfoundry/metric-store-release/src/pkg/tls"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type rulesClientTestContext struct {
	tlsConfig *tls.Config
	rulesApi  *testing.RulesApiSpy
}

var _ = Describe("RulesClient", func() {
	var setup = func() *rulesClientTestContext {
		tlsConfig, err := sharedtls.NewMutualTLSConfig(
			testing.Cert("metric-store-ca.crt"),
			testing.Cert("metric-store.crt"),
			testing.Cert("metric-store.key"),
			"metric-store",
		)
		Expect(err).ToNot(HaveOccurred())

		rulesApi, err := testing.NewRulesApiSpy(tlsConfig)
		Expect(err).ToNot(HaveOccurred())

		err = rulesApi.Start()
		Expect(err).ToNot(HaveOccurred())

		return &rulesClientTestContext{
			tlsConfig: tlsConfig,
			rulesApi:  rulesApi,
		}
	}

	Describe("#CreateManager", func() {
		It("create a rules manager", func() {
			tc := setup()

			client := NewRulesClient(tc.rulesApi.Addr(), tc.tlsConfig)
			manager, err := client.CreateManager("app-metrics", "")
			Expect(err).ToNot(HaveOccurred())

			Expect(tc.rulesApi.RequestsReceived()).To(Equal(1))

			Expect(manager).ToNot(BeNil())
			Expect(manager.Id).To(Equal("app-metrics"))
		})

		It("returns an error when manager creation failed", func() {
			tc := setup()

			tc.rulesApi.NextRequestError(&testing.RulesApiHttpError{
				Status: http.StatusConflict,
				Title:  "Error Occurred",
			})

			client := NewRulesClient(tc.rulesApi.Addr(), tc.tlsConfig)
			_, err := client.CreateManager("", "")
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(&ErrorNotCreated{
				Title: "Error Occurred",
			}))
		})

		It("server unavailable", func() {
			tlsConfig, err := sharedtls.NewMutualTLSConfig(
				testing.Cert("metric-store-ca.crt"),
				testing.Cert("metric-store.crt"),
				testing.Cert("metric-store.key"),
				"metric-store",
			)
			Expect(err).ToNot(HaveOccurred())

			client := NewRulesClient("localhost:10", tlsConfig)
			_, err = client.CreateManager("app-metrics", "")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("#UpsertRuleGroup", func() {
		It("creates a new rule group", func() {
			tc := setup()

			client := NewRulesClient(tc.rulesApi.Addr(), tc.tlsConfig)
			_, err := client.CreateManager("app-metrics", "")
			Expect(err).ToNot(HaveOccurred())

			ruleGroup, err := client.UpsertRuleGroup("app-metrics", RuleGroup{
				Name: "groupOne",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(tc.rulesApi.RequestsReceived()).To(Equal(2))
			Expect(ruleGroup).ToNot(BeNil())
		})

		It("returns an error when group creation fails", func() {
			tc := setup()
			client := NewRulesClient(tc.rulesApi.Addr(), tc.tlsConfig)

			tc.rulesApi.NextRequestError(&testing.RulesApiHttpError{
				Status: http.StatusConflict,
				Title:  "Error Occurred",
			})

			_, err := client.UpsertRuleGroup("app-metrics", RuleGroup{
				Name: "groupOne",
			})
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(&ErrorNotCreated{
				Title: "Error Occurred",
			}))
		})

		It("server unavailable", func() {
			tlsConfig, err := sharedtls.NewMutualTLSConfig(
				testing.Cert("metric-store-ca.crt"),
				testing.Cert("metric-store.crt"),
				testing.Cert("metric-store.key"),
				"metric-store",
			)
			Expect(err).ToNot(HaveOccurred())

			client := NewRulesClient("localhost:10", tlsConfig)
			_, err = client.UpsertRuleGroup("app-metrics", RuleGroup{
				Name: "groupOne",
			})
			Expect(err).To(HaveOccurred())
		})
	})
})