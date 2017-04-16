n = gets

stack = []
result = []

n.to_i.times do |i|
  opecode = gets.split.map(&:to_i)
  if opecode[0] == 1
    stack.push(opecode[1])
  elsif opecode[0] == 2
    stack.pop
  elsif opecode[0] == 3
    result.push(stack.max)
  end
end

puts result
