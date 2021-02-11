package cad

//Based on https://github.com/supereggbert/proctree.js/blob/master/proctree.js

import (
	"math"
)

//Tree is a procedurally generated tree.
type Tree struct {
	root *branch

	properties TreeProperties

	mesh, twig struct {
		verts   [][3]float64
		faces   [][3]int
		normals [][3]float64
		uvs     [][2]float64
	}
}

type Leaves Tree

func (leaves Leaves) Indicies() []uint32 {
	tree := Tree(leaves)
	var retArray = make([]uint32, 0, len(tree.twig.faces)*3)
	for _, face := range tree.twig.faces {
		for _, index := range face {
			retArray = append(retArray, uint32(index))
		}
	}
	return retArray
}

func (leaves Leaves) UVs() []float32 {
	tree := Tree(leaves)
	var retArray = make([]float32, 0, len(tree.twig.uvs)*3)
	for _, uv := range tree.twig.uvs {
		for _, v := range uv {
			retArray = append(retArray, float32(v))
		}
	}
	return retArray
}

func (leaves Leaves) Vertices() []float32 {
	tree := Tree(leaves)
	var retArray = make([]float32, 0, len(tree.twig.verts)*3)
	for _, triangle := range tree.twig.verts {
		for _, vertex := range triangle {
			retArray = append(retArray, float32(vertex))
		}
	}
	return retArray
}

func (leaves Leaves) Normals() []float32 {
	tree := Tree(leaves)
	var retArray = make([]float32, 0, len(tree.twig.normals)*3)
	for _, vector := range tree.twig.normals {
		for _, v := range vector {
			retArray = append(retArray, float32(v))
		}
	}
	return retArray
}

func DefaultTree() TreeProperties {
	return TreeProperties{
		Seed: 10,

		ClumpMax:            0.8,
		ClumpMin:            0.5,
		LengthFalloffFactor: 0.85,
		LengthFalloffPower:  1,
		BranchFactor:        2.0,
		RadiusFalloffRate:   0.6,
		ClimbRate:           1.5,
		TrunkKink:           0.00,
		MaxRadius:           0.25,
		TreeSteps:           2,
		TaperRate:           0.95,
		TwistRate:           13,
		Segments:            6,
		Levels:              3,
		SweepAmount:         0,
		InitalBranchLength:  0.85,
		TrunkLength:         2.5,
		DropAmount:          0.0,
		GrowAmount:          0.0,
		VMultiplier:         0.2,
		TwigScale:           2.0,
	}
}

//TreeProperties when generating the tree, check out https://gltf-trees.donmccurdy.com/ as a reference.
type TreeProperties struct {
	Seed float64

	ClumpMin, ClumpMax float64

	LengthFalloffFactor float64
	LengthFalloffPower  float64

	BranchFactor float64

	RadiusFalloffRate float64
	ClimbRate         float64

	TrunkKink float64
	MaxRadius float64

	TreeSteps float64
	TaperRate float64

	TwistRate float64

	Segments int

	Levels float64

	SweepAmount float64

	InitalBranchLength float64

	TrunkLength float64

	DropAmount float64
	GrowAmount float64

	//UV Multiplier
	VMultiplier float64
	TwigScale   float64

	random func(float64) float64
}

func NewTree(properties TreeProperties) Tree {
	var tree Tree

	properties.random = func(a float64) float64 {
		return math.Abs(math.Cos(a + a*a))
	}

	tree.root = newBranch([3]float64{0, properties.TrunkLength, 0}, nil)
	tree.root.length = properties.InitalBranchLength

	tree.properties = properties
	tree.root.split(properties.Levels, properties.TreeSteps, properties, 0, 0)

	tree.createForks(tree.root, properties.MaxRadius)
	tree.createTwigs(tree.root)
	tree.doFaces(tree.root)
	tree.calcNormals()

	return tree
}

func (tree Tree) Leaves() Leaves {
	return Leaves(tree)
}

func (tree Tree) Indicies() []uint32 {
	var retArray = make([]uint32, 0, len(tree.mesh.faces)*3)
	for _, face := range tree.mesh.faces {
		for _, index := range face {
			retArray = append(retArray, uint32(index))
		}
	}
	return retArray
}

