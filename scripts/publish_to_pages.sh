#!/bin/sh

if [ "`git status -s`" ]
then
    echo "The working directory is dirty. Please commit any pending changes."
    git status
    exit 1;
fi

echo "Deleting old publication"
mkdir tmp
cp ../CNAME ./tmp/CNAME
rm -rf public
mkdir public
git worktree prune
rm -rf .git/worktrees/public/

echo "Checking out gh-pages branch into public"
git worktree add -B gh-pages public origin/gh-pages

echo "Removing existing files"
rm -rf public/*


echo "Generating site"
hugo
mv ./tmp/CNAME ./public/CNAME

echo "Updating gh-pages branch"
cd public && git add --all && git commit -m "Publishing to gh-pages (publish.sh) [ci skip]" || true

echo "Pushing to github"
git push origin gh-pages