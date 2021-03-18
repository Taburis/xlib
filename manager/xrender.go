package manager 

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	"fmt"
)

func NewAttr(key string, val string) html.Attribute{
	att:= html.Attribute{Namespace:"", Key:key, Val:val}
	return att
}

func AddAttr(node *html.Node, key string, val string){
	node.Attr = append(node.Attr, NewAttr(key, val))
	return
}

func Lookup(node *html.Node, key0 string) (exist bool, val0 string){
	for _, attr :=range node.Attr {
		if attr.Key == key0{
			return true, attr.Val
		}
	}
	return false, ""
}

func NewNode(data string) *html.Node {
	var attr []html.Attribute
	node := &html.Node{Parent: nil, FirstChild:nil, LastChild:nil,PrevSibling:nil,
		NextSibling :nil,
		Type: html.ElementNode,
		DataAtom: atom.Lookup([]byte(strings.Title(data))),
		Data: data,
		Namespace: "",
		Attr: attr,
	}
	return node
}

func LinkNode(n1 *html.Node, n2 *html.Node, isSibling bool){
	//rel > 0: n1 parent and n2 is child,
	//rel ==0: n1 and n2 are siblings
	//rel < 0: n2 is the parent of n1
	
	if isSibling {
		node:= n1
		for c:=n1.NextSibling; c!=nil; c=c.NextSibling {
			node=c
		}
		node.NextSibling = n2
		n2.PrevSibling = n1
		n2.Parent = n1.Parent
		if n1.Parent!=nil {n1.Parent.LastChild = n2}
	}else {
		if n1.FirstChild == nil{
			n1.FirstChild = n2
			n1.LastChild = n2
			n2.Parent= n1
		} else {
			LinkNode(n1.FirstChild, n2, true)
		}
	}
}

func AppendTOCNode(node *html.Node, label string, tocnum string, txt string, id string, isnew bool) (*html.Node, *html.Node) {
	// append nodes to node
	tnode := node
	if  isnew {
		tnode = NewNode("ul")
		LinkNode(node, tnode, false)
	}
	nodeil := NewNode("li")
	AddAttr(nodeil, "class", label)
	LinkNode(tnode, nodeil, false)
	nodea := NewNode("a")
	LinkNode(nodeil, nodea, false)
	AddAttr(nodea, "href", id)
	nodec := NewNode("span")
	AddAttr(nodec, "class", "tocnumber")
	LinkNode(nodea, nodec, false)
	nodetxt := NewNode(tocnum)
	nodetxt.Type = 1
	LinkNode(nodec, nodetxt, false)

	nodec = NewNode("span")
	AddAttr(nodec, "class", "toctext")
	LinkNode(nodea, nodec, false)
	nodetxt = NewNode(txt)
	nodetxt.Type = 1
	LinkNode(nodec, nodetxt, false)


	return tnode, nodeil
}

