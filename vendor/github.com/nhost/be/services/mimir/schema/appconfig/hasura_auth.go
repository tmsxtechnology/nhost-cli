package appconfig

import (
	"encoding/json"
	"fmt"

	"github.com/nhost/be/services/mimir/model"
)

const (
	secretHasuraAuthGraphqlDatabaseURL      = "databaseUrl"
	secretHasuraAuthDatabaseMigrationsURL   = "databaseMigrationsUrl"
	secretHasuraAuthHasuraAdminSecret       = "adminSecret"
	secretHasuraAuthJWTSecret               = "jwtSecret"
	secretHasuraAuthGithubClientSecret      = "githubClientSecret" //nolint: gosec
	secretHasuraAuthGoogleClientSecret      = "googleClientSecret"
	secretHasuraAuthFacebookClientSecret    = "facebookClientSecret"
	secretHasuraAuthSpotifyClientSecret     = "spotifyClientSecret"
	secretHasuraAuthLinkedinClientSecret    = "linkedinClientSecret"
	secretHasuraAuthDiscordClientSecret     = "discordClientSecret"
	secretHasuraAuthTwitchClientSecret      = "twitchClientSecret"
	secretHasuraAuthWindowsLiveClientSecret = "windowsLiveClientSecret"
	secretHasuraAuthWorkOsClientSecret      = "workOsClientSecret"
	secretHasuraAuthTwitterConsumerSecret   = "twitterConsumerSecret"
	secretHasuraAuthAppleKeyID              = "appleKeyID"
	secretHasuraAuthApplePrivateKey         = "applePrivateKey"
	secretHasuraAuthAzureADClientSecret     = "azureADClientSecret"
	secretHasuraAuthGitlabClientSecret      = "gitlabClientSecret" //nolint: gosec
	secretHasuraAuthStravaClientSecret      = "stravaClientSecret" //nolint: gosec
	secretHasuraAuthBitbucketClientSecret   = "bitbucketClientSecret"
	secretHasuraAuthSMTPPassword            = "smtpPassword"
)

