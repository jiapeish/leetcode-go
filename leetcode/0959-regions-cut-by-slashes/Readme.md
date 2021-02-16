
https://www.cnblogs.com/MrSaver/p/9607552.html#_label2

*** 并查集
function MakeSet(x)
	x.parent := x

function Find(x)
    if x.parent == x
        return x
    else
        return Find(x.parent)

function Union(x, y)
    xRoot := Find(x)
    yRoot := Find(y)
    xRoot.parent := yRoot
    
*** 优化一：按秩合并
funciont MakeSet(x)
    x.parent := x
    x.rank := 0
    
function Union(x, y)
    xRoot := Find(x)
    yRoot := Find(y)
    if xRoot == yRoot
        return
        
    // x和y不在同一个集合，进行合并
    if xRoot.rank < yRoot.rank
        xRoot.parent := yRoot
    elif xRoot.rank > yRoot.rank
        yRoot.parent := xRoot
    else
        yRoot.parent := xRoot
        xRoot.rank++
        
*** 优化二：路径压缩
function Find(x)
    if x.parent != x
        x.parent := Find(x.parent)
    return x.parent
        
        
        
        
        
        
        
        