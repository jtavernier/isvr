# Contribution Guidelines
---
Want to participate to the project ? Be sure to be familiar with theses guidelines before starting.

##### Git Branching Model
This project follows the "Successful Git Branching Model" described  [here](https://nvie.com/posts/a-successful-git-branching-model/?)
##### Application Versioning
We used a **VERSION file** at the root level of the repository to track the version of the application.
The version number is defined as follow: 
> **MAJOR.MINOR.PATCH**

You should **always update the version** file before merging a pull request or committing to release/master branches.

### Summary
* [How to create a new feature ?](#how-to-create-a-new-feature)
* [How to deploy to QA, UAT ?](#how-to-deploy-to-qa-uat)
* [How to deploy to Production ?](#how-to-deploy-to-production)
* [How to create a hotfix ?](#how-to-create-a-hotfix)


### How to create a new Feature ?
---
Whenever you want to develop a new feature you need to follow theses steps:
##### 1 - Create a feature branch -  feature/XXXX
First, create a feature branch from the latest version of the *develop* branch.
By convention feature branches are named **feature/XXXXX** where XXXX is you work item ID or a short description of your work. If you're fixing a bug use **bugfix/XXXXX**.
You can now work locally on your feature branch. 

##### 2 - Push your changes
While you're working or when you're done you can push your changes to the server. It will trigger the 
CI Build that will run the tests against your branch. Obviously, all the tests need to pass before going further.

##### 3 - Create a Pull Request on develop 
Whenever you branch is ready you will need to merge it to the *develop* branch. To do so create a new pull request from 
your feature branch to the *develop* branch. Please provide a quick description of you work and link you PR to a work item
if you got one. 

##### 4 - Update the Version Number
Before merging your pull request, don't forget to **update the version number** defined in the VERSION file. If it's a feature branch increment the MINOR version; If it's a bugfix increment the PATCH version; If you're introducing breaking change you increment MAJOR version.

##### 5 - Review and merge 
After your pull request have been approved you will be able to merge it. Be sure that you're commits are properly squashed 
before merging. Once it's done, don't forget to delete you're feature branch.

### How to deploy to QA, UAT ?
---
You have successfully added a new feature to the development branch, you now need to deploy your changes to a test environment.

##### 1 - Create a release branch - release/MAJOR.MINOR
First, you need to create a release branch from the develop branch. 
Please use the following naming convention :
- **release/MAJOR.MINOR**

With MAJOR and MINOR the Major and Minor version numbers *(eg: release/1.1)*

##### 2 - Creating Artifacts
Every update on a release branch will trigger the 'Publish' build that create artifacts in the following directory:
[\\vsdeploy\BuildOutputs\drops\Security\IdentityServer]{\\vsdeploy\BuildOutputs\drops\Security\IdentityServer}

##### 3 - Deploy applications using UDeploy
At the time of writting no continuous deployment is configured. Therefore, you will need to manually deploy applications
using UDeploy.

##### 4 - Fix Bugs on your local release branch
While your release candidate will go through QA and UAT you may need to fix some bugs.
In that case, work you can work directly on your release branch. 
Before pushing your changes don't forget to :
 - Squash you commits
 - **Increment the PATCH version number in the VERSION file**
 
Then you can finally push your fixes, it will generate a new artifact that you can redeploy.
You should only commit bugfix on your release branch, please refer to [How to create a new feature ?](#how-to-create-a-new-feature) if you want to add a feature.

##### 5 - Backport your Bugfix to develop branch
If you have committed some bugfixes on your release branch don't forget to backport them to the *develop* branch.
You wouldn't want to fix them twice.

### How to deploy to production ?
---
Once your changes have been through QA and UAT and your release candidate have been signed off, you will need to deploy your changes to production. 

##### 1 - Create a Pull Request on Master
You can now open a new pull request from your release branch to the *maste*r branch. 
Be sure that you add a description and squash your commits.

##### 2 - Merge to Master
Once your pull request is approved you can merge it to production and delete your release branch.

### How to create a hotfix ?
---
If you need to resolve an urgent bug in production you will need to write a hotfix. It's the only case where
you can direclty work on master without going through a feature or bugfix branch.

##### 1 - Pull the Master branch
First, get the latest master branch locally and start to develop your fix.

##### 2 - Increase the version number 
Increment the PATCH version number 

##### 3 - Push your changes 
Squash you commits if necessary and then push your commit to the remote master branch. It will create
an artifact with your updated version number.

##### 3 - Backport your changes
Don't forget to backport your changes to the release and develop branches


