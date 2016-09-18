# 20160916現在のinstagramのスクレイピング
require 'mechanize'
require 'json'

class ScrapingInstagram
  BASE_URL = 'https://www.instagram.com'

  # instagramはcurl叩いたらわかるようにhtmlのbodyが返ってこずにjsonが返ってくる
  # このjsonは<script type="text/javascript">にいる。
  # window._shareData = {}の形のjsonで送られている。
  # これをhashにしている。
  # 以下のような形式のurlが含まれるarrayが返ってくる
  # "https://scontent-nrt1-1.cdninstagram.com/t51.2885-15/e35/14240825_175988942810352_2132301197_n.jpg?ig_cache_key=MTM0MDU5NDUwNTg3NDgzMDc0NQ%3D%3D.2"
  def self.picture_urls_from_tag(tag)
    page_response = Mechanize.new.get("#{BASE_URL}/explore/tags/#{tag}/")
                    .search('script').inner_text
                    .slice(/window\._sharedData\ =(.*)$/)
                    .gsub("window._sharedData\ = ", '')
                    .gsub(';', '')

    tag_contents = JSON.parse(page_response)["entry_data"]["TagPage"]
    # puts instagram_tag_content.length
    # => 1
    # パラメーターなしでただtagをurlで指定しただけだと1ページしか表示されない

    tag_contents.first["tag"]["media"]["nodes"]
    .map { |post| post["display_src"] }
    # ?以降のパラメーターいらなかったら以下のコメントを外してください
    #.map { |url| url.gsub(/\?(.*)$/, '') }
  end
end

puts ScrapingInstagram.picture_urls_from_tag('udon')
