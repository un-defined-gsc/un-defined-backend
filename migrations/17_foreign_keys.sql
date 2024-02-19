-- +goose Up

ALTER TABLE t_action_log  ADD CONSTRAINT fk_action_logs_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE NO ACTION;
ALTER TABLE t_login_logs ADD CONSTRAINT fk_login_logs_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE NO ACTION;
ALTER TABLE t_banned ADD CONSTRAINT fk_banned_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE NO ACTION;
ALTER TABLE t_mfa_settings ADD CONSTRAINT fk_mfa_settings_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;


ALTER TABLE t_completed_maps ADD CONSTRAINT fk_completed_map_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;
ALTER TABLE t_completed_maps ADD CONSTRAINT fk_completed_map_paths FOREIGN KEY (path_id) REFERENCES t_path_ways(id) ON DELETE CASCADE;

ALTER TABLE t_path_ways ADD CONSTRAINT fk_path_ways_roadmaps FOREIGN KEY (roadmap_id) REFERENCES t_roadmaps (id) ON DELETE CASCADE;
ALTER TABLE t_path_ways ADD CONSTRAINT fk_path_ways_path_ways FOREIGN KEY (parent_path) REFERENCES t_path_ways (id) ON DELETE CASCADE;

ALTER TABLE t_tags_posts ADD CONSTRAINT fk_tags_posts_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;
ALTER TABLE t_tags_posts ADD CONSTRAINT fk_tags_posts_posts FOREIGN KEY (post_id) REFERENCES t_posts (id) ON DELETE CASCADE;


ALTER TABLE t_comments ADD CONSTRAINT fk_comments_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;
ALTER TABLE t_comments ADD CONSTRAINT fk_comments_posts FOREIGN KEY (post_id) REFERENCES t_posts (id) ON DELETE CASCADE;


ALTER TABLE t_likes ADD CONSTRAINT fk_likes_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;
ALTER TABLE t_likes ADD CONSTRAINT fk_likes_posts FOREIGN KEY (post_id) REFERENCES t_posts (id) ON DELETE CASCADE;


ALTER TABLE t_posts ADD CONSTRAINT fk_posts_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;
ALTER TABLE t_posts ADD CONSTRAINT fk_posts_categories FOREIGN KEY (category_id) REFERENCES t_categories (id) ON DELETE CASCADE;

ALTER TABLE t_images ADD CONSTRAINT fk_images_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;
ALTER TABLE t_images ADD CONSTRAINT fk_images_posts FOREIGN KEY (post_id) REFERENCES t_posts (id) ON DELETE CASCADE;
-- +goose Down

ALTER TABLE t_action_log DROP CONSTRAINT fk_action_logs_users;
ALTER TABLE t_login_logs DROP CONSTRAINT fk_login_logs_users;
ALTER TABLE t_banned DROP CONSTRAINT fk_banned_users;
ALTER TABLE t_mfa_settings DROP CONSTRAINT fk_mfa_settings_users;

ALTER TABLE t_completed_maps DROP CONSTRAINT fk_completed_map_users;
ALTER TABLE t_completed_maps DROP CONSTRAINT fk_completed_map_paths;

ALTER TABLE t_path_ways DROP CONSTRAINT fk_path_ways_roadmaps;
ALTER TABLE t_path_ways DROP CONSTRAINT fk_path_ways_path_ways;

ALTER TABLE t_tags_posts DROP CONSTRAINT fk_tags_posts_users;
ALTER TABLE t_tags_posts DROP CONSTRAINT fk_tags_posts_posts;

ALTER TABLE t_comments DROP CONSTRAINT fk_comments_users;
ALTER TABLE t_comments DROP CONSTRAINT fk_comments_posts;

ALTER TABLE t_likes DROP CONSTRAINT fk_likes_users;
ALTER TABLE t_likes DROP CONSTRAINT fk_likes_posts;

ALTER TABLE t_posts DROP CONSTRAINT fk_posts_users;
ALTER TABLE t_posts DROP CONSTRAINT fk_posts_categories;

ALTER TABLE t_images DROP CONSTRAINT fk_images_users;
ALTER TABLE t_images DROP CONSTRAINT fk_images_posts;


