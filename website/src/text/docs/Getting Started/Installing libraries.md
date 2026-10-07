# Installing a sculk library*

Just like how a developer may download a particular package, say for example: the pandas library* of python, the command for it in command-line would be: `pip install pandas`

Similarly, sculk-cli uses acronyms/identifiers to identify certain libraries/projects. Sculk doesn't have its own package storage and hosting platform, so it sources the packages from Git Hosting platforms like GitHub and Codeberg. Internally, each unique "identifier" is mapped to a source repository, which is cloned and seemlessly merged with the current project. Alternatively, you can also provide a link directly.

Run the following command:
```
sculk add id-system
```

Sculk will then reach in its internal hashmap, get the source link and get the contents of the library*, and merge them with your sculk-project as is. Your libraries.json will change, and a new library* will be added in the 'libraries[]' field.

```json
{
   "author": "USER",
   "version": "1.0.0",
   "game_version": "26.2",
   "libraries": [
      {
         "identifier": "id-system",
         "source": "http://github.com/officialbarden/id-system",
         "version": "1.0.0",
         "game_version": "26.2"
      }
   ]
}
```

Here, you can see that the library* information is stored inside your `libraries.json` file. You can share this .json file with other sculk developers, and ask them to run the command `sculk install` to install all packages specified in the 'libraries[]' field.

If you're working on an unstable version, like a snapshot: you can use the `--ignore` flag to ignore Game Version Mismatch Checking and install packages directly meant for other Game Versions into your sculk project. 

Alternatively, you can also run the command:
Run the following command:
```
sculk add https://github.com/officialbarden/id-system
```

Sculk will then install the contents of the repository into your datapack, and will update your `libraries.json` to look like the following:

```json
{
   "author": "USER",
   "version": "1.0.0",
   "game_version": "26.2",
   "libraries": [
      {
         "identifier": "http://github.com/officialbarden/id-system",
         "source": "http://github.com/officialbarden/id-system",
         "version": "1.0.0",
         "game_version": "26.2"
      }
   ]
}
```

This allows Sculk to have a *decentralized library ecosystem*.

*Note: By using the --ignore flag, you are to take full responsibility of how the library interacts with your sculk project, as some additional stuff, like tags, advancements, predicates and more may get installed and merged into your project without a warning!*

## Why have a seperate libraries.json if the library* contents are merged in the final datapack, anyways?

Great Question! Sculk isn't just a 'library* installer & merger', it's a full blown package manager! (well, a barebones version of a traditional package manager). By having a seperate 'libraries.json' file in the sculk-project's root directory, sculk developers can keep track of what libraries exist in their codebase! At the same time, say a sculk-library* developer pushed a new version of the library* onto their git repository, by running one command in the shell (`sculk update`), sculk developers can pull the changes into their files without the hassle of merging new updates by hand.


Notes:
- library: in npm & pip, they're called packages. In cargo, they're called crates and in sculk, they're called libraries. Though, we don't object to interchangability of the two terms: library and package.
