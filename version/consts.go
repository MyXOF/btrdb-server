package version

import "fmt"

const Major = 4
const Minor = 9
const Subminor = 1

var VersionString = fmt.Sprintf("%d.%d.%d", Major, Minor, Subminor)
var BuildDate string
