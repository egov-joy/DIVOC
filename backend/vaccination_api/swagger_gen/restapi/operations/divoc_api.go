// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/divoc/api/swagger_gen/models"
	"github.com/divoc/api/swagger_gen/restapi/operations/certificate_revoked"
	"github.com/divoc/api/swagger_gen/restapi/operations/certification"
	"github.com/divoc/api/swagger_gen/restapi/operations/configuration"
	"github.com/divoc/api/swagger_gen/restapi/operations/identity"
	"github.com/divoc/api/swagger_gen/restapi/operations/login"
	"github.com/divoc/api/swagger_gen/restapi/operations/report_side_effects"
	"github.com/divoc/api/swagger_gen/restapi/operations/side_effects"
	"github.com/divoc/api/swagger_gen/restapi/operations/vaccination"
)

// NewDivocAPI creates a new Divoc instance
func NewDivocAPI(spec *loads.Document) *DivocAPI {
	return &DivocAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		GetV1PingHandler: GetV1PingHandlerFunc(func(params GetV1PingParams) middleware.Responder {
			return middleware.NotImplemented("operation GetV1Ping has not yet been implemented")
		}),
		LoginPostV1AuthorizeHandler: login.PostV1AuthorizeHandlerFunc(func(params login.PostV1AuthorizeParams) middleware.Responder {
			return middleware.NotImplemented("operation login.PostV1Authorize has not yet been implemented")
		}),
		IdentityPostV1IdentityVerifyHandler: identity.PostV1IdentityVerifyHandlerFunc(func(params identity.PostV1IdentityVerifyParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation identity.PostV1IdentityVerify has not yet been implemented")
		}),
		CertificationBulkCertifyHandler: certification.BulkCertifyHandlerFunc(func(params certification.BulkCertifyParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.BulkCertify has not yet been implemented")
		}),
		CertificateRevokedCertificateRevokedHandler: certificate_revoked.CertificateRevokedHandlerFunc(func(params certificate_revoked.CertificateRevokedParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_revoked.CertificateRevoked has not yet been implemented")
		}),
		CertificationCertifyHandler: certification.CertifyHandlerFunc(func(params certification.CertifyParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.Certify has not yet been implemented")
		}),
		CertificationCertifyV3Handler: certification.CertifyV3HandlerFunc(func(params certification.CertifyV3Params, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.CertifyV3 has not yet been implemented")
		}),
		ReportSideEffectsCreateReportedSideEffectsHandler: report_side_effects.CreateReportedSideEffectsHandlerFunc(func(params report_side_effects.CreateReportedSideEffectsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation report_side_effects.CreateReportedSideEffects has not yet been implemented")
		}),
		EventsHandler: EventsHandlerFunc(func(params EventsParams) middleware.Responder {
			return middleware.NotImplemented("operation Events has not yet been implemented")
		}),
		GetCertificateHandler: GetCertificateHandlerFunc(func(params GetCertificateParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation GetCertificate has not yet been implemented")
		}),
		CertificationGetCertificateByCertificateIDHandler: certification.GetCertificateByCertificateIDHandlerFunc(func(params certification.GetCertificateByCertificateIDParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.GetCertificateByCertificateID has not yet been implemented")
		}),
		CertificationGetCertifyUploadErrorsHandler: certification.GetCertifyUploadErrorsHandlerFunc(func(params certification.GetCertifyUploadErrorsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.GetCertifyUploadErrors has not yet been implemented")
		}),
		CertificationGetCertifyUploadsHandler: certification.GetCertifyUploadsHandlerFunc(func(params certification.GetCertifyUploadsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.GetCertifyUploads has not yet been implemented")
		}),
		ConfigurationGetConfigurationHandler: configuration.GetConfigurationHandlerFunc(func(params configuration.GetConfigurationParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation configuration.GetConfiguration has not yet been implemented")
		}),
		ConfigurationGetCurrentProgramsHandler: configuration.GetCurrentProgramsHandlerFunc(func(params configuration.GetCurrentProgramsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation configuration.GetCurrentPrograms has not yet been implemented")
		}),
		VaccinationGetLoggedInUserInfoHandler: vaccination.GetLoggedInUserInfoHandlerFunc(func(params vaccination.GetLoggedInUserInfoParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation vaccination.GetLoggedInUserInfo has not yet been implemented")
		}),
		VaccinationGetPreEnrollmentHandler: vaccination.GetPreEnrollmentHandlerFunc(func(params vaccination.GetPreEnrollmentParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation vaccination.GetPreEnrollment has not yet been implemented")
		}),
		VaccinationGetPreEnrollmentsForFacilityHandler: vaccination.GetPreEnrollmentsForFacilityHandlerFunc(func(params vaccination.GetPreEnrollmentsForFacilityParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation vaccination.GetPreEnrollmentsForFacility has not yet been implemented")
		}),
		SideEffectsGetSideEffectsMetadataHandler: side_effects.GetSideEffectsMetadataHandlerFunc(func(params side_effects.GetSideEffectsMetadataParams) middleware.Responder {
			return middleware.NotImplemented("operation side_effects.GetSideEffectsMetadata has not yet been implemented")
		}),
		CertificationGetTestCertifyUploadErrorsHandler: certification.GetTestCertifyUploadErrorsHandlerFunc(func(params certification.GetTestCertifyUploadErrorsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.GetTestCertifyUploadErrors has not yet been implemented")
		}),
		CertificationGetTestCertifyUploadsHandler: certification.GetTestCertifyUploadsHandlerFunc(func(params certification.GetTestCertifyUploadsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.GetTestCertifyUploads has not yet been implemented")
		}),
		ConfigurationGetVaccinatorsHandler: configuration.GetVaccinatorsHandlerFunc(func(params configuration.GetVaccinatorsParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation configuration.GetVaccinators has not yet been implemented")
		}),
		CertificationRevokeCertificateHandler: certification.RevokeCertificateHandlerFunc(func(params certification.RevokeCertificateParams) middleware.Responder {
			return middleware.NotImplemented("operation certification.RevokeCertificate has not yet been implemented")
		}),
		CertificationTestBulkCertifyHandler: certification.TestBulkCertifyHandlerFunc(func(params certification.TestBulkCertifyParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.TestBulkCertify has not yet been implemented")
		}),
		CertificationTestCertifyHandler: certification.TestCertifyHandlerFunc(func(params certification.TestCertifyParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.TestCertify has not yet been implemented")
		}),
		CertificationUpdateCertificateHandler: certification.UpdateCertificateHandlerFunc(func(params certification.UpdateCertificateParams, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.UpdateCertificate has not yet been implemented")
		}),
		CertificationUpdateCertificateV3Handler: certification.UpdateCertificateV3HandlerFunc(func(params certification.UpdateCertificateV3Params, principal *models.JWTClaimBody) middleware.Responder {
			return middleware.NotImplemented("operation certification.UpdateCertificateV3 has not yet been implemented")
		}),

		HasRoleAuth: func(token string, scopes []string) (*models.JWTClaimBody, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (hasRole) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*DivocAPI Digital infra for vaccination certificates */
type DivocAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// HasRoleAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	HasRoleAuth func(string, []string) (*models.JWTClaimBody, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GetV1PingHandler sets the operation handler for the get v1 ping operation
	GetV1PingHandler GetV1PingHandler
	// LoginPostV1AuthorizeHandler sets the operation handler for the post v1 authorize operation
	LoginPostV1AuthorizeHandler login.PostV1AuthorizeHandler
	// IdentityPostV1IdentityVerifyHandler sets the operation handler for the post v1 identity verify operation
	IdentityPostV1IdentityVerifyHandler identity.PostV1IdentityVerifyHandler
	// CertificationBulkCertifyHandler sets the operation handler for the bulk certify operation
	CertificationBulkCertifyHandler certification.BulkCertifyHandler
	// CertificateRevokedCertificateRevokedHandler sets the operation handler for the certificate revoked operation
	CertificateRevokedCertificateRevokedHandler certificate_revoked.CertificateRevokedHandler
	// CertificationCertifyHandler sets the operation handler for the certify operation
	CertificationCertifyHandler certification.CertifyHandler
	// CertificationCertifyV3Handler sets the operation handler for the certify v3 operation
	CertificationCertifyV3Handler certification.CertifyV3Handler
	// ReportSideEffectsCreateReportedSideEffectsHandler sets the operation handler for the create reported side effects operation
	ReportSideEffectsCreateReportedSideEffectsHandler report_side_effects.CreateReportedSideEffectsHandler
	// EventsHandler sets the operation handler for the events operation
	EventsHandler EventsHandler
	// GetCertificateHandler sets the operation handler for the get certificate operation
	GetCertificateHandler GetCertificateHandler
	// CertificationGetCertificateByCertificateIDHandler sets the operation handler for the get certificate by certificate Id operation
	CertificationGetCertificateByCertificateIDHandler certification.GetCertificateByCertificateIDHandler
	// CertificationGetCertifyUploadErrorsHandler sets the operation handler for the get certify upload errors operation
	CertificationGetCertifyUploadErrorsHandler certification.GetCertifyUploadErrorsHandler
	// CertificationGetCertifyUploadsHandler sets the operation handler for the get certify uploads operation
	CertificationGetCertifyUploadsHandler certification.GetCertifyUploadsHandler
	// ConfigurationGetConfigurationHandler sets the operation handler for the get configuration operation
	ConfigurationGetConfigurationHandler configuration.GetConfigurationHandler
	// ConfigurationGetCurrentProgramsHandler sets the operation handler for the get current programs operation
	ConfigurationGetCurrentProgramsHandler configuration.GetCurrentProgramsHandler
	// VaccinationGetLoggedInUserInfoHandler sets the operation handler for the get logged in user info operation
	VaccinationGetLoggedInUserInfoHandler vaccination.GetLoggedInUserInfoHandler
	// VaccinationGetPreEnrollmentHandler sets the operation handler for the get pre enrollment operation
	VaccinationGetPreEnrollmentHandler vaccination.GetPreEnrollmentHandler
	// VaccinationGetPreEnrollmentsForFacilityHandler sets the operation handler for the get pre enrollments for facility operation
	VaccinationGetPreEnrollmentsForFacilityHandler vaccination.GetPreEnrollmentsForFacilityHandler
	// SideEffectsGetSideEffectsMetadataHandler sets the operation handler for the get side effects metadata operation
	SideEffectsGetSideEffectsMetadataHandler side_effects.GetSideEffectsMetadataHandler
	// CertificationGetTestCertifyUploadErrorsHandler sets the operation handler for the get test certify upload errors operation
	CertificationGetTestCertifyUploadErrorsHandler certification.GetTestCertifyUploadErrorsHandler
	// CertificationGetTestCertifyUploadsHandler sets the operation handler for the get test certify uploads operation
	CertificationGetTestCertifyUploadsHandler certification.GetTestCertifyUploadsHandler
	// ConfigurationGetVaccinatorsHandler sets the operation handler for the get vaccinators operation
	ConfigurationGetVaccinatorsHandler configuration.GetVaccinatorsHandler
	// CertificationRevokeCertificateHandler sets the operation handler for the revoke certificate operation
	CertificationRevokeCertificateHandler certification.RevokeCertificateHandler
	// CertificationTestBulkCertifyHandler sets the operation handler for the test bulk certify operation
	CertificationTestBulkCertifyHandler certification.TestBulkCertifyHandler
	// CertificationTestCertifyHandler sets the operation handler for the test certify operation
	CertificationTestCertifyHandler certification.TestCertifyHandler
	// CertificationUpdateCertificateHandler sets the operation handler for the update certificate operation
	CertificationUpdateCertificateHandler certification.UpdateCertificateHandler
	// CertificationUpdateCertificateV3Handler sets the operation handler for the update certificate v3 operation
	CertificationUpdateCertificateV3Handler certification.UpdateCertificateV3Handler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *DivocAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *DivocAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *DivocAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DivocAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DivocAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DivocAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DivocAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DivocAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DivocAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DivocAPI
func (o *DivocAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.HasRoleAuth == nil {
		unregistered = append(unregistered, "HasRoleAuth")
	}

	if o.GetV1PingHandler == nil {
		unregistered = append(unregistered, "GetV1PingHandler")
	}
	if o.LoginPostV1AuthorizeHandler == nil {
		unregistered = append(unregistered, "login.PostV1AuthorizeHandler")
	}
	if o.IdentityPostV1IdentityVerifyHandler == nil {
		unregistered = append(unregistered, "identity.PostV1IdentityVerifyHandler")
	}
	if o.CertificationBulkCertifyHandler == nil {
		unregistered = append(unregistered, "certification.BulkCertifyHandler")
	}
	if o.CertificateRevokedCertificateRevokedHandler == nil {
		unregistered = append(unregistered, "certificate_revoked.CertificateRevokedHandler")
	}
	if o.CertificationCertifyHandler == nil {
		unregistered = append(unregistered, "certification.CertifyHandler")
	}
	if o.CertificationCertifyV3Handler == nil {
		unregistered = append(unregistered, "certification.CertifyV3Handler")
	}
	if o.ReportSideEffectsCreateReportedSideEffectsHandler == nil {
		unregistered = append(unregistered, "report_side_effects.CreateReportedSideEffectsHandler")
	}
	if o.EventsHandler == nil {
		unregistered = append(unregistered, "EventsHandler")
	}
	if o.GetCertificateHandler == nil {
		unregistered = append(unregistered, "GetCertificateHandler")
	}
	if o.CertificationGetCertificateByCertificateIDHandler == nil {
		unregistered = append(unregistered, "certification.GetCertificateByCertificateIDHandler")
	}
	if o.CertificationGetCertifyUploadErrorsHandler == nil {
		unregistered = append(unregistered, "certification.GetCertifyUploadErrorsHandler")
	}
	if o.CertificationGetCertifyUploadsHandler == nil {
		unregistered = append(unregistered, "certification.GetCertifyUploadsHandler")
	}
	if o.ConfigurationGetConfigurationHandler == nil {
		unregistered = append(unregistered, "configuration.GetConfigurationHandler")
	}
	if o.ConfigurationGetCurrentProgramsHandler == nil {
		unregistered = append(unregistered, "configuration.GetCurrentProgramsHandler")
	}
	if o.VaccinationGetLoggedInUserInfoHandler == nil {
		unregistered = append(unregistered, "vaccination.GetLoggedInUserInfoHandler")
	}
	if o.VaccinationGetPreEnrollmentHandler == nil {
		unregistered = append(unregistered, "vaccination.GetPreEnrollmentHandler")
	}
	if o.VaccinationGetPreEnrollmentsForFacilityHandler == nil {
		unregistered = append(unregistered, "vaccination.GetPreEnrollmentsForFacilityHandler")
	}
	if o.SideEffectsGetSideEffectsMetadataHandler == nil {
		unregistered = append(unregistered, "side_effects.GetSideEffectsMetadataHandler")
	}
	if o.CertificationGetTestCertifyUploadErrorsHandler == nil {
		unregistered = append(unregistered, "certification.GetTestCertifyUploadErrorsHandler")
	}
	if o.CertificationGetTestCertifyUploadsHandler == nil {
		unregistered = append(unregistered, "certification.GetTestCertifyUploadsHandler")
	}
	if o.ConfigurationGetVaccinatorsHandler == nil {
		unregistered = append(unregistered, "configuration.GetVaccinatorsHandler")
	}
	if o.CertificationRevokeCertificateHandler == nil {
		unregistered = append(unregistered, "certification.RevokeCertificateHandler")
	}
	if o.CertificationTestBulkCertifyHandler == nil {
		unregistered = append(unregistered, "certification.TestBulkCertifyHandler")
	}
	if o.CertificationTestCertifyHandler == nil {
		unregistered = append(unregistered, "certification.TestCertifyHandler")
	}
	if o.CertificationUpdateCertificateHandler == nil {
		unregistered = append(unregistered, "certification.UpdateCertificateHandler")
	}
	if o.CertificationUpdateCertificateV3Handler == nil {
		unregistered = append(unregistered, "certification.UpdateCertificateV3Handler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DivocAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DivocAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "hasRole":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.HasRoleAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *DivocAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *DivocAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *DivocAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DivocAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the divoc API
func (o *DivocAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DivocAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/ping"] = NewGetV1Ping(o.context, o.GetV1PingHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/authorize"] = login.NewPostV1Authorize(o.context, o.LoginPostV1AuthorizeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/identity/verify"] = identity.NewPostV1IdentityVerify(o.context, o.IdentityPostV1IdentityVerifyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/bulkCertify"] = certification.NewBulkCertify(o.context, o.CertificationBulkCertifyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/certificate/revoked"] = certificate_revoked.NewCertificateRevoked(o.context, o.CertificateRevokedCertificateRevokedHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/certify"] = certification.NewCertify(o.context, o.CertificationCertifyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/certify"] = certification.NewCertifyV3(o.context, o.CertificationCertifyV3Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/report-side-effects"] = report_side_effects.NewCreateReportedSideEffects(o.context, o.ReportSideEffectsCreateReportedSideEffectsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/events"] = NewEvents(o.context, o.EventsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/certificates"] = NewGetCertificate(o.context, o.GetCertificateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/certificates/{certificateId}"] = certification.NewGetCertificateByCertificateID(o.context, o.CertificationGetCertificateByCertificateIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/certify/uploads/{uploadId}/errors"] = certification.NewGetCertifyUploadErrors(o.context, o.CertificationGetCertifyUploadErrorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/certify/uploads"] = certification.NewGetCertifyUploads(o.context, o.CertificationGetCertifyUploadsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/divoc/configuration"] = configuration.NewGetConfiguration(o.context, o.ConfigurationGetConfigurationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/programs/current"] = configuration.NewGetCurrentPrograms(o.context, o.ConfigurationGetCurrentProgramsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/users/me"] = vaccination.NewGetLoggedInUserInfo(o.context, o.VaccinationGetLoggedInUserInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/preEnrollments/{preEnrollmentCode}"] = vaccination.NewGetPreEnrollment(o.context, o.VaccinationGetPreEnrollmentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/preEnrollments"] = vaccination.NewGetPreEnrollmentsForFacility(o.context, o.VaccinationGetPreEnrollmentsForFacilityHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/sideEffects"] = side_effects.NewGetSideEffectsMetadata(o.context, o.SideEffectsGetSideEffectsMetadataHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/test/certify/uploads/{uploadId}/errors"] = certification.NewGetTestCertifyUploadErrors(o.context, o.CertificationGetTestCertifyUploadErrorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/test/certify/uploads"] = certification.NewGetTestCertifyUploads(o.context, o.CertificationGetTestCertifyUploadsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/vaccinators"] = configuration.NewGetVaccinators(o.context, o.ConfigurationGetVaccinatorsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/certificates/{preEnrollmentCode}"] = certification.NewRevokeCertificate(o.context, o.CertificationRevokeCertificateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/test/bulkCertify"] = certification.NewTestBulkCertify(o.context, o.CertificationTestBulkCertifyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/test/certify"] = certification.NewTestCertify(o.context, o.CertificationTestCertifyHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/certificate"] = certification.NewUpdateCertificate(o.context, o.CertificationUpdateCertificateHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v3/certificate"] = certification.NewUpdateCertificateV3(o.context, o.CertificationUpdateCertificateV3Handler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DivocAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *DivocAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DivocAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DivocAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *DivocAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
