## Write a function to see if a binary tree is "superbalanced". A tree is "superbalanced"
## if the difference between the depths of any two leaf nodes is no greater than one.
## Here is a smple binary node tree class:

class BinaryTreeNode
  attr_accessor :value
  attr_reader :left, :right

  def initialize(value)
    @value = value
    @left = nil
    @right = nil
  end

  def insert_left(value)
    @left = BinaryTreeNode.new(value)
    return @left
  end

  def insert_right(value)
    @right = BinaryTreeNode.new(value)
    return @right
  end
end

## Gotchas: this is solvable in O(n) time and space

def binary_tree_is_balanced?(tree_root)
  depths = []
  nodes = [] # this will be a stack that will store pairs of nodes: [node, depth]
  nodes.push([tree_root, 0])

  while !nodes.empty?
    node, depth = nodes.pop
    if !node.left && !node.right # this is a leaf node
      if !depths.include?(depth)
        depths.push(depth)

        if (depths.length > 2) || (depths.length == 2 && depths[0] - depths[1].abs > 1)
          return false
        end
      end
    else # this is not a leaf node (branch node), so push on its leaves
      if node.left
        nodes.push([node.left, depth + 1])
      end

      if node.right
        nodes.push([node.right, depth + 1])
      end
    end
  end

  true
end
