package decoder

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func decoder_() {
	data := "ecu_trace.bbb"

	rawFile, err := os.ReadFile(data)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	packetSize := 1024
	packets := len(rawFile)/packetSize + 1
	fmt.Println("total packets: \n", packets)

	remain := []byte{}

	for i := 400; i < 600; i++ {
		rawBytes := append(remain, rawFile[i*packetSize:(i+1)*packetSize]...)
		remain, out := lgeFullReading_(rawBytes)
		if len(out) > 0 {
			fmt.Println("=================remain, out=================")
			fmt.Println("remain: \n", string(remain))
			fmt.Println("out: \n", out)
		}
	}
}

//	func lgeFullReading(rawBytes []byte) ([]byte, []map[string]interface{}) {
//		var out []map[string]interface{}
func lgeFullReading_(rawBytes []byte) ([]byte, []byte) {
	var out []byte

	// if 128 in rawBytes:
	// fmt.Println(rawBytes)
	length := len(rawBytes)
	fmt.Println("rawBytes length:", length)

	// re, err := regexp.Compile("(.{4})\x80\x4d\x4c\x50.{4}\x00\x00_(.{5})\x06\x00(.{4}).*")
	re := regexp.MustCompile(`(.{4})\\0x80\\0x4d\\0x4c\\0x50.{4}\\0x00\\0x00_(.{5})\\0x06\\0x00(.{4}).*`)
	// re := regexp.MustCompile(`(.{4})(\x{80})`)
	header_idx := re.FindSubmatch(rawBytes)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	if len(header_idx) > 0 {
		fmt.Printf("============Matched: %s\n", header_idx[0])
		fmt.Printf("============Matched group 1: %s\n", header_idx[1])
		fmt.Printf("Matched group 2: %sn", header_idx[2])
		fmt.Printf("Matched group 3: %sn", header_idx[3])
	} else {
		fmt.Println("No match found.")
	}
	return rawBytes[length-25:], out

	// header_idx := re.FindIndex(rawBytes)
	// if header_idx != nil {
	// 	fmt.Println("matched at:", header_idx[0])
	// 	out = rawBytes[header_idx[0]:header_idx[1]]
	// 	return rawBytes[header_idx[1]:], out

	// } else {
	// 	return rawBytes[length-25:], out

	// }

	// for _, header := range headers {
	// 	headerStr := string(header)
	// 	fmt.Println(headerStr)

	// 	// Extract the header fields
	// 	var headerFields map[string]interface{}
	// 	headerFields = parseHeader(headerStr)

	// 	// Extract the trace type and collector
	// 	traceType, traceCollector := extractTraceTypeAndCollector(headerFields)

	// 	// Extract the trace payload
	// 	tracePayload, err := extractTracePayload(headerFields, traceCollector)
	// 	if err != nil {
	// 		fmt.Println("Error extracting trace payload:", err)
	// 		continue
	// 	}

	// 	// Create a map to store the trace information
	// 	trace := map[string]interface{}{
	// 		"type":      traceType,
	// 		"collector": traceCollector,
	// 		"payload":   tracePayload,
	// 	}

	// 	// Add the trace to the output list
	// 	out = append(out, trace)
	// }

}

// func parseHeader(headerStr string) map[string]interface{} {
// 	matches := regexp.FindAllString(headerStr, `(.{4})(.{4})(.{5})(.{4})(.{4})(.{4})(.{4})`)

// 	headerFields := make(map[string]interface{})

// 	for _, match := range matches {
// 		fieldName := strings.TrimSpace(fmt.Sprintf("%s %s %s %s %s %s %s", match[0], match[1], match[2], match[3], match[4], match[5], match[6]))
// 		fieldValue, err := strconv.ParseInt(match[7], 16, 64)
// 		if err != nil {
// 			fmt.Println("Error parsing header field value:", err)
// 			continue
// 		}

// 		headerFields[fieldName] = fieldValue
// 	}

// 	return headerFields
// }

// func extractTraceTypeAndCollector(headerFields map[string]interface{}) (string, string) {
// 	traceCollector := headerFields["Collector"]
// 	traceType := headerFields["TraceType"]

// 	return
// }
