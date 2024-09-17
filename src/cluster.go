/**
 * © 2024 Jerry Tan. All Rights Reserved.
 *
 * This software is the confidential and proprietary information of Jerry Tan
 * ("Confidential Information"). You shall not disclose such Confidential Information
 * and shall use it only in accordance with the terms under which this software
 * was distributed or otherwise published, and solely for the prior express purposes
 * explicitly communicated and agreed upon, and only for those specific permissible purposes.
 *
 * This software is provided "AS IS," without a warranty of any kind. All express or implied
 * conditions, representations, and warranties, including any implied warranty of merchantability,
 * fitness for a particular purpose, or non-infringement, are disclaimed, except to the extent
 * that such disclaimers are held to be legally invalid.
 *
 * This file contains logic for managing the cluster and node operations.
 */

package main

import (
	"log"
)

type Cluster struct {
	Nodes []*Node
}

func NewCluster() *Cluster {
	return &Cluster{
		Nodes: []*Node{},
	}
}

func (c *Cluster) AddNode(node *Node) {
	c.Nodes = append(c.Nodes, node)
	log.Printf("Node %s added to cluster successfully\n", node.ID) // Properly formatted log message
}

func (c *Cluster) RemoveNode(node *Node) {
	log.Printf("Removing node %s from the cluster\n", node.ID)
	// Code to remove node from the cluster
}

func (c *Cluster) ElectLeader() {
	log.Printf("Electing leader among %d nodes\n", len(c.Nodes))
	if len(c.Nodes) > 0 {
		log.Printf("Leader elected: %s\n", c.Nodes[0].ID)
	} else {
		log.Printf("No nodes available for leader election\n")
	}
}
