void levelOrder(Node * root){
        queue<Node*> q;
    
    if( root == NULL)
        return;
    else
        q.push(root);
        
    
    int height(-1);
    int count(0); 
    
    while(1)
    {
        count = q.size();
        
        if(count == 0)
            return;
        
        while( count > 0)
        {
            Node*p = q.front();
            cout << p->data <<" ";
            if( p->left)
                q.push(p->left);
            if( p->right)
                q.push(p->right);
            count--;
            q.pop();
        }
        
        
        
        
    }
  	
  
	}
