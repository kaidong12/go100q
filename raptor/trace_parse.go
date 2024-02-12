package decoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
	"gopkg.in/yaml.v3"
)

func readCodebook(file_path string) map[string]interface{} {
	map_codebook := make(map[string]interface{})
	file_name := filepath.Base(file_path)
	yaml_file := filepath.Join(filepath.Dir(file_path), strings.TrimSuffix(file_name, filepath.Ext(file_name))+".yaml")

	slog.Info("Convert codebook from xml to YAML file", "YAML file path:", yaml_file)

	if _, err := os.Stat(yaml_file); err == nil {
		slog.Info("YAML file already exists!")

		data, err := os.ReadFile(yaml_file)
		if err != nil {
			slog.Error("Error reading YAML file: %v", err)
		}
		// Parse the YAML data
		err = yaml.Unmarshal(data, &map_codebook)
		if err != nil {
			slog.Error("Error unmarshalling YAML data: %v", err)
		}
		// Use the map
		// fmt.Printf("YAML data as map:\n%+v\n", map_codebook)
	} else {
		// Read XML file content
		raw_codebook := etree.NewDocument()
		if err := raw_codebook.ReadFromFile(file_path); err != nil {
			slog.Error("Error reading file:", err)
			// panic(err)
		}

		root := raw_codebook.SelectElement("codebook")
		fmt.Println("ROOT element:", root.Tag)

		module_map := make(map[string]string)
		for _, module := range root.SelectElement("content").SelectElements("module") {
			fmt.Println("CHILD element:", module.Tag)
			tag := module.SelectElement("tag")
			id := module.SelectElement("id")
			if id != nil && tag != nil {
				slog.Info("module:", "module", tag.Text())
				module_map[tag.Text()] = id.Text()
			}
		}

		for _, single_module := range root.SelectElements("single_module_description") {
			trace_slice := []map[string]string{}
			for _, trace := range single_module.SelectElements("trace") {
				if line := trace.SelectElement("line"); line != nil {
					// fmt.Printf("      line: %s\n", line.Text())
					if str := trace.SelectElement("str"); str != nil {
						// fmt.Printf("      str: %s\n", str.Text())
						trace_slice = append(trace_slice, map[string]string{strings.TrimSpace(line.Text()): str.Text()})
					}
				}
			}

			trace_map := make(map[string]interface{})
			if module_name := single_module.SelectElement("module_identification").SelectElement("module_name"); module_name != nil {
				trace_map["name"] = module_name.Text()
				if trace_slice != nil {
					trace_map["trace"] = trace_slice
					map_codebook[module_map[module_name.Text()]] = trace_map
				}
			}
		}

		// Convert map to YAML
		yamlContent, err := yaml.Marshal(&map_codebook)
		if err != nil {
			slog.Error("Error marshaling yaml file:", err)
		}

		// Write YAML to file
		err = os.WriteFile(yaml_file, yamlContent, 0644)
		if err != nil {
			slog.Error("Error writing yaml file:", err)
		}
	}

	return map_codebook
}

func iterateMap(data map[string]interface{}) {
	for key, value := range data {
		switch value.(type) {
		case map[interface{}]interface{}:
			fmt.Printf("Key: %vn", key)
			iterateMap(value.(map[string]interface{}))
		default:
			fmt.Printf("Key: %v, Value: %v\n", key, value)
		}
	}
}

