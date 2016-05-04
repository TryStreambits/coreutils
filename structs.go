package codeutilsShared

// ProjectConfig is the configuration for a Project
type ProjectConfig struct {
	Name         string
	ContentTypes []string
	Go           GoCompilerOptions         // Go key maps to GoCompilerOptions struct
	HTML         HTMLCompilerOptions       // HTML key maps to HTMLCompilerOptions struct
	LESS         LESSCompilerOptions       // LESS key maps to LessCompilerOptions struct
	TypeScript   TypeScriptCompilerOptions // TypeScript maps to TypeScriptCompilerOptions struct
	UsesTests    bool
}

// #region Go

// GoCompilerOptions are options available for compiling the entire Go source tree
type GoCompilerOptions struct {
	CompileFromRoot bool                       // Defaults to true, compiles everything in root into binary
	Branches        map[string]GoBranchOptions // Maps each string to a GoBranchOptions struct
}

// GoBranchOptions are options available for compiling individual branches in the Go source tree
type GoBranchOptions struct {
	BinaryName       string   // Name of the binary to output (will default to project name)
	DoInstall        bool     // Set DoInstall as a bool (BinaryName will be ignored) Implies this branch is a go package.
	SourcesToInclude []string // Sources to include, you can be very selective here
}

// #endregion

// #region HTML

// HTMLCompilerOptions are options provided when compiling the HTML content
type HTMLCompilerOptions struct {
	EnableFrala bool // Defaults to false, enables Frala file parsing
}

// #endregion

// #region LESS

// LESSCompilerOptions are options available for compiling the entire LESS source tree
type LESSCompilerOptions struct {
	Branches map[string]LESSBranchOptions // Maps each string to a LESSBranchOptions struct
}

// LESSBranchOptions are options available for compiling individual branches in the LESS source tree
type LESSBranchOptions struct {
	AdditionalCompileOptions []string // Additional compile options to pass to lessc
	FileName                 string   // Name of the file of the outputted CSS
	SourcesToInclude         []string // Sources to include, you can be very selective here
	UniqueHash               bool     // Define UniqueHash as a bool, determines whether we should append a fragment of the hash of the file to the file name
	UseGlob                  bool
}

// #endregion

// #region TypeScript

// TypeScriptCompilerOptions are options available for compiling TypeScript
type TypeScriptCompilerOptions struct {
	LibreJSLicense   string // Define LibreJSLicense as the valid Spdx license that is also valid with LibreJS Gopher / LibreJS
	MinifyContent    bool   // Defaults to true, compiles using Google closure compiler
	Target           string // Defaults to ES5 - ES3 and ES6 should be options via TS
	UniqueHash       bool   // Define UniqueHash as a bool, determines whether we should append a fragment of the hash of the file to the file name
	UseLibreJSHeader bool   // Define LibreJSHeader, adds compliant headers to JS code
}

// #endregion
