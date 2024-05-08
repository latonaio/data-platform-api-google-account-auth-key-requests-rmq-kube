package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-google-account-auth-key-requests-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-google-account-auth-key-requests-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-google-account-auth-key-requests-rmq-kube/config"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"net/url"
)

func (c *DPFMAPICaller) GoogleAccountAuthKey(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.GoogleAccountAuthKey {
	var googleAccountAuthKey []dpfm_api_output_formatter.GoogleAccountAuthKey

	inputURL := input.GoogleAccountAuthKey.URL

	parsedURL, err := url.Parse(inputURL)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	code := parsedURL.Query().Get("code")

	if code == "" {
		*errs = append(*errs, xerrors.Errorf("URL does not contain Code"))
		return nil
	}

	generatedGoogleAccountAuthKey := conf.OAuth.GenerateOAuthTokenRequestURL(code)

	googleAccountAuthKey = append(googleAccountAuthKey, dpfm_api_output_formatter.GoogleAccountAuthKey{
		URL: generatedGoogleAccountAuthKey,
	})

	return &googleAccountAuthKey
}