func (tree Tree) UVs() []float32 {
	var retArray = make([]float32, 0, len(tree.mesh.uvs)*3)
	for _, uv := range tree.mesh.uvs {
		for _, v := range uv {
			retArray = append(retArray, float32(v))
		}
	}
	return retArray
}

func (tree Tree) Vertices() []float32 {
	var retArray = make([]float32, 0, len(tree.mesh.verts)*3)
	for _, triangle := range tree.mesh.verts {
		for _, vertex := range triangle {
			retArray = append(retArray, float32(vertex))
		}
	}
	return retArray
}

func (tree Tree) Normals() []float32 {
	var retArray = make([]float32, 0, len(tree.mesh.normals)*3)
	for _, vector := range tree.mesh.normals {
		for _, v := range vector {
			retArray = append(retArray, float32(v))
		}
	}
	return retArray
}

func (tree *Tree) calcNormals() {
	var allNormals = make([][][3]float64, len(tree.mesh.verts))

	for _, face := range tree.mesh.faces {
		var norm = normalize(
			cross(
				subVec(tree.mesh.verts[face[1]], tree.mesh.verts[face[2]]),
				subVec(tree.mesh.verts[face[1]], tree.mesh.verts[face[0]]),
			),
		)
		allNormals[face[0]] = append(allNormals[face[0]], norm)
		allNormals[face[1]] = append(allNormals[face[1]], norm)
		allNormals[face[2]] = append(allNormals[face[2]], norm)
	}

	tree.mesh.normals = make([][3]float64, len(allNormals))

	for i := range allNormals {
		var total = [3]float64{0, 0, 0}
		var l = len(allNormals[i])
		for j := 0; j < l; j++ {
			total = addVec(total, scaleVec(allNormals[i][j], 1/float64(l)))
		}
		tree.mesh.normals[i] = total
	}
}

