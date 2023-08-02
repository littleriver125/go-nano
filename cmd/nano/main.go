package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"littleriver.cc/go-nano/cmd/utils"
	"littleriver.cc/go-nano/console/prompt"
	"littleriver.cc/go-nano/internal/debug"
	"littleriver.cc/go-nano/internal/flags"
	"littleriver.cc/go-nano/metrics"
	"os"
	"sort"
	"time"
)

const (
	clientIdentifier = "nano" // Client identifier to advertise over the network
)

var (
	nodeFlags = flags.Merge([]cli.Flag{
		//utils.IdentityFlag,
		//configFileFlag,
	}, utils.NetworkFlags, utils.DatabasePathFlags)

	rpcFlags = []cli.Flag{
		//utils.HTTPEnabledFlag,
		//utils.HTTPListenAddrFlag,
		//utils.HTTPPortFlag,
		//utils.HTTPCORSDomainFlag,
		//utils.AuthListenFlag,
		//utils.AuthPortFlag,
		//utils.AuthVirtualHostsFlag,
		//utils.JWTSecretFlag,
		//utils.HTTPVirtualHostsFlag,
		//utils.GraphQLEnabledFlag,
		//utils.GraphQLCORSDomainFlag,
		//utils.GraphQLVirtualHostsFlag,
		//utils.HTTPApiFlag,
		//utils.HTTPPathPrefixFlag,
		//utils.WSEnabledFlag,
		//utils.WSListenAddrFlag,
		//utils.WSPortFlag,
		//utils.WSApiFlag,
		//utils.WSAllowedOriginsFlag,
		//utils.WSPathPrefixFlag,
		//utils.IPCDisabledFlag,
		//utils.IPCPathFlag,
		//utils.InsecureUnlockAllowedFlag,
		//utils.RPCGlobalGasCapFlag,
		//utils.RPCGlobalEVMTimeoutFlag,
		//utils.RPCGlobalTxFeeCapFlag,
		//utils.AllowUnprotectedTxs,
		//utils.BatchRequestLimit,
		//utils.BatchResponseMaxSize,
	}

	metricsFlags = []cli.Flag{
		//utils.MetricsEnabledFlag,
		//utils.MetricsEnabledExpensiveFlag,
		//utils.MetricsHTTPFlag,
		//utils.MetricsPortFlag,
		//utils.MetricsEnableInfluxDBFlag,
		//utils.MetricsInfluxDBEndpointFlag,
		//utils.MetricsInfluxDBDatabaseFlag,
		//utils.MetricsInfluxDBUsernameFlag,
		//utils.MetricsInfluxDBPasswordFlag,
		//utils.MetricsInfluxDBTagsFlag,
		//utils.MetricsEnableInfluxDBV2Flag,
		//utils.MetricsInfluxDBTokenFlag,
		//utils.MetricsInfluxDBBucketFlag,
		//utils.MetricsInfluxDBOrganizationFlag,
	}
)

var app = flags.NewApp("the go-nano command line interface")

