package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	chk "gopkg.in/check.v1"
)

type Day7TestSuite struct{}
var _ = chk.Suite(&Day7TestSuite{})

func (s *Day7TestSuite) TestPassPart1(c *chk.C) {
	solution := &solutions.Day7Solution{}

	solution.Prepare(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	c.Assert(solution.Part1(), chk.Equals, "4")
}

func (s *Day7TestSuite) TestPassPart2(c *chk.C) {
	solution := &solutions.Day7Solution{}

	solution.Prepare(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	c.Assert(solution.Part2(), chk.Equals, "32")

	solution = &solutions.Day7Solution{}

	solution.Prepare(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`)

	c.Assert(solution.Part2(), chk.Equals, "126")
}