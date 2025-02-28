package togo

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/allisson/go-assert"
)

var (
	buf            bytes.Buffer
	configFilename = "c.json"
)

func TestConfig(t *testing.T) {
	t.Run("trying to parse a nonexistent config file", func(t *testing.T) {
		buf.Reset()
		Logger.SetOutput(&buf)

		LoadJSONFile("invalid.json")
		assert.Equal(t, "Unable to read file \"invalid.json\": open invalid.json: no such file or directory\n", buf.String())
	})

	t.Run("trying to parse a invalid config file", func(t *testing.T) {
		buf.Reset()
		Logger.SetOutput(&buf)

		err := ioutil.WriteFile(configFilename, nil, 0644)
		assert.Nil(t, err)

		LoadJSONFile(configFilename)
		assert.Equal(t, "Unable to parse JSON from file \"c.json\": unexpected end of JSON input\n", buf.String())

		err = os.Remove(configFilename)
		assert.Nil(t, err)
	})

	t.Run("parsing empty file to check default values", func(t *testing.T) {
		err := ioutil.WriteFile(configFilename, []byte("{}\n"), 0644)
		assert.Nil(t, err)

		config := LoadJSONFile(configFilename)
		assert.Equal(t, "0.0.0.0", config.HTTPAddr)
		assert.Equal(t, 3000, config.HTTPPort)
		assert.Equal(t, 30, config.IdleTimeout)
		assert.Equal(t, 5, config.ReadTimeout)
		assert.Equal(t, 10, config.WriteTimeout)

		err = os.Remove(configFilename)
		assert.Nil(t, err)
	})
}
