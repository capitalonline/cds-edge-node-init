package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

//noinspection ALL
type TargetDef struct {
	Version             string `json: "version"`
	Download_url        string `json: "download_url"`
	Checksum_md5_url    string `json: "checksum_md5_url"`
	Checksum_sha256_url string `json: "checksum_sha256_url"`
}

type CloudRequest struct {
	method      string
	params      map[string]string
	action      string
	productType string
	body        io.Reader
}

func NewCCKRequest(action, method string, params map[string]string, body io.Reader) (*CloudRequest, error) {
	return NewRequest(action, method, params, cckProductType, body)
}

func NewRequest(action, method string, params map[string]string, productType string, body io.Reader) (*CloudRequest, error) {
	method = strings.ToUpper(method)
	req := &CloudRequest{
		method:      method,
		params:      params,
		action:      action,
		productType: productType,
		body:        body,
	}
	return req, nil
}

func DoRequest(req *CloudRequest) (resp *http.Response, err error) {
	if AccessKeyID != "" || AccessKeySecret != "" {
		return nil, fmt.Errorf("AccessKeyID or accessKeySecret is empty")
	}

	reqUrl := getUrl(req)
	sendRequest, err := http.NewRequest(req.method, reqUrl, req.body)
	if err != nil {
		return
	}
	log.Infof("send request url: %s", reqUrl)
	resp, err = http.DefaultClient.Do(sendRequest)
	return
}

func getUrl(req *CloudRequest) string {
	urlParams := map[string]string{
		"Action":           req.action,
		"AccessKeyId":      AccessKeyID,
		"SignatureMethod":  signatureMethod,
		"SignatureNonce":   uuid.New().String(),
		"SignatureVersion": signatureVersion,
		"Timestamp":        time.Now().UTC().Format(timeStampFormat),
		"Version":          version,
	}
	if req.params != nil {
		for k, v := range req.params {
			urlParams[k] = v
		}
	}
	var paramSortKeys sort.StringSlice
	for k, _ := range urlParams {
		paramSortKeys = append(paramSortKeys, k)
	}
	sort.Sort(paramSortKeys)
	var urlStr string
	for _, k := range paramSortKeys {
		urlStr += "&" + percentEncode(k) + "=" + percentEncode(urlParams[k])
	}
	urlStr = req.method + "&%2F&" + percentEncode(urlStr[1:])

	h := hmac.New(sha1.New, []byte(AccessKeySecret))
	h.Write([]byte(urlStr))
	signStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	urlParams["Signature"] = signStr

	urlVal := url.Values{}
	for k, v := range urlParams {
		urlVal.Add(k, v)
	}
	urlValStr := urlVal.Encode()
	reqUrl := fmt.Sprintf("%s/%s?%s", apiHost, req.productType, urlValStr)
	return reqUrl
}

func percentEncode(str string) string {
	str = url.QueryEscape(str)
	strings.Replace(str, "+", "%20", -1)
	strings.Replace(str, "*", "%2A", -1)
	strings.Replace(str, "%7E", "~", -1)
	return str
}

func SendRequest(method, url string, data_bytes []byte, headers []string) ([]byte, error) {
	var data io.Reader
	if data_bytes == nil {
		data = nil
	} else {
		data = bytes.NewReader(data_bytes)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for _, header := range headers {
			terms := strings.SplitN(header, " ", 2)
			if len(terms) == 2 {
				req.Header.Add(terms[0], terms[1])
			}
		}
	}
	if *FlagDebugMode {
		Logger.Println("=======Request Info ======")
		Logger.Println("=> URL:", url)
		Logger.Println("=> Method:", method)
		Logger.Println("=> Headers:", req.Header)
		Logger.Println("=> Body:", string(data_bytes))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200, 201, 202:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if *FlagDebugMode {
			Logger.Println("=======Response Info ======")
			Logger.Println("=> Headers:", resp.Header)
			Logger.Println("=> Body:", string(body))
		}
		return body, nil
	default:
		if *FlagDebugMode {
			Logger.Println("=======Response Info (ERROR) ======")
			Logger.Println("=> Headers:", resp.Header)
			b, _ := ioutil.ReadAll(resp.Body)
			Logger.Println("=> Body:", string(b))
		}
		err_msg := fmt.Sprintf("%d", resp.StatusCode)
		return nil, errors.New(err_msg)
	}
}

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		//SendError(err, "HTTP get error", nil)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DownloadFile(url, filepath, name string) {
	Logger.Printf("Downloading %s definition from %s", name, url)
	def := downloadTargetDef(url)

	Logger.Printf("Downloading %s from %s", name, def.Download_url)
	data := downloadTarget(def)
	Logger.Printf("Saving %s to %s", name, filepath)
	uncompress(data, filepath)
}

