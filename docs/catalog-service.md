# Catalog Service

## Use Cases

`Mower`:

- `CreateMower`: create a new Mower
- `UpdateMower`: update a Mower
- `GetMower`: get a Mower
- `GetAvailableMowers`: get all Mowers

`Store`: a store provides mowers to customers

- `CreateStore`: create new store
- `UpdateStore`: update a store
- `GetStoreMowers`: get all mowers provided by a store

`Catalog`: list of all mower models available

- `GetCatalog`: list of all mower models

## Entites

`Mower`:
| Field     | Type       | Description |
| --------- | ---------- | ----------- |
| id        | uuid       |             |
| createdAt | timestampz |             |
| updatedAt | timestampz |             |
| deletedAt | timestampz |             |
| name      | string     |             |

`StoreInventory`:
| Field     | Type       | Description   |
| --------- | ---------- | ------------- |
| id        | uuid       |               |
| createdAt | timestampz |               |
| updatedAt | timestampz |               |
| deletedAt | timestampz |               |
| ns        | string     | serial number |
| model     | Mower.[id] |               |

`Store`:
| Field     | Type       | Description |
| --------- | ---------- | ----------- |
| id        | uuid       |             |
| createdAt | timestampz |             |
| updatedAt | timestampz |             |
| deletedAt | timestampz |             |
| name      | string     |             |
