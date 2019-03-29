package main

// TenorResponse the response from a tenor query.
type TenorResponse struct {
	Results []struct {
		Tags  []interface{} `json:"tags"`
		URL   string        `json:"url"`
		Media []struct {
			Nanomp4 struct {
				URL      string  `json:"url"`
				Dims     []int   `json:"dims"`
				Duration float64 `json:"duration"`
				Preview  string  `json:"preview"`
				Size     int     `json:"size"`
			} `json:"nanomp4"`
			Nanowebm struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"nanowebm"`
			Tinygif struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"tinygif"`
			Tinymp4 struct {
				URL      string  `json:"url"`
				Dims     []int   `json:"dims"`
				Duration float64 `json:"duration"`
				Preview  string  `json:"preview"`
				Size     int     `json:"size"`
			} `json:"tinymp4"`
			Tinywebm struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"tinywebm"`
			Webm struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"webm"`
			Gif struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"gif"`
			Mp4 struct {
				URL      string  `json:"url"`
				Dims     []int   `json:"dims"`
				Duration float64 `json:"duration"`
				Preview  string  `json:"preview"`
				Size     int     `json:"size"`
			} `json:"mp4"`
			Loopedmp4 struct {
				URL      string  `json:"url"`
				Dims     []int   `json:"dims"`
				Duration float64 `json:"duration"`
				Preview  string  `json:"preview"`
				Size     int     `json:"size"`
			} `json:"loopedmp4"`
			Mediumgif struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"mediumgif"`
			Nanogif struct {
				URL     string `json:"url"`
				Dims    []int  `json:"dims"`
				Preview string `json:"preview"`
				Size    int    `json:"size"`
			} `json:"nanogif"`
		} `json:"media"`
		Created   float64     `json:"created"`
		Shares    int         `json:"shares"`
		Itemurl   string      `json:"itemurl"`
		Composite interface{} `json:"composite"`
		Hasaudio  bool        `json:"hasaudio"`
		Title     string      `json:"title"`
		ID        string      `json:"id"`
	} `json:"results"`
	Next string `json:"next"`
}
