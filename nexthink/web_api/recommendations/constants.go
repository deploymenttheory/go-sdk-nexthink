// Package recommendations implements the Forge recommendations UI contracts.
package recommendations

const Endpoint = "/apigateway/recommendations-be/v1/forge/knowledge-recommendations"

const (
	StatusNew                 = "new"
	StatusInProgress          = "in-progress"
	StatusDone                = "done"
	StatusDismissed           = "dismissed"
	CategoryKnowledgeMissing  = "KNOWLEDGE_MISSING"
	CategoryKnowledgeOutdated = "KNOWLEDGE_OUTDATED"
	CategoryActionDisabled    = "ACTION_DISABLED"
	CategoryActionMissing     = "ACTION_MISSING"
	CategoryUnspecified       = "UNSPECIFIED"
)
