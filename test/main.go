package main

import (
	"bytes"
	"container/list"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(os.Getenv("programdata"))
	return
	http.ListenAndServe(":9090", new(x))
}

var _list = list.New()

type x struct{}

func (x x) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RequestURI, "pull") {
		x.pull(w, r)
		return
	}
	if http.MethodGet == r.Method {
		c := NewWXBizMsgCrypt("B61HCDBO4N", "IivlYoyxpK8ErIQzokaV1DlXanTIpndSdL2DxPeZOzR", "ww48fb21eab5cc8802")
		data, err := c.VerifyRequest(r)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(data))
		w.Write(data)
	} else {
		h := make(map[string]interface{})
		h["method"] = r.Method
		h["content-type"] = r.Header.Get("Content-Type")
		h["data"], _ = ioutil.ReadAll(r.Body)
		_list.PushBack(h)
	}
}

func (x) pull(w http.ResponseWriter, r *http.Request) {
	e := _list.Front()
	if e != nil {
		_list.Remove(e)
		rr := e.Value.(map[string]interface{})
		data, _ := json.Marshal(rr)
		w.Write(data)
	} else {
		w.WriteHeader(400)
	}
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	ValidateSignatureError int = -40001
	ParseJsonError         int = -40002
	ComputeSignatureError  int = -40003
	IllegalAesKey          int = -40004
	ValidateCorpidError    int = -40005
	EncryptAESError        int = -40006
	DecryptAESError        int = -40007
	IllegalBuffer          int = -40008
	EncodeBase64Error      int = -40009
	DecodeBase64Error      int = -40010
	GenJsonError           int = -40011
	IllegalProtocolType    int = -40012
)

type ProtocolType int

type CryptError struct {
	ErrCode int
	ErrMsg  string
}

func NewCryptError(errCode int, errMsg string) *CryptError {
	return &CryptError{ErrCode: errCode, ErrMsg: errMsg}
}

func (c CryptError) Error() string {
	return fmt.Sprintf("[%d] %s", c.ErrCode, c.ErrMsg)
}

type WxBizMsgReceive struct {
	ToUsername string `json:"tousername"`
	Encrypt    string `json:"encrypt"`
	AgentId    string `json:"agentid"`
}

