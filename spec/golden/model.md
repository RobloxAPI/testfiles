# Model format
The model format displays a tree of instances and their properties.

The format uses the primitive types described
[here](../README.md#user-content-primitive-value-types).

The root value is a [Root][Root] object.

[bool]:   ../README.md#user-content-bool
[int]:    ../README.md#user-content-int
[uint]:   ../README.md#user-content-uint
[float]:  ../README.md#user-content-float
[bytes]:  ../README.md#user-content-bytes
[string]: ../README.md#user-content-string

## Types
Each subsection describes a value type.

### Root
An object with the following fields:

Field     | Type
----------|-----
Metadata  | array of [MetadataPair][MetadataPair]
Instances | array of [Instance][Instance]

Values of the Metadata field are ordered by their Key field. The value of each
Key is unique within the array.

[Root]: #user-content-root

### MetadataPair
An object with the following fields:

Field     | Type
----------|-----
Key       | [string][string]
Value     | [string][string]

[MetadataPair]: #user-content-metadatapair

### Instance
An object with the following fields:

Field      | Type
-----------|-----
ClassName  | [string][string]
IsService  | [bool][bool]
Reference  | [int][int]
Properties | array of [Property][Property]
Children   | array of [Instance][Instance]

The value of the Reference field is the index of the instance if the tree were
traversed top-down.

Values of the Properties field are ordered by their Name field. The value of
each Name is unique within the array.

[Instance]: #user-content-instance

### Property
An object with the following fields:

Field      | Type
-----------|-----
Name       | [string][string]
Type       | [string][string]
Value      | ...

The type of the Value field is dependent on the value of the Type field,
corresponding to the following types.

[Property]: #user-content-property

### String
Uses [string][string].

### BinaryString
Uses [bytes][bytes].

### ProtectedString
Uses [string][string].

### Content
Uses [string][string].

### Bool
Uses [bool][bool].

### Int
Uses [int][int].

### Float
Uses [float][float].

### Double
Uses [float][float].

### UDim
An object with the following fields:

Field  | Type
-------|-----
Scale  | [float][float]
Offset | [int][int]

[UDim]: #user-content-udim

### UDim2
An object with the following fields:

Field | Type
------|-----
X     | [UDim][UDim]
Y     | [UDim][UDim]

[UDim2]: #user-content-udim2

### Ray
An object with the following fields:

Field     | Type
----------|-----
Origin    | [Vector3][Vector3]
Direction | [Vector3][Vector3]

[Ray]: #user-content-ray

### Faces
An object with the following fields:

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
An object with the following fields:

Field | Type
------|-----
X     | [bool][bool]
Y     | [bool][bool]
Z     | [bool][bool]

[Axes]: #user-content-axes

### BrickColor
Uses [uint][uint].

[BrickColor]: #user-content-brickcolor

### Color3
An object with the following fields:

Field | Type
------|-----
R     | [float][float]
G     | [float][float]
B     | [float][float]

[Color3]: #user-content-color3

### Vector2
An object with the following fields:

Field | Type
------|-----
X     | [float][float]
Y     | [float][float]

[Vector2]: #user-content-vector2

### Vector3
An object with the following fields:

Field | Type
------|-----
X     | [float][float]
Y     | [float][float]
Z     | [float][float]

[Vector3]: #user-content-vector3

### CFrame
An object with the following fields:

Field    | Type
---------|-----
Position | [Vector3][Vector3]
Rotation | [Rotation][Rotation]

[CFrame]: #user-content-cframe

#### Rotation
An object with the following fields:

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
Uses [uint][uint].

[Token]: #user-content-token

### Reference
An [int][int] corresponding to the Reference field of an [Instance][Instance],
or `null`, indicating no reference.

[Reference]: #user-content-reference

### Vector3int16
An object with the following fields:

Field | Type
------|-----
X     | [int][int]
Y     | [int][int]
Z     | [int][int]

[Vector3int16]: #user-content-vector3int16

### Vector2int16
An object with the following fields:

Field | Type
------|-----
X     | [int][int]
Y     | [int][int]

[Vector2int16]: #user-content-vector2int16

### NumberSequence
An array of [NumberSequenceKeypoint][NumberSequenceKeypoint] objects.

[NumberSequence]: #user-content-numbersequence

#### NumberSequenceKeypoint
An object with the following fields:

Field    | Type
---------|-----
Time     | [float][float]
Value    | [float][float]
Envelope | [float][float]

[NumberSequenceKeypoint]: #user-content-numbersequencekeypoint

### ColorSequence
An array of [ColorSequenceKeypoint][ColorSequenceKeypoint] objects.

[ColorSequence]: #user-content-colorsequence

#### ColorSequenceKeypoint
An object with the following fields:

Field    | Type
---------|-----
Time     | [float][float]
Value    | [Color3][Color3]
Envelope | [float][float]

[ColorSequenceKeypoint]: #user-content-colorsequencekeypoint

### NumberRange
An object with the following fields:

Field | Type
------|-----
Min   | [float][float]
Max   | [float][float]

[NumberRange]: #user-content-numberrange

### Rect
An object with the following fields:

Field | Type
------|-----
Min   | [Vector2][Vector2]
Max   | [Vector2][Vector2]

[Rect]: #user-content-rect

### PhysicalProperties
An object with the following fields:

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
An object with the following fields:

Field | Type
------|-----
R     | [uint][uint]
G     | [uint][uint]
B     | [uint][uint]

[Color3uint8]: #user-content-color3uint8

### Int64
Uses [int][int].

[Int64]: #user-content-int64

### SharedString
Uses [bytes][bytes].

[SharedString]: #user-content-sharedstring
