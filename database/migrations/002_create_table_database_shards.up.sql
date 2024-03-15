CREATE TABLE database_shards (
   id BIGINT AUTO_INCREMENT PRIMARY KEY,
   ip_address TEXT NOT NULL,
   cluster_id BIGINT DEFAULT NULL,
   port INT UNSIGNED NOT NULL,
   user_name TEXT NOT NULL,
   password TEXT NOT NULL,
   database_name TEXT NOT NULL,
   start_range TEXT NOT NULL,
   end_range TEXT DEFAULT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   INDEX database_shards_cluster_id_idx (cluster_id)
);
