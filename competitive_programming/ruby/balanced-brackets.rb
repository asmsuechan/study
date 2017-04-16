n = gets.to_i

correspond_table = { '{': '}', '[': ']', '(': ')' }
left_brackets = ['{', '[', '(']

results = []

n.times do |i|
  open_stack = []
  brackets = gets.split('')
  flag = true
  brackets.each do |bracket|
    if left_brackets.include?(bracket)
      open_stack.push(bracket)
    else
      val = open_stack.pop
      if val && bracket != correspond_table[val.to_sym]
        results.push('NO')
        flag = false
        break
      end
    end
  end
  results.push('YES') if flag
end

puts results
