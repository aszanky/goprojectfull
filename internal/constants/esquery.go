package constants

const (
	UpdateLevelLeadsGen string = `
		{
			"script": {
				"source": "ctx._source.name = params.name",
				"params": {"name": "%s"},
				"lang": "painless"
			},
			"query": {
				"match": {
					"wp_code": "%s"
				}
			}
		}
	`

	UpsertLeadsGen string = `
		{
			"script": {
				"source": "ctx._source.level = params.level",
				"params": {"level": 0},
				"lang": "painless"
			},
			"upsert": %s
		}
	`
)
