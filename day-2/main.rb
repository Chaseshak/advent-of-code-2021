puts "-----\nPart 1"

depth = 0
horizontal_pos = 0

instructions = File.readlines('input.txt').map(&:split)

instructions.each do |instruction|
  command, num = instruction
  num = num.to_i

  case command
  when 'forward'
    horizontal_pos += num
  when 'down'
    depth += num
  when 'up'
    depth -= num
  end
end

puts "Horizontal Position: #{horizontal_pos}"
puts "Depth: #{depth}"
puts "Depth * Horizontal Position = #{depth * horizontal_pos}"
puts "-----\nPart Two"

horizontal_pos = 0
depth = 0
aim = 0

instructions.each do |instruction|
  command, num = instruction
  num = num.to_i

  case command
  when 'forward'
    horizontal_pos += num
    depth += aim * num
  when 'down'
    aim += num
  when 'up'
    aim -= num
  end
end

puts "Aim: #{aim}"
puts "Horizontal Position: #{horizontal_pos}"
puts "Depth: #{depth}"
puts "Depth * Horizontal Position = #{depth * horizontal_pos}"
