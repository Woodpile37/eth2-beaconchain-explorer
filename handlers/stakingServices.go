package handlers

import (
	"eth2-exporter/templates"
	"eth2-exporter/types"
	"eth2-exporter/utils"
	"net/http"
)

func StakingServices(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "stakingServices.html")
	var stakingServicesTemplate = templates.GetTemplate(templateFiles...)

	var err error

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/stakingServices", "Ethereum Staking Services Overview", templateFiles)

	pageData := &types.StakeWithUsPageData{}
	pageData.RecaptchaKey = utils.Config.Frontend.RecaptchaSiteKey
	pageData.StakingServices = getStakingServices()
	pageData.FlashMessage, err = utils.GetFlash(w, r, "stake_flash")
	if err != nil {
		logger.Errorf("error retrieving flashes for advertisewithusform %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data.Data = pageData

	if handleTemplateError(w, r, "stakingServices.go", "StakingServices", "", stakingServicesTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func getStakingServices() []types.StakingService {
	isTrue := true
	isFalse := false
	return []types.StakingService{
		{
			Name:                     "Ethpool",
			OperatedBy:               "by beaconcha.in",
			Link:                     "https://ethpool.org/",
			HasHardwareWalletSupport: true,
			HasNoSupermajorityClient: &isTrue,
			ValidatorKeyOwner:        "User & Service",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "32 ETH",
			Fee:                      "2% of rewards",
			Discord:                  "https://discord.io/ethermine",
			Twitter:                  "https://twitter.com/ethermine_org",
			IsHighlighted:            true,
		},
		{
			Name:                  "Abyss Finance",
			Link:                  "https://abyss.finance/hosting",
			ValidatorKeyOwner:     "User & Service",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "32 ETH",
			Twitter:               "https://twitter.com/abyssfinance",
			Telegram:              "https://t.me/abyssfinance",
			EMail:                 "hello@abyss.finance",
		},
		{
			Name:                     "Allnodes",
			Link:                     "https://www.allnodes.com/eth2/staking",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "User & Service",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "32 ETH",
			Fee:                      "US$10/validator/month",
			Discord:                  "https://discord.allnodes.com/",
			Twitter:                  "https://twitter.com/allnodes",
			Telegram:                 "https://t.me/allnodes",
			EMail:                    "support@allnodes.com",
		},
		{
			Name:                     "Ankr Staking",
			Link:                     "https://www.ankr.com/staking/stake/",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "Service",
			HasPoolToken:             &isTrue,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isTrue,
			MinStake:                 "Any amount",
			Fee:                      "10% of rewards",
			Discord:                  "https://discord.gg/2BSdr2g",
			Twitter:                  "https://twitter.com/ankr",
			Telegram:                 "https://t.me/ankrnetwork",
		},
		{
			Name:                  "Atomic Wallet",
			Link:                  "https://atomicwallet.io",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "Any amount",
			Twitter:               "https://twitter.com/atomicwallet",
			Telegram:              "https://t.me/AtomicWallet",
			EMail:                 "support@atomicwallet.io",
		},
		{
			Name:                     "Attestant",
			Link:                     "https://www.attestant.io/",
			HasNoSupermajorityClient: &isTrue,
			ValidatorKeyOwner:        "User & Service",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isTrue,
			MinStake:                 "1000 ETH",
			Fee:                      "12.5% of rewards",
			Discord:                  "https://invite.gg/attestant",
			Twitter:                  "https://twitter.com/attestantio",
			Telegram:                 "https://t.me/attestant",
			EMail:                    "info@attestant.io",
		},
		{
			Name:                     "Binance Exchange",
			Link:                     "https://www.binance.com/en/staking",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "Service",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "Any amount",
		},
		{
			Name:                     "Bitcoin Suisse Exchange",
			Link:                     "https://www.bitcoinsuisse.com/staking/ethereum-2",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "Service",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "Any amount",
			Fee:                      "15% of rewards",
			Twitter:                  "https://twitter.com/bitcoinsuisseag",
			Telegram:                 "https://t.me/bitcoinsuisse_announcements",
			EMail:                    "info@bitcoinsuisse.com",
		},
		{
			Name:                     "Bitfinex Exchange",
			Link:                     "https://staking.bitfinex.com/",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "Service",
			IsOpenSource:             &isFalse,
			MinStake:                 "0.00001 ETH",
			Fee:                      "25% of rewards",
			Twitter:                  "https://twitter.com/bitfinex",
		},
		{
			Name:                  "Bitfrost",
			Link:                  "https://bifrost.finance/",
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "0.1 ETH",
			Fee:                   "15% of rewards",
			Discord:               "https://discord.com/invite/XjnjdKBNXj",
			Twitter:               "https://twitter.com/bifrost_finance",
			Telegram:              "https://t.me/bifrost_finance",
		},
		{
			Name:                  "BloqStake",
			Link:                  "https://stake.bloq.com/",
			ValidatorKeyOwner:     "Service",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "32 ETH",
			Fee:                   "4-8% of rewards",
			LinkedIn:              "https://www.linkedin.com/company/bloq-inc/",
			Twitter:               "https://twitter.com/bloqinc",
			EMail:                 "support@bloq.com",
		},
		{
			Name:                     "Bloxstaking",
			Link:                     "https://www.bloxstaking.com/",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "User",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isTrue,
			IsOpenSource:             &isTrue,
			MinStake:                 "32 ETH",
			Fee:                      "Free",
			Discord:                  "https://discord.gg/cvWkmDg",
			Twitter:                  "https://twitter.com/Blox_Staking",
			EMail:                    "contact@blox.io",
		},
		{
			Name:                     "Coinbase Exchange",
			Link:                     "https://www.coinbase.com/earn/staking/ethereum",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "Service",
			IsOpenSource:             &isFalse,
			MinStake:                 "0.00001 ETH",
			Fee:                      "25% of rewards",
			Discord:                  "https://twitter.com/coinbase",
		},
		{
			Name:                  "Codefi Staking",
			Link:                  "https://consensys.net/codefi/staking/",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "32 ETH",
			Fee:                   "5% of rewards",
			Discord:               "https://discord.gg/YTcA2mxKCm",
			Twitter:               "https://twitter.com/ConsenSysCodefi",
		},
		{
			Name:                  "Cybavo",
			Link:                  "https://www.cybavo.com/eth2-staking/",
			ValidatorKeyOwner:     "Service & User",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "32 ETH",
			Fee:                   "10% of reward",
			Discord:               "https://twitter.com/cybavo",
		},
		{
			Name:         "Dragonstake",
			Link:         "https://dragonstake.io/",
			IsOpenSource: &isFalse,
			MinStake:     "Any amount",
		},
		{
			Name:         "Everstake",
			Link:         "https://everstake.one/",
			IsOpenSource: &isFalse,
			MinStake:     "Any amount",
		},
		{
			Name:                  "Guarda Wallet",
			Link:                  "https://guarda.com/",
			ValidatorKeyOwner:     "Service",
			WithdrawalKeyOwner:    "Service",
			HasPoolToken:          &isTrue,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "0.1 ETH",
			Fee:                   "9.99% of rewards",
			Twitter:               "https://twitter.com/GuardaWallet",
			Telegram:              "https://t.me/Guarda_community",
			EMail:                 "guarda.com",
		},
		{
			Name:                     "Ebunker",
			Link:                     "https://www.ebunker.io/",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "User & Service",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "32 ETH",
			Fee:                      "0-25% of rewards",
			Discord:                  "https://discord.com/invite/nuvw6hmvnK",
			Twitter:                  "https://twitter.com/ebunker_eth",
			Telegram:                 "https://t.me/ebunkerio",
		},
		{
			Name:                  "imToken & InfStones",
			Link:                  "https://token.im/",
			ValidatorKeyOwner:     "Service",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "32 ETH",
			Fee:                   "US$100/validator",
			Twitter:               "https://twitter.com/imTokenOfficial",
			EMail:                 "support@token.im",
		},
		{
			Name:                     "Kraken Exchange",
			Link:                     "https://www.kraken.com/en-us/features/staking-coins",
			HasNoSupermajorityClient: &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "0.00001 ETH",
			Fee:                      "15% of rewards",
			Twitter:                  "https://twitter.com/krakenfx",
		},
		{
			Name:                  "Launchnodes",
			Link:                  "https://www.launchnodes.com/",
			ValidatorKeyOwner:     "User",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isTrue,
			OpenSourceLink:        "https://github.com/launchnodes/ValidatorNodeProduct",
			MinStake:              "32 ETH",
			Fee:                   "US$20/validator/month",
			AWS:                   "https://aws.amazon.com/marketplace/seller-profile?id=0eee06cf-d077-4219-9d46-642372b9df1b",
			Twitter:               "https://twitter.com/launchnodes",
			LinkedIn:              "https://www.linkedin.com/company/launchnodes",
			EMail:                 "launchnodes.com",
		},
		{
			Name:                     "Lido Finance",
			Link:                     "https://www.lido.fi",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "Service",
			HasPoolToken:             &isTrue,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isTrue,
			OpenSourceLink:           "https://github.com/lidofinance",
			MinStake:                 "Any amount",
			Fee:                      "10% of rewards",
			Discord:                  "https://discord.com/invite/vgdPfhZ",
			Twitter:                  "https://twitter.com/lidofinance",
			Telegram:                 "https://t.me/lidofinance",
			EMail:                    "info@lido.fi",
		},
		{
			Name:                  "MIDL.dev",
			Link:                  "https://midl.dev/ethereum/",
			ValidatorKeyOwner:     "Service",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isTrue,
			MinStake:              "Any amount",
			Fee:                   "US$45USD/validator/month",
			Twitter:               "https://twitter.com/midl_dev",
			EMail:                 "hello@midl.dev",
		},
		{
			Name:                  "MyCointainer",
			Link:                  "https://www.mycointainer.com/",
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "0.1 ETH",
			Fee:                   "2.27% of rewards",
			Twitter:               "https://twitter.com/MyCointainerCom",
			Telegram:              "https://t.me/mycointainer",
		},
		{
			Name:                     "P2P.org",
			Link:                     "https://ethereum-staking.p2p.org",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "32 ETH",
			Fee:                      "10% of rewards",
			Twitter:                  "https://twitter.com/P2Pvalidator",
			Telegram:                 "https://t.me/P2Pstaking",
			EMail:                    "letsgo@p2p.org",
		},
		{
			Name:                     "Rocket Pool (Pool)",
			Link:                     "https://www.rocketpool.net/",
			HasNoSupermajorityClient: &isTrue,
			ValidatorKeyOwner:        "User",
			WithdrawalKeyOwner:       "Smart Contract",
			HasPoolToken:             &isTrue,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isTrue,
			OpenSourceLink:           "https://github.com/rocket-pool",
			MinStake:                 "0.01 ETH",
			Fee:                      "15% of rewards",
			Discord:                  "https://discordapp.com/invite/tCRG54c",
			Twitter:                  "https://twitter.com/Rocket_Pool",
		},
		{
			Name:                     "Rocket Pool (Validator)",
			Link:                     "https://www.rocketpool.net/",
			HasNoSupermajorityClient: &isTrue,
			ValidatorKeyOwner:        "User",
			WithdrawalKeyOwner:       "Smart Contract",
			HasPoolToken:             &isTrue,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isTrue,
			MinStake:                 "16 ETH",
			Fee:                      "Free",
			Discord:                  "https://discordapp.com/invite/tCRG54c",
			Twitter:                  "https://twitter.com/Rocket_Pool",
		},
		{
			Name:                  "StaFi",
			Link:                  "https://www.stafi.io/",
			HasThirdPartySoftware: &isTrue,
			IsOpenSource:          &isTrue,
			MinStake:              "0.01 ETH<",
			Fee:                   "20% of rewards",
			Twitter:               "https://twitter.com/Stafi_Protocol",
		},
		{
			Name:                  "StakedUs",
			Link:                  "https://staked.us/",
			ValidatorKeyOwner:     "Service",
			WithdrawalKeyOwner:    "User",
			HasPoolToken:          &isFalse,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isFalse,
			MinStake:              "32 ETH",
			Fee:                   "US$5/validator/month<",
			Discord:               "https://discord.gg/48sPdbm",
			Twitter:               "https://twitter.com/staked_us",
			Telegram:              "https://t.me/staked_official",
			EMail:                 "staked@staked.us",
		},
		{
			Name:                     "stakefish",
			Link:                     "https://stake.fish/",
			HasNoSupermajorityClient: &isFalse,
			ValidatorKeyOwner:        "Service",
			WithdrawalKeyOwner:       "User",
			HasPoolToken:             &isFalse,
			HasThirdPartySoftware:    &isFalse,
			IsOpenSource:             &isFalse,
			MinStake:                 "0.1 ETH",
			Fee:                      "10% of rewards",
			Twitter:                  "https://twitter.com/stakefish",
			Telegram:                 "https://t.me/stakefish",
			EMail:                    "hi@stake.fish",
		},
		{
			Name:                  "Stakewise Pool",
			Link:                  "https://stakewise.io/app/pool",
			ValidatorKeyOwner:     "Service",
			WithdrawalKeyOwner:    "Smart Contract",
			HasPoolToken:          &isTrue,
			HasThirdPartySoftware: &isFalse,
			IsOpenSource:          &isTrue,
			MinStake:              "Any amount",
			Fee:                   "10% of rewards",
			Discord:               "https://discord.gg/8Zf7tKyXeZ",
			Twitter:               "https://twitter.com/stakewise_io",
			Telegram:              "https://t.me/stakewise_io",
			EMail:                 "kirill@stakewise.io",
		},
		{
			Name:         "Staking Facilities",
			Link:         "https://stakingfacilities.com/",
			IsOpenSource: &isFalse,
			Twitter:      "https://twitter.com/StakingFac",
			Telegram:     "https://t.me/stakingfacilities",
		},
		{
			Name:                  "Swissborg",
			Link:                  "https://swissborg.com/",
			HasThirdPartySoftware: &isTrue,
			IsOpenSource:          &isFalse,
			Fee:                   "1% of rewards",
			Twitter:               "https://twitter.com/swissborg",
		},
	}
}
