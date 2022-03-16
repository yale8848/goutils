git add -A
git commit -a -m "update"
git push origin master

git tag -a %1 -m "update"
git push --tags