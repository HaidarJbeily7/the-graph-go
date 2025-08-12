# NFT

## Ownerships

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTOwnershipGetByAddressResponse">NFTOwnershipGetByAddressResponse</a>

Methods:

- <code title="get /nft/ownerships/evm/{address}">client.NFT.Ownerships.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTOwnershipService.GetByAddress">GetByAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTOwnershipGetByAddressParams">NFTOwnershipGetByAddressParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTOwnershipGetByAddressResponse">NFTOwnershipGetByAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Collections

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTCollectionGetByContractResponse">NFTCollectionGetByContractResponse</a>

Methods:

- <code title="get /nft/collections/evm/{contract}">client.NFT.Collections.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTCollectionService.GetByContract">GetByContract</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contract <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTCollectionGetByContractParams">NFTCollectionGetByContractParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTCollectionGetByContractResponse">NFTCollectionGetByContractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

### Evm

#### Contract

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTItemEvmContractGetByTokenResponse">NFTItemEvmContractGetByTokenResponse</a>

Methods:

- <code title="get /nft/items/evm/contract/{contract}/token_id/{token_id}">client.NFT.Items.Evm.Contract.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTItemEvmContractService.GetByToken">GetByToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTItemEvmContractGetByTokenParams">NFTItemEvmContractGetByTokenParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTItemEvmContractGetByTokenResponse">NFTItemEvmContractGetByTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Activities

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTActivityListResponse">NFTActivityListResponse</a>

Methods:

- <code title="get /nft/activities/evm">client.NFT.Activities.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTActivityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTActivityListParams">NFTActivityListParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTActivityListResponse">NFTActivityListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Holders

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTHolderGetByContractResponse">NFTHolderGetByContractResponse</a>

Methods:

- <code title="get /nft/holders/evm/{contract}">client.NFT.Holders.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTHolderService.GetByContract">GetByContract</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contract <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTHolderGetByContractParams">NFTHolderGetByContractParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTHolderGetByContractResponse">NFTHolderGetByContractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sales

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTSaleListResponse">NFTSaleListResponse</a>

Methods:

- <code title="get /nft/sales/evm">client.NFT.Sales.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTSaleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTSaleListParams">NFTSaleListParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NFTSaleListResponse">NFTSaleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Balances

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceListSvmResponse">BalanceListSvmResponse</a>
- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceGetEvmResponse">BalanceGetEvmResponse</a>

Methods:

- <code title="get /balances/svm">client.Balances.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceService.ListSvm">ListSvm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceListSvmParams">BalanceListSvmParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceListSvmResponse">BalanceListSvmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /balances/evm/{address}">client.Balances.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceService.GetEvm">GetEvm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceGetEvmParams">BalanceGetEvmParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#BalanceGetEvmResponse">BalanceGetEvmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transfers

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferListEvmResponse">TransferListEvmResponse</a>
- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferListSvmResponse">TransferListSvmResponse</a>

Methods:

- <code title="get /transfers/evm">client.Transfers.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferService.ListEvm">ListEvm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferListEvmParams">TransferListEvmParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferListEvmResponse">TransferListEvmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transfers/svm">client.Transfers.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferService.ListSvm">ListSvm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferListSvmParams">TransferListSvmParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TransferListSvmResponse">TransferListSvmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tokens

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TokenGetMetadataResponse">TokenGetMetadataResponse</a>

Methods:

- <code title="get /tokens/evm/{contract}">client.Tokens.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TokenService.GetMetadata">GetMetadata</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contract <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TokenGetMetadataParams">TokenGetMetadataParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#TokenGetMetadataResponse">TokenGetMetadataResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Holders

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HolderGetByContractResponse">HolderGetByContractResponse</a>

Methods:

- <code title="get /holders/evm/{contract}">client.Holders.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HolderService.GetByContract">GetByContract</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contract <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HolderGetByContractParams">HolderGetByContractParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HolderGetByContractResponse">HolderGetByContractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Swaps

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapGetEvmResponse">SwapGetEvmResponse</a>
- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapGetSvmResponse">SwapGetSvmResponse</a>

Methods:

- <code title="get /swaps/evm">client.Swaps.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapService.GetEvm">GetEvm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapGetEvmParams">SwapGetEvmParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapGetEvmResponse">SwapGetEvmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swaps/svm">client.Swaps.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapService.GetSvm">GetSvm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapGetSvmParams">SwapGetSvmParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#SwapGetSvmResponse">SwapGetSvmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pools

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#PoolGetResponse">PoolGetResponse</a>

Methods:

- <code title="get /pools/evm">client.Pools.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#PoolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#PoolGetParams">PoolGetParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#PoolGetResponse">PoolGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ohlc

## Pools

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPoolGetResponse">OhlcPoolGetResponse</a>

Methods:

- <code title="get /ohlc/pools/evm/{pool}">client.Ohlc.Pools.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPoolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pool <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPoolGetParams">OhlcPoolGetParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPoolGetResponse">OhlcPoolGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Prices

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPriceGetResponse">OhlcPriceGetResponse</a>

Methods:

- <code title="get /ohlc/prices/evm/{contract}">client.Ohlc.Prices.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPriceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contract <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPriceGetParams">OhlcPriceGetParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#OhlcPriceGetResponse">OhlcPriceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Historical

## Balances

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HistoricalBalanceGetResponse">HistoricalBalanceGetResponse</a>

Methods:

- <code title="get /historical/balances/evm/{address}">client.Historical.Balances.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HistoricalBalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HistoricalBalanceGetParams">HistoricalBalanceGetParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HistoricalBalanceGetResponse">HistoricalBalanceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HealthCheckParams">HealthCheckParams</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Version

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#VersionGetResponse">VersionGetResponse</a>

Methods:

- <code title="get /version">client.Version.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#VersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#VersionGetResponse">VersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Networks

Response Types:

- <a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NetworkListResponse">NetworkListResponse</a>

Methods:

- <code title="get /networks">client.Networks.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NetworkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go">thegraph</a>.<a href="https://pkg.go.dev/github.com/HaidarJbeily7/the-graph-go#NetworkListResponse">NetworkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
