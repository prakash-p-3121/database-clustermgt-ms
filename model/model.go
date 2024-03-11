package model

type EmailIDLookUpResp struct {
	UserID    string `json:"user-id"`
	ClusterID string `json:"cluster-id"`
}

type UserIDLookUpResp struct {
	EmailID   string `json:"email-id"`
	ClusterID string `json:"cluster-id"`
}

type DatabaseShard struct {
	ID        *uint64 `json:"id"` //shardID
	IPAddress *string `json:"string"`
}

type DatabaseCluster struct {
	ID        *uint64 `json:"id"` // clusterID
	TableName *string `json:"table-name"`
	ShardSize *uint64 `json:"shard-size"`
}
