package buildkit

import (
	"regexp"
	"strings"
)

type node struct {
	part      string
	isDynamic bool
	regex     *regexp.Regexp
	children  []*node
	isEnd     bool
	route     *route
}

type routeMatch struct {
	route  *route
	params map[string]string
}

type trie struct {
	root *node
}

func newNode() *node {
	return &node{}
}

func (t *trie) insert(pattern string, r *route) {
	parts := strings.Split(pattern, "/")
	n := t.root

	for _, part := range parts {
		matched := false
		for _, child := range n.children {
			if child.part == part || child.isDynamic {
				n = child
				matched = true
				break
			}
		}

		if !matched {
			isDynamic := strings.HasPrefix(part, ":")
			regexPattern := ""
			if isDynamic && strings.Contains(part, "(") && strings.HasSuffix(part, ")") {
				regexPattern = part[strings.Index(part, "(")+1 : len(part)-1]
				part = part[0:strings.Index(part, "(")]
			}
			child := &node{part: part, isDynamic: isDynamic}
			if regexPattern != "" {
				child.regex = regexp.MustCompile(regexPattern)
			}
			n.children = append(n.children, child)
			n = child
		}
	}
	n.isEnd = true
	n.route = r
}

func (t *trie) search(path string) *routeMatch {
	parts := strings.Split(path, "/")
	n := t.root
	params := make(map[string]string)

	for _, part := range parts {
		matched := false
		for _, child := range n.children {
			if child.part == part || child.isDynamic || (child.regex != nil && child.regex.MatchString(part)) {
				if child.isDynamic || child.regex != nil {
					params[strings.TrimPrefix(child.part, ":")] = part
				}
				n = child
				matched = true
				break
			}
		}

		if !matched {
			return nil
		}
	}

	if n.isEnd {
		return &routeMatch{route: n.route, params: params}
	}
	return nil
}
