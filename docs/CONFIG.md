# Configuration

This document describes the configuration schema for AIHUNTER-X.

## `log`

| Name | Type | Description | Default |
|---|---|---|---|
| `level` | `string` | The logging level. | `info` |
| `format` | `string` | The logging format. | `console` |

## `queue`

| Name | Type | Description |
|---|---|---|
| `kafka` | `object` | The Kafka configuration. |
| `redis` | `object` | The Redis configuration. |

### `queue.kafka`

| Name | Type | Description |
|---|---|---|
| `brokers` | `[]string` | A list of Kafka brokers. |

### `queue.redis`

| Name | Type | Description |
|---|---|---|
| `address` | `string` | The Redis address. |
