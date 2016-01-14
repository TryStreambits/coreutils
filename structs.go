package codeutilsShared

// Structure of Project configuration
type ProjectConfig struct {
	ProjectName  string
	ContentTypes []string
	Go           GoCompilerOptions         // Go key maps to GoCompilerOptions struct
	HTML         HTMLCompilerOptions       // HTML maps to HTMLCompilerOptions struct
	LESS         LESSCompilerOptions       // LESS key maps to LessCompilerOptions struct
	TypeScript   TypeScriptCompilerOptions // TypeScript maps to TypeScriptCompilerOptions struct
	UsesTests    string
}

// #region Go

// Options available for compiling the entire Go source tree
type GoCompilerOptions struct {
	CompileFromRoot bool                       // Defaults to true, compiles everything in root into binary
	Branches        map[string]GoBranchOptions // Maps each string to a GoBranchOptions struct
}

// Options available for compiling individual branches in the Go source tree
type GoBranchOptions struct {
	BinaryName       string   // Name of the binary to output (will default to project name)
	DoInstall        bool     // Set DoInstall as a bool (BinaryName will be ignored) Implies this branch is a go package.
	SourcesToInclude []string // Sources to include, you can be very selective here
}

// #endregion

// #region HTML

// Options available for "compile" HTML
type HTMLCompilerOptions struct {
	Compress bool // Defaults to true, despite bool defaults to false
	CopyAll  bool // Defaults to true, copies everything in src/html
}

// #endregion

// #region LESS

// Options available for compiling the entire LESS source tree
type LESSCompilerOptions struct {
	CompileFromRoot bool                         // Defaults to true, compiles everything in root into binary
	Branches        map[string]LESSBranchOptions // Maps each string to a LESSBranchOptions struct
}

// Options available for compiling individual branches in the LESS source tree
type LESSBranchOptions struct {
	AdditionalCompileOptions []string // Additional compile options to pass to lessc
	FileName                 string   // Name of the file of the outputted CSS
	UniqueHash               bool     // Define UniqueHash as a bool, determines whether we should append a fragment of the hash of the file to the file name
	UseGlob                  bool
}

// #endregion

// #region TypeScript

// Options available for compiling TypeScript
type TypeScriptCompilerOptions struct {
	LibreJSLicense    string // Define LibreJSLicense as the valid Spdx license that is also valid with LibreJS Gopher / LibreJS
	MinifyContent     bool   // Defaults to true, compiles using Google closure compiler
	OutputDeclaration bool   // Defaults to false, like a normal bool type
	RemoveComments    bool   // Defaults to true
	Target            string // Defaults to ES5 - ES3 and ES6 should be options via TS
	UniqueHash        bool   // Define UniqueHash as a bool, determines whether we should append a fragment of the hash of the file to the file name
	UseLibreJSHeader  bool   // Define LibreJSHeader, adds compliant headers to JS code
}

// #endregion
