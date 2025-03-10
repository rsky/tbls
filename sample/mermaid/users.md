# users

## Description

Users table

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `email` varchar(355) NOT NULL COMMENT 'ex. user@example.com',
  `created` timestamp NOT NULL,
  `updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=[Redacted by tbls] DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Users table'
```

</details>

## Columns

| Name | Type | Default | Nullable | Extra Definition | Children | Parents | Comment | Labels |
| ---- | ---- | ------- | -------- | ---------------- | -------- | ------- | ------- | ------ |
| id | int |  | false | auto_increment | [comment_stars](comment_stars.md) [comments](comments.md) [posts](posts.md) [user_options](user_options.md) [logs](logs.md) |  |  |  |
| username | varchar(50) |  | false |  |  |  |  |  |
| password | varchar(50) |  | false |  |  |  |  | `secure` `encrypted` |
| email | varchar(355) |  | false |  |  |  | ex. user@example.com | `secure` |
| created | timestamp |  | false |  |  |  |  |  |
| updated | timestamp |  | true |  |  |  |  |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| email | UNIQUE | UNIQUE KEY email (email) |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |
| username | UNIQUE | UNIQUE KEY username (username) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| PRIMARY | PRIMARY KEY (id) USING BTREE |
| email | UNIQUE KEY email (email) USING BTREE |
| username | UNIQUE KEY username (username) USING BTREE |

## Relations

```mermaid
erDiagram

"comment_stars" }o--|| "users" : "FOREIGN KEY (comment_user_id) REFERENCES users (id)"
"comments" }o--|| "users" : "FOREIGN KEY (user_id) REFERENCES users (id)"
"posts" }o--|| "users" : "FOREIGN KEY (user_id) REFERENCES users (id)"
"user_options" |o--|| "users" : "FOREIGN KEY (user_id) REFERENCES users (id)"
"logs" }o--|| "users" : "logs-&gt;users"

"users" {
  int id PK
  varchar_50_ username
  varchar_50_ password
  varchar_355_ email
  timestamp created
  timestamp updated
}
"comment_stars" {
  bigint id PK
  int user_id
  bigint comment_post_id FK
  int comment_user_id FK
  timestamp created
  timestamp updated
}
"comments" {
  bigint id PK
  bigint post_id FK
  int user_id FK
  text comment
  bigint post_id_desc
  datetime created
  datetime updated
}
"posts" {
  bigint id PK
  int user_id FK
  varchar_255_ title
  text body
  enum__public___private___draft__ post_type
  datetime created
  datetime updated
}
"user_options" {
  int user_id PK
  tinyint_1_ show_email
  timestamp created
  timestamp updated
}
"logs" {
  bigint id PK
  int user_id
  bigint post_id
  bigint comment_id
  bigint comment_star_id
  text payload
  datetime created
}
```

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
