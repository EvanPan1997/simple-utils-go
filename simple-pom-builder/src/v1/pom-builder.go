package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"simple-pom-builder/src/v1/util"
	"strings"
)

/*
备注:
-s=<sourceDir> 指定依赖目录
-t=<targetDir> 指定生成pom的目标目录
*/
func main() {
	// 常量
	const (
		targetName = "pom"
		postfix    = "xml"
		gid        = "com.gingkoo.xxx"
		packaging  = "war"

		defaultGroupId    = "com.gingkoo.xxx"
		defaultArtifactId = "urp-hsbc"
		version           = "1.0.0-SNAPSHOT"

		jarVersion = "1.0.0"
		scope      = "system"
	)
	// 声明变量
	var (
		defaultSource = "./libs"
		defaultTarget = "./"

		source     string
		target     string
		groupId    string
		artifactId string
	)
	// Receive Arguments From Command Line-接收命令行传参
	/*
			-s=<源lib目录>
			-t=<pom.xml输出目录>
			-g=<groupId>
			-a=<artifactId>

		Example:
			-s=<源lib目录> -t=./ -g=com.gingkoo.xxx -a=1.0.0-SNAPSHOT
	*/
	flag.StringVar(&source, "s", defaultSource, "")
	flag.StringVar(&target, "t", defaultTarget, "")
	flag.StringVar(&groupId, "g", defaultGroupId, "")
	flag.StringVar(&artifactId, "a", defaultArtifactId, "")
	flag.Parse()

	basedir := os.Getenv("PWD")
	// path := util.GetRunPath2()
	// fmt.Println(path)

	// 构造DependencyManagement结构数据
	ds := &util.Dependencies{}
	// 遍历指定目录所有文件
	fs, err := util.ListDir(source)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	for _, v := range fs {
		valLen := len(v)
		strArr := strings.Split(v, ".")
		arrLen := len(strArr)

		suffixLen := len(strArr[arrLen-1])
		// 获取后缀
		suffix := util.PakcageType(strArr[arrLen-1])
		// fmt.Println(suffix)
		// fmt.Println("name:",v,";suffix:",strArr[arrLen-1])
		// fmt.Println("arrlen:",arrLen,";suffixLen:",suffixLen)
		// fmt.Println("valLen:",valLen, ";suffixLen:",suffixLen)

		// 根据后缀进行处理
		fn := v[0 : valLen-suffixLen-1]
		dep := util.Dependency{}
		if suffix == util.JAR {
			dep = util.Dependency{
				GroupId:    gid,
				ArtifactId: fn,
				Version:    jarVersion,
				Scope:      scope,
				SystemPath: basedir + "/" + source + "/" + v,
			}
		} else if suffix == util.WAR {
			dep = util.Dependency{
				GroupId:    gid,
				ArtifactId: fn,
				Version:    jarVersion,
				Type:       string(util.WAR),
				Scope:      scope,
				SystemPath: basedir + "/" + source + "/" + v,
			}
		}
		ds.Deps = append(ds.Deps, dep)
	}

	ps := &util.Plugins{}
	args := []string{"-verbose", "-Xlint:unchecked", "-Xlint:deprecation", "-extdirs", "${project.basedir}/lib"}
	// maven plugin
	ps.Plgs = append(ps.Plgs, util.Plugin{
		GroupId:    "org.apache.maven.plugins",
		ArtifactId: "maven-compiler-plugin",
		Version:    "3.1",
		Configuration: &util.Configuration{
			Source:   "1.8",
			Target:   "1.8",
			Encoding: "UTF-8",
			CompilerArgs: &util.CompilerArgs{
				Args: args,
			},
		},
	})
	// spring boot plugin
	ps.Plgs = append(ps.Plgs, util.Plugin{
		GroupId:    "org.springframework.boot",
		ArtifactId: "spring-boot-maven-plugin",
		Version:    "2.4.0",
		// Configuration: util.Configuration{
		// 	IncludeSystemScope: "true",
		// },
	})
	// project
	p := &util.Project{
		ModelVersion: "4.0.0",
		// properties
		Properties: &util.Properties{
			ProjectBuildSourceEncoding: "UTF-8",
			MavenCompilerEncoding:      "UTF-8",
			JavaVersion:                "11",
			MavenCompilerSource:        "11",
			MavenCompilerTarget:        "11",
		},
		GroupId:    groupId,
		ArtifactId: artifactId,
		Version:    version,
		Packaging:  packaging,
		// dependencies
		Dependencies: ds,
		// build
		Build: &util.Build{
			Plugins: ps,
		},
	}

	output, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	data := []byte(xml.Header + string(output))

	filePath := target + "/" + targetName + "." + postfix
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		os.Exit(-1)
	}
	// 及时关闭file句柄
	defer file.Close()

	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.Write(data)
	write.Flush()

	os.Stdout.Write(data)
}
