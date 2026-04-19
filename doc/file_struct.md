```tree
[user@vbox energy_rms]$ tree .
.
├── doc
│   ├── db_struct.md
│   ├── file_struct.md
│   └── instruction.md
├── docker-compose.yml
├── postgres
│   ├── config
│   │   ├── pg_hba.conf
│   │   └── postgresql.conf
│   └── scripts
│       ├── 0001_create_table_user_up.sql
│       └── initdb
│           ├── 1_user_ctx_create.sql
│           └── 2_user_ctx_data.sql
├── README.md
├── rest-api
│   ├── cmd
│   │   └── app
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── handlers
│       │   └── user_handlers.go
│       ├── models
│       │   └── user.go
│       └── repository
│           └── user_repository.go
└── web-service
    ├── apache-config
    │   └── 000-default.conf
    ├── Dockerfile
    ├── php-config
    │   └── php.ini
    └── www
        └── index.php

16 directories, 21 files
```

<table width=100%>
    <tr>
        <td width=30%>
            [Страница проекта](../)
        </td>
        <td>
            Наверх
        </td>
        <td width=30%>
            &nbsp;
        </td>
    </tr>
</table>