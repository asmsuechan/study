require 'ffi'

module Hello
  extend FFI::Library
  ffi_lib FFI::Library::LIBC
  attach_function :puts, [ :string ], :string
end

Hello.puts("Hello, World")
