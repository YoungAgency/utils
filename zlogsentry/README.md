# zerolog-sentry

### Example
```go
import (
	"errors"
	"io"
	stdlog "log"
	"os"

	"github.com/YoungAgency/utils/zlogsentry"
	"github.com/rs/zerolog"
)

func main() {
	var sentryClient *sentry.Client
	{
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "cannot get hostname"
		}

		sentryClient, err = sentry.NewClient(sentry.ClientOptions{
			Dsn:        "YOUR_DSN",
			ServerName: hostname,
			Release:    Version,
		})
		if err != nil {
			stdlog.Printf("cannot create sentry client (sentry will be disabled): %v", err)
		}

		defer sentryClient.Flush(5 * time.Second)
	}

	// logging
	var log zerolog.Logger
	{
		var writers = []io.Writer{
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
		}

		if sentryClient != nil {
			w, err := zlogsentry.New(sentryClient)
			if err != nil {
				stdlog.Printf("sentry initialization failed: %v", err)
			}

			defer w.Close()

			writers = append(writers, w)
		}

		log = zerolog.
            New(zerolog.MultiLevelWriter(writers...)).
            With().
            Timestamp().
            Logger()
	}

	log.Error().
        Err(errors.New("error")).
        Str("component", "example").
        Msg("can't load this")
}
```

