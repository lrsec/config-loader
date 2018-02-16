# config-loader

> load configs in different env

# Usage

1. define config files for different env in conf/
```
src
 |--your prject
        |--conf
             |--test
                 config.toml
             |--product
                 config.toml
             |--stage
                 config.toml
```


2.  define environment in servers:
```
// for test env
RUN_MODE = "test"

// for product env
RUN_MODE = "product"

// for stage env
RUN_MODE = "stage"
```

3. load config in code

```
cfg, err := config.LoadToml(c, "config.toml")
```

product/config.toml will be load on product env server

test/config.toml for test env server

stage/config.toml for stage env server
