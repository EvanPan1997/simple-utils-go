package util

import "encoding/xml"

// Project xml整体结构
type Project struct {
	XMLName      xml.Name `xml:"project"`
	ModelVersion string   `xml:"modelVersion"`
	GroupId      string   `xml:"groupId"`
	ArtifactId   string   `xml:"artifactId"`
	Version      string   `xml:"version"`
	Packaging    string   `xml:"packaging,omitempty"`

	Properties *Properties
	// DependencyManagement	*DependencyManagement
	Dependencies *Dependencies
	Build        *Build
}

type Properties struct {
	XMLName                    xml.Name `xml:"properties"`
	ProjectBuildSourceEncoding string   `xml:"project.build.sourceEncoding"`
	MavenCompilerEncoding      string   `xml:"maven.compiler.encoding"`
	JavaVersion                string   `xml:"java.version"`
	MavenCompilerSource        string   `xml:"maven.compiler.source"`
	MavenCompilerTarget        string   `xml:"maven.compiler.target"`
}

// type DependencyManagement struct {
// 	XMLName	xml.Name               `xml:"dependencyManagement"`
// 	Dependencies		*Dependencies `xml:"dependencies"`
// }

type Dependencies struct {
	XMLName xml.Name     `xml:"dependencies"`
	Deps    []Dependency `xml:"dependency"`
}

type Dependency struct {
	XMLName xml.Name `xml:"dependency"`

	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Type       string `xml:"type,omitempty"`
	Scope      string `xml:"scope,omitempty"`
	SystemPath string `xml:"systemPath,omitempty"`
}

type Build struct {
	XMLName xml.Name `xml:"build"`
	Plugins *Plugins
}

type Plugins struct {
	XMLName xml.Name `xml:"plugins"`
	Plgs    []Plugin
}

type Plugin struct {
	XMLName       xml.Name `xml:"plugin"`
	GroupId       string   `xml:"groupId,omitempty"`
	ArtifactId    string   `xml:"artifactId,omitempty"`
	Version       string   `xml:"version,omitempty"`
	Configuration *Configuration
}

type Configuration struct {
	XMLName            xml.Name      `xml:"configuration"`
	Source             string        `xml:"source,omitempty"`
	Target             string        `xml:"target,omitempty"`
	Encoding           string        `xml:"encoding,omitempty"`
	CompilerArgs       *CompilerArgs `xml:"compilerArgs,omitempty"`
	IncludeSystemScope string        `xml:"includeSystemScope,omitempty"`
}

type CompilerArgs struct {
	Args []string `xml:"arg,omitempty"`
}
