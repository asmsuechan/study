x1,v1,x2,v2 = gets.strip.split(' ')
x1 = x1.to_i
v1 = v1.to_i
x2 = x2.to_i
v2 = v2.to_i

gap = x2 - x1

if ((x2 + v2) - (x1 + v1)) < gap
  i = 0
  prev_gap = 0
  loop do
    current_gap = (x2 + v2 * i) - (x1 + v1 * i)
    if current_gap > 0
      prev_gap = current_gap
    elsif current_gap == 0
      puts "YES"
      break
    else
      puts "NO"
      break
    end
    i += 1
  end
else
  puts "NO"
end
