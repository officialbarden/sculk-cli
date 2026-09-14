package add

import (
)


// Steps
// 1. Clone the "Github Repo" into memory
// 2. Walk through all the files and merge them one by one

func Main(ignoreVersionMismatch bool, args []string) error { 
	libraries := args

	
	// 1. Clone the "Github Repo" into memory
	GetLibraries(libraries)

	return nil
}
