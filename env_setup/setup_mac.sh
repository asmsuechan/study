# brewのインストール
# brewが存在しない場合のみインストール
if [[ `type brew` =~ (?!found)$ ]]; then
  echo "starting install brew"
  /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
  echo "install finished!"
fi

# xcodeは手でインストールする必要がある・・・？
xcode-select --install
brew install git
brew tap caskroom/cask
brew tap codekitchen/dinghy

# brew caskからアプリケーションのインストール
# cask外からインストールして既存の場合は上書き禁止にしたい

brew cask install flux
brew cask install iterm2
brew cask install vagrant
brew cask install slack
brew cask install skype
brew cask install wireshark
brew cask install evernote

if [[ -f /Application/Google\ Chrome.app ]];then
  echo "starting install google-chrome"
  brew cask install google-chrome
  echo "install finished!"
fi

# Boostnoteのインストール。.dmgをマウントしたらBoostnote.appがいる
# ここ参照
# https://hnakamur.github.io/blog/2015/04/06/install-apps-without-homebrew-cask/
if [[ -f /Application/Boostnote.app ]];then
  echo "starting install Boostnote"
  wget https://github.com/BoostIO/boost-releases/releases/download/v0.6.5/Boostnote-mac.dmg
  mount_dir=`hdiutil attach $dmg_file | awk -F '\t' 'END{print $NF}'`
  cp $mount_dir/Boostnote.app /Application
  hdiutil detach "$mount_dir"
  rm $dmg_file
  echo "install finished"
fi

# zshのインストール、設定
echo "installing zsh unless exist..."
if [[ $SHELL =~ (?!zsh)$ ]]; then
  brew install --without-etcdir zsh
  brew install zsh-completions
  git clone https://github.com/b4b4r07/zplug ~/.zplug
fi

# 設定ファイルをasmsuechanのリポジトリからダウンロードして設置
echo "putting zsh config files unless exist..."
if [ ! -f ~/.zprofile ];then
  curl https://raw.githubusercontent.com/asmsuechan/dotfiles/master/.zprofile > ~/.zprofile
elif [ ! -f ~/.zshrc ];then
  curl https://raw.githubusercontent.com/asmsuechan/dotfiles/master/.zshrc > ~/.zshrc
fi

echo 'autoload -Uz compinit' >> ~/.zshenv

# vimのインストール
if [ -f /usr/bin/vim ];then
  echo "cool! vim is waiting you!"
else
  echo "fuck! your PC doesn't have vim! fuck! fuck! fuck!"
fi

if [ -f ~/.vim/bundle ];then
  mkdir -p ~/.vim/bundle
  git clone git://github.com/Shougo/neobundle.vim ~/.vim/bundle/neobundle.vim
fi

# ~/.vimrcの設置
if [ ! -f ~/.vimrc ];then
  curl https://raw.githubusercontent.com/asmsuechan/dotfiles/master/.vimrc > ~/.vimrc
fi

# 便利ツール達のインストール
# elcapitanにmysql5.7インストールすると「「「ヤバい」」」が構わず5.7をインストールする
brew install tig
brew install rbenv
brew install mysql
brew install tree
brew install docker
brew install docker-machine
brew install dinghy
brew install nginx

# rbenvを使ってrubyのインストール
rbenv install 2.3.1
rbenv global 2.3.1
rbenv rehash

gem install rails

# 最後にインストールチェック
# ruby rails virtualbox iterm2 zsh
# あとで書く

# ログインシェルをzshに変更
sudo dscl . -create /Users/$USER UserShell `which zsh`
