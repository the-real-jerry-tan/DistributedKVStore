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
 * This file contains logic for managing individual nodes in the distributed system.
 */

package main

import (
	"log"
)

type Node struct {
	ID   string
	Addr string
}

func (n *Node) JoinCluster(cluster *Cluster) {
	log.Printf("Node %s at address %s joining the cluster\n", n.ID, n.Addr) // Properly formatted log message
}

func (n *Node) LeaveCluster(cluster *Cluster) {
	log.Printf("Node %s at address %s leaving the cluster\n", n.ID, n.Addr)
	// Code to remove the node from the cluster
}
