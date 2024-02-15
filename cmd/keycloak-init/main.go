package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gantrycd/backend/internal/utils/random"
)

func init() {
	// フラグの読み込み
}

const (
	RealmMaster     string = "master"
	DefaultRealm    string = "gantry"
	DefaultPtorocol string = "openid-connect"
	DefaultAuthType string = "client-secret"
)

func main() {
	// TODO:エラーハンドリング
	keycloakHost := os.Getenv("KEYCLOAK_HOST")
	adminUser := os.Getenv("ADMIN_USER")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	client := new(keycloakHost, adminUser, adminPassword)

	_, err := client.getrealm(DefaultRealm)
	if err != nil {
		if err.Error()[:3] != strconv.Itoa(http.StatusNotFound) {
			panic(err)
		}
		_, err := client.createRealm(DefaultRealm)
		if err != nil {
			fmt.Println("already exists")
			os.Exit(1)
		}
	}

	secret, err := random.RandomString(20)
	if err != nil {
		panic(err)
	}

	c, err := client.createClient(DefaultRealm, secret)
	if err != nil {
		panic(err)
	}

	admin, err := client.createInitialAdmin(adminPassword)
	if err != nil {
		panic(err)
	}

	println("Realm: ", DefaultRealm)
	println("ClientID: ", c)
	println("ClientSecret: ", secret)
	println("AdminID: ", admin)
	println("AdminPassword: ", adminPassword)

	file, err := os.Create("keycloak.env")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := file.WriteString(
		fmt.Sprintf(` # Keycloak Configuration Generated by keycloak-init
KEYCLOAK_REALM=%s
KEYCLOAK_CLIENT_ID=%s
KEYCLOAK_CLIENT_SECRET=%s
KEYCLOAK_ADMIN_ID=%s
KEYCLOAK_ADMIN_PASSWORD=%s
`, DefaultRealm, c, secret, admin, adminPassword)); err != nil {
		panic(err)
	}
}

type kcClient struct {
	client *gocloak.GoCloak
	token  string
}

func new(host, user, password string) *kcClient {
	client := gocloak.NewClient(host)
	//  master の Admin でログインする
	token, err := client.LoginAdmin(context.Background(), user, password, RealmMaster)
	if err != nil {
		panic(err)
	}

	return &kcClient{
		client: client,
		token:  token.AccessToken,
	}
}

func (c *kcClient) getrealm(name string) (*gocloak.RealmRepresentation, error) {
	return c.client.GetRealm(context.Background(), c.token, name)
}

func (c *kcClient) createRealm(name string) (string, error) {
	return c.client.CreateRealm(context.Background(), c.token, gocloak.RealmRepresentation{
		Realm:   toPtr(name),
		Enabled: toPtr(true),
	})
}

func (c *kcClient) createClient(realm, secret string) (string, error) {
	keyCloakClientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	keyCloakRedirectURI := os.Getenv("KEYCLOAK_REDIRECT_URIS")
	keyCloakWebOrigins := os.Getenv("KEYCLOAK_WEB_ORIGIN")
	KeyCloakRootURL := os.Getenv("KEYCLOAK_ROOT_URL")
	return c.client.CreateClient(context.Background(), c.token, realm, gocloak.Client{
		ClientID:                     toPtr(keyCloakClientID),
		Secret:                       toPtr(secret),
		Name:                         toPtr("frontend-client"),
		Description:                  toPtr("frontend only client"),
		PublicClient:                 toPtr(false),
		AuthorizationServicesEnabled: toPtr(true),
		ServiceAccountsEnabled:       toPtr(true),
		StandardFlowEnabled:          toPtr(true),
		ClientAuthenticatorType:      toPtr(DefaultAuthType),
		Protocol:                     toPtr(DefaultPtorocol),
		RedirectURIs:                 toPtr(strings.Split(keyCloakRedirectURI, ",")),
		WebOrigins:                   toPtr(strings.Split(keyCloakWebOrigins, ",")),
		RootURL:                      toPtr(KeyCloakRootURL),
	})

}

func (c *kcClient) createInitialAdmin(password string) (string, error) {
	ctx := context.Background()
	clientID, err := c.client.CreateUser(ctx, c.token, DefaultRealm, gocloak.User{
		Username: toPtr("admin"),
		ID:       toPtr("admin"),
		Enabled:  toPtr(true),
	})
	if err != nil {
		return "", err
	}

	return clientID, c.client.SetPassword(ctx, c.token, clientID, DefaultRealm, password, false)
}

func toPtr[T any](i T) *T {
	return &i
}