package main

import "github.com/solarlune/resolv"

var space *resolv.Space
var playerObj *resolv.Object

func init() {

	space = resolv.NewSpace(1000, 600, 16, 16)

	space.Add( //x y Width height
		resolv.NewObject(93.9394, 3.0303, 2775.76, 87.8788),
		resolv.NewObject(9.09091, 100, 84.8485, 1815.15),
		resolv.NewObject(1836.36, 109.091, 2769.7, 84.8485),
		resolv.NewObject(2790.91, 6.606061, 87.8788, 1915.15),
		resolv.NewObject(114, 737, 487, 25),
		resolv.NewObject(112, 550, 488, 24),
		resolv.NewObject(111, 354, 480, 24),
		resolv.NewObject(100, 102, 572, 84),
		resolv.NewObject(193, 602, 68, 573),
		resolv.NewObject(601, 769, 70, 188),
		resolv.NewObject(52, 75, 57, 880),
		resolv.NewObject(53, 958, 170, 864),
		resolv.NewObject(356, 1251, 124, 580),
		resolv.NewObject(1173, 1349, 76, 383,),
		resolv.NewObject(598, 1537, 155, 12),
		resolv.NewObject(980,876, 330, 179),
		resolv.NewObject(675, 717, 161, 339),
		resolv.NewObject(870, 1234, 286, 12),
		resolv.NewObject(835, 672, 33, 573),
		resolv.NewObject(1443, 1131, 476, 34),
		resolv.NewObject(1827, 1169, 90, 367),
		resolv.NewObject(1444, 675, 109, 181),
		resolv.NewObject(838, 677, 73, 178),
		resolv.NewObject(1812, 1536, 26, 289,),
		resolv.NewObject(2433, 865, 59, 408),
		resolv.NewObject(2305, 577, 542, 278),
		resolv.NewObject(1726, 388, 37, 284),
		resolv.NewObject(1539, 96, 199, 94),
		resolv.NewObject(1726, 197, 18, 194),
		resolv.NewObject(1728, 6, 513, 78),
		resolv.NewObject(677, 6, 282, 185),
		resolv.NewObject(1904, 868, 16, 264),
		resolv.NewObject(1712, 856, 82, 5),
		resolv.NewObject(1827, 853, 72, 9),
		resolv.NewObject(1000, 769, 345, 38),

	)
	playerObj = resolv.NewObject(32, 32, 16, 16)

	space.Add(playerObj)

}

func upt() {
	dx := 2.0

	if collision := playerObj.Check(dx, 0); collision != nil {

	dx = collision.ContactWithObject(collision.Objects[0]).X()
	}

	playerObj.X += dx
	playerObj.Update()

	


}