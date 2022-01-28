package common

type Blockchain = string

const (
	ETHEREUM Blockchain = "ethereum"
	POLYGON  Blockchain = "polygon"
	ARBITRUM Blockchain = "arbitrum"
	OPTIMISM Blockchain = "optimism"
)

type Interface = string

const (
	ERC721              Interface = "ERC721"
	ERC1155             Interface = "ERC1155"
	ERC20               Interface = "ERC20"
	ENS_REGISTRAR       Interface = "ENS_REGISTRAR"
	SUSHISWAP_EXCHANGE  Interface = "SUSHISWAP_EXCHANGE"
	UNISWAP_V2_EXCHANGE Interface = "UNISWAP_V2_EXCHANGE"
	UNISWAP_V3_EXCHANGE Interface = "UNISWAP_V3_EXCHANGE"
)

type Address = string

const (
	ZERO_ADDRESS Address = "0x0000000000000000000000000000000000000000"
)

// testnet erc20 addresses
const (
	UNI_GOERLI  Address = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	LINK_GOERLI Address = "0x14b7ba66139c234b1be9a157d4f8b985b8a7f762"
	HEX_GOERLI  Address = "0x08249c12c66c76ea384cf851bb3e274a2bb1874a"
	DAI_GOERLI  Address = "0x11fe4b6ae13d2a6055c8d9cf65c55bac32b5d844"
	BUSD_GOERLI Address = "0xb809b9b2dc5e93cb863176ea2d565425b03c0540"
	USDC_GOERLI Address = "0x07865c6e87b9f70255377e024ace6630c1eaa37f"
	USDT_GOERLI Address = "0xe583769738b6dd4e7caf8451050d1948be717679"
	WETH_GOERLI Address = "0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6"
)

// mainnet erc20 addresses
const (
	USDT_MAINNET   Address = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	BNB_MAINNET    Address = "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"
	USDC_MAINNET   Address = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	HEX_MAINNET    Address = "0x2b591e99afe9f32eaa6214f7b7629768c40eeb39"
	MATIC_MAINNET  Address = "0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0"
	SHIB_MAINNET   Address = "0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"
	BUSD_MAINNET   Address = "0x4fabb145d64652a948d72533023f6e7a623c7c53"
	LINK_MAINNET   Address = "0x514910771af9ca656af840dff83e8264ecf986ca"
	CRO_MAINNET    Address = "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b"
	WBTC_MAINNET   Address = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	UST_MAINNET    Address = "0xa47c8bf37f92abed4a126bda807a7b7498661acd"
	DAI_MAINNET    Address = "0x6b175474e89094c44da98b954eedeac495271d0f"
	FTM_MAINNET    Address = "0x4e15361fd6b4bb609fa63c81a2be19d873717870"
	UNI_MAINNET    Address = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	OKB_MAINNET    Address = "0x75231f58b43240c9718dd58b4967c5114342a86c"
	TRX_MAINNET    Address = "0xe1be5d3f34e89de342ee97e6e90d405884da6c67"
	STETH_MAINNET  Address = "0xae7ab96520de3a18e5e111b5eaab095312d7fe84"
	WSTETH_MAINNET Address = "0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"
	THETA_MAINNET  Address = "0x3883f5e181fccaf8410fa61e12b59bad963fb645"
	VEN_MAINNET    Address = "0xd850942ef8811f2a866692a623011bde52a462c1"
)
