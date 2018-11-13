///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing testing package
import (
	"testing"
)

// NewKnowledgeGraph test method
func TestNewKnowledgeGraph(test *testing.T) {

	var knowledgeGraph *KnowledgeGraph

	knowledgeGraph = NewKnowledgeGraph()

	test.Log(knowledgeGraph)

	if knowledgeGraph == nil {

		test.Errorf("error in creating a knowledgeGraph")
	}

}

// NewMapLayout test method
func TestNewMapLayout(test *testing.T) {

	var mapLayout *MapLayout

	mapLayout = NewMapLayout()

	test.Log(mapLayout)

	if mapLayout == nil {

		test.Errorf("error in creating a mapLayout")
	}

}

// NewSocialGraph test method
func TestNewSocialGraph(test *testing.T) {

	var socialGraph *SocialGraph

	socialGraph = NewSocialGraph(1)

	if socialGraph == nil {

		test.Errorf("error in creating a socail Graph")
	}

}
