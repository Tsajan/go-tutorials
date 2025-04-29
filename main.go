package main

import (
	"fmt"
	"tutorials/tutorials"
)

func main() {
	fmt.Println("This is main")

	fmt.Println("---------Calling fn Hello() from tutorials package!---------")
	tutorials.Hello()

	fmt.Println("--------Calling fn VarDeclare() from tutorials package!---------")
	tutorials.VarDeclare()

	fmt.Println("--------Calling fn DefaultVariableValues() from tutorials package!---------")
	tutorials.DefaultVariableValues()

	fmt.Println("---------Calling fn VariableReassignment() from the tutorials package!--------")
	tutorials.VariableReassignment()

	fmt.Println("---------Calling fn MultiVariablesAssignment() from the tutorials package!--------")
	tutorials.MultiVariablesAssignment()

	fmt.Println("--------Calling fn FnCallWithPointers() from tutorials package!---------")
	tutorials.FnCallWithPointers()

	fmt.Println("---------Calling fn StructDemo() from the tutorials package!---------")
	tutorials.StructsDemo()

	fmt.Println("---------Calling fn DeferExample() from the tutorials package!--------")
	tutorials.DeferExample()

	fmt.Println("---------Calling fn ArrayExample() from the tutorials package!--------")
	tutorials.ArrayExample()

	fmt.Println("---------Calling fn SliceDemo() from the tutorials package!--------")
	tutorials.SliceDemo()

	fmt.Println("--------Calling fn FnCallWithNamedReturns() from tutorials package!---------")
	tutorials.FnCallWithNamedReturns()

	fmt.Println("--------Calling fn FnCallsWithSlice() from tutorials package!---------")
	tutorials.FnCallsWithSlice()

	fmt.Println("---------Calling fn MultiDimensionalSlice() from the tutorials package!--------")
	tutorials.MultiDimensionalSlice()

	fmt.Println("---------Calling fn FnCallByValue() from the tutorials package!--------")
	tutorials.FnCallByValue()

	fmt.Println("---------Calling fn FnCallWithArray() from the tutorials package!--------")
	tutorials.FnCallWithArray()

	fmt.Println("--------Calling fn FnCallsWithArrayPointers() from tutorials package!---------")
	tutorials.FnCallsWithArrayPointers()

	fmt.Println("---------Calling fn FnCallWithArray() from the tutorials package!--------")
	// tutorials.InfiniteForLoop() // commenting out because I don't want to be stuck in this infinite loop

	fmt.Println("---------Calling fn FnAdd() from the tutorials package!--------")
	tutorials.FnAdd()

	fmt.Println("---------Calling fn SwitchDemo() from the tutorials package!--------")
	tutorials.SwitchDemo()

	fmt.Println("---------Calling fn SwitchWithFallThrough() from the tutorials package!--------")
	tutorials.SwithWithFallThrough()

	fmt.Println("---------Calling fn VariadicAddDemo() from the tutorials package!--------")
	tutorials.VariadicAddDemo()

	fmt.Println("---------Calling fn LogicalOperatorsDemo() from the tutorials package!--------")
	tutorials.LogicalOperatorsDemo()

	fmt.Println("---------Calling fn MapsDemo() from the tutorials package!--------")
	tutorials.MapsDemo()

	fmt.Println("---------Calling fn MatrixMultiplication() from the tutorials package!--------")
	tutorials.MatrixMultiplication()

	fmt.Println("---------Calling fn MatrixScalarProduct() from the tutorials package!--------")
	tutorials.MatrixScalarProduct()

	fmt.Println("---------Calling fn StructExportsDemo() from the tutorials package!-----------")
	tutorials.StructExportsDemo()
	// also we can initlize the Bicycle struct that was exported, but we don't get the price
	b2 := tutorials.Bicycle{Model: "Trek Excalibur", ReleaseYear: 2024}
	fmt.Println("Exported Bicycle b2 is: ", b2)

	// structs are passed by value (unless of course we use reference operators)
	// so changing a value of an element of original struct object doesn't essentially alter the original struct
	b3 := b2
	b3.Model = "Trek Excalibur II"

	fmt.Println("Original struct b2 is still: ", b2)
	fmt.Println("Copied and updated struct b3 is: ", b3)

	fmt.Println("--------------Calling fn AnonymousFnCalls() from the tutorials package!-------------")
	tutorials.AnonymousFnCall()

	fmt.Println("Separately just printing ExportedVar within the AnonymousFnCall()")
	fmt.Println("ExportedVar circum-navigated is: ", tutorials.ExportedVar)

	fmt.Println("------------Calling Fn ClosureDemo() from the tutorials package!---------------")
	tutorials.ClosureDemo()

	fmt.Println("------------Calling Fn MethodsDemo() from the tutorials package!---------------")
	tutorials.MethodsDemo()

	fmt.Println("------------Calling Fn MethodsWithFnAsReceiversDemo() from the tutorials package!-----------")
	tutorials.MethodsWithFnAsReceiversDemo()

	fmt.Println("------------Calling Fn ObjectEmbeddingDemo() from the tutorials package!---------------")
	tutorials.ObjectEmbeddingDemo()

	fmt.Println("------------Calling Fn computeStableMatching() from the tutorials package!-------")
	tutorials.ComputeStableMatching()

}
