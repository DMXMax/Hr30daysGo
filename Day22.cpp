/*NOTE: I really wanted to use the queue solution here, and the way Hackerrank locks code prohibited me until 
I found a way to write around it*/

#include <iostream>
#include <cstddef>

using namespace std;	


class Node{
    public:
        int data;
        Node *left;
        Node *right;
        Node(int d){
            data = d;
            left = NULL;
            right = NULL;
        }
};
class Solution{
    public:
  		Node* insert(Node* root, int data) {
            if(root == NULL) {
                return new Node(data);
            }
            else {
                Node* cur;
                if(data <= root->data){
                    cur = insert(root->left, data);
                    root->left = cur;
                }
                else{
                    cur = insert(root->right, data);
                    root->right = cur;
               }

               return root;
           }
        }
        
        int getHeight(Node* root);
};

#include <queue>
        
int Solution::getHeight(Node* root){
    queue<Node*> q;
    
    if( root == NULL)
        return 0;
    else
        q.push(root);
        
    
    int height(-1);
    int count(0); 
    
    while(1)
    {
        count = q.size();
        
        if(count == 0)
            return height;
        else
            height++;
        
        while( count > 0)
        {
            Node*p = q.front();
            if( p->left)
                q.push(p->left);
            if( p->right)
                q.push(p->right);
            count--;
            q.pop();
        }
        
        //return 1;
        
        
    }

    //return 0;
        
}
int test(){
    
    return 1;
    
