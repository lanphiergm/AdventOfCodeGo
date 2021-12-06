read -p "Year: " year
read -p "Day: " day
read -p "Puzzle Name: " name

filename=`echo "year$year/day_${day}_$name.go" | awk '{print tolower($0)}' | sed "s/ /_/g"`
testfilename=`echo "year$year/day_${day}_${name}_test.go" | awk '{print tolower($0)}' | sed "s/ /_/g"`
classname=`echo "$name" | sed "s/ //g"`

sed "s/{year}/$year/g" puzzle.template | sed "s/{puzzlename}/$name/g" | sed "s/{classname}/$classname/g" > $filename
sed "s/{year}/$year/g" test.template | sed "s/{puzzlename}/$name/g" | sed "s/{classname}/$classname/g" | sed "s/{day}/$day/" > $testfilename
touch ../../data/year$year/day_${day}_puzzle_data.txt
touch ../../data/year$year/day_${day}_sample_data.txt

echo ""
echo "Add to main.go:"
echo "    fmt.Printf(\"$year Day $day Part 1: %v\\n\","
echo "        year$year.${classname}Part1(\"data/year$year/day_${day}_puzzle_data.txt\"))"
echo "    fmt.Printf(\"$year Day $day Part 2: %v\\n\","
echo "        year$year.${classname}Part2(\"data/year$year/day_${day}_puzzle_data.txt\"))"
