package manager
import (
	//
	"bytes"
	"fmt"
	"io/ioutil"
	"io"
	"strings"
	"os"
	//"errors"
	"path/filepath"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
    "github.com/gomarkdown/markdown/html"
    html2 "golang.org/x/net/html"
)

type  Manager struct {
	parser *parser.Parser
	renderer *html.Renderer
}

func New() *Manager{
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	t := &Manager{parser: parser.NewWithExtensions(extensions),
	 renderer: html.NewRenderer(opts),
	}
	return t
}

func (p *Manager) Config0() {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p.parser = parser.NewWithExtensions(extensions)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	p.renderer = html.NewRenderer(opts)
}


func (p *Manager) MD2html(in_name string, out_name string)(error, error, error){
	// converting MD files to html format file
	// in_name : abs path of the input md file
	// out_name: abs path of the output html file
	content , error := ioutil.ReadFile(in_name)
	if error != nil {
        return error, nil, nil
    }	
    p.Config0()
    //fmt.Printf("here -------------------------------------------\n")
    myhtml := markdown.ToHTML(content, p.parser, p.renderer)
    toc_str:=p.TOCGenerator(string(myhtml))
    toc := []byte(toc_str)
    //fmt.Printf("down -------------------------------------------\n")
    error1 := ioutil.WriteFile(out_name, myhtml, 0644)
    error2 := ioutil.WriteFile(fmt.Sprintf("%s_toc",out_name), toc, 0644)
    if error1 != nil {panic (error1)}
    if error2 != nil {panic (error2)}
    return nil, error1, error2
}

func (p *Manager) ListDir(fileType string, path string) string{
	// iterating over the folders (path) and return the list of 
	// paths of files that satisfies the given type (fileType)
	var list strings.Builder
	filepath.Walk(path, func (path string, info os.FileInfo, err error) error{
			if err != nil {
				return err
			}
			s := strings.Split(path, ".")
			if fileType == s[len(s)-1]{
				list.WriteString(path)
				list.WriteString("\n")
				//fmt.Printf("found: %s\n",path)
			}
			return nil
		})
	return list.String()
}

func (p *Manager) DetachPath(in string) (path string, name string, ftype string){
	// strip off the path and return the file name
	fds := strings.Split(in, "/")
	filename := fds[len(fds)-1]
	filepart := strings.Split(filename, ".")
	ftype = filepart[len(filepart)-1]
	name = filepart[0]
	if len(fds) > 1{
		path = in[:len(in)-len(name)-len(ftype)-2]
	} else {
		path = ""
	}
	return 
}

func (p *Manager) ProduceHTML(root string, dest string) error {
	// root: root path of the inputs;
	// dest: the root path of outputs.
	if root[len(root)-1] == '/' {
		root = root[:len(root)-1]
	}
	if dest[len(dest)-1] == '/' {
		dest = dest[:len(dest)-1]
	}
	candidates := p.ListDir("md", root)
	queue := strings.Split(candidates,"\n")
	nroot := len(root)
	for _,f := range queue{
		if len(f) < nroot {continue}
		ipath, ifile, _ := p.DetachPath(f[nroot-1:])
		relpath := fmt.Sprintf("%s/%s",dest, ipath)
		ofile   := fmt.Sprintf("%s/%s.%s", relpath, ifile, "html")
		fmt.Println(relpath)
		fmt.Printf("Compiling: %s...\n",f)
		os.MkdirAll(relpath, 0755)
		p.MD2html(f, ofile)
		fmt.Printf("finished, dump to %s\n", ofile)
	}
	return nil
}

func (p *Manager) RenderNode(n *html2.Node) string {
    var buf bytes.Buffer
    w := io.Writer(&buf)
    html2.Render(w, n)
    return buf.String()
}

func (p *Manager) TOCGenerator(doc string) (toc string){
	ns,_:=html2.Parse(strings.NewReader(doc))
	return p.RenderNode( MakeTOCNodes(ns))
}

func (p *Manager) ProduceIndex(path string, output string) error{
	// for instance, for notes with path "./dist/notes/subjects/articles.html"
	// the path should be path "./dist/notes"
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	sbuf := strings.Split(path, "/")
	h1title := sbuf[len(sbuf)-1]
	candidates:=p.ListDir("html", path)
	htmlstr := p.RenderNode(MakeIndexPage(h1title, path, candidates))
	htmlbyte := []byte(htmlstr)
	error := ioutil.WriteFile(fmt.Sprintf("%s/%sIndex.html",output,h1title), htmlbyte, 0644)
	return error
}

