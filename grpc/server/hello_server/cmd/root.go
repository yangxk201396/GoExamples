package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yang201396/GoExamples/grpc/server/hello_server/server"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run the gRPC hello-world server",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error : %v", err)
			}
		}()

		err := server.Run()
		if err != nil {
			log.Println("Run err: ", err.Error())
			return
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.Port, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "../../conf/server.pem", "cert-pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "../../conf/server-key.pem", "cert-key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "server-name", "", "127.0.0.1", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "../../proto", "path to the directory which contains swagger definitions")

	rootCmd.AddCommand(serverCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("Execute err: ", err.Error())
		os.Exit(-1)
	}
}
