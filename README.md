# SpotlightControl

This activity allows you point a spot light at a target location.

## Configuration

### Input:

| Name      | Type   | Description                                          |
| :-------- | :----- | :--------------------------------------------------- |
| LightHost | string | IP address followed by port. e.g. "32.41.13.112:322" |
| X         | string | X position coordinate. e.g. "43"                     |
| Y         | string | Y position coordinate. e.g. "42"                     |
| StartX    | string | StartX position coordinate. e.g. "43"                |
| StartY    | string | StartY position coordinate. e.g. "42"                |
| Color     | string | Color for spotlight. e.g. "10"=Red or "30"=Blue      |
| ResetTime | string | Time in Minutes before light reset. e.g. "3"         |

### Output:

| Name   | Type    | Description                         |
| :----- | :------ | :---------------------------------- |
| Status | Boolean | True if sucessful, false if failure |
