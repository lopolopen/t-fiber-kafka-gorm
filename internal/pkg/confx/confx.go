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

	if err := k.Load(env.Provider("APP_", ".", func(key string) string {
		return strings.ToLower(strings.TrimLeft(key, "APP_"))
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
	if strings.HasPrefix(ext, ".") {
		ext = strings.TrimLeft(ext, ".")
	}
	return strings.TrimRight(f, ext) + fmt.Sprintf("%s.%s", env, ext)
}
