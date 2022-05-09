package main

import (
	"fmt"
	"github.com/sebge2/test/application"
	"github.com/sebge2/test/channels"
	"github.com/sebge2/test/concurrency"
	"github.com/sebge2/test/controlFlow"
	"github.com/sebge2/test/error"
	"github.com/sebge2/test/exampleDoc"
	"github.com/sebge2/test/fundamentals"
	"github.com/sebge2/test/groupingData"
	"github.com/sebge2/test/pointer"
	"github.com/sebge2/test/test"
	"github.com/sebge2/test/variableValueType"
)
import "github.com/sebge2/test/functions"
import "github.com/sebge2/test/struct"

func main() {
	printCategory("Variables, values & Type - FMT package")
	variableValueType.FmtPackageExample()

	printCategory("Variables, values & Type - Own TYpe")
	variableValueType.OwnTypeExample()

	printCategory("Variables, values & Type - Conversion")
	variableValueType.ConversionExample()

	// TODO ninja

	printCategory("Programming Fundamentals - Constants")
	fundamentals.ConstantsExample()

	printCategory("Programming Fundamentals - Iota")
	fundamentals.IotaExamples()

	printCategory("Programming Fundamentals - Bit Shifting")
	fundamentals.BitShiftingExample()

	// TODO ninja

	printCategory("Control Flow - Loop - Init, condition, post")
	controlFlow.LoopInitConditionPostExample()

	printCategory("Control Flow - Loop - Nested Loop")
	controlFlow.LoopNestedLoop()

	printCategory("Control Flow - Loop - For Statement")
	controlFlow.LoopForStatement()

	printCategory("Control Flow - Conditional - If, else if, else")
	controlFlow.ConditionalIfElseIfElse()

	printCategory("Control Flow - Switch")
	controlFlow.SwitchExample()

	// TODO ninja

	printCategory("Grouping Data - Array")
	groupingData.ArrayExample()

	printCategory("Grouping Data - Slice for Range")
	groupingData.SliceForRangeExample()

	printCategory("Grouping Data - Slicing a Slice")
	groupingData.SliceSlicingExample()

	printCategory("Grouping Data - Slice Append")
	groupingData.SliceAppendExample()

	printCategory("Grouping Data - Slice Delete")
	groupingData.SliceDeleteExample()

	printCategory("Grouping Data - Slice Make")
	groupingData.SliceMakeExample()

	printCategory("Grouping Data - Slice Multi-Dimensional")
	groupingData.SliceMultiDimensionalExample()

	printCategory("Grouping Data - Map")
	groupingData.MapExample()

	printCategory("Grouping Data - Map Add and Range")
	groupingData.MapAddAndRangeExample()

	printCategory("Grouping Data - Map Delete")
	groupingData.MapDeleteExample()

	// TODO ninja

	//x := map[string][]string{
	//	"James": {"toto", "tata"},
	//	"Toto":  {"txxoto", "taxxta"},
	//}
	//
	//x["tata"] = []string{"titi", "toto"}
	//
	//delete(x, "tata")
	//
	//for k, v := range x {
	//	fmt.Println("This is record for", k)
	//
	//	for _, v2 := range v {
	//		fmt.Printf("%v ", v2)
	//	}
	//
	//	fmt.Print("\n")
	//}

	//
	//p1 := secretAgent{
	//	person: person{
	//		firstname: "James",
	//		lastname:  "Bond",
	//	},
	//	rightToKill: true,
	//}
	//
	//fmt.Println(p1.firstname)
	//
	//p2 := struct {
	//	firstname string
	//	lastname  string
	//}{
	//	firstname: "James",
	//	lastname:  "Bond",
	//}
	//
	//fmt.Println(p2.firstname)

	//
	//x := map[string]person{
	//	p1.lastname: p1,
	//}
	//
	//for k, v := range x {
	//	fmt.Println(k, v)
	//}

	printCategory("Struct - Anonymous")
	_struct.AnonymousExample()

	printCategory("Struct - Embedded")
	_struct.EmbeddedExample()

	printCategory("Struct - Ninja Level 05 - 01")
	_struct.Ninja01()

	printCategory("Struct - Ninja Level 05 - 02")
	_struct.Ninja02()

	printCategory("Struct - Ninja Level 05 - 03")
	_struct.Ninja03()

	printCategory("Struct - Ninja Level 05 - 04")
	_struct.Ninja04()

	printCategory("Functions - Defer")
	functions.DeferExample()

	printCategory("Functions - Method")
	functions.MethodExample()

	printCategory("Functions - Interface")
	functions.InterfaceExample()

	printCategory("Functions - Switch Type")
	functions.SwitchTypeExample()

	printCategory("Functions - Anonymous")
	functions.AnonymousFunctionExample()

	printCategory("Functions - Expression")
	functions.FuncExpressionExample()

	printCategory("Functions - Return Function")
	functions.ReturnFuncExample()

	printCategory("Functions - Callback")
	functions.CallbackExample()

	printCategory("Functions - Ninja Level 06 - 01")
	functions.Ninja01()

	printCategory("Functions - Ninja Level 06 - 02")
	functions.Ninja02()

	printCategory("Functions - Ninja Level 06 - 03")
	functions.Ninja03()

	printCategory("Functions - Ninja Level 06 - 04")
	functions.Ninja04()

	printCategory("Functions - Ninja  Level 06 - 05")
	functions.Ninja05()

	printCategory("Functions - Ninja  Level 06 - 06")
	functions.Ninja06()

	printCategory("Functions - Ninja  Level 06 - 07")
	functions.Ninja07()

	printCategory("Functions - Ninja  Level 06 - 08")
	functions.Ninja08()

	printCategory("Functions - Ninja  Level 06 - 09")
	functions.Ninja09()

	printCategory("Functions - Ninja  Level 06 - 10")
	functions.Ninja10()

	printCategory("Pointers - Pointers")
	pointer.Pointer()

	printCategory("Pointers - When to use Pointers?")
	pointer.PointerWhenToUseExample()

	printCategory("Pointers - Method Set")
	pointer.PointerMethodSetExample()

	printCategory("Pointers - Ninja  Level 07 - 01")
	pointer.Ninja01()

	printCategory("Pointers - Ninja  Level 07 - 02")
	pointer.Ninja02()

	printCategory("Application - JSON")
	application.JsonExample()

	printCategory("Application - Writer Interface")
	application.WriterInterfaceExample()

	printCategory("Application - Sort")
	application.SortExample()

	printCategory("Application - Sort Custom")
	application.SortCustomExample()

	printCategory("Application - Bcrypt")
	application.BcryptExample()

	printCategory("Pointers - Ninja  Level 08 - 01")
	application.Ninja01()

	printCategory("Pointers - Ninja  Level 08 - 02")
	application.Ninja02()

	printCategory("Pointers - Ninja  Level 08 - 03")
	application.Ninja03()

	printCategory("Pointers - Ninja  Level 08 - 04")
	application.Ninja04()

	printCategory("Pointers - Ninja  Level 08 - 05")
	application.Ninja05()

	printCategory("Concurrency - WaitGroup")
	concurrency.WaitGroupExample()

	printCategory("Pointers - Ninja  Level 09 - 01")
	concurrency.Ninja01()

	printCategory("Pointers - Ninja  Level 09 - 02")
	concurrency.Ninja02()

	printCategory("Pointers - Ninja  Level 09 - 03")
	concurrency.Ninja03()

	printCategory("Pointers - Ninja  Level 09 - 04")
	concurrency.Ninja04()

	printCategory("Pointers - Ninja  Level 09 - 05")
	concurrency.Ninja05()

	printCategory("Pointers - Ninja  Level 09 - 06")
	concurrency.Ninja06()

	printCategory("Channels - Buffered")
	channels.BufferedExamples()

	printCategory("Channels - Directional")
	channels.DirectionalExamples()

	printCategory("Channels - Range")
	channels.RangeExample()

	printCategory("Channels - Select")
	channels.SelectExample()

	printCategory("Channels - Comma OK idiom")
	channels.CommaOkIdiomExample()

	printCategory("Channels - Fan IN")
	channels.FanInExample()

	printCategory("Channels - Context")
	channels.ContextExample()

	printCategory("Channels - Ninja  Level 10 - 01")
	channels.Ninja01()

	printCategory("Channels - Ninja  Level 10 - 02")
	channels.Ninja02()

	printCategory("Channels - Ninja  Level 10 - 03")
	channels.Ninja03()

	printCategory("Channels - Ninja  Level 10 - 04")
	channels.Ninja04()

	printCategory("Channels - Ninja  Level 10 - 05")
	channels.Ninja05()

	printCategory("Channels - Ninja  Level 10 - 06")
	channels.Ninja06()

	printCategory("Channels - Ninja  Level 10 - 07")
	channels.Ninja07()

	printCategory("Error Handling - Checking Errors")
	error.CheckingErrorsExample()

	printCategory("Error Handling - Printing & Logging")
	error.PrintingLoggingExamples()

	printCategory("Error Handling - Recover")
	error.RecoverExamples()

	printCategory("Error Handling - Error Info")
	error.ErrorInfoExample()

	printCategory("Error Handling - Ninja  Level 11 - 01")
	error.Ninja01()

	printCategory("Error Handling - Ninja  Level 11 - 02")
	error.Ninja02()

	printCategory("Error Handling - Ninja  Level 11 - 03")
	error.Ninja03()

	printCategory("Error Handling - Ninja  Level 11 - 04")
	error.Ninja04()

	printCategory("Error Handling - Ninja  Level 11 - 05")
	error.Ninja05()

	printCategory("Documentation - Go Doc")
	exampleDoc.GoDocExample()

	printCategory("Documentation - Ninja Level 12 - 01")
	exampleDoc.Ninja01()

	printCategory("Testing & Benchmarking - Ninja Level 13 - 01")
	test.Ninja01()

	printCategory("Testing & Benchmarking - Ninja Level 13 - 02")
	test.Ninja02()

	printCategory("Testing & Benchmarking - Ninja Level 13 - 03")
	test.Ninja03()

	//myFigure := figure.NewColorFigure("My Lovely Lady", "", "red", true)
	//myFigure.Print()
}

func printCategory(category string) {
	fmt.Printf(
		"\n%s\n%v\n",
		category,
		func() string {
			v := ""

			for i := 0; i < len(category); i++ {
				v += "="
			}
			return v
		}(),
	)
}

func toto() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(fmt.Sprintf("i = %d", i))
		}
	}
}

func varArg() {
	val := []string{"x", "y"}

	function(val...)
}

func function(test ...string) (string, string) {
	return "a", "b"
}
