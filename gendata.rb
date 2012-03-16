#!/usr/bin/env ruby

if ARGV[0] == nil
	puts "Please enter a output filename"
	Process.exit
end

op_file = ARGV[0]

File.open(op_file, "w") { |f|
	for c in 1..1000000
		f.write(rand(10000))
		f.write("\n")
	end
}

