n = gets.strip.to_i
arr = gets.strip
arr = arr.split(' ').map(&:to_i)

n -= 1

while n>= 0
  print arr[n]
  print " " unless n==0
  n -= 1
end
