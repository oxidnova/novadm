# Novadm

Assistive Admin system, to help locate issues and provide a visual user UI.

### Run it local

**run service**

```
pnpm i
pnpm build:antd

go run main.go serve -c contrib/serve.yaml all --verbose

# open http://127.0.0.1:5320/
```

**you want to modify UI, please run the following command:**

```
pnpm i
pnpm dev:antd

# open http://127.0.0.1:5666/
```

### Configuration

#### Config Endpoint

```
serve:
  api:
    base_url: http://127.0.0.1:5320/
```

#### Config Users & Allowed Menus

```
  credentials:
    - username: super@example.com
      realname: Super
      # bcrypt hash of the string "password"
      password: "$2a$10$2b2cU8CPhOTaGrs1HRQuAueS7JTT5ZHsHSzYiFPm1leZck7Mc8T4W"
      menus:
        - "*"
```

- Super Can view all menus

Please view [more details](./backend/contrib/serve.yaml)
