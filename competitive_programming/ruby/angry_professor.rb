def count_attendance_on_time(a)
  students = 0
  a.each do |time|
    if time <= 0
      students += 1
    end
  end
  students
end

student_attendances = []
ks = []

t = gets.strip.to_i
for a0 in (0..t-1)
  n,k = gets.strip.split(' ')
  n = n.to_i
  k = k.to_i
  a = gets.strip
  a = a.split(' ').map(&:to_i)
  student_attendances.push(a)
  ks.push(k)
end

student_attendances.each_with_index do |attendance, i|
  if count_attendance_on_time(attendance) < ks[i]
    puts "YES"
  else
    puts "NO"
  end
end
