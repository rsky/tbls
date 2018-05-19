# posts

## Columns

| Name | Type | Default | NOT NULL | Comment |
| ---- | ---- | ------- | -------- | ------- |
| id | bigint | nextval('posts_id_seq'::regclass) | true |  |
| user_id | integer |  | true |  |
| title | varchar(255) |  | true |  |
| body | text |  | true |  |
| post_type | post_types |  | true |  |
| labels | ARRAY |  | false |  |
| created | timestamp without time zone |  | true |  |
| updated | timestamp without time zone |  | false |  |