func MakeTOCTree(doc []*html.Node) *html.Node{
	//headtag := []string{"h2","h3","h4","h5","h6"}
	//counter := []int   {0,   0,    0,   0,   0}
	headtag := []string{"h2","h3"}
	counter := []int   {0,   0}
	root_node := NewNode("div")
	AddAttr(root_node, "class", "toc-container")
	//titlenode := NewNode("div")
	//AddAttr(titlenode, "class", "toc-caption")
	//LinkNode(root_node, titlenode, false)
	titlenode_a := NewNode("a")
	AddAttr(titlenode_a, "href", "#")
	LinkNode(root_node, titlenode_a, false)
	titlenode_head :=NewNode("h4")
	AddAttr(titlenode_head, "style", "color: white")
	LinkNode(titlenode_a, titlenode_head, false)
	titlenode_doc :=NewNode("Content")
	titlenode_doc.Type = 1
	LinkNode(titlenode_head, titlenode_doc, false)
	

	
	tailnodes:=[]*html.Node{root_node, nil,nil}
	pindex:= -1
	index := -1
	for _, node := range doc {
		//fmt.Println("!")
		for i, tag := range headtag{
			if node.Data == tag{
				index = i
				counter[i]++
				//fmt.Printf("%d, %d\n",i, counter[i])
				break
			}
		}
		if index < 0 { continue }
		isnew := false
		if pindex < index {
			isnew = true
			counter[index] = 1 //rest counter for new section
		}
		label := fmt.Sprintf("toclevel-%d", counter[0])
		if index > 0 {
			label = fmt.Sprintf("%s tocsection-%d", label, counter[1])
		}
		txt:=node.FirstChild.Data
		_, id:= Lookup(node, "id")
		tocnum := fmt.Sprintf("%d", counter[0])	

		if index == 1 { 
			tocnum = fmt.Sprintf("%d.%d", counter[0], counter[1])
		}
		id = fmt.Sprintf("#%s",id)
		tailnodes[index], tailnodes[index+1]=AppendTOCNode(tailnodes[index], label, tocnum, txt, id, isnew)
		pindex = index
	}
	return root_node
}

func GetHTree(doc *html.Node) (nodes []*html.Node){
 	//headtag := []string{"h1","h2","h3","h4","h5","h6"}
 	headtag := []string{"h1","h2","h3"} //now only support upto depth 3

	var walker func(*html.Node) 

	walker=func(node *html.Node){
		if node.Type == html.ElementNode{
			for _, tag := range headtag{
				if node.Data == tag {
					nodes = append(nodes,node)
					break
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			walker(child)
		}
	}
	walker(doc)
	return
}

func MakeTOCNodes(html_node *html.Node) (*html.Node){
	//input node should be the nodes only contained headers
	nodes:=GetHTree(html_node)
	return MakeTOCTree(nodes)
}

func MakeIndexPage(h1title string, path string, candi string) (*html.Node){
	queue := strings.Split(candi, "\n")
	tmp := strings.Split(queue[0], "/")
	h3title := tmp[len(tmp)-2]
	root_node:= NewNode("div")
	AddAttr(root_node, "class", "category-wrapper")
	node := NewNode("h1")
	AddAttr(node, "id", "category_notes")
	LinkNode(root_node, node, false)
	txtnd := NewNode(strings.ToUpper(h1title))
	txtnd.Type = 1
	LinkNode(node, txtnd, false)

	node = NewNode("h3")
	AddAttr(node, "id", h3title)
	LinkNode(root_node, node, false)
	txtnd = NewNode(strings.ReplaceAll(h3title, "_", " "))
	txtnd.Type = 1
	LinkNode(node, txtnd, false)

	node = NewNode("ul")
	LinkNode(root_node, node, false)

	for _, filepath:= range queue{
		stmp := strings.Split(filepath, "/")
		if len(stmp) < 2 { continue }
		checkTitle := stmp[len(stmp)-2]
		if h3title != checkTitle{
			h3title = checkTitle
			tmpnode := NewNode("h3")
			AddAttr(node, "id", h3title)
			LinkNode(root_node, tmpnode, false)
			txtnd = NewNode(strings.ReplaceAll(h3title, "_", " "))
			txtnd.Type = 1
			LinkNode(tmpnode, txtnd, false)

			node = NewNode("ul")
			LinkNode(root_node, node, false)
		}
		fileName := stmp[len(stmp)-1]
		stmp = strings.Split(fileName, ".")
		fileName = strings.ReplaceAll(stmp[0], "_", " ")
		inode := NewNode("li")
		LinkNode(node, inode, false)
		anode := NewNode("a")
		AddAttr(anode, "href", fmt.Sprintf("/%s/%s/%s",h1title, h3title, stmp[0]))
		LinkNode(inode, anode, false)
		txtnode := NewNode(fileName)
		txtnode.Type=1
		LinkNode(anode, txtnode, false)
	}
	return root_node
}
