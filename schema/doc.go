// Package schema parses Apple's [device management] YAML schemas (commands,
// profiles, and declarative management) into a Go AST of structs, enums, and
// maps.
//
// Parse one or more schema files with [New] or [NewFromFile], build a [File]
// AST with [NewFile], then encode to Go source with [Encoder]:
//
//	s, err := schema.NewFromFile("path/to/profile.yaml")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	f := schema.NewFile([]*schema.Schema{s})
//
//	out := jen.NewFile("mypackage")
//	enc := schema.NewEncoder(out, schema.WithTags([]string{"json", "plist"}))
//	enc.Encode(f)
//
//	fmt.Println(out.GoString())
//
// [device management]: https://github.com/apple/device-management
package schema
