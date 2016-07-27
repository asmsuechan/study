class Response
  attr_accessor :request, :headers, :response_except_body, :status, :body
  def initialize(response)
    @response = response
    @body = @response.slice!(/(\n\r\n).+/m)
    @response_except_body = @response.slice(/.+/m)
    #なぜsliceしないと動かないのだろうか・・・・
    @request = @response.slice!(/.+/)
    @status = @request.slice!(/\d{3}.+/)

    #ヘッダーの切り出し
    @headers = []
    while @response != ""
      header = @response.slice!(/.+/)
      @headers.push(header) unless header.nil?
      @response.sub!(/(\r\n|\r|\n)/, '')
    end

    #切り出したヘッダーからインスタンス変数、
    #アクセサメソッドを作成
    @headers.each do |header|
      header_name_raw = header.slice!(/^.+: /)
      unless header_name_raw.nil?
        header_name = header_name_raw.downcase
                      .gsub("-", "_")
                      .gsub(":", "")
                      .gsub(" ", "")
        self.instance_variable_set("@#{header_name}",header)
        add_accessor header_name
      end
    end
  end

  def all_response
    self.response_except_body
  end

  def disp_body
    self.body
  end

  private

    def add_accessor instance_variable_name
      self.class.send(:attr_accessor, instance_variable_name)
    end

end
