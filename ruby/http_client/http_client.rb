require 'socket'
require './response'

class HttpClient
  attr_accessor :uri, :port, :path

  def initialize(options = {})
    @uri = options[:uri]
    @path = options[:path]
    @port = options[:port]
  end

  [:get, :post, :put, :delete, :head, :options].each do |method|
    define_method(method) do |content = nil, type = "text/plain"|
      @socket = TCPSocket.open(uri, port)
      @socket << "#{method.to_s.swapcase} #{self.path} HTTP/1.1\r\n"
      @socket << "Host: #{self.uri}\r\n"
      @socket << "Accept: */*"
      @socket << "Content-Type: #{type}; charset=utf-8\r\n"

      unless content.nil?
        @socket << "Content-Length: #{content.length}\r\n"
        @socket << "\r\n#{content}"
      end

      @socket << "Connection: close\r\n"
      @socket << "\r\n" 
      response = Response.new(@socket.read)
      return response
      @socket.close
    end
  end
end

#request = HttpClient.new(uri:"example.com", path:"/", port:80)
#request = HttpClient.new(uri:"my-http-response.herokuapp.com", path:"/", port:80)
request = HttpClient.new(uri:"http://challenge-your-limits.herokuapp.com", path:"/challenge_users", port:80)
puts request.post("asmsuechan").all_response
#puts "===================="
#puts request.get.all_response
#puts "===================="
#puts HttpClient.new(uri:"example.com", path:"/", port:80).post("こんにちは").all_response
#puts "===================="
#puts HttpClient.new(uri:"www.google.com", path:"/", port:80).post("こんにちは").all_response
#puts "===================="
#puts HttpClient.new(uri:"example.com", path:"/", port:80).options.all_response
#puts "===================="
