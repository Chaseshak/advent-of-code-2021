depths = File.readlines('depths.txt')
depths = depths.map(&:to_i)

increased_count = 0
depths.each_with_index do |depth, i|
  previous_depth = depths[i - 1]
  next if previous_depth.nil?

  increased_count += 1 if depth > previous_depth
end

puts "------\nPart 1 - Simple Increases"
puts "Answer: #{increased_count}"

puts "------\nPart 2 - Sliding Window Increases"

increased_window_count = 0

sliding_windows = depths.map.with_index { |_depth, i| depths[i..(i + 2)] }.select { |window| window.size == 3 }
sliding_windows.each_with_index do |window, i|
  previous_window = sliding_windows[i - 1]
  next if previous_window.nil?

  increased_window_count += 1 if window.sum > previous_window.sum
end

puts "Answer: #{increased_window_count}"
