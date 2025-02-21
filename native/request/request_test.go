package request

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/bsm/openrtb/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	It("should parse correctly", func() {
		req := fixture("testdata/request1.json")
		Expect(req).To(Equal(&Request{
			Version:          "1.1",
			ContextTypeID:    ContextTypeSocial,
			ContextSubTypeID: ContextSubTypeSocial,
			PlacementTypeID:  PlacementTypeID(11),
			PlacementCount:   "1",
			Sequence:         2,
			Assets: []Asset{
				{ID: 123, Required: 1, Title: &Title{Length: 140}},
				{ID: 128, Image: &Image{TypeID: ImageTypeMain, WidthMin: 836, HeightMin: 627, Width: 1000, Height: 800, MIMEs: []string{"image/jpg"}}},
				{ID: 126, Required: 1, Data: &Data{TypeID: DataTypeSponsored, Length: 25}},
				{ID: 127, Required: 1, Data: &Data{TypeID: DataTypeDesc, Length: 140}},
				{ID: 4, Video: &Video{MinDuration: 15, MaxDuration: 30, Protocols: []openrtb.Protocol{openrtb.ProtocolVAST2, openrtb.ProtocolVAST3}, MIMEs: []string{"video/mp4"}}},
			},
			IsWrapped: false,
		}))
	})

	It("should parse prior to 1.1 correctly", func() {
		req := fixture("testdata/request2.json")
		Expect(req).To(Equal(&Request{
			Version:          "1.1",
			ContextTypeID:    ContextTypeSocial,
			ContextSubTypeID: ContextSubTypeSocial,
			PlacementTypeID:  PlacementTypeID(11),
			PlacementCount:   "1",
			Sequence:         2,
			Assets: []Asset{
				{ID: 123, Required: 1, Title: &Title{Length: 140}},
				{ID: 128, Image: &Image{TypeID: ImageTypeMain, WidthMin: 836, HeightMin: 627, Width: 1000, Height: 800, MIMEs: []string{"image/jpg"}}},
				{ID: 126, Required: 1, Data: &Data{TypeID: DataTypeSponsored, Length: 25}},
				{ID: 127, Required: 1, Data: &Data{TypeID: DataTypeDesc, Length: 140}},
				{ID: 4, Video: &Video{MinDuration: 15, MaxDuration: 30, Protocols: []openrtb.Protocol{openrtb.ProtocolVAST2, openrtb.ProtocolVAST3}, MIMEs: []string{"video/mp4"}}},
			},
			IsWrapped: true,
		}))
	})

	It("should parse json string correctly", func() {
		req := fixture("testdata/request3.json")
		Expect(req).To(Equal(&Request{
			Version:   "1.0",
			Assets:    []Asset{},
			IsWrapped: true,
		}))
	})

	It("should parse string plcmtcnt correctly", func() {
		req := fixture("testdata/request4.json")
		Expect(req).To(Equal(&Request{
			Version:        "1.1",
			PlacementCount: "12",
			Assets: []Asset{
				{ID: 123, Required: 1, Title: &Title{Length: 140}},
			},
			IsWrapped: false,
		}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openrtb/native/request")
}

func fixture(path string) *Request {
	var subject Request
	enc, err := ioutil.ReadFile(path)
	Expect(err).ToNot(HaveOccurred())
	Expect(json.Unmarshal(enc, &subject)).To(Succeed())
	return &subject
}
