n = gets.strip.to_i

n.times do |i|
  space = n - i -1
  sharp = n - space
  print " " * space 
  puts "#" * sharp
end
