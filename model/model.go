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
	ID        *int64  `json:"id"` //shardID
	IPAddress *string `json:"ip-address"`
}

type DatabaseCluster struct {
	ID        *int64  `json:"id"` // clusterID
	TableName *string `json:"table-name"`
	ShardSize *int64  `json:"shard-size"`
}
