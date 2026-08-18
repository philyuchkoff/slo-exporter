package config

import (
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	return logger
}

// resetViper возвращает чистое глобальное состояние viper между тестами,
// так как пакет конфигурации опирается на глобальный синглтон.
func resetViper(t *testing.T) {
	t.Helper()
	viper.Reset()
}

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.yml")
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
	return path
}

func TestLoadFromFileValidConfig(t *testing.T) {
	resetViper(t)
	path := writeTempConfig(t, `
pipeline:
  - tailer
  - sloEventProducer
logLevel: debug
webServerListenAddress: "0.0.0.0:9090"
maximumGracefulShutdownDuration: 30s
afterPipelineShutdownDelay: 5s
modules:
  tailer:
    tailedFile: /var/log/access.log
`)

	cfg := New(testLogger())
	err := cfg.LoadFromFile(path)
	require.NoError(t, err)
	assert.Equal(t, []string{"tailer", "sloEventProducer"}, cfg.Pipeline)
	assert.Equal(t, "debug", cfg.LogLevel)
	assert.Equal(t, "0.0.0.0:9090", cfg.WebServerListenAddress)
	assert.Equal(t, 30*time.Second, cfg.MaximumGracefulShutdownDuration)
	assert.Equal(t, 5*time.Second, cfg.AfterPipelineShutdownDelay)
}

func TestLoadFromFileAppliesDefaults(t *testing.T) {
	resetViper(t)
	path := writeTempConfig(t, "pipeline: []\n")

	cfg := New(testLogger())
	err := cfg.LoadFromFile(path)
	require.NoError(t, err)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Equal(t, "0.0.0.0:8080", cfg.WebServerListenAddress)
	assert.Equal(t, 20*time.Second, cfg.MaximumGracefulShutdownDuration)
	assert.Equal(t, 0*time.Second, cfg.AfterPipelineShutdownDelay)
}

func TestLoadFromFileMissingFile(t *testing.T) {
	resetViper(t)
	cfg := New(testLogger())
	err := cfg.LoadFromFile("/nonexistent/path/config.yml")
	assert.Error(t, err)
}

func TestLoadFromFileInvalidYaml(t *testing.T) {
	resetViper(t)
	path := writeTempConfig(t, "pipeline: [\n")

	cfg := New(testLogger())
	err := cfg.LoadFromFile(path)
	assert.Error(t, err)
}

func TestLoadFromFileUnknownField(t *testing.T) {
	resetViper(t)
	path := writeTempConfig(t, "thisFieldDoesNotExist: true\n")

	cfg := New(testLogger())
	err := cfg.LoadFromFile(path)
	assert.Error(t, err)
}

func TestModuleConfig(t *testing.T) {
	resetViper(t)
	path := writeTempConfig(t, `
pipeline:
  - tailer
modules:
  tailer:
    tailedFile: /var/log/access.log
    follow: true
`)

	cfg := New(testLogger())
	require.NoError(t, cfg.LoadFromFile(path))

	moduleConfig, err := cfg.ModuleConfig("tailer")
	require.NoError(t, err)
	assert.Equal(t, "/var/log/access.log", moduleConfig.GetString("tailedFile"))
	assert.Equal(t, true, moduleConfig.GetBool("follow"))
}

func TestModuleConfigMissing(t *testing.T) {
	resetViper(t)
	path := writeTempConfig(t, "pipeline:\n  - tailer\nmodules:\n  tailer: {}\n")

	cfg := New(testLogger())
	require.NoError(t, cfg.LoadFromFile(path))

	_, err := cfg.ModuleConfig("kafkaIngester")
	assert.Error(t, err)
}
