package confx

import (
	"fmt"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func MustLoad(envName string, confFile string, v any) {
	if err := k.Load(file.Provider(confFile), yaml.Parser()); err != nil {
		panic(err)
	}

	envf := withEnv(confFile, ".yaml", envName)
	if _, err := os.Stat(envf); err == nil {
		if err := k.Load(file.Provider(envf), yaml.Parser()); err != nil {
			panic(err)
		}
	}

	if err := k.Load(env.ProviderWithValue("APP_", ".", func(key, value string) (string, any) {
		if value == "" {
			return "", nil
		}
		key = strings.ToLower(strings.TrimPrefix(key, "APP_"))
		key = smartKey(key)
		if strings.Contains(value, ",") {
			parts := strings.Split(value, ",")
			var sliceValues []string
			for _, p := range parts {
				sliceValues = append(sliceValues, strings.TrimSpace(p))
			}
			return key, sliceValues
		}
		return key, value
	}), nil); err != nil {
		panic(err)
	}

	if err := k.UnmarshalWithConf("", v, koanf.UnmarshalConf{
		Tag: "yaml",
	}); err != nil {
		panic(err)
	}
}

func withEnv(f string, ext string, env string) string {
	ext = strings.TrimPrefix(ext, ".")
	return strings.TrimSuffix(f, ext) + fmt.Sprintf("%s.%s", env, ext)
}

func smartKey(input string) string {
	var builder strings.Builder
	builder.Grow(len(input))

	runes := []rune(input)
	n := len(runes)

	for i := 0; i < n; i++ {
		if runes[i] == '_' {
			if i+1 < n && runes[i+1] == '_' {
				builder.WriteRune('_')
				i++
			} else {
				builder.WriteRune('.')
			}
		} else {
			builder.WriteRune(runes[i])
		}
	}
	return builder.String()
}