type WxBizMsgSend struct {
	Encrypt   string `json:"encrypt"`
	Signature string `json:"msgsignature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
}

func NewWxBizMsgSend(encrypt, signature, timestamp, nonce string) *WxBizMsgSend {
	return &WxBizMsgSend{Encrypt: encrypt, Signature: signature, Timestamp: timestamp, Nonce: nonce}
}

type ProtocolProcessor interface {
	parse(sec []byte) (*WxBizMsgReceive, *CryptError)
	serialize(data *WxBizMsgSend) ([]byte, *CryptError)
}

type WXBizMsgCrypt struct {
	token             string
	encodingAesKey    string
	receiverId        string
	protocolProcessor ProtocolProcessor
}

type JsonProcessor struct{}

func (p *JsonProcessor) parse(ata []byte) (*WxBizMsgReceive, *CryptError) {
	var msg WxBizMsgReceive
	err := json.Unmarshal(ata, &msg)
	if nil != err {
		return nil, NewCryptError(ParseJsonError, "json to msg fail")
	}
	return &msg, nil
}

func (p *JsonProcessor) serialize(sendMsg *WxBizMsgSend) ([]byte, *CryptError) {
	msg, err := json.Marshal(sendMsg)
	if nil != err {
		return nil, NewCryptError(GenJsonError, err.Error())
	}

	return msg, nil
}

func NewWXBizMsgCrypt(token, aesKey, receiverId string) *WXBizMsgCrypt {
	return &WXBizMsgCrypt{token: token, encodingAesKey: aesKey + "=", receiverId: receiverId, protocolProcessor: new(JsonProcessor)}
}

func (c *WXBizMsgCrypt) randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (c *WXBizMsgCrypt) pKCS7Padding(plaintext string, block_size int) []byte {
	padding := block_size - (len(plaintext) % block_size)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	var buffer bytes.Buffer
	buffer.WriteString(plaintext)
	buffer.Write(padtext)
	return buffer.Bytes()
}

func (c *WXBizMsgCrypt) pKCS7Unpadding(plaintext []byte, block_size int) ([]byte, *CryptError) {
	plaintext_len := len(plaintext)
	if nil == plaintext || plaintext_len == 0 {
		return nil, NewCryptError(DecryptAESError, "pKCS7Unpadding error nil or zero")
	}
	if plaintext_len%block_size != 0 {
		return nil, NewCryptError(DecryptAESError, "pKCS7Unpadding text not a multiple of the block size")
	}
	padding_len := int(plaintext[plaintext_len-1])
	return plaintext[:plaintext_len-padding_len], nil
}

func (c *WXBizMsgCrypt) cbcEncrypter(plaintext string) ([]byte, *CryptError) {
	aeskey, err := base64.StdEncoding.DecodeString(c.encodingAesKey)
	if nil != err {
		return nil, NewCryptError(DecodeBase64Error, err.Error())
	}
	const block_size = 32
	pad_msg := c.pKCS7Padding(plaintext, block_size)

	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, NewCryptError(EncryptAESError, err.Error())
	}

	ciphertext := make([]byte, len(pad_msg))
	iv := aeskey[:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext, pad_msg)
	base64_msg := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(base64_msg, ciphertext)

	return base64_msg, nil
}

func (c *WXBizMsgCrypt) cbcDecrypter(base64_encrypt_msg string) ([]byte, *CryptError) {
	aeskey, err := base64.StdEncoding.DecodeString(c.encodingAesKey)
	if nil != err {
		return nil, NewCryptError(DecodeBase64Error, err.Error())
	}

	encrypt_msg, err := base64.StdEncoding.DecodeString(base64_encrypt_msg)
	if nil != err {
		return nil, NewCryptError(DecodeBase64Error, err.Error())
	}

	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, NewCryptError(DecryptAESError, err.Error())
	}

	if len(encrypt_msg) < aes.BlockSize {
		return nil, NewCryptError(DecryptAESError, "encrypt_msg size is not valid")
	}

	iv := aeskey[:aes.BlockSize]

	if len(encrypt_msg)%aes.BlockSize != 0 {
		return nil, NewCryptError(DecryptAESError, "encrypt_msg not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encrypt_msg, encrypt_msg)

	return encrypt_msg, nil
}

func (c *WXBizMsgCrypt) calSignature(timestamp, nonce, data string) string {
	sort_arr := []string{c.token, timestamp, nonce, data}
	sort.Strings(sort_arr)
	var buffer bytes.Buffer
	for _, value := range sort_arr {
		buffer.WriteString(value)
	}

	sha := sha1.New()
	sha.Write(buffer.Bytes())
	signature := fmt.Sprintf("%x", sha.Sum(nil))
	return signature
}

func (c *WXBizMsgCrypt) ParsePlainText(plaintext []byte) ([]byte, uint32, []byte, []byte, *CryptError) {
	const block_size = 32
	plaintext, err := c.pKCS7Unpadding(plaintext, block_size)
	if nil != err {
		return nil, 0, nil, nil, err
	}

	text_len := uint32(len(plaintext))
	if text_len < 20 {
		return nil, 0, nil, nil, NewCryptError(IllegalBuffer, "plain is to small 1")
	}
	random := plaintext[:16]
	msg_len := binary.BigEndian.Uint32(plaintext[16:20])
	if text_len < (20 + msg_len) {
		return nil, 0, nil, nil, NewCryptError(IllegalBuffer, "plain is to small 2")
	}

	msg := plaintext[20 : 20+msg_len]
	receiver_id := plaintext[20+msg_len:]

	return random, msg_len, msg, receiver_id, nil
}

func (c *WXBizMsgCrypt) VerifyRequest(r *http.Request) ([]byte, error) {
	sig := r.URL.Query().Get("msg_signature")
	timestamp := r.URL.Query().Get("timestamp")
	nonce := r.URL.Query().Get("nonce")
	echostr := r.URL.Query().Get("echostr")
	if sig == "" || timestamp == "" || nonce == "" || echostr == "" {
		return nil, NewCryptError(ValidateSignatureError, "empty message")
	}
	signature := c.calSignature(timestamp, nonce, echostr)
	if strings.Compare(signature, sig) != 0 {
		return nil, NewCryptError(ValidateSignatureError, "signature not equal")
	}

	plaintext, err := c.cbcDecrypter(echostr)
	if nil != err {
		return nil, err
	}
	_, _, msg, receiver_id, err := c.ParsePlainText(plaintext)
	if nil != err {
		return nil, err
	}
	if len(c.receiverId) > 0 && strings.Compare(string(receiver_id), c.receiverId) != 0 {
		return nil, NewCryptError(ValidateCorpidError, "receiverId is not equil")
	}
	return msg, nil
}

func (c *WXBizMsgCrypt) EncryptMsg(replyMsg, timestamp, nonce string) ([]byte, *CryptError) {
	rand_str := c.randString(16)
	var buffer bytes.Buffer
	buffer.WriteString(rand_str)

	msg_len_buf := make([]byte, 4)
	binary.BigEndian.PutUint32(msg_len_buf, uint32(len(replyMsg)))
	buffer.Write(msg_len_buf)
	buffer.WriteString(replyMsg)
	buffer.WriteString(c.receiverId)

	tmp_ciphertext, err := c.cbcEncrypter(buffer.String())
	if nil != err {
		return nil, err
	}
	ciphertext := string(tmp_ciphertext)

	signature := c.calSignature(timestamp, nonce, ciphertext)

	msg4_send := NewWxBizMsgSend(ciphertext, signature, timestamp, nonce)
	return c.protocolProcessor.serialize(msg4_send)
}

func (c *WXBizMsgCrypt) DecryptMsg(sig, timestamp, nonce string, data []byte) ([]byte, *CryptError) {
	msg4_recv, crypt_err := c.protocolProcessor.parse(data)
	if nil != crypt_err {
		return nil, crypt_err
	}

	signature := c.calSignature(timestamp, nonce, msg4_recv.Encrypt)

	if strings.Compare(signature, sig) != 0 {
		return nil, NewCryptError(ValidateSignatureError, "signature not equal")
	}

	plaintext, crypt_err := c.cbcDecrypter(msg4_recv.Encrypt)
	if nil != crypt_err {
		return nil, crypt_err
	}

	_, _, msg, receiver_id, crypt_err := c.ParsePlainText(plaintext)
	if nil != crypt_err {
		return nil, crypt_err
	}

	if len(c.receiverId) > 0 && strings.Compare(string(receiver_id), c.receiverId) != 0 {
		return nil, NewCryptError(ValidateCorpidError, "receiverId is not equil")
	}

	return msg, nil
}
