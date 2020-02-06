package index

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestBackwardIncompatibleFeatures_EnableURLEncodingOfIndexName(t *testing.T) {
	t.Parallel()

	client := search.NewClientWithConfig(search.Configuration{
		AppID:  os.Getenv("ALGOLIA_APPLICATION_ID_1"),
		APIKey: os.Getenv("ALGOLIA_ADMIN_KEY_1"),
		CompatibilityBreakingFeatures: search.CompatibilityBreakingFeatures{
			EnableURLEncodingOfIndexName: true,
		},
	})

	indexNames := []string{
		cts.GenerateIndexName(t) + "_index 1",
		cts.GenerateIndexName(t) + "_index+1",
		cts.GenerateIndexName(t) + "_index-1",
		cts.GenerateIndexName(t) + "_index.1",
		cts.GenerateIndexName(t) + "_index%1",
		cts.GenerateIndexName(t) + "_index^1",
	}

	g := wait.NewGroup()

	for _, indexName := range indexNames {
		res, err := client.InitIndex(indexName).SaveObject(map[string]string{
			"objectID": indexName,
		})
		require.NoError(t, err, "should save object in index %q", indexName)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	res, err := client.ListIndices()
	require.NoError(t, err)

	var listedIndexNames []string
	for _, index := range res.Items {
		listedIndexNames = append(listedIndexNames, index.Name)
	}

	for _, indexName := range indexNames {
		require.Contains(t, listedIndexNames, indexName)
	}
}
