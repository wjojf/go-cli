package key

import (
	"log"
	"testing"
)

func TestGetKeyLoaders(t *testing.T) {
	t.Run("Test Loading", func(t *testing.T) {
		keys := GetPrivateKeyLoaders()

		if len(keys) == 0 {
			t.Error("No keys found")
		}

		for _, loader := range keys {
			key, err := loader.Load()
			if err != nil {
				t.Error(err)
			}

			log.Println(key + "\n\n\n")
		}

	})
}
