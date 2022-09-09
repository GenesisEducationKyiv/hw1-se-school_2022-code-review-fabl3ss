package gmail_api

import (
	"context"
	"encoding/json"
	"fmt"
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/types/filemodes"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func getClient(authConfig *oauth2.Config) *http.Client {
	tokFile := os.Getenv(config.EnvGmailTokenPath)
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		saveToken(tokFile, getTokenFromWeb(authConfig))
	}
	return authConfig.Client(context.Background(), tok)
}

func getTokenFromWeb(authConfig *oauth2.Config) *oauth2.Token {
	authURL := authConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := authConfig.Exchange(context.TODO(), authCode, oauth2.AccessTypeOffline)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = f.Close()
	}()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	log.Printf("Saving credential file to: %s\n", path)
	tokenFileMode := os.ModeDir | filemodes.OS_USER_RW
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, tokenFileMode)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalf("Unable to close oauth token file")
		}
	}()
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Unable to encode oauth token")
	}
}

func UpdateService(credentialsPath string, tokenPath string) (*gmail.Service, error) {
	config, err := getClientFromFile(credentialsPath)
	if err != nil {
		return nil, err
	}

	tok, err := tokenFromFile(tokenPath)
	if err != nil {
		return nil, err
	}
	tokenSource := config.TokenSource(context.TODO(), tok)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	if newToken.AccessToken != tok.AccessToken {
		saveToken(tokenPath, newToken)
	}
	return GetGmailService()
}

func getClientFromFile(path string) (*oauth2.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read client secret file")
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse client secret file to config")
	}
	return config, nil
}

func GetGmailService() (*gmail.Service, error) {
	gmailConfig, err := getClientFromFile(
		os.Getenv(config.EnvGmailCredentialsPath),
	)
	if err != nil {
		return nil, err
	}

	srv, err := gmail.NewService(
		context.Background(),
		option.WithHTTPClient(getClient(gmailConfig)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve Gmail client")
	}
	return srv, nil
}