func (tree *Tree) doFaces(branch *branch) {
	var segments = tree.properties.Segments

	if branch.parent == nil {
		tree.mesh.uvs = make([][2]float64, len(tree.mesh.verts))

		var tangent = normalize(
			cross(
				subVec(branch.child[0].head, branch.head),
				subVec(branch.child[1].head, branch.head),
			),
		)

		var normal = normalize(branch.head)
		var angle = math.Acos(dot(tangent, [3]float64{-1, 0, 0}))
		if dot(cross([3]float64{-1, 0, 0}, tangent), normal) > 0 {
			angle = 2*math.Pi - angle
		}
		var segOffset = int(math.Round(angle / math.Pi / 2 * float64(segments)))

		for i := 0; i < segments; i++ {
			var v1 = branch.ring[0][i]
			var v2 = branch.root[(i+segOffset+1)%segments]
			var v3 = branch.root[(i+segOffset)%segments]
			var v4 = branch.ring[0][(i+1)%segments]

			tree.mesh.faces = append(tree.mesh.faces, [3]int{v1, v4, v3})
			tree.mesh.faces = append(tree.mesh.faces, [3]int{v1, v4, v3})
			tree.mesh.faces = append(tree.mesh.faces, [3]int{v4, v2, v3})

			tree.mesh.uvs[(i+segOffset)%segments] = [2]float64{math.Abs(float64(i)/float64(segments)-0.5) * 2, 0}
			var l = length(subVec(tree.mesh.verts[branch.ring[0][i]], tree.mesh.verts[branch.root[(i+segOffset)%segments]])) * tree.properties.VMultiplier

			tree.mesh.uvs[branch.ring[0][i]] = [2]float64{math.Abs(float64(i)/float64(segments)-0.5) * 2, l}
			tree.mesh.uvs[branch.ring[2][i]] = [2]float64{math.Abs(float64(i)/float64(segments)-0.5) * 2, l}
		}
	}

	if branch.child[0].ring[0] != nil {
		var segOffset0, segOffset1 int
		var match0, match1 float64

		var first0, first1 bool = true, true

		var v1 = normalize(subVec(tree.mesh.verts[branch.ring[1][0]], branch.head))
		var v2 = normalize(subVec(tree.mesh.verts[branch.ring[2][0]], branch.head))

		v1 = scaleInDirection(v1, normalize(subVec(branch.child[0].head, branch.head)), 0)
		v2 = scaleInDirection(v2, normalize(subVec(branch.child[1].head, branch.head)), 0)

		for i := 0; i < segments; i++ {
			var d = normalize(subVec(tree.mesh.verts[branch.child[0].ring[0][i]], branch.child[0].head))
			var l = dot(d, v1)

			if first0 || l > match0 {
				match0 = l
				segOffset0 = segments - i
				first0 = false

			}
			d = normalize(subVec(tree.mesh.verts[branch.child[1].ring[0][i]], branch.child[1].head))
			l = dot(d, v2)
			if first1 || l > match1 {
				match1 = l
				segOffset1 = segments - i
				first1 = false
			}
		}

		var UVScale = tree.properties.MaxRadius / branch.radius

		for i := 0; i < segments; i++ {

			v1 := branch.child[0].ring[0][i]
			v2 := branch.ring[1][(i+segOffset0+1)%segments]
			v3 := branch.ring[1][(i+segOffset0)%segments]
			v4 := branch.child[0].ring[0][(i+1)%segments]
			tree.mesh.faces = append(tree.mesh.faces, [3]int{v1, v4, v3})
			tree.mesh.faces = append(tree.mesh.faces, [3]int{v4, v2, v3})
			v1 = branch.child[1].ring[0][i]
			v2 = branch.ring[2][(i+segOffset1+1)%segments]
			v3 = branch.ring[2][(i+segOffset1)%segments]
			v4 = branch.child[1].ring[0][(i+1)%segments]

			tree.mesh.faces = append(tree.mesh.faces, [3]int{v1, v2, v3})
			tree.mesh.faces = append(tree.mesh.faces, [3]int{v1, v4, v2})

			var len1 = length(subVec(tree.mesh.verts[branch.child[0].ring[0][i]], tree.mesh.verts[branch.ring[1][(i+segOffset0)%segments]])) * UVScale
			var uv1 = tree.mesh.uvs[branch.ring[1][(i+segOffset0-1)%segments]]

			tree.mesh.uvs[branch.child[0].ring[0][i]] = [2]float64{uv1[0], uv1[1] + len1*tree.properties.VMultiplier}
			tree.mesh.uvs[branch.child[0].ring[2][i]] = [2]float64{uv1[0], uv1[1] + len1*tree.properties.VMultiplier}

			var len2 = length(subVec(tree.mesh.verts[branch.child[1].ring[0][i]], tree.mesh.verts[branch.ring[2][(i+segOffset1)%segments]])) * UVScale
			var uv2 = tree.mesh.uvs[branch.ring[2][(i+segOffset1-1)%segments]]

			tree.mesh.uvs[branch.child[1].ring[0][i]] = [2]float64{uv2[0], uv2[1] + len2*tree.properties.VMultiplier}
			tree.mesh.uvs[branch.child[1].ring[2][i]] = [2]float64{uv2[0], uv2[1] + len2*tree.properties.VMultiplier}
		}

		tree.doFaces(branch.child[0])
		tree.doFaces(branch.child[1])
	} else {
		for i := 0; i < segments; i++ {
			tree.mesh.faces = append(tree.mesh.faces, [3]int{branch.child[0].end, branch.ring[1][(i+1)%segments], branch.ring[1][i]})
			tree.mesh.faces = append(tree.mesh.faces, [3]int{branch.child[1].end, branch.ring[2][(i+1)%segments], branch.ring[2][i]})

			var len = length(subVec(tree.mesh.verts[branch.child[0].end], tree.mesh.verts[branch.ring[1][i]]))
			tree.mesh.uvs[branch.child[0].end] = [2]float64{math.Abs(float64(i)/float64(segments)-1-0.5) * 2, len * tree.properties.VMultiplier}
			len = length(subVec(tree.mesh.verts[branch.child[1].end], tree.mesh.verts[branch.ring[2][i]]))
			tree.mesh.uvs[branch.child[1].end] = [2]float64{math.Abs(float64(i)/float64(segments)-0.5) * 2, len * tree.properties.VMultiplier}
		}
	}
}

