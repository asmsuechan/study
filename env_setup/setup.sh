# brewのインストール
# brewが存在しない場合のみインストール
if [[ `type brew` =~ (?!found)$ ]]; then
  echo "starting install brew"
  sudo /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
  echo "install finished!"
fi

xcode-select --install
brew install git
brew tap caskroom/cask

# brew caskからアプリケーションのインストール
# cask外からインストールして既存の場合は上書き禁止にしたい

BREW_CASK_LIST=`brew cask list`

if [[ $BREW_CASK_LIST =~ (?!flux) ]];then
  echo "starting install flux..."
  brew cask install flux
  echo "install finished!"
elif [[ $BREW_CASK_LIST =~ (?!iterm2) ]];then
  echo "starting install iterm2"
  brew cask install iterm2
  echo "install finished!"
elif [[ -f /Application/Google\ Chrome.app && $BREW_CASK_LIST =~ (?!google-chrome) ]];then
  echo "starting install google-chrome"
  brew cask install google-chrome
  echo "install finished!"
elif [[ $BREW_CASK_LIST =~ (?!virtualbox) ]];then
  echo "starting install virtualbox"
  brew cask install virtualbox
  echo "install finished!"
fi

# zshのインストール、設定
echo "installing zsh unless exist..."
if [[ $SHELL =~ (?!zsh)$ ]]; then
  brew install --without-etcdir zsh
  brew install zsh-completions
fi

# 設定ファイルをasmsuechanのリポジトリからダウンロードして設置
echo "putting zsh config files unless exist..."
if [ ! -f ~/.zprofile ];then
  curl https://raw.githubusercontent.com/asmsuechan/dotfiles/master/.zprofile > ~/.zprofile
elif [ ! -f ~/.zshrc ];then
  curl https://raw.githubusercontent.com/asmsuechan/dotfiles/master/.zshrc > ~/.zshrc
fi

# ログインシェルをzshに変更
sudo dscl . -create /Users/$USER UserShell `which zsh`

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
# elcapitanにmysql5.7インストールすると「「「ヤバい」」」が構わず5.7(edge)をインストールする
brew install tig
brew install rbenv
brew install mysql
brew install tree
brew install docker
brew install docker-machine
brew install nginx
brew install vagrant

# rbenvを使ってrubyのインストール
rbenv install 2.2.3
rbenv global 2.2.3
rbnev rehash

gem install rails

# 最後にインストールチェック
# ruby rails virtualbox iterm2 zsh
# あとで書く
