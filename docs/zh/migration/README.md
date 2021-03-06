# IRIShub 从 v0.16 迁移到 v1.0.1

## 1. 导出 genesis 文件

停止 irishub v0.16 守护程序，并使用 `irishub v0.16.4（已修复导出中的错误）` 在升级高度导出主网状态并指定 `--for-zero-height`。

```bash
iris export --home [v0.16_node_home] --height 9146455 --for-zero-height
```

## 2. 迁移 genesis 文件

使用 `irishub v1.0.1` 迁移导出的 genesis.json 文件

```bash
iris migrate genesis.json --chain-id irishub-1 > genesis_v1.0.1.json
```

将升级高度+1指定为irishub v1.0.1的初始高度

```bash
# export initial_height=$[${upgrade block height} + 1]
export initial_height=9146456
jq --arg v "$initial_height" '.initial_height=$v' genesis_v1.0.1.json | sponge genesis_v1.0.1.json
```

校验 sha256sum 哈希值是否正确

```bash
sha256sum genesis_v1.0.1.json
62ef600a3cd7c6a94a47e1af66e9d6a6a0c9b224171e3e70e8ecf6497f47716b  genesis_v1.0.1.json
```

## 3. 初始化新节点

使用 `irishub v1.0.1` 初始化新的节点

```bash
iris init [moniker] --home [v1.0.1_node_home]
```

## 4. 迁移私钥文件

使用 `irishub v1.0.1` 迁移私钥文件。

### `KMS`用户

如果你使用的是KMS部署节点，请先升级`tmkms`，然后修改相关配置，详细内容请参考[kms](../tools/kms.md)

### 非`KMS`用户

如果你没有使用的是KMS部署节点，节点配置文件存在的情况下，可以使用以下两种方式中的一种迁移配置文件：

- 修改文件名称

```bash
cp [v0.16_node_home]/config/priv_validator.json [v1.0.1_node_home]/config/priv_validator_key.json
```

- 使用脚本

```bash
go run migrate/scripts/privValUpgrade.go [v0.16_node_home]/config/priv_validator.json [v1.0.1_node_home]/config/priv_validator_key.json [v1.0.1_node_home]/data/priv_validator_state.json
```

## 5. 迁移节点密钥文件

使用 `irishub v1.0.1` 迁移节点密钥文件

```bash
cp [v0.16_node_home]/config/node_key.json [v1.0.1_node_home]/config/node_key.json
```

## 6. 复制迁移过的 genensis 文件到新的节点目录

复制 genesis_v1.0.1.json 到新节点的配置文件目录

```bash
cp genesis_v1.0.1.json [v1.0.1_node_home]/config/genesis.json
```

## 7. 配置新的节点

在 `[v1.0.1_node_home]/config/app.toml` 中配置 `minimum-gas-prices`

```toml

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 0.25token1;0.0001token2).
minimum-gas-prices = "0.2uiris"

```

复制 `[v0.16_node_home]/config/config.toml` 中的 `persistent_peers` 到 `[v1.0.1_node_home]/config/config.toml` 中

```toml

# Comma separated list of nodes to keep persistent connections to
persistent_peers = ""

```

参考原节点配置文件 `[v0.16_node_home]/config/config.toml` 配置新的节点

## 8. 启动新的节点

使用 irishub v1.0.1 启动新的节点

```bash
iris unsafe-reset-all --home [v1.0.1_node_home]
iris start --home [v1.0.1_node_home]
```
