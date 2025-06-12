package logging

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

var provider *sdklog.LoggerProvider

// InitLog configura o exporter OTLP, cria o LoggerProvider
// e retorna um *slog.Logger que já envia registros para o collector.
// `serviceName` será usado como nome de instrumentação.
//
// Exemplos de configuração de ambiente:
//
//	OTEL_SERVICE_NAME="meu-servico"
//	OTEL_EXPORTER_OTLP_ENDPOINT="meu-collector:4317"
//	OTEL_EXPORTER_OTLP_INSECURE=true
func InitLog(ctx context.Context, serviceName string, tags map[string]string) (*slog.Logger, error) {
	exp, err := otlploghttp.New(
		ctx,
		otlploghttp.WithEndpoint("localhost:4318"),
		otlploghttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	provider = sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
	)
	global.SetLoggerProvider(provider)

	// Monta atributos a partir do mapa de tags
	attrs := make([]any, 0, len(tags)+1)
	attrs = append(attrs, slog.String("service.name", serviceName))
	for k, v := range tags {
		attrs = append(attrs, slog.String(k, v))
	}

	logger := otelslog.NewLogger(serviceName).With(attrs...)
	return logger, nil
}

// Shutdown garante envio de todos os logs pendentes antes de fechar a aplicação.
func Shutdown(ctx context.Context) error {
	if provider == nil {
		return nil
	}
	return provider.Shutdown(ctx)
}
