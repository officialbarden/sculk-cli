package add


// Steps
// 1. Clone the "Github Repo" into memory
// 2. Walk through all the files and merge them one by one

func Main(ignoreVersionMismatch bool, args []string) error { 

	for i := range args {
		// check if its already installed
		if IsPreinstalled(args[i]) {
			continue
		} else {
			InstallLibraries(args[i], ignoreVersionMismatch)
		}
	}
	return nil
}
