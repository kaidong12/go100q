package main

import "go100q/im_system"

func main() {

	//===============================
	//basic.Variable()

	//basic.Array_demo()
	//
	//basic.Slice_and_range_demo()

	//basic.Map_demo()

	//===============================
	//basic.Pointer()
	//a := 10
	//b := 20
	//basic.Swap(&a, &b)
	//fmt.Println(a, b)
	//
	//basic.Print_prime(100)
	//
	//basic.Recursion_demo()

	//===============================
	//basic.Struct_demo()
	//basic.Interface_demo()

	//===============================
	//basic.InterfaceDemo1_type_assertion()
	//basic.InterfaceDemo1_type_assertion_with_ok("select")

	//===============================
	//basic.TypeDemo1_value_receiver_vs_pointer_receiver()
	//basic.TypeDemo2_method_value_and_receiver()
	//basic.TypeDemo2_method_expression_and_receiver()
	//basic.TypeDemo2_method_selector_and_method_receiver()

	//===============================
	//basic.ReflectPair_pass_1()
	//basic.ReflectPair_pass_2()
	//basic.ReflectDemo_1()
	//basic.ReflectDemo_with_convert()
	//basic.ReflectDemo_unknown_type()

	//===============================
	//basic.StructTagDemo_1()
	//basic.Struct_to_json()

	//===============================
	//basic.GoroutineDemo1()
	//basic.GoroutineDemo2()
	//basic.GoroutineDemo3()
	//basic.GoroutineDemo4()

	//===============================
	//basic.GoroutineChannelDemo1()
	//basic.GoroutineChannelDemo2_no_buffer()
	//basic.GoroutineChannelDemo3_buffered()
	//basic.GoroutineChannelDemo4_close_chan()
	//basic.GoroutineChannelDemo4_range()
	//basic.GoroutineChannelDemo5_one_direction()
	//basic.GoroutineChannelDemo6_select_multiple_channel()

	server := im.NewServer("127.0.0.1", 8888)
	server.Start()

}
