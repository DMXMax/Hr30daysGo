          Node* removeDuplicates(Node *head)
          {
            Node* pcurNode = head;
            int curr; 
              
            if( pcurNode == NULL)
                return head;
            else
            {
                curr = head->data;
            } 
              
            while( pcurNode->next != NULL )
            {
                if( pcurNode->next->data == curr) //delete
                {
                    pcurNode->next = pcurNode->next->next;
                    
                }
                else
                {
                    pcurNode = pcurNode->next;
                    curr = pcurNode->data;
                }
            }
              return head;
          }
