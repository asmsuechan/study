# ファイルの命名規則
# .mdファイルの先頭に'# [title]'でタイトルを示す
find ./ -type f -print | xargs grep -E "^# \[" |sed -e 's/^/(https:\/\/github.com\/asmsuechan\/study\/blob\/master\//g' | sed -e 's/:#/)/g' | awk '{ print $2" "$1}' | sed -e 's/ //g' | sed -e 's/\.\///g' | (cat ./README.md | grep "* \[" | sed -e 's/* //g' |diff /dev/fd/3 -) 3<&0 | sed -e 's/< //g' |  sed -e '1,1d' | sed -e 's/^/* /g' >> README.md
