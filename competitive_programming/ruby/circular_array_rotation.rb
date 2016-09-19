first_line = gets
nkq = first_line.split(" ").map(&:to_i)
n = nkq[0]
k = nkq[1]
q = nkq[2]

second_line = gets
ai = second_line.split(" ").map(&:to_i)

# rotate ai k times
rotated_ai = ai.rotate(-k)

indexes = []
q.times do
  indexes.push gets
end

indexes.map(&:to_i).each do |index|
  puts rotated_ai[index]
end
