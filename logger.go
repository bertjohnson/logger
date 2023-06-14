// Package logger provides methods for logging events.
package logger

import (
	"context"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	contexttype "github.com/bertjohnson/logger/types/context"
	"github.com/bertjohnson/logger/types/env"
	"github.com/bertjohnson/startup"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// PackageType reflects the object type.
	PackageType = "logprovider"
)

var (
	// Local hostname.
	hostname string

	// Logger.
	logger *zap.Logger

	// Logger for third party components
	externalLogger *zap.Logger
)

type (
	// LogProvider structure to expose logger to third party components.
	LogProvider struct{}
)

// Output method used by mgo debug logger.
func (l LogProvider) Output(depth int, data string) error {
	externalLogger.Debug(data)
	return nil
}

// init initializes logging.
func init() {
	go func() {
		// Wait for configuration to initialize.
		ctx := context.Background()
		startup.Wait(ctx, startup.PackageType)

		// Remember local hostname.
		var err error
		hostname, err = os.Hostname()
		if err != nil {
			log.Fatalln("Unable to resolve hostname for worker.")
		}

		// Create logger.
		var loggerConfig zap.Config
		var externalLoggerConfig zap.Config
		if os.Getenv(env.Debug) != "" {
			loggerConfig = zap.NewDevelopmentConfig()
			externalLoggerConfig = zap.NewDevelopmentConfig()

			// Unless we're on Windows, color output.
			if runtime.GOOS != "windows" {
				loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			}
		} else {
			loggerConfig = zap.NewProductionConfig()
			externalLoggerConfig = zap.NewProductionConfig()
		}
		switch strings.ToLower(os.Getenv(env.LoggingLevel)) {
		case "debug", "verbose", "1":
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info", "2":
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warn", "warning", "3":
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error", "4":
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		case "critical", "fatal", "5":
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
		}
		externalLoggerConfig.Level = loggerConfig.Level
		logger, err = loggerConfig.Build()
		if err != nil {
			log.Fatalln("Error building log configuration: " + err.Error())
		}
		externalLogger, err = externalLoggerConfig.Build()
		if err != nil {
			log.Fatalln("Error building external log configuration: " + err.Error())
		}
		logger = logger.WithOptions(zap.AddCallerSkip(1))
		if err != nil {
			log.Fatalln("Error when creating logger:", err)
		}
		externalLogger = externalLogger.WithOptions(zap.AddCallerSkip(1))
		if err != nil {
			log.Fatalln("Error when creating external logger:", err)
		}
		defer func() {
			logger.Sync()         // nolint
			externalLogger.Sync() // nolint
		}()

		Verbose(ctx, "Logging service is ready.")

		startup.Publish(ctx, PackageType)
	}()
}

// Error logs an error.
func Error(ctx context.Context, data string) {
	logger.Error(data, getFields(ctx)...)
}

// Fatal logs a fatal message and then exits.
func Fatal(ctx context.Context, data string) {
	logger.Fatal(data, getFields(ctx)...)
	err := logger.Sync()
	if err != nil {
		log.Println("Error flushing logs for fatal error: " + err.Error())
	}
	os.Exit(1)
}

// Flush flushes logs.
func Flush(ctx context.Context) error {
	return logger.Sync()
}

// Info logs an informational message.
func Info(ctx context.Context, data string) {
	logger.Info(data, getFields(ctx)...)
}

// Verbose logs a verbose message.
func Verbose(ctx context.Context, data string) {
	logger.Debug(data, getFields(ctx)...)
}

// Warn logs a warning.
func Warn(ctx context.Context, data string) {
	logger.Warn(data, getFields(ctx)...)
}

