package schema


// Definition for a generator
#HofGenerator: {
  // Base directory for the output
  Outdir: string | *"./"

  // "Global" input, merged with out replacing onto the files
  In: { ... } | * {...}

  // The list fo files for hof to generate
  Out: [...#HofGeneratorFile] | *[...]

  //  Attribute expansions are used to turn the @attributes
	//  into something you can use in the templates
	//  (i.e. we need to add it to the data model after Cue processing)
	//  
	AttributeExpansions?: {
		[AttrName=string]: {
			InjectName: string | *AttrName
			// keep the value (by default), replace with anything
			[KeyName=string]?: _
		}
	}

  // The following will be automatically added to the template context
  // under its name for reference in GenFiles  and partials in templates
  NamedTemplates: { [Name=string]: string }
  NamedPartials:  { [Name=string]: string }
  // Static files are available for pure cue generators that want to have static files
  // These should be named by their filepath, but be the content of the file
  StaticFiles: { [Name=string]:  string }

  //
  // For file based generators
  //   files here will be automatically added to the template context
  //   under its filepath for reference in GenFiles and partials in templates

  // Used for indexing into the vendor directory...
  PackageName: string

  // Filepath globs for static files to load
  StaticGlobs: [...string] | *[]

  // Base directory of partial templatess to load
  PartialsDir: string | * "/partials"

  // Base directory of entrypoint templates to load
  TemplatesDir: string | * "/templates"

  TemplatesDirConfig: [Filename=string]: #TemplateConfigReplacible

  TemplateConfig: #DefaultTemplateConfig

  //
  // Open for whatever else you may need as a generator writer
  ...
} 
