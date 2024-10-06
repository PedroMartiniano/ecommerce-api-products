package configs

import (
	"fmt"
	"os"
)

func GetEnv(name string) string {
	env, exists := os.LookupEnv(name)
	if !exists {
		panic(fmt.Sprintf("Não foi possível encontrar a variável de ambiente %s", name))
	}

	return env
}