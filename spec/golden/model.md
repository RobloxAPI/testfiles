# Model format
The model format displays a tree of instances and their properties.

- **Output:** `model`
- **Directive:** `model`
- **Formats:**
	- `rbxl`
	- `rbxm`
	- `rbxlx`
	- `rbxmx`

## Remarks
The format uses the standard struct with Format, Output, and Data fields,
described [here](../README.md#user-content-common). The type of the Data field
is a [Root][Root] struct.

The format uses the primitive types described
[here](../README.md#user-content-primitive-value-types).

[bool]:   ../README.md#user-content-bool
[int]:    ../README.md#user-content-int
[uint]:   ../README.md#user-content-uint
[float]:  ../README.md#user-content-float
[bytes]:  ../README.md#user-content-bytes
[string]: ../README.md#user-content-string

## Types
Each subsection describes a value type.

### Root
Field     | Type
----------|-----
Metadata  | list of [MetadataPair][MetadataPair]
Instances | list of [Instance][Instance]

Elements of the Metadata field are ordered by their Key value. The value of each
Key is unique within the list.

[Root]: #user-content-root

### MetadataPair
Field     | Type
----------|-----
Key       | [string][string]
Value     | [string][string]

[MetadataPair]: #user-content-metadatapair

### Instance
Field      | Type
-----------|-----
ClassName  | [string][string]
IsService  | [bool][bool]
Reference  | [int][int]
Properties | list of [Property][Property]
Children   | list of [Instance][Instance]

The value of the Reference field is the index of the instance if the tree were
traversed top-down.

[Instance]: #user-content-instance

### Property
Field      | Type
-----------|-----
Name       | [string][string]
Type       | [string][string]
Value      | ...

The type of the Value field is dependent on the value of the Type field,
corresponding to the following types.

[Property]: #user-content-property

### String
See [string][string].

### BinaryString
See [bytes][bytes].

### ProtectedString
See [string][string].

### Content
See [string][string].

### Bool
See [bool][bool].

### Int
See [int][int].

### Float
See [float][float].

### Double
See [float][float].

### UDim
Field  | Type
-------|-----
Scale  | [float][float]
Offset | [int][int]

[UDim]: #user-content-udim

### UDim2
Field | Type
------|-----
X     | [UDim][UDim]
Y     | [UDim][UDim]

[UDim2]: #user-content-udim2

### Ray
Field     | Type
----------|-----
Origin    | [Vector3][Vector3]
Direction | [Vector3][Vector3]

[Ray]: #user-content-ray

### Faces
Field  | Type
-------|-----
Right  | [bool][bool]
Top    | [bool][bool]
Back   | [bool][bool]
Left   | [bool][bool]
Bottom | [bool][bool]
Front  | [bool][bool]

[Faces]: #user-content-faces

### Axes
Field | Type
------|-----
X     | [bool][bool]
Y     | [bool][bool]
Z     | [bool][bool]

[Axes]: #user-content-axes

### BrickColor
See [uint][uint].

[BrickColor]: #user-content-brickcolor

### Color3
Field | Type
------|-----
R     | [float][float]
G     | [float][float]
B     | [float][float]

[Color3]: #user-content-color3

### Vector2
Field | Type
------|-----
X     | [float][float]
Y     | [float][float]

[Vector2]: #user-content-vector2

### Vector3
Field | Type
------|-----
X     | [float][float]
Y     | [float][float]
Z     | [float][float]

[Vector3]: #user-content-vector3

### CFrame
Field    | Type
---------|-----
Position | [Vector3][Vector3]
Rotation | [Rotation][Rotation]

[CFrame]: #user-content-cframe

#### Rotation
Field | Type
------|-----
R00   | [float][float]
R01   | [float][float]
R02   | [float][float]
R10   | [float][float]
R11   | [float][float]
R12   | [float][float]
R20   | [float][float]
R21   | [float][float]
R22   | [float][float]

[Rotation]: #user-content-rotation

### Token
See [uint][uint].

[Token]: #user-content-token

### Reference
An [int][int] corresponding to the Reference field of an [Instance][Instance],
or `nil` indicating no reference.

[Reference]: #user-content-reference

### Vector3int16
Field | Type
------|-----
X     | [int][int]
Y     | [int][int]
Z     | [int][int]

[Vector3int16]: #user-content-vector3int16

### Vector2int16
Field | Type
------|-----
X     | [int][int]
Y     | [int][int]

[Vector2int16]: #user-content-vector2int16

### NumberSequence
A list of [NumberSequenceKeypoint][NumberSequenceKeypoint] structs.

[NumberSequence]: #user-content-numbersequence

#### NumberSequenceKeypoint
Field    | Type
---------|-----
Time     | [float][float]
Value    | [float][float]
Envelope | [float][float]

[NumberSequenceKeypoint]: #user-content-numbersequencekeypoint

### ColorSequence
A list of [ColorSequenceKeypoint][ColorSequenceKeypoint] structs.

[ColorSequence]: #user-content-colorsequence

#### ColorSequenceKeypoint
Field    | Type
---------|-----
Time     | [float][float]
Value    | [Color3][Color3]
Envelope | [float][float]

[ColorSequenceKeypoint]: #user-content-colorsequencekeypoint

### NumberRange
Field | Type
------|-----
Min   | [float][float]
Max   | [float][float]

[NumberRange]: #user-content-numberrange

### Rect
Field | Type
------|-----
Min   | [Vector2][Vector2]
Max   | [Vector2][Vector2]

[Rect]: #user-content-rect

### PhysicalProperties
Field            | Type
-----------------|-----
CustomPhysics    | [bool][bool]
Density          | [float][float]
Friction         | [float][float]
Elasticity       | [float][float]
FrictionWeight   | [float][float]
ElasticityWeight | [float][float]

If `CustomPhysics` is false, then the other fields are omitted.

[PhysicalProperties]: #user-content-physicalproperties

### Color3uint8
Field | Type
------|-----
R     | [uint][uint]
G     | [uint][uint]
B     | [uint][uint]

[Color3uint8]: #user-content-color3uint8

### Int64
See [int][int].

[Int64]: #user-content-int64

### SharedString
See [string][string].

[SharedString]: #user-content-sharedstring
