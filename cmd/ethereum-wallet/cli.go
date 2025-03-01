package main

import (
	"ethereum-wallet/common/opio"
	"ethereum-wallet/database"
	"ethereum-wallet/flags"
	"ethereum-wallet/tools"

	"ethereum-wallet/config"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// func runEthWallet(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
// 	log.Info("exec wallet sync")
// 	cfg, err := config.LoadConfig(ctx)
// 	if err != nil {
// 		log.Error("failed to load config", "err", err)
// 		return nil, err
// 	}
// 	return eth_wallet.NewEthWallet(ctx.Context, &cfg, shutdown)
// }

// func runRestApi(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
// 	log.Info("running api...")
// 	cfg, err := config.LoadConfig(ctx)
// 	if err != nil {
// 		log.Error("failed to load config", "err", err)
// 		return nil, err
// 	}
// 	return api.NewApi(ctx.Context, &cfg)
// }

// func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
// 	fmt.Println("running grpc server...")
// 	cfg, err := config.LoadConfig(ctx)
// 	if err != nil {
// 		log.Error("failed to load config", "err", err)
// 		return nil, err
// 	}
// 	grpcServerCfg := &services.RpcServerConfig{
// 		GrpcHostname: cfg.RpcServer.Host,
// 		GrpcPort:     cfg.RpcServer.Port,
// 	}
// 	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
// 	if err != nil {
// 		log.Error("failed to connect to database", "err", err)
// 		return nil, err
// 	}
// 	return services.NewRpcServer(db, grpcServerCfg)
// }

func runGenerateAddress(ctx *cli.Context) error {
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	return tools.CreateAddressTools(ctx, db)
}

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	log.Info("running migrations...")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Error("fail to close database", "err", err)
		}
	}(db)
	return db.ExecuteSQLMigration(cfg.Migrations)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData),
		Description:          "An exchange wallet scanner services with rpc and rest api server",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			// {
			// 	Name:        "api",
			// 	Flags:       flags,
			// 	Description: "Run api services",
			// 	Action:      cliapp.LifecycleCmd(runRestApi),
			// },
			// {
			// 	Name:        "rpc",
			// 	Flags:       flags,
			// 	Description: "Run rpc services",
			// 	Action:      cliapp.LifecycleCmd(runRpc),
			// },
			{
				Name:        "generate-address",
				Flags:       flags,
				Description: "Run grenerate adddress tools",
				Action:      runGenerateAddress,
			},
			// {
			// 	Name:        "wallet",
			// 	Flags:       flags,
			// 	Description: "Run rpc scanner wallet services",
			// 	Action:      cliapp.LifecycleCmd(runEthWallet),
			// },
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Run database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
