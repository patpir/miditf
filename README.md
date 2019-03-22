# MIDITF

Framework to apply transformations to and print visualizations of MIDI files.

## Features

- [ ] Transpose notes to another key
- [ ] Visualize key presses for Piano
- [ ] Visualize button presses for Steirische Harmonika
- [ ] Control playback speed and repeating range for key press visualization
- [ ] Show notes in various formats

## Development

This software is written in Go.

The project is intended to be used as a framework for other projects.
Therefore, compilation to an executable binary is not supported.

Use one of the following commands to execute the unit tests:

 - `make test` executes all tests and displays a summary for every package
 - `make verbose` executes and displays details for every test
 - `make coverage` executes all tests and displays coverage per package


## Concept

Transformations and visualizations are applied in a pipeline.
A pipeline consists of a source, a transform and a visualization stage.

The source stage consists of one or more individual sources, each of which
is later processed independently of the other sources.
The transform stage consists of one or more transformations, which are applied
in the specified order **for every source**.
The output stage consists of one or more visualizations, which are generated
for every result.

Consider an example of:
 - three sources S1, S2, and S3
 - two transformations T1 and T2
 - and two visualizations V1 and V2

This pipeline would be processed as follows:

```
                                              +------+
                                         ,->  |  V1  |
+------+      +------+      +------+    /     +------+
|  S1  |  ->  |  T1  |  ->  |  T2  |  -<
+------+      +------+      +------+    \     +------+
                                         `->  |  V2  |
                                              +------+

                                              +------+
                                         ,->  |  V1  |
+------+      +------+      +------+    /     +------+
|  S2  |  ->  |  T1  |  ->  |  T2  |  -<
+------+      +------+      +------+    \     +------+
                                         `->  |  V2  |
                                              +------+

                                              +------+
                                         ,->  |  V1  |
+------+      +------+      +------+    /     +------+
|  S3  |  ->  |  T1  |  ->  |  T2  |  -<
+------+      +------+      +------+    \     +------+
                                         `->  |  V2  |
                                              +------+
```

### Blocks

Anything in a pipeline is called a *Block*.
The `github.com/patpir/miditf/blocks.Block` interface provides a recipe to
instantiate a `Source`, a `Transformation`, or a `Visualization`.

Every `Block` has a `TypeId`, zero or more `Arguments` and a `Comment`.
The `TypeId` identifies the type of `Source` / `Transformation` /
`Visualization` to use.
The `Arguments` are passed when creating the `Source` / `Transformation` /
`Visualization`.
The `Comment` is any user-provided string for traceability.


### Sources

Sources implement the `github.com/patpir/miditf/blocks.Source` interface:

```go
type Source interface {
	Piece() (*core.Piece, error)
}
```

`Source.Piece()` must return either a `Piece` or an `error`.

The `Piece()` method of any `Source` instance will only be called once during
its whole lifetime.
This enables the use of (file) readers as arguments.


### Transformations

Transformations implement the `github.com/patpir/miditf/blocks.Transformation`
interface:

```go
type Transformation interface {
	Transform(piece *core.Piece) (*core.Piece, error)
}
```

`Transformation.Transform()` must return either a `Piece` or an `error`.
Feel free to modify and return the `Piece` instance received as the `piece`
parameter or to construct a new instance.

The `Transform()` method of any `Transformation` instance may be called
multiple times during its lifetime.
This means that `Transform()` should not return a Piece which is kept as a
member variable. Doing so could lead to unexpected results if the next
`Transformation` in the pipeline modifies the `piece` parameter's instance.


### Visualizations

Visualizations implement the `github.com/patpir/miditf/blocks.Visualization`
interface:

```go
type Visualization interface {
	Visualize(piece *core.Piece) (string, error)
}
```

`Visualization.Visualize()` must return either a `string` representation of the
given `Piece` or an `error`.
*Visualization* only describes the typical use-case of the last pipeline stage.
A `Visualization` may, for example, produce an audio file (nothing visual
about that).

The `Visualize()` method of any `Visualization` instance may be called
multiple times during its lifetime.


### Registration

`Source`s, `Transformation`s and `Visualization`s are managed via a
registration mechanism.

All integrated `Source`s, `Transformation`s and `Visualization`s will be
registered by including the respective package:

```go
import (
	_ "github.com/patpir/miditf/sources"
	_ "github.com/patpir/miditf/transform"
	_ "github.com/patpir/miditf/visualize"
)
```

All registrations will normally take place on the `DefaultRegistrator()`
provided by the `github.com/patpir/miditf/blocks` package.
Additional `Source`s, `Transformation`s and `Visualization`s can be registered
by calling `RegisterSource()`, `RegisterTransformation()`, or
`RegisterVisualization()` on the `DefaultRegistrator()`.
These methods require a unique identifier (*unique* means not used for any
other registration of the same type) and a function to create a new `Source` /
`Transformation` / `Visualization` instance.

The factory function types are defined as:

```go
type SourceFactory func(map[string]interface{}) (Source, error)
type TransformationFactory func(map[string]interface{}) (Transformation, error)
type VisualizationFactory func(map[string]interface{}) (Visualization, error)
```

The map parameter contains the arguments specified by the user.
Usually, they will be stored in the instance created by the factory function.

To receive a list of all registered `Source`s / `Transformation`s /
`Visualization`s, use the following methods of the `DefaultRegistrator()`:

```go
Sources() []BlockInfo
Transformations() []BlockInfo
Visualizations() []BlockInfo
```