func downloadTargetDef(url string) *TargetDef {
	def, err := getTargetDef(url)
	for i := 1; ; i *= 2 {
		if i > MaxWaitingTime {
			i = 1
		}
		if err != nil || def == nil {
			Logger.Printf("Cannot get target definition: %s. Retry in %d second", err, i)
			time.Sleep(time.Duration(i) * time.Second)
			def, err = getTargetDef(url)
		} else {
			break
		}
	}
	return def
}

func getTargetDef(url string) (*TargetDef, error) {
	var def TargetDef
	body, err := HttpGet(url)
	if err != nil {
		//SendError(err, "HTTP get error", nil)
		return nil, err
	}
	if err = json.Unmarshal(body, &def); err != nil {
		//SendError(err, "json unmarshal error", nil)
		return nil, err
	}
	if def == (TargetDef{}) {
		//SendError(errors.New("Wrong target definition"), "Wrong target definition", nil)
		return nil, errors.New("Wrong target definition")
	}
	return &def, nil
}

func downloadTarget(def *TargetDef) []byte {
	b, err := getTarget(def)
	for i := 1; ; i *= 2 {
		if i > MaxWaitingTime {
			i = 1
		}
		if err != nil {
			Logger.Printf("Cannot get target: %s. Retry in %d second", err, i)
			time.Sleep(time.Duration(i) * time.Second)
			b, err = getTarget(def)
		} else {
			break
		}
	}
	return b
}

func getTarget(def *TargetDef) ([]byte, error) {
	b, err := HttpGet(def.Download_url)
	if err != nil {
		//SendError(err, "HTTP get error", nil)
		return nil, err
	}

	//validate md5 checksum of the target
	md5hasher := md5.New()
	md5hasher.Write(b)
	md5s := hex.EncodeToString(md5hasher.Sum(nil))
	md5b, err := HttpGet(def.Checksum_md5_url)
	if err != nil {
		//SendError(err, "HTTP get error", nil)
		Logger.Println("Failed to get md5 for the target")
		return nil, err
	} else {
		if !strings.Contains(string(md5b), md5s) {
			//SendError(errors.New("Failed to pass md5 checksum test"), "Failed on md5 checksum test", nil)
			return nil, errors.New("Failed to pass md5 checksum test")
		}
	}

	//validate sha256 checksum of the target
	sha256hasher := sha256.New()
	sha256hasher.Write(b)
	sha256s := hex.EncodeToString(sha256hasher.Sum(nil))
	sha256b, err := HttpGet(def.Checksum_sha256_url)
	if err != nil {
		//SendError(err, "HTTP error", nil)
		Logger.Println("Failed to get sha256 for the target")
		return nil, err
	} else {
		if !strings.Contains(string(sha256b), sha256s) {
			//SendError(errors.New("Failed to pass sha256 checksum test"), "Failed on sha256 checksum test", nil)
			return nil, errors.New("Failed to pass sha256 checksum test")
		}
	}

	return b, nil
}

func uncompress(data []byte, filefolder string) {
	byteReader := bytes.NewReader(data)
	gzipReader, err := gzip.NewReader(byteReader)
	if err != nil {
		//SendError(err, "Failed to uncompress file", nil)
		Logger.Print("Failed to uncompress file", err)
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			//SendError(err, "Failed to uncompress file", nil)
			Logger.Print("Failed to uncompress file", err)
		}
		filepath := path.Join(filefolder, path.Base(header.Name))

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			Logger.Print("Uncompressing: ", filepath)
			f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, header.FileInfo().Mode())
			if err != nil {
				//SendError(err, "Failed to open file", nil)
				Logger.Print("Failed to open file", err)
			}
			defer f.Close()
			for i := 1; ; i *= 2 {
				if i > MaxWaitingTime {
					i = 1
				}
				_, err := io.Copy(f, tarReader)
				if err != nil {
					//SendError(err, "Failed to uncompress file", nil)
					Logger.Printf("Failed to uncompress file: %s. Retrying in %d second", err, i)
					time.Sleep(time.Duration(i) * time.Second)
				} else {
					break
				}
			}
		default:
			Logger.Print("Can't: %c, %s\n", header.Typeflag, filepath)
		}
	}
}

