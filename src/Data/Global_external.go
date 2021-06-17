package DataGlobal

import "gocurry"
import "curry2go/Prelude"

var global_map = make(map[string]*gocurry.Node)
var func_name = "arg_eval"

func ExternalData_Global_prim_globalT(task *gocurry.Task){
    root := task.GetControl()
    curry_name := root.GetChild(0)
    value := root.GetChild(1)
    name := gocurry.ReadString(curry_name)
    
    // test if an entry already exists
    _, ok := global_map[name]
    
    // initialize the entry
    if !ok{
        value = Prelude.Prelude__CREATE_DolHashHash(root.NewNode(), Prelude.Prelude__CREATE_id(root.NewNode()), value)
        global_map[name] = value
        
        // evaluate the entry
        gocurry.FuncCreate(root, func(task *gocurry.Task){
            DataGlobal__CREATE_GlobalT(task.GetControl(), curry_name)
        }, &func_name, 1, 0, value)
        return
    }
    
    // return GlobalT
    DataGlobal__CREATE_GlobalT(root, curry_name)
}

func ExternalData_Global_prim_readGlobalT(task *gocurry.Task){
    root := task.GetControl()
    global := root.GetChild(0)
    curry_name := global.GetChild(0)
    name := gocurry.ReadString(curry_name)
    
    // retrieve the value
    value, ok := global_map[name]
    
    if !ok{
        panic("Data.Global.readGlobalT: read on uninitialized entry '" + name + "'")
    }
    
    // return the value
    gocurry.IOCreate(root, value)
}

func ExternalData_Global_prim_writeGlobalT(task *gocurry.Task){
    root := task.GetControl()
    global := root.GetChild(0)
    curry_name := global.GetChild(0)
    name := gocurry.ReadString(curry_name)
    value := root.GetChild(1)
    
    // set entry to value
    global_map[name] = value
    
    // return
    gocurry.IOCreate(root, Prelude.Prelude__CREATE_LbRb(root.NewNode()))
}

