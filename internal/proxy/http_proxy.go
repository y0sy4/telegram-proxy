// Package proxy provides HTTP proxy server functionality.
package proxy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

// HTTPProxy represents an HTTP proxy server.
type HTTPProxy struct {
	port          int
	verbose       bool
	logger        *log.Logger
	upstreamProxy *url.URL
}

// dialWithUpstream creates a connection, optionally routing through an upstream proxy.
func (h *HTTPProxy) dialWithUpstream(network, addr string) (net.Conn, error) {
	if h.upstreamProxy == nil {
		return net.Dial(network, addr)
	}

	switch h.upstreamProxy.Scheme {
	case "socks5", "socks":
		// Use proxy package for SOCKS5
		proxyDialer, err := proxy.FromURL(h.upstreamProxy, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("create SOCKS5 dialer: %w", err)
		}
		return proxyDialer.Dial(network, addr)

	case "http", "https":
		// Use http.Transport with Proxy for HTTP CONNECT
		transport := &http.Transport{
			Proxy: http.ProxyURL(h.upstreamProxy),
		}
		return transport.Dial(network, addr)

	default:
		return nil, fmt.Errorf("unsupported upstream proxy scheme: %s", h.upstreamProxy.Scheme)
	}
}

// NewHTTPProxy creates a new HTTP proxy server.
func NewHTTPProxy(port int, verbose bool, logger *log.Logger, upstreamProxyURL string) (*HTTPProxy, error) {
	var upstreamProxy *url.URL
	var err error
	
	if upstreamProxyURL != "" {
		upstreamProxy, err = url.Parse(upstreamProxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid upstream proxy URL: %v", err)
		}
	}
	
	return &HTTPProxy{
		port:         port,
		verbose:      verbose,
		logger:       logger,
		upstreamProxy: upstreamProxy,
	}, nil
}

// Start starts the HTTP proxy server.
func (h *HTTPProxy) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", h.port))
	if err != nil {
		return err
	}
	defer listener.Close()
	
	if h.verbose {
		h.logger.Printf("[HTTP] Listening on port %d", h.port)
	}
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go h.handleConnection(conn)
	}
}

func (h *HTTPProxy) handleConnection(conn net.Conn) {
	defer conn.Close()
	
	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		return
	}
	defer req.Body.Close()
	
	// Handle CONNECT method (for HTTPS)
	if req.Method == http.MethodConnect {
		h.handleConnect(conn, req)
		return
	}
	
	// Handle HTTP requests
	h.handleHTTP(conn, req)
}

func (h *HTTPProxy) handleConnect(conn net.Conn, req *http.Request) {
	// Parse host:port
	host := req.URL.Host
	if !strings.Contains(host, ":") {
		host = host + ":80"
	}

	// Connect to target (with upstream proxy if configured)
	target, err := h.dialWithUpstream("tcp", host)
	if err != nil {
		conn.Write([]byte("HTTP/1.1 502 Bad Gateway\r\n\r\n"))
		return
	}
	defer target.Close()

	// Send success response
	conn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	// Bridge connections
	go io.Copy(target, conn)
	io.Copy(conn, target)
}

func (h *HTTPProxy) handleHTTP(conn net.Conn, req *http.Request) {
	// For now, just return error - full HTTP proxy is complex
	conn.Write([]byte("HTTP/1.1 501 Not Implemented\r\n\r\n"))
}
