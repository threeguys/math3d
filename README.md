math3d
======

The math3d library is a very bare bones 3d math library to support using OpenGL in go. Only a few 3d operations will be implemented using a couple of simple primitives. I also plan on adding some level of integration with either go-gl or gogl. The plan is to use float32 since that's the floating point type supported by OpenGL.

Most of the math has been lifted from either pages around the net and reading the code in glm. I've done some simple parallelization of the matrix multiplication as well as some loop unrolling as a performance improvement but right now I'm focusing on functionality over performance. After I've got it working and written a few OpenGL programs with it, I may go back and do some refactoring to gain boost performance.

Also, I'm brand new to go so if you see anything odd, please let me know. Any input is appreciated.

Here's a list of the implemented primitives and related operations:

Vector
------
A simple 3d vector. Basically just a 3 element array of *float32*

<table>
	<thead>
		<tr>
			<th>Operation</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>NewVector</td>
			<td>Creates a new instance of Vector and initializes with the passed in points</td>
		</tr>
		<tr>
			<td>X</td>
			<td>The X component of the vector</td>
		</tr>
		<tr>
			<td>Y</td>
			<td>The Y component of the vector</td>
		</tr>
		<tr>
			<td>Z</td>
			<td>The Z component of the vector</td>
		</tr>
		<tr>
			<td>Length</td>
			<td>The magnitude of the vector</td>
		</tr>
		<tr>
			<td>Add</td>
			<td>Adds the argument to the current vector and returns a new instance containing the sum</td>
		</tr>
		<tr>
			<td>Normalize</td>
			<td>Returns the current vector as a unit vector</td>
		</tr>
		<tr>
			<td>Dot</td>
			<td>Returns the dot product of the vector and the passed in vector</td>
		</tr>
		<tr>
			<td>Cross</td>
			<td>Returns the cross product of the current vector and the passed in vector, as a new vector</td>
		</tr>
	</tbody>
</table>

Matrix
------
3d matrix (4x4) implemented as an 16 element array of float32

<table>
	<thead>
		<tr>
			<th>Function</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>NewMatrix</td>
			<td>Creates a new instance of a matrix and initializes it with the passed in values</td>
		</tr>
		<tr>
			<td>Identity</td>
			<td>Returns a new instance of the identity matrix</td>
		</tr>
		<tr>
			<td>Print</td>
			<td>Prints the matrix to stdout, basically just a debug function</td>
		</tr>
		<tr>
			<td>SetValues</td>
			<td>Sets the values of the matrix</td>
		</tr>
		<tr>
			<td>MultipleMatrices</td>
			<td>Multiplies multiple matrices together and returns the resulting matrix</td>
		</tr>
		<tr>
			<td>NaiveMultiply</td>
			<td>Multiplies the current matrix with the passed in matrix and returns a new matrix with the result. Naive implementation with nested loops. Should probably be private but it's currently public in order to do some performance testing</td>
		</tr>
		<tr>
			<td>Multiply</td>
			<td>Currently just calls NaiveMultiply. Intended to be the public face of matrix multiplication</td>
		</tr>
	</tbody>
</table>

License
-------

*Released under Apache 2 license*

Copyright 2013 Ray Cole

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.