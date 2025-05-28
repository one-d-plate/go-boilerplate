package cmd

import (
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "rc [config-file]",
	Short: "Register a new client service",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()

		// db := bootstrap.NewDatabase(ctx)
		// dbConnect, err := db.Connect()
		// if err != nil {
		// 	pkg.Logger.Error("Database connection failed ", err)
		// 	return err
		// }

		// authClientRepo := repository.NewAuthClientRepo(*dbConnect)
		// authClientSvc := service.NewAuthClientService(authClientRepo, nil)

		// configPath := args[0]
		// var req dto.CreateAuthClientRequest

		// if err := pkg.LoadJSONFile(configPath, &req); err != nil {
		// 	pkg.LogError("failed to load config: ", err)
		// 	return fmt.Errorf("failed to load config: %w", err)
		// }

		// fmt.Printf("Are you sure you want to register client '%s'? (y/n): ", req.Name)
		// var confirm string
		// fmt.Scanln(&confirm)
		// confirm = strings.ToLower(confirm)
		// if confirm != "y" && confirm != "yes" {
		// 	fmt.Println("❌ Aborted.")
		// 	return errors.New("aborted")
		// }

		// res, err := authClientSvc.Insert(ctx, req)
		// if err != nil {
		// 	return fmt.Errorf("insert failed: %w", err)
		// }

		// fmt.Println("✅ Client created successfully")
		// fmt.Println("Client ID:     ", res.ClientID)
		// fmt.Println("Client Secret: ", res.ClientSecret)

		// return nil
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
