package stdlist

import (
	"strings"
)

var rawdata = map[string][]string{
	"compress/bzip2":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/build":                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/testdata":                      []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"fmt":                                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"io/ioutil":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/debug":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/testdata/testprognet":          []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/dns":           []string{"1.11"},
	"crypto":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/sha512":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"regexp":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/plan9obj":                        []string{"1.10", "1.11", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/gccgoimporter":             []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal":                              []string{"1.10", "1.11", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/pprof/internal/protopprof":     []string{"1.8"},
	"crypto/subtle":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/gif":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"strconv":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"math/rand":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/httptest":                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"syscall/js":                            []string{"1.11"},
	"compress/flate":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/des":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"log/syslog":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/golang.org/x/net/http2":       []string{"1.6"},
	"image/color/palette":                   []string{"1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/testlog":                      []string{"1.10", "1.11"},
	"liblink":                               []string{"1.4"},
	"vendor/golang_org/x/text":              []string{"1.10", "1.11", "1.8", "1.9"},
	"image/png":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/testdata/testprogcgo/windows":  []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org":                     []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"unicode/utf16":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/plan9obj/testdata":               []string{"1.10", "1.11", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/idna":          []string{"1.10", "1.11", "1.8", "1.9"},
	"hash":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"mime/multipart":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"text/template":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/internal":                      []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"compress/gzip":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/gosym":                           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"strings":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/race":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/internal":                     []string{"1.10", "1.11", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"archive/tar/testdata":                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/pe/testdata":                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/mail":                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/lif":           []string{"1.10", "1.11", "1.8", "1.9"},
	"crypto/ecdsa":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/png/testdata/pngsuite":           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/testdata":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/parser/testdata":                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/scanner":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"libbio":                                []string{"1.4"},
	"net/internal":                          []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/testdata/testprog":             []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/md5":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"mime":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/nettrace":                     []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"math":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"os/user":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"path/filepath":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/srcimporter":               []string{"1.10", "1.11", "1.9"},
	"archive/tar":                           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/printer":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image":                                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"io/ioutil/testdata":                    []string{"1.10", "1.11", "1.9"},
	"image/draw":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/pprof/internal/profile":        []string{"1.10", "1.11", "1.9"},
	"text/tabwriter":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/importer":                           []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/nettest":       []string{"1.10", "1.11", "1.9"},
	"lib9":                                  []string{"1.4"},
	"internal/pprof/profile":                []string{"1.8"},
	"net/http/httputil":                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"os/exec":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/pprof":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/doc":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/http/httpguts": []string{"1.11"},
	"runtime/pprof/testdata":                []string{"1.10", "1.11", "1.9"},
	"vendor/golang_org/x/net":               []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"go/build/testdata/doc":                 []string{"1.11"},
	"crypto/sha256":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/elf":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"path":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/race":                         []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/internal/socktest":                 []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"math/bits":                             []string{"1.10", "1.11", "1.9"},
	"vendor/golang_org/x/crypto/poly1305":   []string{"1.10", "1.11", "1.8", "1.9"},
	"vendor/golang_org/x/text/unicode":      []string{"1.10", "1.11", "1.8", "1.9"},
	"crypto/rsa":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/testdata":                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/rpc/jsonrpc":                       []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"sync":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/x509/testdata":                  []string{"1.10", "1.11", "1.9"},
	"crypto/internal/randutil":              []string{"1.11"},
	"internal/pprof":                        []string{"1.8"},
	"crypto/sha1":                           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"database/sql/driver":                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/cgi/testdata":                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"os/signal/internal/pty":                []string{"1.10", "1.11"},
	"internal/format":                       []string{"1.5"},
	"encoding/gob":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/build/testdata/other":               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/internal":                        []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"compress/gzip/testdata":                []string{"1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/http2/hpack":   []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"expvar":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"html":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"sort":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"reflect":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/syscall/windows":              []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/internal/nettest":    []string{"1.10", "1.11"},
	"vendor/golang_org/x/net/route":               []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"container/heap":                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/json":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"mime/multipart/testdata":                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/cpu":                                []string{"1.10", "1.11", "1.9"},
	"internal/syscall/windows/registry":           []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/internal/sys":                        []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"flag":                                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/jpeg":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/srcimporter/testdata/issue20855": []string{"1.10", "1.11"},
	"go/internal":                                 []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"os/signal/internal":                          []string{"1.10", "1.11"},
	"vendor":                                      []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto":                  []string{"1.10", "1.11", "1.8", "1.9"},
	"go/printer/testdata":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/textproto":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"strconv/testdata":                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"mime/quotedprintable":                        []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x":                         []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/text/width":              []string{"1.8"},
	"compress/lzw":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"html/template":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http":                                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"text/scanner":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/internal/subtle":                      []string{"1.11"},
	"internal/golang.org/x":                       []string{"1.6"},
	"hash/fnv":                                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"bufio":                                       []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/hmac":                                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/base64":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto/internal/chacha20": []string{"1.11"},
	"text":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"time":                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/internal": []string{"1.10", "1.11"},
	"internal/trace":                   []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"plugin":                           []string{"1.10", "1.11", "1.8", "1.9"},
	"go/internal/gcimporter/testdata":  []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/types/testdata":                []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/pprof":                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"regexp/testdata":                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"unicode":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"lib9/fmt":                         []string{"1.4"},
	"crypto/elliptic":                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/smtp":                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"testing/quick":                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/msan":                     []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/format":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/color":                      []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"unsafe":                           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/constant":                      []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto/cryptobyte/asn1":  []string{"1.10", "1.11"},
	"internal/syscall":                            []string{"1.10", "1.11", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"syscall":                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/gcimporter/testdata/versions":    []string{"1.10", "1.11", "1.8", "1.9"},
	"internal/trace/testdata":                     []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"net/http/httptrace":                          []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"runtime/trace":                               []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"archive/zip/testdata":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"log":                                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/rpc":                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/lex":                 []string{"1.10", "1.7", "1.8", "1.9", "1"},
	"go/internal/srcimporter/testdata/issue23092": []string{"1.11"},
	"runtime/pprof/testdata/mappingtest":          []string{"1.11"},
	"mime/testdata":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/url":                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/gcimporter":                      []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"builtin":                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/text/transform":          []string{"1.10", "1.11", "1.8", "1.9"},
	"hash/adler32":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"index":                                       []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/cgo":                                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"testing/internal/testdeps":                   []string{"1.10", "1.11", "1.8", "1.9"},
	"vendor/golang_org/x/crypto/internal":         []string{"1.11"},
	"compress/testdata":                           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/ecdsa/testdata":                       []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/dwarf":                                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/poll":                               []string{"1.10", "1.11", "1.9"},
	"testing/internal":                            []string{"1.10", "1.11", "1.8", "1.9"},
	"vendor/golang_org/x/crypto/cryptobyte":       []string{"1.10", "1.11"},
	"vendor/golang_org/x/net/dns/dnsmessage":      []string{"1.11"},
	"container/list":                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/cgi":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/race/testdata":                       []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/text/unicode/bidi":       []string{"1.10", "1.11", "1.9"},
	"crypto/rc4":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/x509":                                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/internal":                             []string{"1.10", "1.11", "1.8", "1.9"},
	"go/token":                                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"lib9/utf":                                    []string{"1.4"},
	"compress":                                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/dwarf/testdata":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/csv":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/gccgoimporter/testdata":          []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/http2":               []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/http/httpproxy":      []string{"1.11"},
	"container/ring":                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/ast":                                      []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"testing":                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"hash/crc32":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"hash/crc64":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/singleflight":                       []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/syscall/windows/sysdll":             []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto/curve25519":       []string{"1.10", "1.11", "1.8", "1.9"},
	"compress/zlib":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/elf/testdata":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/pem":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/proxy":               []string{"1.10", "1.9"},
	"vendor/golang_org/x/text/secure/bidirule":    []string{"1.10", "1.11", "1.9"},
	"archive":                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/dsa":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/golang.org/x/net/http2/hpack":       []string{"1.6"},
	"image/testdata":                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"unicode/utf8":                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/text/unicode/norm":       []string{"1.10", "1.11", "1.8", "1.9"},
	"os/signal":                                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"math/big":                                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net":                                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/cookiejar":                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"os":                                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"compress/flate/testdata":                     []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"context":                                     []string{"1.10", "1.11", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto/chacha20poly1305/internal": []string{"1.10", "1.8", "1.9"},
	"vendor/golang_org/x/net/lex/httplex":                  []string{"1.10", "1.7", "1.8", "1.9", "1"},
	"encoding/asn1":                                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"text/template/parse":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/types":                                             []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/xml":                                         []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/build/testdata/multi":                              []string{"1.10", "1.11", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/rand":                                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/pe":                                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/json/testdata":                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"net/http/fcgi":                                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/internal/srcimporter/testdata":                     []string{"1.10", "1.11"},
	"go/internal/srcimporter/testdata/issue24392":          []string{"1.11"},
	"bytes":                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/hex":                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/png/testdata":           []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"archive/zip":                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/tls":                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/base32":              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/net/http": []string{"1.11"},
	"crypto/cipher":                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"math/cmplx":                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/internal/atomic":      []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto/chacha20poly1305": []string{"1.10", "1.11", "1.8", "1.9"},
	"debug":               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go":                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"index/suffixarray":   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/doc/testdata":     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/tls/testdata": []string{"1.10", "1.11", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20": []string{"1.10", "1.8", "1.9"},
	"internal/golang.org/x/net":                                     []string{"1.6"},
	"debug/macho":                                                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/ascii85":                                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding/binary":                                               []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"debug/macho/testdata":                                          []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"io":                                                            []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/rsa/testdata":                                           []string{"1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"runtime/pprof/internal":                                        []string{"1.10", "1.11", "1.8", "1.9"},
	"runtime/testdata/testprogcgo":                                  []string{"1.10", "1.11", "1.6", "1.7", "1.8", "1.9", "1"},
	"regexp/syntax":                                                 []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"sync/atomic":                                                   []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"compress/bzip2/testdata":                                       []string{"1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/internal/cipherhw":                                      []string{"1.10", "1.8", "1.9"},
	"internal/syscall/unix":                                         []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"container":                                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"crypto/aes":                                                    []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/build/testdata":                                             []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"vendor/golang_org/x/text/secure":                               []string{"1.10", "1.11", "1.9"},
	"go/build/testdata/empty":                                       []string{"1.10", "1.11", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"database/sql":                                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"errors":                                                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"testing/iotest":                                                []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"encoding":                                                      []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/parser":                                                     []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/golang.org":                                           []string{"1.6"},
	"crypto/x509/pkix":                                              []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"go/build/testdata/other/file":                                  []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"text/template/testdata":                                        []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/bytealg":                                              []string{"1.11"},
	"database":                                                      []string{"1.1", "1.10", "1.11", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"image/internal/imageutil":                                      []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
	"internal/testenv":                                              []string{"1.10", "1.11", "1.5", "1.6", "1.7", "1.8", "1.9", "1"},
}

func init() {
	for k, v := range rawdata {
		data[k] = v
	}
}
