package cmd

import (
	"context"
	"github.com/AdguardTeam/golibs/timeutil"
	"log/slog"
)

// Configuration 1:1复制内部configuration
type Configuration struct {
	ConfigPath             string
	LogOutput              string            `yaml:"output"`
	TLSCertPath            string            `yaml:"tls-crt"`
	TLSKeyPath             string            `yaml:"tls-key"`
	HTTPSServerName        string            `yaml:"https-server-name"`
	HTTPSUserinfo          string            `yaml:"https-userinfo"`
	DNSCryptConfigPath     string            `yaml:"dnscrypt-config"`
	EDNSAddr               string            `yaml:"edns-addr"`
	UpstreamMode           string            `yaml:"upstream-mode"`
	ListenAddrs            []string          `yaml:"listen-addrs"`
	ListenPorts            []int             `yaml:"listen-ports"`
	HTTPSListenPorts       []int             `yaml:"https-port"`
	TLSListenPorts         []int             `yaml:"tls-port"`
	QUICListenPorts        []int             `yaml:"quic-port"`
	DNSCryptListenPorts    []int             `yaml:"dnscrypt-port"`
	Upstreams              []string          `yaml:"upstream"`
	BootstrapDNS           []string          `yaml:"bootstrap"`
	Fallbacks              []string          `yaml:"fallback"`
	PrivateRDNSUpstreams   []string          `yaml:"private-rdns-upstream"`
	DNS64Prefix            []string          `yaml:"dns64-prefix"`
	PrivateSubnets         []string          `yaml:"private-subnets"`
	BogusNXDomain          []string          `yaml:"bogus-nxdomain"`
	HostsFiles             []string          `yaml:"hosts-files"`
	Timeout                timeutil.Duration `yaml:"timeout"`
	CacheMinTTL            uint32            `yaml:"cache-min-ttl"`
	CacheMaxTTL            uint32            `yaml:"cache-max-ttl"`
	CacheSizeBytes         int               `yaml:"cache-size"`
	Ratelimit              int               `yaml:"ratelimit"`
	RatelimitSubnetLenIPv4 int               `yaml:"ratelimit-subnet-len-ipv4"`
	RatelimitSubnetLenIPv6 int               `yaml:"ratelimit-subnet-len-ipv6"`
	UDPBufferSize          int               `yaml:"udp-buf-size"`
	MaxGoRoutines          uint              `yaml:"max-go-routines"`
	TLSMinVersion          float32           `yaml:"tls-min-version"`
	TLSMaxVersion          float32           `yaml:"tls-max-version"`
	help                   bool
	HostsFileEnabled       bool `yaml:"hosts-file-enabled"`
	Pprof                  bool `yaml:"pprof"`
	Version                bool `yaml:"version"`
	Verbose                bool `yaml:"verbose"`
	Insecure               bool `yaml:"insecure"`
	IPv6Disabled           bool `yaml:"ipv6-disabled"`
	HTTP3                  bool `yaml:"http3"`
	CacheOptimistic        bool `yaml:"cache-optimistic"`
	Cache                  bool `yaml:"cache"`
	RefuseAny              bool `yaml:"refuse-any"`
	EnableEDNSSubnet       bool `yaml:"edns"`
	DNS64                  bool `yaml:"dns64"`
	UsePrivateRDNS         bool `yaml:"use-private-rdns"`
}

func (conf *Configuration) toInternalConf() *configuration {
	c := &configuration{
		ConfigPath:             conf.ConfigPath,
		LogOutput:              conf.LogOutput,
		TLSCertPath:            conf.TLSCertPath,
		TLSKeyPath:             conf.TLSKeyPath,
		HTTPSServerName:        conf.HTTPSServerName,
		HTTPSUserinfo:          conf.HTTPSUserinfo,
		DNSCryptConfigPath:     conf.DNSCryptConfigPath,
		EDNSAddr:               conf.EDNSAddr,
		UpstreamMode:           conf.UpstreamMode,
		ListenAddrs:            conf.ListenAddrs,
		ListenPorts:            conf.ListenPorts,
		HTTPSListenPorts:       conf.HTTPSListenPorts,
		TLSListenPorts:         conf.TLSListenPorts,
		QUICListenPorts:        conf.QUICListenPorts,
		DNSCryptListenPorts:    conf.DNSCryptListenPorts,
		Upstreams:              conf.Upstreams,
		BootstrapDNS:           conf.BootstrapDNS,
		Fallbacks:              conf.Fallbacks,
		PrivateRDNSUpstreams:   conf.PrivateRDNSUpstreams,
		DNS64Prefix:            conf.DNS64Prefix,
		PrivateSubnets:         conf.PrivateSubnets,
		BogusNXDomain:          conf.BogusNXDomain,
		HostsFiles:             conf.HostsFiles,
		Timeout:                conf.Timeout,
		CacheMinTTL:            conf.CacheMinTTL,
		CacheMaxTTL:            conf.CacheMaxTTL,
		CacheSizeBytes:         conf.CacheSizeBytes,
		Ratelimit:              conf.Ratelimit,
		RatelimitSubnetLenIPv4: conf.RatelimitSubnetLenIPv4,
		RatelimitSubnetLenIPv6: conf.RatelimitSubnetLenIPv6,
		UDPBufferSize:          conf.UDPBufferSize,
		MaxGoRoutines:          conf.MaxGoRoutines,
		TLSMinVersion:          conf.TLSMinVersion,
		TLSMaxVersion:          conf.TLSMaxVersion,
		HostsFileEnabled:       conf.HostsFileEnabled,
		Pprof:                  conf.Pprof,
		Version:                conf.Version,
		Verbose:                conf.Verbose,
		Insecure:               conf.Insecure,
		IPv6Disabled:           conf.IPv6Disabled,
		HTTP3:                  conf.HTTP3,
		CacheOptimistic:        conf.CacheOptimistic,
		Cache:                  conf.Cache,
		RefuseAny:              conf.RefuseAny,
		EnableEDNSSubnet:       conf.EnableEDNSSubnet,
		DNS64:                  conf.DNS64,
		UsePrivateRDNS:         conf.UsePrivateRDNS,
	}

	return c
}

func RunProxy(ctx context.Context, l *slog.Logger, conf *Configuration) error {
	return runProxy(ctx, l, conf.toInternalConf())
}
