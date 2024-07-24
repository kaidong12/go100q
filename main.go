package main

import "go100q/examples/etree"

func main() {

	//===============================
	//basic.Variable()
	//basic.Bytes_demo()

	//basic.Array_demo()
	//
	//basic.Slice_and_range_demo()

	//basic.Map_demo()
	//basic.SlideWindow1()

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
	//net.Net_demo_To4()

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
	//basic.GoroutineChannelDemo3_buffered_int()

	//start := time.Now()
	//basic.GoroutineChannelDemo3_slide_window1()
	//basic.GoroutineChannelDemo3_slide_window2()
	//basic.GoroutineChannelDemo3_slide_window3()
	//elapsed := time.Since(start)

	//fmt.Printf("GoroutineChannelDemo3_buffered_bytes() took %s\n", elapsed)

	//basic.GoroutineChannelDemo4_close_chan()
	//basic.GoroutineChannelDemo4_range()
	//basic.GoroutineChannelDemo5_one_direction()
	//basic.GoroutineChannelDemo6_select_multiple_channel()

	//basic.GoroutineChannelDemo7_WithTimeout_ctx_done()
	//basic.GoroutineChannelDemo8_WithTimeout_call_cancel()

	//===============================
	//===============================
	//basic.Channel_cool_down_demo_1()

	//cooldownService := basic.NewCooldownService(10 * time.Second)
	//basic.Channel_cool_down_demo_2_stop(cooldownService)
	//basic.Channel_cool_down_demo_2_start(cooldownService)

	//===============================
	//===============================
	//server := im.NewServer("127.0.0.1", 8888)
	//server.Start()

	//===============================
	//===============================
	//===============================
	//etree.ParseXMLDemo1()
	//etree.ParseXMLDemo2()
	etree.ParseXMLDemo3()

	//===============================
	//===============================
	//===============================
	//gin_demo.PingPong()
	//gin_demo.RESTfulDemo()
	//gin_demo.LoadStaticPage()

	//gin_demo.GetParametersInRequests1()
	//gin_demo.GetParametersInRequests2()

	//gin_demo.GetJsonInRequests()
	//gin_demo.FormDemo_submit_form_data()
	//gin_demo.FormDemo_submit_form_to_new_page()

	//gin_demo.Redirect_to_another_page()
	//gin_demo.Route_group()
	//gin_demo.Redirect_to_404_page()

	//gin_demo.My_Handler()

	//===============================
	//===============================
	//===============================
	//gorm_demo.Sqlite_demo_1()
	//gorm_demo.Sqlite_demo_2()
	//gorm_demo.Sqlite_demo_3()
}