func init() {
	// Initialize the CLI app and start Geth
	app.Action = nano
	app.Copyright = "Copyright 2013-2023 The go-nano Authors"
	app.Commands = []*cli.Command{
		//// See chaincmd.go:
		//initCommand,
		//importCommand,
		//exportCommand,
		//importPreimagesCommand,
		//exportPreimagesCommand,
		//removedbCommand,
		//dumpCommand,
		//dumpGenesisCommand,
		//// See accountcmd.go:
		//accountCommand,
		//walletCommand,
		//// See consolecmd.go:
		//consoleCommand,
		//attachCommand,
		//javascriptCommand,
		//// See misccmd.go:
		//versionCommand,
		//versionCheckCommand,
		//licenseCommand,
		//// See config.go
		//dumpConfigCommand,
		//// see dbcmd.go
		//dbCommand,
		//// See cmd/utils/flags_legacy.go
		//utils.ShowDeprecated,
		//// See snapshot.go
		//snapshotCommand,
		//// See verkle.go
		//verkleCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = flags.Merge(
	//nodeFlags,
	//rpcFlags,
	//consoleFlags,
	//debug.Flags,
	//metricsFlags,
	)

	app.Before = func(ctx *cli.Context) error {
		flags.MigrateGlobalFlags(ctx)
		return debug.Setup(ctx)
	}
	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		prompt.Stdin.Close() // Resets terminal mode.
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func prepare(ctx *cli.Context) {
	//	switch {
	//	case ctx.IsSet(utils.GoerliFlag.Name):
	//		log.Info("Starting Geth on Görli testnet...")
	//
	//	case ctx.IsSet(utils.SepoliaFlag.Name):
	//		log.Info("Starting Geth on Sepolia testnet...")
	//
	//	case ctx.IsSet(utils.DeveloperFlag.Name):
	//		log.Info("Starting Geth in ephemeral dev mode...")
	//		log.Warn(`You are running Geth in --dev mode. Please note the following:
	//
	//`)
	//
	//	case !ctx.IsSet(utils.NetworkIdFlag.Name):
	//		log.Info("Starting Geth on Ethereum mainnet...")
	//	}
	//if ctx.String(utils.SyncModeFlag.Name) != "light" && !ctx.IsSet(utils.CacheFlag.Name) && !ctx.IsSet(utils.NetworkIdFlag.Name) {
	//	// Make sure we're not on any supported preconfigured testnet either
	//	if !ctx.IsSet(utils.SepoliaFlag.Name) &&
	//		!ctx.IsSet(utils.GoerliFlag.Name) &&
	//		!ctx.IsSet(utils.DeveloperFlag.Name) {
	//		// Nope, we're really on mainnet. Bump that cache up!
	//		log.Info("Bumping default cache on mainnet", "provided", ctx.Int(utils.CacheFlag.Name), "updated", 4096)
	//		ctx.Set(utils.CacheFlag.Name, strconv.Itoa(4096))
	//	}
	//}
	//if ctx.String(utils.SyncModeFlag.Name) == "light" && !ctx.IsSet(utils.CacheFlag.Name) {
	//	log.Info("Dropping default light client cache", "provided", ctx.Int(utils.CacheFlag.Name), "updated", 128)
	//	ctx.Set(utils.CacheFlag.Name, strconv.Itoa(128))
	//}

	// Start metrics export if enabled
	//utils.SetupMetrics(ctx)

	// Start system runtime metrics collection
	go metrics.CollectProcessMetrics(3 * time.Second)
}

func nano(ctx *cli.Context) error {
	if args := ctx.Args().Slice(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}

	prepare(ctx)
	//stack, backend := makeFullNode(ctx)
	//defer stack.Close()

	//startNode(ctx, stack, backend, false)
	//stack.Wait()
	return nil
}

//func startNode(ctx *cli.Context, stack *node.Node, backend ethapi.Backend, isConsole bool) {
//	debug.Memsize.Add("node", stack)
//
//	// Start up the node itself
//	utils.StartNode(ctx, stack, isConsole)
//
//	// Unlock any account specifically requested
//	unlockAccounts(ctx, stack)
//
//	// Register wallet event handlers to open and auto-derive wallets
//	events := make(chan accounts.WalletEvent, 16)
//	stack.AccountManager().Subscribe(events)
//
//	// Create a client to interact with local geth node.
//	rpcClient := stack.Attach()
//	ethClient := ethclient.NewClient(rpcClient)
//
//	go func() {
//		// Open any wallets already attached
//		for _, wallet := range stack.AccountManager().Wallets() {
//			if err := wallet.Open(""); err != nil {
//				log.Warn("Failed to open wallet", "url", wallet.URL(), "err", err)
//			}
//		}
//		// Listen for wallet event till termination
//		for event := range events {
//			switch event.Kind {
//			case accounts.WalletArrived:
//				if err := event.Wallet.Open(""); err != nil {
//					log.Warn("New wallet appeared, failed to open", "url", event.Wallet.URL(), "err", err)
//				}
//			case accounts.WalletOpened:
//				status, _ := event.Wallet.Status()
//				log.Info("New wallet appeared", "url", event.Wallet.URL(), "status", status)
//
//				var derivationPaths []accounts.DerivationPath
//				if event.Wallet.URL().Scheme == "ledger" {
//					derivationPaths = append(derivationPaths, accounts.LegacyLedgerBaseDerivationPath)
//				}
//				derivationPaths = append(derivationPaths, accounts.DefaultBaseDerivationPath)
//
//				event.Wallet.SelfDerive(derivationPaths, ethClient)
//
//			case accounts.WalletDropped:
//				log.Info("Old wallet dropped", "url", event.Wallet.URL())
//				event.Wallet.Close()
//			}
//		}
//	}()
//
//	// Spawn a standalone goroutine for status synchronization monitoring,
//	// close the node when synchronization is complete if user required.
//	if ctx.Bool(utils.ExitWhenSyncedFlag.Name) {
//		go func() {
//			sub := stack.EventMux().Subscribe(downloader.DoneEvent{})
//			defer sub.Unsubscribe()
//			for {
//				event := <-sub.Chan()
//				if event == nil {
//					continue
//				}
//				done, ok := event.Data.(downloader.DoneEvent)
//				if !ok {
//					continue
//				}
//				if timestamp := time.Unix(int64(done.Latest.Time), 0); time.Since(timestamp) < 10*time.Minute {
//					log.Info("Synchronisation completed", "latestnum", done.Latest.Number, "latesthash", done.Latest.Hash(),
//						"age", common.PrettyAge(timestamp))
//					stack.Close()
//				}
//			}
//		}()
//	}
//
//	// Start auxiliary services if enabled
//	if ctx.Bool(utils.MiningEnabledFlag.Name) {
//		// Mining only makes sense if a full Ethereum node is running
//		if ctx.String(utils.SyncModeFlag.Name) == "light" {
//			utils.Fatalf("Light clients do not support mining")
//		}
//		ethBackend, ok := backend.(*eth.EthAPIBackend)
//		if !ok {
//			utils.Fatalf("Ethereum service not running")
//		}
//		// Set the gas price to the limits from the CLI and start mining
//		gasprice := flags.GlobalBig(ctx, utils.MinerGasPriceFlag.Name)
//		ethBackend.TxPool().SetGasTip(gasprice)
//		if err := ethBackend.StartMining(); err != nil {
//			utils.Fatalf("Failed to start mining: %v", err)
//		}
//	}
//}

//func unlockAccounts(ctx *cli.Context, stack *node.Node) {
//	var unlocks []string
//	inputs := strings.Split(ctx.String(utils.UnlockedAccountFlag.Name), ",")
//	for _, input := range inputs {
//		if trimmed := strings.TrimSpace(input); trimmed != "" {
//			unlocks = append(unlocks, trimmed)
//		}
//	}
//	// Short circuit if there is no account to unlock.
//	if len(unlocks) == 0 {
//		return
//	}
//	// If insecure account unlocking is not allowed if node's APIs are exposed to external.
//	// Print warning log to user and skip unlocking.
//	if !stack.Config().InsecureUnlockAllowed && stack.Config().ExtRPCEnabled() {
//		utils.Fatalf("Account unlock with HTTP access is forbidden!")
//	}
//	backends := stack.AccountManager().Backends(keystore.KeyStoreType)
//	if len(backends) == 0 {
//		log.Warn("Failed to unlock accounts, keystore is not available")
//		return
//	}
//	ks := backends[0].(*keystore.KeyStore)
//	passwords := utils.MakePasswordList(ctx)
//	for i, account := range unlocks {
//		unlockAccount(ks, account, i, passwords)
//	}
//}

//var (
//	rootCmd = &cobra.Command{
//		Use:   "nano-node",
//		Short: "Nano node",
//		Run:   startNode,
//	}
//
//	man         *Manager
//	cfg         Config
//	cfgDefaults = Config{
//		Addr: node.DefaultOptions.Address,
//		Peers: []string{
//			"rai.raiblocks.net:7075",
//		},
//		Network: proto.NetworkLive,
//	}
//
//	logger = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
//)

//func init() {
//	rootCmd.Flags().StringVar(&cfg.Addr, "addr", cfgDefaults.Addr, "address to listen on for UDP and TCP")
//	rootCmd.Flags().StringVar(&cfg.AddrRPC, "addr-rpc", cfgDefaults.AddrRPC, "address to listen on for RPC")
//	rootCmd.Flags().StringVar(&cfg.AddrPprof, "addr-pprof", cfgDefaults.AddrPprof, "address to listen on for pprof")
//	cobra.OnInitialize(initConfig)
//}

//func initConfig() {
//	var err error
//	if man, err = NewManager(".config/gonano/node", "config.json"); err != nil {
//		logger.Fatalf("config manager error: %s", err)
//	}
//
//	if err = man.Prepare(&cfgDefaults); err != nil {
//		logger.Fatalf("config load error: %s", err)
//	}
//
//	if err = man.Load(&cfg); err != nil {
//		logger.Fatalf("config load error: %s", err)
//	}
//}

//func startNode(cmd *cobra.Command, args []string) {
//	startPprof()
//
//	gen, err := genesis.Get(cfg.Network)
//	if err != nil {
//		logger.Fatalf("error obtaining genesis info: %s", err)
//	}
//	ledgerOpts := store.LedgerOptions{Genesis: gen}
//
//	nodeOpts := node.DefaultOptions
//	nodeOpts.Peers = cfg.Peers
//	nodeOpts.Address = cfg.Addr
//	nodeOpts.Network = cfg.Network
//
//	logger.Printf("opening badger database at %s", man.Dir())
//	db, err := store.NewBadgerStore(path.Join(man.Dir(), "db"))
//	if err != nil {
//		logger.Fatalf("error opening database: %s", err)
//	}
//	defer db.Close()
//
//	logger.Printf("initializing ledger")
//	ledger, err := store.NewLedger(db, ledgerOpts)
//	if err != nil {
//		logger.Fatalf("error initializing ledger: %s", err)
//	}
//
//	logger.Printf("initializing node")
//	nanode, err := node.New(ledger, nodeOpts)
//	if err != nil {
//		logger.Fatalf("error initializing node: %s", err)
//	}
//	go func() {
//		logger.Printf("starting node (network: %s)", nodeOpts.Network)
//		if err := nanode.Run(); err != nil {
//			logger.Fatalf("error starting node: %s", err)
//		}
//	}()
//
//	sigc := make(chan os.Signal, 1)
//	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
//
//	for {
//		switch <-sigc {
//		case syscall.SIGTERM, syscall.SIGINT:
//			logger.Println("exit signal caught")
//
//			logger.Println("stopping the node")
//			if err := nanode.Stop(); err != nil {
//				logger.Printf("error stopping node: %s", err)
//			}
//
//			logger.Println("closing badger database")
//			if err := db.Close(); err != nil {
//				logger.Printf("error closing db: %s", err)
//			}
//
//			os.Exit(0)
//		case syscall.SIGHUP:
//			logger.Println("error reloading config: not implemented")
//		}
//	}
//}
//
//func startPprof() {
//	if cfg.AddrPprof == "" {
//		return
//	}
//
//	logger.Printf("starting pprof http server at %s/debug/pprof", cfg.AddrPprof)
//
//	go func() {
//		logger.Printf("error running pprof: %s", http.ListenAndServe(cfg.AddrPprof, nil))
//	}()
//}
