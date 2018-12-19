package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/models"
	"github.com/turing-ml/turing-api/pkg/dag"

	"github.com/gin-gonic/gin"
	"github.com/turing-ml/turing-api/api/utils"
)

// Play will be running the whole pipeline into NiFi
func Play(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("playground_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodes, err := models.GetNodes(db, playgroundID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	g, err := buildDAG(db, nodes)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	template := traverseDAG(g)

	// send to Apache NiFi
	fmt.Print(template)

	utils.Response(c, http.StatusOK, g.String())
}

func buildDAG(db *gorm.DB, nodes []models.Node) (*dag.DAG, error) {
	graph := dag.NewDAG()

	// add all the vertexes to the graphs
	for _, node := range nodes {
		fields, err := models.GetFields(db, int(node.ID))
		if err != nil {
			return nil, err
		}

		v := buildVertex(node, fields)
		err = graph.AddVertex(v)
		if err != nil {
			return nil, err
		}
	}

	// add the edges to the graphs
	for _, node := range nodes {
		links, err := models.GetLinks(db, int(node.ID))
		if err != nil {
			return nil, err
		}

		var outputFromNodeID int
		for _, ls := range links {
			// distiguish between inputs and output. AUTO_INCREMENT starts from 1
			// TODO: find a smarter way to do this
			if ls.Link.Output.From.NodeID != 0 {
				outputFromNodeID = ls.Link.Output.From.NodeID
			}

			for _, input := range ls.Link.Inputs {
				// create link
				if input.To.NodeID != outputFromNodeID {
					i, err := graph.GetVertex(strconv.Itoa(outputFromNodeID))
					if err != nil {
						return nil, err
					}
					o, err := graph.GetVertex(strconv.Itoa(input.To.NodeID))
					if err != nil {
						return nil, err
					}

					graph.AddEdge(i, o)
				}
			}
		}
	}

	return graph, nil
}

func buildVertex(node models.Node, fields []models.Field) *dag.Vertex {

	ng := &models.NodeGraph{
		Node:   &node,
		Fields: fields,
	}

	return dag.NewVertex(strconv.Itoa(int(node.ID)), ng)
}

func traverseDAG(g *dag.DAG) *models.Template {

	template := models.NewTemplate("test")

	for _, v := range g.GetVertices() {
		if !v.Visited {
			v.Visited = true
			n := v.Value.(*models.NodeGraph)
			template.AddFlow(n, n.Node.Type, n.Node.SubType)

			for _, c := range v.Children.Values() {
				child := c.(*dag.Vertex)
				if !child.Visited {
					child.Visited = true
					n := child.Value.(*models.NodeGraph)
					template.AddFlow(n, n.Node.Type, n.Node.SubType)
				}
			}
		}
	}
	return template
}
