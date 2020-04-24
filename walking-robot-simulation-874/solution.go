package main

type Point struct {
	x int
	y int
}

type Robot struct {
	Point
	// 0 表示北，1 表示东，2 表示南，3表示北
	direction int
}

func robotSim(commands []int, obstacles [][]int) int {
	mapping := make(map[Point]struct{}, len(obstacles))
	for i := 0; i < len(obstacles); i++ {
		p := Point{
			x: obstacles[i][0],
			y: obstacles[i][1],
		}
		mapping[p] = struct{}{}
	}

	robot := Robot{}
	for _, x := range commands {
		if x == -1 {
			robot.direction = (robot.direction + 1) % 4
		} else if x == -2 {
			robot.direction = (robot.direction + 3) % 4
		} else {
			if robot.direction == 0 {
				for i := 1; i <= x; i++ {
					robot.x++
					if _, ok := mapping[robot.Point]; ok {
						robot.x--
						continue
					}
				}
			} else if robot.direction == 1 {
				for i := 1; i <= x; i++ {
					robot.y++
					if _, ok := mapping[robot.Point]; ok {
						robot.y++
						continue
					}
				}

			} else if robot.direction == 2 {
				for i := 1; i <= x; i++ {
					robot.x--
					if _, ok := mapping[robot.Point]; ok {
						robot.x--
						continue
					}
				}
			} else if robot.direction == 3 {
				for i := 1; i <= x; i++ {
					robot.y--
					if _, ok := mapping[robot.Point]; ok {
						robot.y--
						continue
					}
				}
			}
		}
	}
	return robot.x*robot.x + robot.y*robot.y
}
