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

