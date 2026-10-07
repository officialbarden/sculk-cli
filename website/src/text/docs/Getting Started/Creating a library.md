
# Creating/Publishing a sculk library
You can create and submit your sculk libraries in our [sculk-discord](
https://discord.gg/JWZkAgsyry) for approval and direct integration into the cli. As long as your datapack outputs into a standard datapack, you may use whatever pre-compiler you desire however, for seemless integration with sculk, here are some rules you will have to follow while writing your sculk library.

1. Use a unique namespace<br>
Generally, we'll encourage library publishers to use the same as the identifier they wish to associate their sculk library with. However, incase there are issues like the identifier slug being pre-used, we ask developers to change their namespaces to a much more customized and unique slug. You can look at the [wiki-guide](https://minecraft.wiki/w/Identifier#Legal_characters) to learn what characters you can add to your namespace to make it more unique. <br><br>
As best practice, we would encourage developers to follow [Smithed's Conventions](https://docs.smithed.dev/conventions/index.html).

2. Do Not Use Extremely Minimal Namespaces/Filenames <br>
Same as the first rule, to prevent potential merge conflicting with other libraries, we hope that your naming convention, especially if you're writing to namespaces that are extremely common (namespaces like 'minecraft', 'code', 'namespace' etc.) for whatever reason, your file name must be unique.

Once you've followed the above recommendations/rules and created a library, make a `libraries.json` and paste the following code inside it:
```json
{
   "author": "AUTHOR-NAME",
   "version": "1.0.0",
   "game_version": "26.2",
   "libraries": []
}
```

Alternatively, you can run the command: 
```
sculk makelib
```
to create a libraries.json for your existing datapack.

Fields like author are currently only semantic. The 'version' field and 'game_version' are what you should pay attention to, as sculk will look at these values while making decisions during the merging of your library.