func (tree *Tree) createTwigs(branch *branch) {
	if branch.child[0] == nil {
		var tangent = normalize(
			cross(
				subVec(branch.parent.child[0].head, branch.parent.head),
				subVec(branch.parent.child[1].head, branch.parent.head),
			),
		)
		var binormal = normalize(subVec(branch.head, branch.parent.head))
		var normal = cross(tangent, binormal)

		//This can probably be factored into a loop.
		var vert1 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, tree.properties.TwigScale)),
				scaleVec(binormal, tree.properties.TwigScale*2-branch.length),
			),
		)

		var vert2 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, -tree.properties.TwigScale)),
				scaleVec(binormal, tree.properties.TwigScale*2-branch.length),
			),
		)

		var vert3 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, -tree.properties.TwigScale)),
				scaleVec(binormal, -branch.length),
			),
		)

		var vert4 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, tree.properties.TwigScale)),
				scaleVec(binormal, -branch.length),
			),
		)

		var vert8 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, tree.properties.TwigScale)),
				scaleVec(binormal, tree.properties.TwigScale*2-branch.length),
			),
		)

		var vert7 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, -tree.properties.TwigScale)),
				scaleVec(binormal, tree.properties.TwigScale*2-branch.length),
			),
		)

		var vert6 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, -tree.properties.TwigScale)),
				scaleVec(binormal, -branch.length),
			),
		)

		var vert5 = len(tree.twig.verts)
		tree.twig.verts = append(tree.twig.verts,
			addVec(
				addVec(branch.head, scaleVec(tangent, tree.properties.TwigScale)),
				scaleVec(binormal, -branch.length),
			),
		)

		tree.twig.faces = append(tree.twig.faces, [3]int{vert1, vert2, vert3})
		tree.twig.faces = append(tree.twig.faces, [3]int{vert4, vert1, vert3})

		tree.twig.faces = append(tree.twig.faces, [3]int{vert6, vert7, vert8})
		tree.twig.faces = append(tree.twig.faces, [3]int{vert6, vert8, vert5})

		normal = normalize(
			cross(
				subVec(tree.twig.verts[vert1], tree.twig.verts[vert3]),
				subVec(tree.twig.verts[vert2], tree.twig.verts[vert3]),
			),
		)

		var normal2 = normalize(
			cross(
				subVec(tree.twig.verts[vert7], tree.twig.verts[vert6]),
				subVec(tree.twig.verts[vert8], tree.twig.verts[vert6]),
			),
		)

		tree.twig.normals = append(tree.twig.normals, normal)
		tree.twig.normals = append(tree.twig.normals, normal)
		tree.twig.normals = append(tree.twig.normals, normal)
		tree.twig.normals = append(tree.twig.normals, normal)

		tree.twig.normals = append(tree.twig.normals, normal2)
		tree.twig.normals = append(tree.twig.normals, normal2)
		tree.twig.normals = append(tree.twig.normals, normal2)
		tree.twig.normals = append(tree.twig.normals, normal2)

		tree.twig.uvs = append(tree.twig.uvs, [2]float64{0, 1})
		tree.twig.uvs = append(tree.twig.uvs, [2]float64{1, 1})
		tree.twig.uvs = append(tree.twig.uvs, [2]float64{1, 0})
		tree.twig.uvs = append(tree.twig.uvs, [2]float64{0, 0})

		tree.twig.uvs = append(tree.twig.uvs, [2]float64{0, 1})
		tree.twig.uvs = append(tree.twig.uvs, [2]float64{1, 1})
		tree.twig.uvs = append(tree.twig.uvs, [2]float64{1, 0})
		tree.twig.uvs = append(tree.twig.uvs, [2]float64{0, 0})
	} else {
		tree.createTwigs(branch.child[0])
		tree.createTwigs(branch.child[1])
	}
}

