# trees
Go package to implement a number of tree data structures.
*Note: A number of comparators are able to be passed into the trees' constructors. Please see tree.go in each package for more information.*

## Currently available
- Binary Search Tree

Example usage:
```go
import github.com/chancetudor/trees/bst

tree := bst.NewWithIntComparator()

returnedKey, err := tree.Insert(key, value)
exists := tree.Search(returnedKey)
newVal := tree.Update(returnedKey, newVal)
returnedVal, err := tree.ReturnNodeValue(returnedKey)
deletedKey, err := tree.Delete(returnedKey)
rootNode := tree.Root()
treeSize := tree.Size()
emptyFlag := tree.IsEmpty()
tree.Clear()
```

- Red-Black Tree

Example usage:
```go
import github.com/chancetudor/trees/bst

tree := bst.NewWithIntComparator()

returnedKey, err := tree.Insert(key, value)
exists := tree.Search(returnedKey)
newVal := tree.Update(returnedKey, newVal)
returnedVal, err := tree.ReturnNodeValue(returnedKey)
deletedKey, err := tree.Delete(returnedKey)
rootNode := tree.Root()
treeSize := tree.Size()
emptyFlag := tree.IsEmpty()
tree.Clear()
```

- AVL Tree

Example usage:
```go
import github.com/chancetudor/trees/bst

tree := bst.NewWithIntComparator()

returnedKey, err := tree.Insert(key, value)
exists := tree.Search(returnedKey)
newVal := tree.Update(returnedKey, newVal)
returnedVal, err := tree.ReturnNodeValue(returnedKey)
deletedKey, err := tree.Delete(returnedKey)
rootNode := tree.Root()
treeSize := tree.Size()
emptyFlag := tree.IsEmpty()
tree.Clear()
```

## In progress
- Trie
- Min heap
- Max heap
- Binomial heap
- Fibonacci heap
