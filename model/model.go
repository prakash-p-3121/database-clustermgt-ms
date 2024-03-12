package model

import "time"

type EmailIDLookUpResp struct {
	UserID    string `json:"user-id"`
	ClusterID string `json:"cluster-id"`
}

type UserIDLookUpResp struct {
	EmailID   string `json:"email-id"`
	ClusterID string `json:"cluster-id"`
}

type DatabaseShard struct {
	ID           *int64    `json:"id"` //shardID
	IPAddress    *string   `json:"ip-address"`
	ClusterID    *int64    `json:"cluster-id"`
	Port         *int      `json:"port"`
	UserName     *string   `json:"user-name"`
	Password     *string   `json:"password"`
	DatabaseName *string   `json:"string"`
	CreatedAt    time.Time `json:"created-at"`
	UpdatedAt    time.Time `json:"updated-at"`
}

type DatabaseCluster struct {
	ID        *int64    `json:"id"` // clusterID
	TableName *string   `json:"table-name"`
	ShardSize *int64    `json:"shard-size"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}
