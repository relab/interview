# Tasks for interviews

## Important Instructions

Do not fork this repository, since that would create links that become visible to the other applicant’s.
Instead, create a **private repo without any files** on your personal GitHub account, and add a remote pointing to this `interview` repository, and pull from our repo into yours.
You can then push your changes to your personal version of the interview repo.

In the command trace below, replace `meling` with your GitHub username, and make sure that your repo is private.

```console
# Create private repo called interview on GitHub first
% git clone git@github.com:meling/interview.git
Cloning into 'interview'...
warning: You appear to have cloned an empty repository.
% cd interview
% git remote add uis-interview git@github.com:relab/interview.git
% git pull uis-interview main
% git push
```