func (tree *Tree) createForks(branch *branch, radius float64) {
	branch.radius = radius

	if radius > branch.length {
		radius = branch.length
	}

	var segments = tree.properties.Segments

	var segmentAngle = math.Pi * 2 / float64(segments)

	if branch.parent == nil {
		//create the root of the tree
		var axis = [3]float64{0, 1, 0}
		for i := 0; i < segments; i++ {
			var vec = vecAxisAngle([3]float64{-1, 0, 0}, axis, -segmentAngle*float64(i))
			branch.root = append(branch.root, len(tree.mesh.verts))

			tree.mesh.verts = append(tree.mesh.verts, scaleVec(vec, radius/tree.properties.RadiusFalloffRate))
		}
	}

	//cross the branches to get the left
	//add the branches to get the up
	if branch.child[0] != nil {
		var axis [3]float64
		if branch.parent != nil {
			axis = normalize(subVec(branch.head, branch.parent.head))
		} else {
			axis = normalize(branch.head)
		}

		var axis1 = normalize(subVec(branch.head, branch.child[0].head))
		var axis2 = normalize(subVec(branch.head, branch.child[1].head))
		var tangent = normalize(cross(axis1, axis2))
		branch.tangent = tangent

		var axis3 = normalize(cross(tangent, normalize(addVec(scaleVec(axis1, -1), scaleVec(axis2, -1)))))
		var dir = [3]float64{axis2[0], 0, axis2[2]}
		var centerloc = addVec(branch.head, scaleVec(dir, -tree.properties.MaxRadius/2))

		var scale = tree.properties.RadiusFalloffRate

		if branch.child[0].trunk || branch.trunk {
			scale = 1 / tree.properties.TaperRate
		}

		//main segment ring
		var linch0 = len(tree.mesh.verts)

		branch.ring[0] = append(branch.ring[0], linch0)
		branch.ring[2] = append(branch.ring[2], linch0)
		tree.mesh.verts = append(tree.mesh.verts,
			addVec(centerloc, scaleVec(tangent, radius*scale)))

		var start = len(tree.mesh.verts) - 1
		var d1 = vecAxisAngle(tangent, axis2, 1.57)
		var d2 = normalize(cross(tangent, axis))
		var s = 1 / dot(d1, d2)
		for i := 1; i < segments/2; i++ {
			var vec = vecAxisAngle(tangent, axis2, segmentAngle*float64(i))
			branch.ring[0] = append(branch.ring[0], start+i)
			branch.ring[2] = append(branch.ring[2], start+i)
			vec = scaleInDirection(vec, d2, s)
			tree.mesh.verts = append(tree.mesh.verts,
				addVec(centerloc, scaleVec(vec, radius*scale)))
		}

		var linch1 = len(tree.mesh.verts)
		branch.ring[0] = append(branch.ring[0], linch1)
		branch.ring[1] = append(branch.ring[1], linch1)
		tree.mesh.verts = append(tree.mesh.verts,
			addVec(centerloc, scaleVec(tangent, -radius*scale)))
		for i := segments/2 + 1; i < segments; i++ {
			var vec = vecAxisAngle(tangent, axis1, segmentAngle*float64(i))
			branch.ring[0] = append(branch.ring[0], len(tree.mesh.verts))
			branch.ring[1] = append(branch.ring[1], len(tree.mesh.verts))

			tree.mesh.verts = append(tree.mesh.verts,
				addVec(centerloc, scaleVec(vec, radius*scale)))
		}

		branch.ring[1] = append(branch.ring[1], linch0)
		branch.ring[2] = append(branch.ring[2], linch1)

		start = len(tree.mesh.verts) - 1
		for i := 1; i < segments/2; i++ {
			var vec = vecAxisAngle(tangent, axis3, segmentAngle*float64(i))

			branch.ring[1] = append(branch.ring[1], start+i)
			branch.ring[2] = append(branch.ring[2], start+(segments/2-i))

			var v = scaleVec(vec, radius*scale)

			tree.mesh.verts = append(tree.mesh.verts,
				addVec(centerloc, v))
		}

		//child radius is related to the brans direction and the length of the branch
		//var length0 = length(subVec(branch.head, branch.child[0].head))
		//var length1 = length(subVec(branch.head, branch.child[1].head))

		var radius0 = 1 * radius * tree.properties.RadiusFalloffRate
		var radius1 = 1 * radius * tree.properties.RadiusFalloffRate
		if branch.trunk {
			radius0 = radius * tree.properties.TaperRate
		}

		tree.createForks(branch.child[0], radius0)
		tree.createForks(branch.child[1], radius1)
	} else {
		//add points for the ends of braches
		branch.end = len(tree.mesh.verts)
		//branch.head=addVec(branch.head,scaleVec([this.properties.xBias,this.properties.yBias,this.properties.zBias],branch.length*3));

		tree.mesh.verts = append(tree.mesh.verts,
			branch.head,
		)
	}
}

