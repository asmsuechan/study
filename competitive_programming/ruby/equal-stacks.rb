numbers = gets.split('').map(&:to_i)
stacks = []

numbers.length.times do |i|
  values = gets.split('').map(&:to_i)
  stacks.push(create_stack(values))
end

def create_stack(arr)
  arr.reverse
end
