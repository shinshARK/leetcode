/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    new := &ListNode{}
    currNew := new
    curr1 := list1
    curr2 := list2
    for curr1 != nil && curr2 != nil {
        if curr1.Val < curr2.Val {
            currNew.Next = curr1
            curr1 = curr1.Next
        } else {
            currNew.Next = curr2
            curr2 = curr2.Next
        }
        currNew = currNew.Next
    }
    if curr1 != nil {
        currNew.Next = curr1
    } else if curr2 != nil {
        currNew.Next = curr2
    }
    fmt.Println(new.Next)
    return new.Next
}