func HasuraAuthEnv( //nolint:funlen,cyclop,maintidx
	config *model.ConfigConfig,
	hasuraGraphqlURL,
	authServerURL,
	databaseURL string,
	databaseMigrationURL string,
	smtpSettings *model.ConfigSmtp,
	isCustomSMTP bool,
	autoScalerEnabled bool,
	appID string,
) ([]EnvVar, error) {
	customClaims := make(
		map[string]string,
		len(config.GetAuth().Session.AccessToken.CustomClaims),
	)
	for _, c := range config.GetAuth().Session.AccessToken.CustomClaims {
		customClaims[c.Key] = c.Value
	}

	authJwtCustomClaims, err := json.Marshal(customClaims)
	if err != nil {
		return nil, fmt.Errorf("problem marshalling auth jwt custom claims: %w", err)
	}

	jwtSecret, err := marshalJWT(config.GetHasura().GetJwtSecrets()[0])
	if err != nil {
		return nil, fmt.Errorf("could not marshal JWT secret: %w", err)
	}

	replicas := 1
	if !autoScalerEnabled && config.GetAuth().GetResources().GetReplicas() != nil {
		replicas = int(*config.GetAuth().GetResources().GetReplicas())
	}

	dbURL := databaseURL
	version := *config.GetAuth().GetVersion()
	if version != "0.0.0-dev" && CompareVersions(version, "0.30.999999") <= 0 {
		dbURL = databaseMigrationURL
	}
	env := []EnvVar{
		{
			Name:       "HASURA_GRAPHQL_DATABASE_URL",
			SecretName: secretHasuraAuthGraphqlDatabaseURL,
			Value:      dbURL,
			IsSecret:   true,
		},
		{
			Name:       "POSTGRES_MIGRATIONS_CONNECTION",
			SecretName: secretHasuraAuthDatabaseMigrationsURL,
			Value:      databaseMigrationURL,
			IsSecret:   true,
		},
		{
			Name:       "HASURA_GRAPHQL_ADMIN_SECRET",
			SecretName: secretHasuraAuthHasuraAdminSecret,
			Value:      config.GetHasura().GetAdminSecret(),
			IsSecret:   true,
		},
		{
			Name:       "HASURA_GRAPHQL_JWT_SECRET",
			SecretName: secretHasuraAuthJWTSecret,
			Value:      string(jwtSecret),
			IsSecret:   true,
		},
		{
			Name:       "HASURA_GRAPHQL_GRAPHQL_URL",
			Value:      hasuraGraphqlURL,
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ACCESS_TOKEN_EXPIRES_IN",
			Value: Stringify(
				unptr(
					config.GetAuth().
						GetSession().
						GetAccessToken().
						GetExpiresIn(),
				),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_REFRESH_TOKEN_EXPIRES_IN",
			Value: Stringify(
				unptr(
					config.
						GetAuth().
						GetSession().
						GetRefreshToken().
						GetExpiresIn(),
				),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_SERVER_URL",
			Value:      authServerURL,
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_CLIENT_URL",
			Value: unptr(
				config.GetAuth().GetRedirections().GetClientUrl(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_PORT",
			Value:      "4000",
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_HOST",
			Value:      "0.0.0.0",
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_API_PREFIX",
			Value:      "/v1",
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ANONYMOUS_USERS_ENABLED",
			Value: Stringify(
				unptr(
					config.
						GetAuth().
						GetMethod().
						GetAnonymous().
						GetEnabled(),
				),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_USER_DEFAULT_ROLE",
			Value:      *config.GetAuth().GetUser().GetRoles().GetDefault(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_PASSWORD_HIBP_ENABLED",
			Value: Stringify(
				*config.GetAuth().GetMethod().GetEmailPassword().GetHibpEnabled(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_PASSWORD_MIN_LENGTH",
			Value: Stringify(
				*config.GetAuth().GetMethod().GetEmailPassword().GetPasswordMinLength(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_USER_DEFAULT_ALLOWED_ROLES",
			Value: Stringify(
				config.GetAuth().GetUser().GetRoles().GetAllowed(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_MFA_ENABLED",
			Value:      Stringify(*config.GetAuth().GetTotp().GetEnabled()),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_MFA_TOTP_ISSUER",
			Value:      unptr(config.GetAuth().GetTotp().GetIssuer()),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_LOCALE_DEFAULT",
			Value: unptr(
				config.GetAuth().GetUser().GetLocale().GetDefault(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_LOCALE_ALLOWED_LOCALES",
			Value: Stringify(
				config.GetAuth().GetUser().GetLocale().GetAllowed(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ACCESS_CONTROL_ALLOWED_REDIRECT_URLS",
			Value: Stringify(
				config.GetAuth().GetRedirections().GetAllowedUrls(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ACCESS_CONTROL_ALLOWED_EMAILS",
			Value: Stringify(
				config.GetAuth().GetUser().GetEmail().GetAllowed(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ACCESS_CONTROL_ALLOWED_EMAIL_DOMAINS",
			Value: Stringify(
				config.GetAuth().GetUser().GetEmailDomains().GetAllowed(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ACCESS_CONTROL_BLOCKED_EMAILS",
			Value: Stringify(
				config.GetAuth().GetUser().GetEmail().GetBlocked(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_ACCESS_CONTROL_BLOCKED_EMAIL_DOMAINS",
			Value: Stringify(
				config.GetAuth().GetUser().GetEmailDomains().GetBlocked(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_JWT_CUSTOM_CLAIMS",
			Value:      string(authJwtCustomClaims),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_EMAIL_PASSWORDLESS_ENABLED",
			Value: Stringify(
				*config.GetAuth().GetMethod().GetEmailPasswordless().GetEnabled(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED",
			Value: Stringify(
				*config.GetAuth().GetMethod().GetEmailPassword().GetEmailVerificationRequired(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_DISABLE_SIGNUP",
			Value:      Stringify(!unptr(config.GetAuth().GetSignUp().GetEnabled())),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_DISABLE_NEW_USERS",
			Value: Stringify(
				!unptr(config.GetAuth().GetSignUp().GetEnabled()) ||
					unptr(config.GetAuth().GetSignUp().GetDisableNewUsers()),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_REQUIRE_ELEVATED_CLAIM",
			Value: Stringify(
				unptr(config.GetAuth().GetElevatedPrivileges().GetMode()),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_CONCEAL_ERRORS",
			Value:      Stringify(unptr(config.GetAuth().GetMisc().GetConcealErrors())),
			IsSecret:   false,
			SecretName: "",
		},
	}

	env = append(env, []EnvVar{
		{
			Name: "AUTH_GRAVATAR_ENABLED",
			Value: Stringify(
				unptr(
					config.GetAuth().GetUser().GetGravatar().GetEnabled(),
				),
			),
			IsSecret:   false,
			SecretName: "",
		},
	}...)
	if unptr(config.GetAuth().GetUser().GetGravatar().GetEnabled()) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_GRAVATAR_DEFAULT",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetUser().
							GetGravatar().
							GetDefault(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_GRAVATAR_RATING",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetUser().
							GetGravatar().
							GetRating(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetWebauthn().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_WEBAUTHN_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetWebauthn().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_WEBAUTHN_RP_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetWebauthn().
						GetRelyingParty().
						GetId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_WEBAUTHN_RP_NAME",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetWebauthn().
						GetRelyingParty().
						GetName(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_WEBAUTHN_RP_ORIGINS",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetWebauthn().
						GetRelyingParty().
						GetOrigins(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_WEBAUTHN_ATTESTATION_TIMEOUT",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetWebauthn().
							GetAttestation().
							GetTimeout(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetSmsPasswordless().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_SMS_PASSWORDLESS_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetSmsPasswordless().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_SMS_PROVIDER",
				Value: unptr(
					config.GetProvider().GetSms().GetProvider(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_SMS_TWILIO_ACCOUNT_SID",
				Value:      config.GetProvider().GetSms().GetAccountSid(),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_SMS_TWILIO_AUTH_TOKEN",
				Value:      config.GetProvider().GetSms().GetAuthToken(),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_SMS_TWILIO_MESSAGING_SERVICE_ID",
				Value: config.
					GetProvider().
					GetSms().
					GetMessagingServiceId(),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetOauth().GetGithub().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_GITHUB_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetGithub().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_GITHUB_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetGithub().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_GITHUB_CLIENT_SECRET",
				SecretName: secretHasuraAuthGithubClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetGithub().GetClientSecret(),
				),
				IsSecret: true,
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetGoogle().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_GOOGLE_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetGoogle().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_GOOGLE_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetGoogle().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_GOOGLE_CLIENT_SECRET",
				SecretName: secretHasuraAuthGoogleClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetGoogle().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_GOOGLE_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetGoogle().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetFacebook().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			// FACEBOOK
			{
				Name: "AUTH_PROVIDER_FACEBOOK_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetFacebook().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_FACEBOOK_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetFacebook().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_FACEBOOK_CLIENT_SECRET",
				SecretName: secretHasuraAuthFacebookClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetFacebook().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_FACEBOOK_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetFacebook().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetSpotify().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			// SPOTIFY
			{
				Name: "AUTH_PROVIDER_SPOTIFY_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetSpotify().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_SPOTIFY_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetSpotify().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_SPOTIFY_CLIENT_SECRET",
				SecretName: secretHasuraAuthSpotifyClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetSpotify().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_SPOTIFY_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetSpotify().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetLinkedin().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_LINKEDIN_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetLinkedin().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_LINKEDIN_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetLinkedin().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_LINKEDIN_CLIENT_SECRET",
				SecretName: secretHasuraAuthLinkedinClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetLinkedin().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_LINKEDIN_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetLinkedin().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetDiscord().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_DISCORD_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetDiscord().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_DISCORD_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetDiscord().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_DISCORD_CLIENT_SECRET",
				SecretName: secretHasuraAuthDiscordClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetDiscord().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_DISCORD_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetDiscord().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetTwitch().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_TWITCH_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetTwitch().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_TWITCH_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetTwitch().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_TWITCH_CLIENT_SECRET",
				SecretName: secretHasuraAuthTwitchClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetTwitch().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_TWITCH_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetTwitch().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.
			GetAuth().
			GetMethod().
			GetOauth().
			GetWindowslive().
			GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_WINDOWS_LIVE_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetWindowslive().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_WINDOWS_LIVE_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetWindowslive().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_WINDOWS_LIVE_CLIENT_SECRET",
				SecretName: secretHasuraAuthWindowsLiveClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetWindowslive().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_WINDOWS_LIVE_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetWindowslive().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetOauth().GetWorkos().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_WORKOS_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetWorkos().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_WORKOS_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetWorkos().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_WORKOS_CLIENT_SECRET",
				SecretName: secretHasuraAuthWorkOsClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetWorkos().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_WORKOS_DEFAULT_ORGANIZATION",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetWorkos().
						GetOrganization(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_WORKOS_DEFAULT_CONNECTION",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetWorkos().
						GetConnection(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetOauth().GetTwitter().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			// TWITTER
			{
				Name: "AUTH_PROVIDER_TWITTER_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetTwitter().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_TWITTER_CONSUMER_KEY",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetTwitter().
						GetConsumerKey(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_TWITTER_CONSUMER_SECRET",
				SecretName: secretHasuraAuthTwitterConsumerSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetTwitter().GetConsumerSecret(),
				),
				IsSecret: true,
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetOauth().GetApple().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_APPLE_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetApple().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_APPLE_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetApple().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_APPLE_TEAM_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetApple().
						GetTeamId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_APPLE_KEY_ID",
				SecretName: secretHasuraAuthAppleKeyID,
				Value:      unptr(config.GetAuth().GetMethod().GetOauth().GetApple().GetKeyId()),
				IsSecret:   true,
			},
			{
				Name:       "AUTH_PROVIDER_APPLE_PRIVATE_KEY",
				SecretName: secretHasuraAuthApplePrivateKey,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetApple().GetPrivateKey(),
				),
				IsSecret: true,
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetAzuread().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_AZUREAD_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetAzuread().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_AZUREAD_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetAzuread().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_AZUREAD_CLIENT_SECRET",
				SecretName: secretHasuraAuthAzureADClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetAzuread().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_AZUREAD_TENANT",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetAzuread().
						GetTenant(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetGitlab().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_GITLAB_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetGitlab().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_GITLAB_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetGitlab().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_GITLAB_CLIENT_SECRET",
				SecretName: secretHasuraAuthGitlabClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetGitlab().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_GITLAB_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetGitlab().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr( //nolint:dupl
		config.GetAuth().GetMethod().GetOauth().GetStrava().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_STRAVA_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetStrava().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_STRAVA_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetStrava().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_STRAVA_CLIENT_SECRET",
				SecretName: secretHasuraAuthStravaClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetStrava().GetClientSecret(),
				),
				IsSecret: true,
			},
			{
				Name: "AUTH_PROVIDER_STRAVA_SCOPE",
				Value: Stringify(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetStrava().
						GetScope(),
				),
				IsSecret:   false,
				SecretName: "",
			},
		}...)
	}

	if unptr(
		config.GetAuth().GetMethod().GetOauth().GetBitbucket().GetEnabled(),
	) {
		env = append(env, []EnvVar{
			{
				Name: "AUTH_PROVIDER_BITBUCKET_ENABLED",
				Value: Stringify(
					unptr(
						config.
							GetAuth().
							GetMethod().
							GetOauth().
							GetBitbucket().
							GetEnabled(),
					),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name: "AUTH_PROVIDER_BITBUCKET_CLIENT_ID",
				Value: unptr(
					config.
						GetAuth().
						GetMethod().
						GetOauth().
						GetBitbucket().
						GetClientId(),
				),
				IsSecret:   false,
				SecretName: "",
			},
			{
				Name:       "AUTH_PROVIDER_BITBUCKET_CLIENT_SECRET",
				SecretName: secretHasuraAuthBitbucketClientSecret,
				Value: unptr(
					config.GetAuth().GetMethod().GetOauth().GetBitbucket().GetClientSecret(),
				),
				IsSecret: true,
			},
		}...)
	}

	env = append(env, []EnvVar{
		{
			Name:       "AUTH_SMTP_HOST",
			Value:      smtpSettings.GetHost(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_SMTP_SECURE",
			Value: Stringify(
				smtpSettings.GetSecure(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_SMTP_PORT",
			Value: Stringify(
				smtpSettings.GetPort(),
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_SMTP_USER",
			Value:      smtpSettings.GetUser(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_SMTP_SENDER",
			Value:      smtpSettings.GetSender(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_SMTP_AUTH_METHOD",
			Value:      smtpSettings.GetMethod(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_SMTP_PASS",
			SecretName: secretHasuraAuthSMTPPassword,
			Value:      smtpSettings.GetPassword(),
			IsSecret:   true,
		},
		{
			Name:       "AUTH_RATE_LIMIT_ENABLE",
			Value:      "true",
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_RATE_LIMIT_GLOBAL_BURST",
			Value: Stringify(
				config.GetAuth().GetRateLimit().GetGlobal().GetLimit() /
					uint32(replicas), //nolint:gosec
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_RATE_LIMIT_GLOBAL_INTERVAL",
			Value:      config.GetAuth().GetRateLimit().GetGlobal().GetInterval(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_RATE_LIMIT_EMAIL_IS_GLOBAL",
			Value:      Stringify(!isCustomSMTP),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_RATE_LIMIT_EMAIL_BURST",
			Value: Stringify(
				config.GetAuth().GetRateLimit().GetEmails().GetLimit() /
					uint32(replicas), //nolint:gosec
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_RATE_LIMIT_EMAIL_INTERVAL",
			Value:      config.GetAuth().GetRateLimit().GetEmails().GetInterval(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_RATE_LIMIT_SMS_BURST",
			Value: Stringify(
				config.GetAuth().GetRateLimit().GetSms().GetLimit() /
					uint32(replicas), //nolint:gosec
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_RATE_LIMIT_SMS_INTERVAL",
			Value:      config.GetAuth().GetRateLimit().GetSms().GetInterval(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_RATE_LIMIT_BRUTE_FORCE_BURST",
			Value: Stringify(
				config.GetAuth().GetRateLimit().GetBruteForce().GetLimit() /
					uint32(replicas), //nolint:gosec
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_RATE_LIMIT_BRUTE_FORCE_INTERVAL",
			Value:      config.GetAuth().GetRateLimit().GetBruteForce().GetInterval(),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name: "AUTH_RATE_LIMIT_SIGNUPS_BURST",
			Value: Stringify(
				config.GetAuth().GetRateLimit().GetSignups().GetLimit() /
					uint32(replicas), //nolint:gosec
			),
			IsSecret:   false,
			SecretName: "",
		},
		{
			Name:       "AUTH_RATE_LIMIT_SIGNUPS_INTERVAL",
			Value:      config.GetAuth().GetRateLimit().GetSignups().GetInterval(),
			IsSecret:   false,
			SecretName: "",
		},
	}...)

	if autoScalerEnabled {
		env = append(env,
			EnvVar{
				Name:       "AUTH_RATE_LIMIT_MEMCACHE_PREFIX",
				Value:      appID + "/hasura-auth/",
				SecretName: "",
				IsSecret:   false,
			},
			EnvVar{
				Name:       "AUTH_RATE_LIMIT_MEMCACHE_SERVER",
				Value:      "memcached.ingress-nginx.svc.cluster.local:11211",
				SecretName: "",
				IsSecret:   false,
			})
	}

	for _, e := range config.GetGlobal().GetEnvironment() {
		env = append(env, EnvVar{ //nolint:exhaustruct
			Name:  e.Name,
			Value: e.Value,
		})
	}

	return env, nil
}
