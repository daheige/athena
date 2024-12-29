package commands

// Init 初始化所有的命令
func Init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newUserCommand())
}
