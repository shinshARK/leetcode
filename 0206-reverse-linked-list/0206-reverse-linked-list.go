/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    current := head
    var prev *ListNode = nil
    for current != nil {
        next := current.Next
        // fmt.Println(next)
        current.Next = prev
        // fmt.Println(current.Next)
        prev = current
        // fmt.Println(prev)
        current = next
        // fmt.Println(current)
        // fmt.Println("-------")
    }
    return prev
}