type branch struct {
	head, tangent  [3]float64
	length, radius float64

	end  int
	root []int
	ring [3][]int

	//Is this branch the main trunk of the tree?
	trunk bool

	child  [2]*branch
	parent *branch
}

func newBranch(head [3]float64, parent *branch) *branch {
	return &branch{
		head:   head,
		parent: parent,
		length: 1,
	}
}

func (b *branch) mirrorBranch(vec, norm [3]float64, properties TreeProperties) [3]float64 {
	var v = cross(norm, cross(vec, norm))
	var s = properties.BranchFactor * dot(v, vec)
	return [3]float64{vec[0] - v[0]*s, vec[1] - v[1]*s, vec[2] - v[2]*s}
}

func (bra *branch) split(level float64, steps float64, properties TreeProperties, l1, l2 float64) {
	if l1 == 0 {
		l1 = 1
	}
	if l2 == 0 {
		l2 = 1
	}

	var rLevel = properties.Levels - level
	var po [3]float64
	if bra.parent != nil {
		po = bra.parent.head
	} else {
		bra.trunk = true
	}
	var so = bra.head
	var dir = normalize(subVec(so, po))

	var normal = cross(dir, [3]float64{dir[2], dir[0], dir[1]})
	var tangent = cross(dir, normal)
	var r = properties.random(rLevel*10 + l1*5 + l2 + properties.Seed)
	//var r2 = properties.random(rLevel*10 + l1*5 + l2 + 1 + properties.Seed)
	var clumpmax = properties.ClumpMax
	var clumpmin = properties.ClumpMin

	var adj = addVec(scaleVec(normal, r), scaleVec(tangent, 1-r))
	if r > 0.5 {
		adj = scaleVec(adj, -1)
	}

	var clump = (clumpmax-clumpmin)*r + clumpmin
	var newdir = normalize(addVec(scaleVec(adj, 1-clump), scaleVec(dir, clump)))

	var newdir2 = bra.mirrorBranch(newdir, dir, properties)
	if r > 0.5 {
		var tmp = newdir
		newdir = newdir2
		newdir2 = tmp
	}
	if steps > 0 {
		var angle = steps / properties.TreeSteps * 2 * math.Pi * properties.TwistRate
		newdir2 = normalize([3]float64{math.Sin(angle), r, math.Cos(angle)})
	}

	var growAmount = level * level / (properties.Levels * properties.Levels) * properties.GrowAmount
	var dropAmount = rLevel * properties.DropAmount
	var sweepAmount = rLevel * properties.SweepAmount
	newdir = normalize(addVec(newdir, [3]float64{sweepAmount, dropAmount + growAmount, 0}))
	newdir2 = normalize(addVec(newdir2, [3]float64{sweepAmount, dropAmount + growAmount, 0}))

	var head0 = addVec(so, scaleVec(newdir, bra.length))
	var head1 = addVec(so, scaleVec(newdir2, bra.length))
	bra.child[0] = newBranch(head0, bra)
	bra.child[1] = newBranch(head1, bra)
	bra.child[0].length = math.Pow(bra.length, properties.LengthFalloffPower) * properties.LengthFalloffFactor
	bra.child[1].length = math.Pow(bra.length, properties.LengthFalloffPower) * properties.LengthFalloffFactor
	if level > 0 {
		if steps > 0 {
			bra.child[0].head = addVec(bra.head,
				[3]float64{(r - 0.5) * 2 * properties.TrunkKink, properties.ClimbRate, (r - 0.5) * 2 * properties.TrunkKink})
			bra.child[0].trunk = true
			bra.child[0].length = bra.length * properties.TaperRate
			bra.child[0].split(level, steps-1, properties, l1+1, l2)
		} else {
			bra.child[0].split(level-1, 0, properties, l1+1, l2)
		}
		bra.child[1].split(level-1, 0, properties, l1, l2+1)
	}
}
