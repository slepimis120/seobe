# Why GitHub Actions Importer?
***
There are multiple reasons we chose to go with GitHub Actions Importer, main ones being their extensive documentation and ability to expand the tool. The tool isn't perfect, but GitHub Actions Importer developers mention a number of around 80% of the work being done for you with just the tool.

## GitHub Actions Importer Documentation
We won't go too much in depth about the GitHub Actions Importer documentation, the whole docs can be seen [here](https://docs.github.com/en/actions). It was one of the main resources I came back to while working on this project, because it even describes the flaws with the whole system and how we can expand upon it.
## Extensibility with Transformers
As I just mentioned, Github Actions importer offers us the ability to expand upon it. We can do this by creating transformers, which are basically scripts that can modify the GitHub Actions workflow file. For the sake of this project, we only wrote a single "script" transformer that seems to do good enough of a job.
## Flaws with GitHub Actions Importer
Even though GitHub Actions importer can do most of the stuff we're looking for, there are some limitations that we need to address. 
- Lack of sharing workspaces
  - In Jenkinsfile, all of the jobs share the same workspace, which would mean you only need to build the project once in the entire Jenkinsfile.
  - In GitHub Actions, each job has its own workspace, which means you need to build the project multiple times.
  - We tried to fix this by using the cache actions, which only goes so far to cover Java projects.
- Not everything in Jenkins can be converted 1:1 to GitHub Actions
  - There are some things that we can't convert to GitHub Actions, like `post` and `excludes` in Jenkins.
  - This mostly comes down to manual implementation, implementing post and excludes in the GitHub Actions workflow file.
  - Good thing is that GH Actions Importer leaves a comment if some specific part of Jenkinsfile wasn't converted properly.
  - Along with that, GitHub left resources that explain what can and can't be converted 1:1, you can read more about it [here](https://docs.github.com/en/actions/migrating-to-github-actions/manually-migrating-to-github-actions/migrating-from-jenkins-to-github-actions)

## Other options
One of the first options we considered was to write a custom tool that would parse the Jenkinsfile and create a GitHub Actions workflow file. That would mean creating a AST tree for the entirety of Jenkinsfile and then converting it to a GitHub Actions workflow file. That would also include covering every single gramatical rule that Jenkinsfile has, which is an insane amount of work to do for a project like this.

Another option we had was to use ANTLR4 to parse the Jenkinsfile and create a GitHub Actions workflow file. This would be a bit easier than writing a custom tool, but it would still require a lot of work to cover every single rule that Jenkinsfile has, and as of right now, there aren't any official ANTLR4 grammars for Groovy syntax.

In the end, we decided to go with GitHub Actions Importer because it seemed like the best option for us, even with its flaws. We tried to expand upon it and make it better with custom Go script that adds what was missing to the file.
