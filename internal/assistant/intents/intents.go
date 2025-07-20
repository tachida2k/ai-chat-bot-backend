package intent

type IntentDefinition struct {
	Name        string
	Description string
	Examples    []string
	Keywords    []string
}

var IntentList = []IntentDefinition{
	{
		Name:        "swap_token",
		Description: "Executes a token swap from Coinhall Routes.",
		Examples: []string{
			"I want to trade SEI for USDT.",
			"Swap 50 SEI to BNB.",
			"Convert my SEI into ETH.",
			"Trade 10 USDT into SEI.",
		},
		Keywords: []string{
			"swap", "exchange", "convert", "trade",
			"swap tokens", "exchange tokens", "convert tokens", "trade tokens",
			"swap SEI", "swap to USDT", "trade for", "convert my",
			"exchange my", "where can I swap", "how to swap",
		},
	},
	{
		Name:        "stake",
		Description: "Stake SEI tokens to validator.",
		Examples: []string{
			"I want to stake my SEI tokens.",
			"How do I delegate SEI?",
			"Stake 100 SEI with a validator.",
		},
		Keywords: []string{
			"stake", "staking", "delegate", "validator", "earn rewards",
			"staking rewards", "staking pool",
		},
	},
	{
		Name:        "unstake",
		Description: "Unstake SEI tokens.",
		Examples: []string{
			"I want to unstake my SEI tokens.",
			"How do I undelegate my SEI?",
			"Unstake 50 SEI from validator.",
		},
		Keywords: []string{
			"unstake", "unstaking", "undelegate", "withdraw stake",
			"remove stake", "stop staking", "withdraw staked", "exit staking",
		},
	},
	{
		Name:        "send_token",
		Description: "Send tokens to an EVM address.",
		Examples: []string{
			"Send 5 SEI to my friend.",
			"Transfer SEI to this wallet.",
			"Move 10 USDT to another account.",
		},
		Keywords: []string{
			"send", "transfer", "move", "send SEI", "transfer USDT",
			"send funds", "move tokens", "send crypto", "send to address",
		},
	},
	{
		Name:        "check_balance",
		Description: "Check my wallet balance.",
		Examples: []string{
			"Check my wallet balance.",
			"What is my current SEI balance?",
			"Show me my token holdings.",
		},
		Keywords: []string{
			"my balance", "wallet balance", "portfolio", "how much", "check my funds",
			"my tokens", "holdings", "how much sei", "current balance",
		},
	},
	{
		Name:        "tx_search",
		Description: "Search for a transaction.",
		Examples: []string{
			"Find this transaction hash on Sei.",
			"Check this transaction ID: 0x1234abcd.",
		},
		Keywords: []string{
			"tx", "transaction", "txid", "tx hash", "check tx",
		},
	},
	{
		Name:        "analyze_wallet",
		Description: "Analyze wallet portfolio.",
		Examples: []string{
			"Analyze this SEI wallet.",
			"What tokens does this address hold?",
			"Give me the portfolio breakdown.",
		},
		Keywords: []string{
			"analyze", "wallet overview", "token distribution", "asset summary",
			"wallet breakdown", "portfolio analysis",
		},
	},
	{
		Name:        "forbidden_topics",
		Description: "Topics unrelated to Sei or out of scope.",
		Examples: []string{
			"Write a Python script.",
			"What are the latest updates in Bitcoin?",
			"Can you explain AI and machine learning?",
		},
		Keywords: []string{
			"python", "AI", "machine learning", "bitcoin", "stock", "solana", "ethereum",
			"script", "trading bots", "smart contract", "chatgpt", "llama", "openai",
		},
	},
	{
		Name:        "default",
		Description: "Greetings, general questions about Sei.",
		Examples: []string{
			"Who is Sonia?",
			"Hey there!",
			"Tell me about Sei.",
			"I am new to Sei—where should I start?",
		},
		Keywords: []string{
			"hi", "hello", "thanks", "good morning", "good evening", "sei", "what is sei",
			"how does sei work", "tell me about sei", "getting started",
		},
	},
}
