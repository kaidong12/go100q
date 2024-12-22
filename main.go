package main

import (
	"fmt"
	"go100q/examples/basic"
)

func main() {

	//===============================
	//basic.Variable()
	//basic.Bytes_demo()

	//basic.Array_demo()
	//basic.Slice_and_range_demo()
	//basic.Slice_reverse_demo()
	//basic.Slice_compare_demo()
	//basic.Slice_append_demo()
	//basic.Slice_nonempty_demo()
	//basic.Slice_remove_demo()

	//basic.Map_demo()
	//basic.SlideWindow1()

	//===============================
	basic.Pointer()
	a := 10
	b := 20
	basic.Swap(&a, &b)
	fmt.Println(a, b)
	//
	//basic.Print_prime(100)
	//
	//basic.Recursion_demo()

	//===============================
	//===============================
	//===============================
	//basic.IO_reader_demo_1()
	//basic.IO_reader_demo_2()
	//basic.IO_reader_demo_3()
	//basic.IO_reader_demo_4_read()
	//basic.IO_reader_demo_4_write()
	//basic.IO_reader_demo_5(gin.Context{})

	//===============================
	//basic.Struct_demo_0()
	//basic.Struct_demo_1()
	//basic.Struct_demo_2()
	//basic.Struct_demo_3()
	//basic.Struct_to_json()
	//basic.StructTagDemo_1()
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

	//basic.GoroutineChannelDemo4_close_chan_read()
	//basic.GoroutineChannelDemo4_close_chan_range()

	//basic.GoroutineChannelDemo5_one_direction()
	//basic.GoroutineChannelDemo6_select_multiple_channel()

	//basic.GoroutineChannelDemo7_WithTimeout_ctx_done()
	//basic.GoroutineChannelDemo7_WithTimeout_call_cancel()
	//basic.GoroutineChannelDemo8_select_timeout()

	//===============================
	//===============================
	//basic.GoroutineChannelDemo9_cool_down_demo_1_start()
	//basic.GoroutineChannelDemo9_cool_down_demo_1_wait()

	//cooldownService := basic.NewCooldownService(10 * time.Second)
	//basic.GoroutineChannelDemo9_cool_down_demo_2_stop(cooldownService)
	//basic.GoroutineChannelDemo9_cool_down_demo_2_start(cooldownService)

	//basic.GoroutineChannelDemo9_2_cool_down_demo_1_start()
	//basic.GoroutineChannelDemo9_2_cool_down_demo_1_wait()

	//===============================
	//===============================
	//server := im.NewServer("127.0.0.1", 8888)
	//server.Start()

	//===============================
	//===============================
	//===============================
	//etree.ParseXMLDemo1()
	//etree.ParseXMLDemo2()
	//etree.ParseXMLDemo3()

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

	//===============================
	//===============================
	//===============================
	// function.Function_demo_1()
	// pointer.Pointer_demo_1()

}
