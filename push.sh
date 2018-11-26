
git config --global user.email "travis@travis-ci.org"
git config --global user.name "Travis CI"
git add -A
git commit --message "Travis build: go ${TRAVIS_GO_VERSION}"
git remote add origin2 https://clementauger:${GH_TOKEN}@github.com/${TRAVIS_REPO_SLUG}.git > /dev/null 2>&1
git push --set-upstream origin2 master
