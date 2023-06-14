package logger

import (
	"context"
	"log"
	"os"
	"testing"

	contexttype "github.com/bertjohnson/logger/types/context"
	"github.com/bertjohnson/startup"
)

// TestMain runs tests.
func TestMain(m *testing.M) {
	// Declare that the configuration is ready.
	err := startup.Ready()
	if err != nil {
		log.Fatalln("Error loading configuration values: " + err.Error())
	}

	// Wait for the initializer to finish.
	Wait(context.Background())

	// Run tests.
	retCode := m.Run()

	// Exit.
	os.Exit(retCode)
}

// TestLog tests Verbose(), Info(), Warn(), Error(), and Fatal().
func TestLog(t *testing.T) {
	requestIDContext := context.WithValue(context.Background(), contexttype.AccountID, "5d67d1af-204a-44e7-be70-e4e35059e1fb")                                                                                                                                                  // nolint
	accountIDContext := context.WithValue(requestIDContext, contexttype.AccountID, "5d67d1af-204a-44e7-be70-e4e35059e1fb")                                                                                                                                                      // nolint
	accountEmailContext := context.WithValue(accountIDContext, contexttype.AccountEmail, "user@example.com")                                                                                                                                                                    // nolint
	accountJWTContext := context.WithValue(accountEmailContext, contexttype.AccountJWT, "sampleJWT")                                                                                                                                                                            // nolint
	accountXSRFStateContext := context.WithValue(accountJWTContext, contexttype.AccountXSRFState, "sampleXSRF")                                                                                                                                                                 // nolint
	hostnameContext := context.WithValue(accountXSRFStateContext, contexttype.Hostname, "example.com")                                                                                                                                                                          // nolint
	objectIDsContext := context.WithValue(hostnameContext, contexttype.ObjectIDs, "tenantid=21b13f98-feee-4832-9ca2-2c5f5e18cf8c&workspaceid=de0cc560-1830-4d09-a8e8-c5fffe9e518d&databaseid=bc001ef9-3a5c-4bba-bb57-0c9d6fd19479&itemid=7a684a4a-12d1-403e-abb2-d8db11df791a") // nolint
	requestIPContext := context.WithValue(objectIDsContext, contexttype.RequestIP, "127.0.0.1")                                                                                                                                                                                 // nolint
	requestURLContext := context.WithValue(requestIPContext, contexttype.RequestURL, "https://bertjohnson.com")                                                                                                                                                                 // nolint

	Verbose(requestURLContext, "Successfully logged this verbose message with context data.")
	Info(requestURLContext, "Successfully logged this info message with context data.")
	Warn(requestURLContext, "Successfully logged this warning message with context data.")
	Error(requestURLContext, "Successfully logged this error message with context data.")
}
