#!/bin/bash
cd `dirname $0`

# デフォルト値20.0はCelsiusFlag()の第2引数に与えられており、
# その引数はCelsius型であるため、何もしなくてもStringメソッドが得られる。
# CelsiusのStringメソッドにより"℃"が付与されるため、ヘルプメッセージにも"℃"が含まれる。

