package main

import (
	"log"
	"net/http"
	"proto-server/protobuf"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func protoHandler(w http.ResponseWriter, r *http.Request) {

	phone := protobuf.Person_PhoneNumber_builder{
		Number: "09...",
		Type:   protobuf.PhoneType_PHONE_TYPE_MOBILE,
	}.Build()

	person := protobuf.Person_builder{
		Id:          1,
		Name:        "test-name",
		Email:       "test-email",
		Phones:      []*protobuf.Person_PhoneNumber{phone},
		LastUpdated: timestamppb.Now(),
	}.Build()

	// Serialize to protobuf
	data, err := proto.Marshal(person)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}

	// decerializedData := new(protobuf.Person)

	// err = proto.Unmarshal(data, decerializedData)
	// if err != nil {
	// 	log.Fatalf("Failed to decode data: %v", err)
	// }

	// log.Printf("Deserialized Person: %v", decerializedData)

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/proto", protoHandler)

	port := ":3000"
	log.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