// getFields gets fields associated with a logging context.
func getFields(ctx context.Context) []zapcore.Field {
	// Initialize fields.
	fields := []zapcore.Field{
		zap.String("server", hostname),
	}

	if ctx != nil {
		now := time.Now().UTC()

		// Account Email.
		if accountEmail, ok := ctx.Value(contexttype.AccountEmail).(string); ok {
			if accountEmail != "" {
				fields = append(fields, zap.String(contexttype.AccountEmail, accountEmail))
			}
		}

		// Account ID.
		if accountID, ok := ctx.Value(contexttype.AccountID).(string); ok {
			if accountID != "" {
				fields = append(fields, zap.String(contexttype.AccountID, accountID))
			}
		}

		// BLOB Provider ID.
		if blobID, ok := ctx.Value(contexttype.BLOBID).(string); ok {
			if blobID != "" {
				fields = append(fields, zap.String(contexttype.BLOBID, blobID))
			}
		}

		// BLOB Provider ID.
		if blobProviderID, ok := ctx.Value(contexttype.BLOBProviderID).(string); ok {
			if blobProviderID != "" {
				fields = append(fields, zap.String(contexttype.BLOBProviderID, blobProviderID))
			}
		}

		// Client Date.
		if clientDate, ok := ctx.Value(contexttype.ClientDate).(string); ok {
			if clientDate != "" {
				fields = append(fields, zap.String(contexttype.ClientDate, clientDate))
			}
		}

		// Client ID.
		if clientID, ok := ctx.Value(contexttype.ClientID).(string); ok {
			if clientID != "" {
				fields = append(fields, zap.String(contexttype.ClientID, clientID))
			}
		}

		// Collection.
		if collection, ok := ctx.Value(contexttype.Collection).(string); ok {
			if collection != "" {
				fields = append(fields, zap.String(contexttype.Collection, collection))
			}
		}

		// Database.
		if database, ok := ctx.Value(contexttype.Database).(string); ok {
			if database != "" {
				fields = append(fields, zap.String(contexttype.Database, database))
			}
		}

		// Data Provider ID.
		if dataProviderID, ok := ctx.Value(contexttype.DataProviderID).(string); ok {
			if dataProviderID != "" {
				fields = append(fields, zap.String(contexttype.DataProviderID, dataProviderID))
			}
		}

		// Exchange.
		if exchange, ok := ctx.Value(contexttype.Exchange).(string); ok {
			if exchange != "" {
				fields = append(fields, zap.String(contexttype.Exchange, exchange))
			}
		}

		// File.
		if file, ok := ctx.Value(contexttype.File).(string); ok {
			if file != "" {
				fields = append(fields, zap.String(contexttype.File, file))
			}
		}

		// Hostname.
		if hostname, ok := ctx.Value(contexttype.Hostname).(string); ok {
			if hostname != "" {
				fields = append(fields, zap.String(contexttype.Hostname, hostname))
			}
		}

		// Index ID.
		if indexID, ok := ctx.Value(contexttype.IndexID).(string); ok {
			if indexID != "" {
				fields = append(fields, zap.String(contexttype.IndexID, indexID))
			}
		}

		// Line.
		if line, ok := ctx.Value(contexttype.Line).(uint); ok {
			fields = append(fields, zap.Uint(contexttype.Line, line))
		}

		// Node ID.
		if nodeID, ok := ctx.Value(contexttype.NodeID).(string); ok {
			if nodeID != "" {
				fields = append(fields, zap.String(contexttype.NodeID, nodeID))
			}
		}

		// Node Type.
		if nodeType, ok := ctx.Value(contexttype.NodeType).(string); ok {
			if nodeType != "" {
				fields = append(fields, zap.String(contexttype.NodeType, nodeType))
			}
		}

		// Node URI.
		if nodeURI, ok := ctx.Value(contexttype.NodeURI).(string); ok {
			if nodeURI != "" {
				fields = append(fields, zap.String(contexttype.NodeURI, nodeURI))
			}
		}

		// Object IDs.
		if objectIDs, ok := ctx.Value(contexttype.ObjectIDs).(string); ok {
			objectIDsParts := strings.Split(objectIDs, "&")
			for _, objectID := range objectIDsParts {
				objectIDParts := strings.SplitN(objectID, "=", 2)
				if len(objectIDParts) == 2 {
					fields = append(fields, zap.String(objectIDParts[0], objectIDParts[1]))
				}
			}
		}

		// Referer.
		if referer, ok := ctx.Value(contexttype.Referer).(string); ok {
			if referer != "" {
				fields = append(fields, zap.String(contexttype.Referer, referer))
			}
		}

		// Request Duration.
		if requestStart, ok := ctx.Value(contexttype.RequestStart).(time.Time); ok {
			if !requestStart.IsZero() {
				fields = append(fields, zap.Duration("requestDuration", now.Sub(requestStart)))
			}
		}

		// Request ID.
		if requestID, ok := ctx.Value(contexttype.RequestID).(string); ok {
			if requestID != "" {
				fields = append(fields, zap.String(contexttype.RequestID, requestID))
			}
		}

		// Request IP.
		if requestIP, ok := ctx.Value(contexttype.RequestIP).(string); ok {
			if requestIP != "" {
				fields = append(fields, zap.String(contexttype.RequestIP, requestIP))
			}
		}

		// Request Method.
		if requestMethod, ok := ctx.Value(contexttype.RequestMethod).(string); ok {
			if requestMethod != "" {
				fields = append(fields, zap.String(contexttype.RequestMethod, requestMethod))
			}
		}

		// Request URL.
		if requestURL, ok := ctx.Value(contexttype.RequestURL).(string); ok {
			if requestURL != "" {
				fields = append(fields, zap.String(contexttype.RequestURL, requestURL))
			}
		}

		// Response Code.
		if responseCode, ok := ctx.Value(contexttype.ResponseCode).(uint); ok {
			fields = append(fields, zap.Uint(contexttype.ResponseCode, responseCode))
		}

		// Response Length.
		if responseLength, ok := ctx.Value(contexttype.ResponseLength).(uint); ok {
			fields = append(fields, zap.Uint(contexttype.ResponseLength, responseLength))
		}

		// Search Engine ID.
		if searchEngineID, ok := ctx.Value(contexttype.SearchEngineID).(string); ok {
			if searchEngineID != "" {
				fields = append(fields, zap.String(contexttype.SearchEngineID, searchEngineID))
			}
		}

		// Secret Provider ID.
		if secretProviderID, ok := ctx.Value(contexttype.SecretProviderID).(string); ok {
			if secretProviderID != "" {
				fields = append(fields, zap.String(contexttype.SecretProviderID, secretProviderID))
			}
		}

		// Worker ID.
		if workerID, ok := ctx.Value(contexttype.WorkerID).(string); ok {
			if workerID != "" {
				fields = append(fields, zap.String(contexttype.WorkerID, workerID))
			}
		}

		// Gin-specific logging.
		if ginContext, ok := ctx.(*gin.Context); ok {
			if ginContext.Value("bulk") != true {
				if ginContext != nil && ginContext.Writer != nil {
					header := ginContext.Writer.Header()
					if header != nil {
						contentLengthString := header.Get("Content-Length")
						contentLength, err := strconv.Atoi(contentLengthString)
						if err == nil {
							fields = append(fields, zap.Int("contentLength", contentLength))
						}

						// Segment Duration.
						if segmentStart, ok := ctx.Value(contexttype.SegmentStart).(time.Time); ok {
							fields = append(fields, zap.Duration("segmentDuration", time.Now().UTC().Sub(segmentStart)))
						}
						ginContext.Set("segmentStart", now)
					}
				}
			}
		}
	}
	return fields
}

// Wait waits until the logger is available.
func Wait(ctx context.Context) {
	startup.Wait(ctx, PackageType)
}