func decoder() {

	// codebook_ap := "codebook_ap_s.xml"
	codebook_ap := "../testdata/codebook_o4xmebsop2cgbvxxrel_ap.xml"
	codebook_cp := "../testdata/codebook_o4xmebsop2cgbvxxrel_cp.xml"
	raw_ecu_trace := "../testdata/ecu_trace.bbb"
	output_file := "../testdata/output.txt"

	codebooks := make(map[string]map[string]interface{})

	codebooks["AP"] = readCodebook(codebook_ap)
	codebooks["CP"] = readCodebook(codebook_cp)

	// iterateMap(readCodebook(codebook_ap))
	// fmt.Printf("readCodebook:\n%+v\n", readCodebook(codebook_ap)["0"].(map[string]interface{})["trace"].([]interface{})[0].(map[string]interface{})["163"])
	fmt.Printf("readCodebook:\n%+v\n", codebooks["AP"]["0"].(map[string]interface{})["trace"].([]interface{})[0].(map[string]interface{})["163"])

	// =====================================================================

	rawFile, err := os.ReadFile(raw_ecu_trace)
	if err != nil {
		slog.Error("Error reading file:", err)
	}

	packetSize := 1024
	packets := len(rawFile)/packetSize + 1
	slog.Info("packets number\n", "total packets:", packets)

	var remain []byte
	// var out [][]byte
	var out []map[string]string

	for i := 400; i < 700; i++ {
		// fmt.Printf("=================remain addr 1, %d=================\n", &remain)
		// fmt.Println("remain1111111111: \n", string(remain))
		rawBytes := append(remain, rawFile[i*packetSize:(i+1)*packetSize]...)
		remain, out = lgeFullReading(rawBytes, codebooks)
		// fmt.Printf("=================remain addr 2, %d=================\n", &remain)
		if len(out) > 0 {
			slog.Info("=================remain, out=================")
			slog.Info("remain:\n", "remain:", string(remain))
			slog.Info("out:\n", "out:", out)

			file, err := os.OpenFile(output_file, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			for _, msg := range out {
				for key, value := range msg {
					_, err = file.WriteString(fmt.Sprintf("%s: %s\n", key, value))
					if err != nil {
						fmt.Println("Error writing to file:", err)
						return
					}
				}
				_, err = file.WriteString("\n")
			}
		}
	}
}

func lgeFullReading(rawBytes []byte, codebooks map[string]map[string]interface{}) ([]byte, []map[string]string) {
	//		var out []map[string]interface{}
	// func lgeFullReading(rawBytes []byte) ([]byte, [][]byte) {
	slice_raw_trace := rawBytes

	// fmt.Printf("=================length=>%d=============\n", length)
	header_start := []byte{0x80, 0x4d, 0x4c, 0x50}
	header_end := []byte{0x06, 0x00}
	// out := [][]byte{}
	var out []map[string]string
	remain := []byte{}

	// length actually is index of the last byte in slice
	length := len(slice_raw_trace) - 1
	idx1 := bytes.Index(slice_raw_trace, header_start)

	if idx1 >= 4 { // found header_start in packet and message lengh is not truncated, only for the beginning of the whole packet
		for {
			// length actually is index of the last byte in slice
			length := len(slice_raw_trace) - 1
			idx1 = bytes.Index(slice_raw_trace, header_start)

			if idx1 >= 4 { // found header_start in slice_raw_trace and message lengh is not truncated
				idx2 := bytes.Index(slice_raw_trace, header_end)
				if idx2 > 0 {
					if idx2-idx1 == 16 { // valid header
						slog.Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
						slog.Info("Index of header start and end\n", "header_start index", idx1, "header_end index", idx2)
						remain = slice_raw_trace[length-idx2:]

						message_len := binary.BigEndian.Uint32(slice_raw_trace[idx1-4 : idx1])
						if length >= idx2+5 { // can get payload_len successfully
							payload_len := binary.BigEndian.Uint32(slice_raw_trace[idx2+2 : idx2+6])

							if message_len == payload_len+22 { // valid payload length
								message_end := idx1 + int(message_len)
								slog.Info("Valid message length and payload length\n", "message_len", message_len, "payload_len", payload_len)
								if length >= message_end { // Payload not truncated
									payload := slice_raw_trace[idx2+6 : idx2+6+int(payload_len)]
									slog.Info("Get Payload successfully\n", "Payload", payload)

									trace_type_id := int(payload[1])
									trace_collector_id := int(payload[2])

									trace_type := traceTypeFromID(trace_type_id)
									trace_collector := collectorFromID(trace_collector_id)

									if trace_type == "NORMAL_TRACE" && (trace_collector == "AP" || trace_collector == "CP") {
										slog.Info("Parse payload with decode_payload start!!!")
										out = decode_payload(payload[20:], trace_collector, codebooks)

										// out = decode_payload(payload[20:], trace_collector)

										// siteMap := make(map[string]string)
										// siteMap[string(payload[10:15])] = string(payload)
										// out = append(out, siteMap)

										slog.Info("Parse payload with decode_payload done!!!")
									}

									slice_raw_trace = slice_raw_trace[message_end+1:]

								} else { // Payload been truncated
									slog.Warn("Payload has been trunctated, suggest to increase packet size!")
									remain = slice_raw_trace[idx1-4:]
									break

								}

							} else { //invalid payload length
								// keep on searching on the right side of current index of the slice
								slog.Error("Message_length != Payload_length + Header_length\n", "Message length", message_len, "Payload length", payload_len)
								slice_raw_trace = slice_raw_trace[idx2+5:]

							}

						} else { // payload_len been trunctated
							slog.Warn("payload_len has been trunctated!")
							remain = slice_raw_trace[length-25:]
							break

						}

					} else { //invalid idx2
						// index of header_end is invalid:
						// a, idx2 < idx1
						// b, idx2 - idx1 != 16
						// keep on searching on the right side of current index of the slice
						slog.Error("Index of header_end is invalid!\n", "header_start index", idx1, "header_end index", idx2)
						slice_raw_trace = slice_raw_trace[max(idx1, idx2):]

					}

				} else {
					// no header_end match at all
					// a, it may be trancated
					// b, no match at all
					slog.Warn("No header_end found in current slice_raw_trace!")
					remain = slice_raw_trace[length-20:]
					break

				}
			} else { // no header_start found in slice_raw_trace
				slog.Info("**In for loop** No header_start found in current slice_raw_trace! **In for loop**")
				remain = slice_raw_trace[length-3:]
				break

			}

		}

	} else {
		// no header_start match at all
		slog.Info("**Out of loop** No header_start found in current rawBytes packet! **Out of loop**")
		remain = slice_raw_trace[length-3:]

	}

	return remain, out

}

func decode_payload(payload []byte, collector string, codebooks map[string]map[string]interface{}) []map[string]string {
	var out_slice_of_map []map[string]string

	for {
		var trace_map map[string]string
		timeGap := int(payload[0])
		traceLen := payload[1]
		continues := false
		var timeArea []byte
		var component string
		var traceLevel string
		var trace string

		if timeGap == 3 {
			timeArea = payload[2:10]
			payload = payload[10:]
		} else if timeGap == 2 {
			timeArea = payload[2:6]
			payload = payload[6:]
		} else {
			slog.Error("Invalid TimeGap:\n", "TimeGap", timeGap)
			break
		}

		if len(payload) >= int(traceLen) {
			if collector == "JAVA" || collector == "SYSLOG" {
				// component, traceLevel, payload := parserJavaSyslog(X, traceLen, tc)
				continues = false
			} else if collector == "AP" || collector == "CP" {
				// component, traceLevel, trace := parserApCp(payload, int(traceLen), collector, MODULE_NAME_BY_ID, TRACE_BY_MODULE_NAME_AND_LINE_NUMBER)
				component, traceLevel, trace = parserApCp(payload, int(traceLen), collector, codebooks)
				continues = false
				slog.Info("In decode_payload\n", "component", component, "traceLevel", traceLevel, "trace", trace)
			} else if collector == "DMESG" {
				// component, traceLevel, payload, continues := parserDmesg(X, traceLen, tc)
			}

			if continues {
				if out_slice_of_map == nil || len(out_slice_of_map[len(out_slice_of_map)-1]) < 4 {
					slog.Error("Bad trace_slice_of_slice:\n", "trace_slice_of_slice", out_slice_of_map, "trace", trace)
				} else {
					out_slice_of_map[len(out_slice_of_map)-1]["Trace"] += trace // if it is a continued trace, it will be appended to the previous trace
				}
			} else {
				// [
				// [10651719, 'MPONENT_ERROR_310', 'Info', 'VAL Msg TX: msgId:0x1a00    recvID:0x001a Opts:0x00000000-0x00000020-0x00000000 data:00.00.00.64  00.00.00.08', 2],
				// [637396, 'MPONENT_ERROR_494', 'Info', '[ ECU-D ] [ ERROR ] Filling wrong: AP_CODEBOOK_ERROR_494_325 [0]', 2]
				// ]
				var currentTime int64
				if timeGap == 3 {
					currentTime = 0
				}
				currentTime += timeStampComp(timeArea)
				slog.Info("In decode_payload\n", "currentTime", currentTime)
				trace_map = make(map[string]string)
				trace_map["DateTime"] = formatTime(currentTime, collector)
				trace_map["Component"] = component
				trace_map["TraceLevel"] = traceLevel
				trace_map["Trace"] = trace

				out_slice_of_map = append(out_slice_of_map, trace_map)
				payload = payload[traceLen:]
			}
		} else {
			slog.Error("Bad payload:\n", "payload", payload, "traceLen", traceLen)
			break
		}
	}

	return out_slice_of_map
}

func parserApCp(payload []byte, length int, collector string, codebooks map[string]map[string]interface{}) (string, string, string) {
	var rowID string
	traceStyle, num := traceStyleComp(payload[0])
	traceLevel := traceLevelComp(4 + (int(payload[0]) & 3))
	compID := strconv.Itoa(int(payload[3])*4 + (int((payload[2] >> 6)) & 3))
	if collector == "CP" {
		rowID = strconv.Itoa(int(payload[1])+(int(payload[2])&63)*256) + ">"
	} else {
		rowID = strconv.Itoa(int(payload[1]) + (int(payload[2])&63)*256)
	}

	slog.Info("[ ECU-D ] [ parserApCp ] Trace debug info!", "traceStyle", traceStyle)
	slog.Info("[ ECU-D ] [ parserApCp ] Trace debug info!", "rowID", rowID)
	slog.Info("[ ECU-D ] [ parserApCp ] Trace debug info!", "compID", compID)
	slog.Info("[ ECU-D ] [ parserApCp ] Trace debug info!", "traceLevel", traceLevel)

	// fmt.Printf("readCodebook:\n%+v\n", readCodebook(codebook_ap)["0"].(map[string]interface{})["trace"].([]interface{})[0].(map[string]interface{})["163"])

	var component string
	var trace string

	if traceStyle == "" {
		slog.Error("[ ECU-D ] [ ERROR ] [ TODO ] traceStyle null!")
		component = "NONE_NO_COMPONENT"
		traceLevel = "NO_TRACE_LEVEL"
		trace = "NO_PAYLOAD"
	} else if traceStyle == "RAW_OLD" {
		slog.Error("[ ECU-D ] [ ERROR ] [ TODO ] traceStyle RAW_OLD!")
		// component = self.__listToChr(X[2:traceLen]).split(" ")[0]
		// traceLevel = self.__traceLevelComp(X[1])
		// trace = " ".join(self.__listToChr(X[2:traceLen]).split(" ")[1:])
	} else if traceStyle == "RAW" {
		slog.Error("[ ECU-D ] [ ERROR ] [ TODO ] traceStyle RAW!")
		// rawLen := X[4]
		// if rawLen != traceLen-5 {
		// 	slog.Error("[ ECU-D ] [ ERROR ] Trace Len Mismatch!")
		// 	trace = ""
		// } else {
		// 	trace = self.__listToChr(X[5:traceLen])
		// }
	} else {
		slog.Info("[ ECU-D ] [ ERROR ] [ TODO ] traceStyle CODE!")
		MODULE_NAME_BY_ID := codebooks[collector]

		if _, ok := MODULE_NAME_BY_ID[compID]; ok {
			slog.Info("[ ECU-D ] [ ERROR ] [ TODO ] MODULE_NAME_BY_ID[compID]")
			component = MODULE_NAME_BY_ID[compID].(map[string]interface{})["name"].(string)
			slog.Info("[ ECU-D ] [ parserApCp ] MODULE_NAME_BY_ID[compID]", "component", component)
		}

		if component != "" {
			slog.Info("[ ECU-D ] [ parserApCp ] TRACE_BY_MODULE_NAME_AND_LINE_NUMBER[rowID]", "component", component)
			payloadInts := ComputePayloadInts(payload, num)
			TRACE_BY_MODULE_NAME_AND_LINE_NUMBER := MODULE_NAME_BY_ID[compID].(map[string]interface{})["trace"].([]interface{})
			for _, single_module := range TRACE_BY_MODULE_NAME_AND_LINE_NUMBER {
				single_traces := single_module.(map[string]interface{})
				if _, ok := single_traces[rowID]; ok {
					trace = single_traces[rowID].(string)
					slog.Info("[ ECU-D ] [ parserApCp ] TRACE_BY_MODULE_NAME_AND_LINE_NUMBER[rowID]", "rowID", rowID)
					break
				}
			}
			slog.Info("[ ECU-D ] [ parserApCp ] TRACE_BY_MODULE_NAME_AND_LINE_NUMBER[rowID]", "trace", trace)

			trace = fillTrace(trace, payloadInts)
		} else {
			component = collector + "_COMPONENT_ERROR_" + compID
			trace = collector + "_CODEBOOK_ERROR_" + compID + "_" + rowID
			slog.Error("[ ECU-D ] [ parserApCp ] [ERROR] component not found", "component", component, "trace", trace)
		}

	}

	return component[5:], traceLevel, trace

}

func fillTrace(s string, L []int32) string {

	if s == "Send POWERMAN_BUDGET_STATUS_IND ( shortTerm=%umAh, longTerm=%umAh, shortTermRemaining=%umAh, longTermRemaining=%umAh ) to MBX%u" {
		s = "Send POWERMAN_BUDGET_STATUS_IND ( shortTerm=%umAh, longTerm=%umAh, shortTermRemaining=%umAh, longTermRemaining=%umAh ) to MBX"
	} else if s == "Eni_TimerStart_Timeout: Start timer: %d, time:%d, msg:0x%x" {
		s = "Eni_TimerStart_Timeout: Start timer: %d, time:%d"
	} else if s == "invalidating current tow_diff due to age" {
		s = "invalidating current tow_diff due to age %d"
	} else if s == "IAA Overruled BCM userchange during initialization by changing to Anonymous Guest! (IAA UserId:%d, BCM Active:%d, BCM Identified:%d, VAL:Active:%d)" {
		s = "IAA Overruled BCM userchange during initialization by changing to Anonymous Guest! (IAA UserId:%d, BCM Active:%d, BCM Identified:%d, VAL:Active:???)"
	}

	placeholders := make([]interface{}, len(L))
	for i, v := range L {
		placeholders[i] = v
	}

	result := fmt.Sprintf(s, placeholders...)
	return result
}

func ComputePayloadInts(X []byte, n int) []int32 {
	L := []int32{}
	for i := 0; i < n; i++ {
		L = append(L, int32(binary.BigEndian.Uint32(X)))
	}
	return L
}

func formatTime(n int64, collector string) string {
	cntMillisecs := n / 100000
	if collector == "CP" {
		cntMillisecs = n / 1000000
	}
	mil := cntMillisecs % 1000
	cntSecs := cntMillisecs / 1000
	return time.Unix(cntSecs, int64(mil)).Format("15:04:05") + ":" + strings.Repeat("0", 3-len(strconv.Itoa(int(mil)))) + strconv.Itoa(int(mil))
}

func timeStampComp(time_slice []byte) int64 {
	if len(time_slice) == 4 {
		return int64(binary.BigEndian.Uint32(time_slice))
	} else if len(time_slice) == 8 {
		return int64(binary.BigEndian.Uint64(time_slice))
	} else {
		slog.Error("[ ECU-D ] TIME_ERROR")
		return 0
	}
}

func traceTypeFromID(n int) string {
	if n == 0 {
		return "INIT_TRACE"
	} else if n == 1 {
		return "ONE_TRACE"
	} else if n == 6 {
		return "NORMAL_TRACE"
	} else if n == 12 {
		return "BOERNE_TRACE"
	} else {
		return ""
	}
}

func collectorFromID(n int) string {
	if n == 1 {
		return "AP"
	} else if n == 2 {
		return "DMESG"
	} else if n == 3 {
		return "SYSLOG"
	} else if n == 4 {
		return "JAVA"
	} else if n == 5 {
		return "CP"
	} else if n == 6 {
		return "COLLECTOR_006"
	} else if n == 7 {
		return "DRLIB"
	} else if n == 8 {
		return "COLLECTOR_008"
	} else if n == 9 {
		return "SYSMON"
	} else {
		return ""
	}
}

func traceLevelComp(n int) string {
	if n == 0 {
		return "All"
	} else if n == 1 {
		return "Trace"
	} else if n == 2 {
		return "Debug"
	} else if n == 4 {
		return "Info"
	} else if n == 5 {
		return "Warn"
	} else if n == 6 {
		return "Error"
	} else if n == 7 {
		return "Fatal"
	} else {
		return "TRACE_LEVEL_NONE"
	}
}

func traceStyleComp(X byte) (string, int) {
	t := (X >> 2) & 15
	if t == 0 {
		return "RAW_OLD", -1
	} else if t == 3 {
		return "RAW", -1
	} else if t == 7 {
		return "CODE", 0
	} else if t == 9 {
		return "CODE", 1
	} else if t == 11 {
		return "CODE", 2
	} else if t == 13 {
		return "CODE", 3
	} else if t == 15 {
		return "CODE", 4
	} else {
		fmt.Printf("[ ECU-D ] [ ERROR ] Style not found %d\n", t)
		return "", -1
	}
}
