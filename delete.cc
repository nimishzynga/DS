
void deleteNode(node *r, int v) {
    node *root = r;
    node *p = r;
    while(1) {
        if (r == NULL) {
            return;
        }
        if (r->val > v) {
            p = r;
            r = r->left;
        } else {
            if (r->val == v) {
                break;
            }
        } else {
            p = r;
            r = r->right;
        }
    }
    if (r->left && r->right) {
        node *t= r->left;
        while (t->right) {
            t = t->right;
        }
        int g = t->val;
        deleteNode(root, t->val); 
        r->val = g; 
    } else if (r->left) {
        if (p->right == r) {
            p->right = r->left;
        } else {
            p->left = r->left;
        } 
        free(r);
    } else if (r->right) {
         if (p->right == r) {
            p->right = r->right;
        } else {
            p->left = r->right;
        } 
        free(r);
    } else {
        if (p->left == r) {
            p->left = NULL;
        } else {
            p->right = NULL;
        }
        free(r);
    }
}


