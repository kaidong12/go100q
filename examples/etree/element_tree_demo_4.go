package etree

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/goccy/go-json"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type CodeBook struct {
	Id      int               `json:"id"`
	Name    string            `json:"name"`
	Pattern map[string]string `json:"pattern"`
}

var CodeLinePatten = regexp.MustCompile(`[0-9]+`)

func parserCodeBook(codeBookPath string) (map[int]*CodeBook, error) {
	res := make(map[int]*CodeBook)
	fileName := filepath.Base(codeBookPath)
	if strings.HasSuffix(fileName, ".xml") {
		doc := etree.NewDocument()
		if err := doc.ReadFromFile(codeBookPath); err != nil {
			return nil, err
		}
		root := doc.Root()
		if root.Tag != "codebook" {
			return nil, fmt.Errorf("not a codebook file")
		}
		if elements := root.FindElements(`./content/module`); elements != nil {
			for _, element := range elements {
				idElm := element.FindElement(`./id`)
				tagElm := element.FindElement(`./tag`)
				if idElm != nil && tagElm != nil {
					name, _ := strings.CutPrefix(tagElm.Text(), "FILE_")
					if id, err := strconv.Atoi(idElm.Text()); err == nil {
						codeBook := &CodeBook{Id: id, Name: name}
						codeBook.Pattern = make(map[string]string)
						if traceElms := root.FindElements(fmt.Sprintf(`./single_module_description/module_identification/module_name[text()='%s']/..//../trace`, tagElm.Text())); traceElms != nil {
							for _, elm := range traceElms {
								lineElm := elm.FindElement(`./line`)
								strElm := elm.FindElement(`./str`)
								if lineElm != nil && strElm != nil {
									codeBook.Pattern[CodeLinePatten.FindString(lineElm.Text())] = strings.Replace(strElm.Text(), "%u", "%d", -1)
								}
							}
						}

						res[id] = codeBook
					}
				}
			}
		}
	} else if strings.HasSuffix(fileName, ".json") {
		file, err := os.ReadFile(codeBookPath)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(file, &res)
		if err != nil {
			slog.Error("[serial2 decoder]err to unmarshal codebook from json", "err", err.Error())
			return nil, err
		}
	}
	return res, nil
}
