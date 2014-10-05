package page

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"strings"
)

type PageSelection struct {
	selectors []string
	page      *Page
}

func (s *PageSelection) Within(selector string, bodies ...SelectionFunc) *PageSelection {
	subSelection := &PageSelection{append(s.selectors, selector), s.page}
	for _, body := range bodies {
		body.Call(subSelection)
	}
	return subSelection
}

func (s *PageSelection) ShouldContainText(text string) {
	selector := strings.Join(s.selectors, " ")
	elements, err := s.page.Driver.GetElements(selector)
	if err != nil {
		ginkgo.Fail("Failed to retrieve elements", 1)
	}
	if len(elements) > 1 {
		ginkgo.Fail("Mutiple items were selected", 1)
	}
	if len(elements) == 0 {
		ginkgo.Fail("No items were selected", 1)
	}
	elementText, err := elements[0].GetText()
	if err != nil {
		ginkgo.Fail("Failed to retrieve text for selection", 1)
	}

	gomega.ExpectWithOffset(1, elementText).To(gomega.ContainSubstring(text))
}

// TODO: unit test both selection and page with fake driver and fake ginkgo.Fail
