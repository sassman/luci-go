// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/milo/manifest.proto

package milo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Manifest attempts to make an accurate accounting of source/data directories
// during the execution of a LUCI task.
//
// These directories are primarily in the form of e.g. git checkouts of
// source, but also include things like isolated hashes and CIPD package
// deployments. In the future, other deployment forms may be supported (like
// other SCMs).
//
// The purpose of this manifest is so that other parts of the LUCI stack (e.g.
// Milo) can work with the descriptions of this deployed data as a first-class
// citizen. Initially this Manifest will be used to allow Milo to display diffs
// between jobs, but it will also be useful for tools and humans to get a
// record of exactly what data went into this LUCI task.
//
// Source Manifests can be emitted from recipes using the
// 'recipe_engine/source_manifest' module.
type Manifest struct {
	// Version will increment on backwards-incompatible changes only. Backwards
	// compatible changes will not alter this version number.
	//
	// Currently, the only valid version number is 0.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Map of local file system directory path (with forward slashes) to
	// a Directory message containing one or more deployments.
	//
	// The local path is relative to some job-specific root. This should be used
	// for informational/display/organization purposes. In particular, think VERY
	// CAREFULLY before you configure remote services/recipes to look for
	// particular filesystem layouts here. For example, if you want to look for
	// "the version of chromium/src checked out by the job", prefer to look for
	// a Directory which checks out "chromium/src", as opposed to assuming this
	// checkout lives in a top-level folder called "src". The reason for this is
	// that jobs SHOULD reserve the right to do their checkouts in any way they
	// please.
	//
	// If you feel like you need to make some service configuration which uses one
	// of these local filesystem paths as a key, please consult with the Chrome
	// Infrastructure team to see if there's a better alternative.
	//
	// Ex.
	//   "": {...}  // root directory
	//   "src/third_party/something": {...}
	Directories          map[string]*Manifest_Directory `protobuf:"bytes,2,rep,name=directories" json:"directories,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}
func (*Manifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_6c171d9612b862d6, []int{0}
}
func (m *Manifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest.Unmarshal(m, b)
}
func (m *Manifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest.Marshal(b, m, deterministic)
}
func (dst *Manifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest.Merge(dst, src)
}
func (m *Manifest) XXX_Size() int {
	return xxx_messageInfo_Manifest.Size(m)
}
func (m *Manifest) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest proto.InternalMessageInfo

func (m *Manifest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Manifest) GetDirectories() map[string]*Manifest_Directory {
	if m != nil {
		return m.Directories
	}
	return nil
}

type Manifest_GitCheckout struct {
	// The canonicalized URL of the original repo that is considered the “source
	// of truth” for the source code.
	//
	// Ex.
	//   https://chromium.googlesource.com/chromium/tools/build
	//   https://chromium.googlesource.com/infra/luci/recipes-py
	RepoUrl string `protobuf:"bytes,1,opt,name=repo_url,json=repoUrl" json:"repo_url,omitempty"`
	// If different from repo_url, this can be the URL of the repo that the source
	// was actually fetched from (i.e. a mirror).
	//
	// If this is empty, it's presumed to be equal to repo_url.
	//
	// Ex.
	//   https://github.com/luci/recipes-py
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl" json:"fetch_url,omitempty"`
	// The fully resolved revision (commit hash) of the source.
	//
	// This must always be a revision on the hosted repo (not any locally
	// generated commit).
	//
	// Ex.
	//   3617b0eea7ec74b8e731a23fed2f4070cbc284c4
	Revision string `protobuf:"bytes,3,opt,name=revision" json:"revision,omitempty"`
	// The ref that the task used to resolve/fetch the revision of the source
	// (if any).
	//
	// This must always be a ref on the hosted repo (not any local alias
	// like 'refs/remotes/...').
	//
	// This must always be an absolute ref (i.e. starts with 'refs/'). An
	// example of a non-absolute ref would be 'master'.
	//
	// Ex.
	//   refs/heads/master
	FetchRef string `protobuf:"bytes,4,opt,name=fetch_ref,json=fetchRef" json:"fetch_ref,omitempty"`
	// If the checkout had a CL associated with it (i.e. a gerrit commit), this
	// is the fully resolved revision (commit hash) of the CL. If there was no
	// CL, this is empty. Typically the checkout application (e.g. bot_update)
	// rebases this revision on top of the `revision` fetched above.
	//
	// If specified, this must always be a revision on the hosted repo (not any
	// locally generated commit).
	//
	// Ex.
	//   6b0b5c12443cfb93305f8d9e21f8d762c8dad9f0
	PatchRevision string `protobuf:"bytes,5,opt,name=patch_revision,json=patchRevision" json:"patch_revision,omitempty"`
	// If the checkout had a CL associated with it, this is the ref that the
	// task used to fetch patch_revision. If `patch_revision` is supplied, this
	// field is required. If there was no CL, this is empty.
	//
	// If specified, this must always be a ref on the hosted repo (not any local
	// alias like 'refs/remotes/...').
	//
	// This must always be an absolute ref (i.e. starts with 'refs/').
	//
	// Ex.
	//   refs/changes/04/511804/4
	PatchFetchRef        string   `protobuf:"bytes,6,opt,name=patch_fetch_ref,json=patchFetchRef" json:"patch_fetch_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest_GitCheckout) Reset()         { *m = Manifest_GitCheckout{} }
func (m *Manifest_GitCheckout) String() string { return proto.CompactTextString(m) }
func (*Manifest_GitCheckout) ProtoMessage()    {}
func (*Manifest_GitCheckout) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_6c171d9612b862d6, []int{0, 0}
}
func (m *Manifest_GitCheckout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_GitCheckout.Unmarshal(m, b)
}
func (m *Manifest_GitCheckout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_GitCheckout.Marshal(b, m, deterministic)
}
func (dst *Manifest_GitCheckout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_GitCheckout.Merge(dst, src)
}
func (m *Manifest_GitCheckout) XXX_Size() int {
	return xxx_messageInfo_Manifest_GitCheckout.Size(m)
}
func (m *Manifest_GitCheckout) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_GitCheckout.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_GitCheckout proto.InternalMessageInfo

func (m *Manifest_GitCheckout) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *Manifest_GitCheckout) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *Manifest_GitCheckout) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *Manifest_GitCheckout) GetFetchRef() string {
	if m != nil {
		return m.FetchRef
	}
	return ""
}

func (m *Manifest_GitCheckout) GetPatchRevision() string {
	if m != nil {
		return m.PatchRevision
	}
	return ""
}

func (m *Manifest_GitCheckout) GetPatchFetchRef() string {
	if m != nil {
		return m.PatchFetchRef
	}
	return ""
}

type Manifest_CIPDPackage struct {
	// The package pattern that was given to the CIPD client (if known).
	//
	// Ex.
	//   infra/tools/luci/led/${platform}
	PackagePattern string `protobuf:"bytes,1,opt,name=package_pattern,json=packagePattern" json:"package_pattern,omitempty"`
	// The fully resolved instance ID of the deployed package.
	//
	// Ex.
	//   0cfafb3a705bd8f05f86c6444ff500397fbb711c
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	// The unresolved version ID of the deployed package (if known).
	//
	// Ex.
	//   git_revision:aaf3a2cfccc227b5141caa1b6b3502c9907d7420
	//   latest
	Version              string   `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest_CIPDPackage) Reset()         { *m = Manifest_CIPDPackage{} }
func (m *Manifest_CIPDPackage) String() string { return proto.CompactTextString(m) }
func (*Manifest_CIPDPackage) ProtoMessage()    {}
func (*Manifest_CIPDPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_6c171d9612b862d6, []int{0, 1}
}
func (m *Manifest_CIPDPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_CIPDPackage.Unmarshal(m, b)
}
func (m *Manifest_CIPDPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_CIPDPackage.Marshal(b, m, deterministic)
}
func (dst *Manifest_CIPDPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_CIPDPackage.Merge(dst, src)
}
func (m *Manifest_CIPDPackage) XXX_Size() int {
	return xxx_messageInfo_Manifest_CIPDPackage.Size(m)
}
func (m *Manifest_CIPDPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_CIPDPackage.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_CIPDPackage proto.InternalMessageInfo

func (m *Manifest_CIPDPackage) GetPackagePattern() string {
	if m != nil {
		return m.PackagePattern
	}
	return ""
}

func (m *Manifest_CIPDPackage) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Manifest_CIPDPackage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Manifest_Isolated struct {
	// The namespace of the isolated document.
	//
	// Ex.
	//   default-gzip
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	// The hash of the isolated document.
	//
	// Ex.
	//   62a7df62ea122380afb306bb4d9cdac1bc7e9a96
	Hash                 string   `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest_Isolated) Reset()         { *m = Manifest_Isolated{} }
func (m *Manifest_Isolated) String() string { return proto.CompactTextString(m) }
func (*Manifest_Isolated) ProtoMessage()    {}
func (*Manifest_Isolated) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_6c171d9612b862d6, []int{0, 2}
}
func (m *Manifest_Isolated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_Isolated.Unmarshal(m, b)
}
func (m *Manifest_Isolated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_Isolated.Marshal(b, m, deterministic)
}
func (dst *Manifest_Isolated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_Isolated.Merge(dst, src)
}
func (m *Manifest_Isolated) XXX_Size() int {
	return xxx_messageInfo_Manifest_Isolated.Size(m)
}
func (m *Manifest_Isolated) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_Isolated.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_Isolated proto.InternalMessageInfo

func (m *Manifest_Isolated) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Manifest_Isolated) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// A Directory contains one or more descriptions of deployed artifacts. Note
// that due to the practical nature of jobs on bots, it may be the case that
// a given directory contains e.g. a git checkout and multiple cipd packages.
type Manifest_Directory struct {
	GitCheckout *Manifest_GitCheckout `protobuf:"bytes,1,opt,name=git_checkout,json=gitCheckout" json:"git_checkout,omitempty"`
	// The canonicalized hostname of the CIPD server which hosts the CIPD
	// packages (if any). If no CIPD packages are in this Directory, this must
	// be blank.
	//
	// Ex.
	//   chrome-infra-packages.appspot.com
	CipdServerHost string `protobuf:"bytes,2,opt,name=cipd_server_host,json=cipdServerHost" json:"cipd_server_host,omitempty"`
	// Maps CIPD package name to CIPDPackage.
	//
	// Ex.
	//   "some/package/name": {...}
	//   "other/package": {...}
	CipdPackage map[string]*Manifest_CIPDPackage `protobuf:"bytes,4,rep,name=cipd_package,json=cipdPackage" json:"cipd_package,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The canonicalized hostname of the isolated server which hosts the
	// isolated. If no Isolated objects are in this Directory, this must be
	// blank.
	//
	// Ex.
	//   isolateserver.appspot.com
	IsolatedServerHost string `protobuf:"bytes,5,opt,name=isolated_server_host,json=isolatedServerHost" json:"isolated_server_host,omitempty"`
	// A list of all isolateds which have been installed in this directory.
	Isolated             []*Manifest_Isolated `protobuf:"bytes,6,rep,name=isolated" json:"isolated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Manifest_Directory) Reset()         { *m = Manifest_Directory{} }
func (m *Manifest_Directory) String() string { return proto.CompactTextString(m) }
func (*Manifest_Directory) ProtoMessage()    {}
func (*Manifest_Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_6c171d9612b862d6, []int{0, 3}
}
func (m *Manifest_Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_Directory.Unmarshal(m, b)
}
func (m *Manifest_Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_Directory.Marshal(b, m, deterministic)
}
func (dst *Manifest_Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_Directory.Merge(dst, src)
}
func (m *Manifest_Directory) XXX_Size() int {
	return xxx_messageInfo_Manifest_Directory.Size(m)
}
func (m *Manifest_Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_Directory proto.InternalMessageInfo

func (m *Manifest_Directory) GetGitCheckout() *Manifest_GitCheckout {
	if m != nil {
		return m.GitCheckout
	}
	return nil
}

func (m *Manifest_Directory) GetCipdServerHost() string {
	if m != nil {
		return m.CipdServerHost
	}
	return ""
}

func (m *Manifest_Directory) GetCipdPackage() map[string]*Manifest_CIPDPackage {
	if m != nil {
		return m.CipdPackage
	}
	return nil
}

func (m *Manifest_Directory) GetIsolatedServerHost() string {
	if m != nil {
		return m.IsolatedServerHost
	}
	return ""
}

func (m *Manifest_Directory) GetIsolated() []*Manifest_Isolated {
	if m != nil {
		return m.Isolated
	}
	return nil
}

// Links to an externally stored Manifest proto.
type ManifestLink struct {
	// The fully qualified url of the Manifest proto. It's expected that this is
	// a binary logdog stream consisting of exactly one Manifest proto. For now
	// this will always be the `logdog` uri scheme, though it's feasible to put
	// other uri schemes here later.
	//
	// Ex.
	//   logdog://logs.chromium.org/infra/build/12345/+/some/path
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// The hash of the Manifest's raw binary form (i.e. the bytes at the end of
	// `url`, without any interpretation or decoding). Milo will use this as an
	// optimization; Manifests will be interned once into Milo's datastore.
	// Future hashes which match will not be loaded from the url, but will be
	// assumed to be identical. If the sha256 doesn't match the data at the URL,
	// Milo may render this build with the wrong manifest.
	//
	// This is the raw sha256, so it must be exactly 32 bytes.
	Sha256               []byte   `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestLink) Reset()         { *m = ManifestLink{} }
func (m *ManifestLink) String() string { return proto.CompactTextString(m) }
func (*ManifestLink) ProtoMessage()    {}
func (*ManifestLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_6c171d9612b862d6, []int{1}
}
func (m *ManifestLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestLink.Unmarshal(m, b)
}
func (m *ManifestLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestLink.Marshal(b, m, deterministic)
}
func (dst *ManifestLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestLink.Merge(dst, src)
}
func (m *ManifestLink) XXX_Size() int {
	return xxx_messageInfo_ManifestLink.Size(m)
}
func (m *ManifestLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestLink.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestLink proto.InternalMessageInfo

func (m *ManifestLink) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ManifestLink) GetSha256() []byte {
	if m != nil {
		return m.Sha256
	}
	return nil
}

func init() {
	proto.RegisterType((*Manifest)(nil), "milo.Manifest")
	proto.RegisterMapType((map[string]*Manifest_Directory)(nil), "milo.Manifest.DirectoriesEntry")
	proto.RegisterType((*Manifest_GitCheckout)(nil), "milo.Manifest.GitCheckout")
	proto.RegisterType((*Manifest_CIPDPackage)(nil), "milo.Manifest.CIPDPackage")
	proto.RegisterType((*Manifest_Isolated)(nil), "milo.Manifest.Isolated")
	proto.RegisterType((*Manifest_Directory)(nil), "milo.Manifest.Directory")
	proto.RegisterMapType((map[string]*Manifest_CIPDPackage)(nil), "milo.Manifest.Directory.CipdPackageEntry")
	proto.RegisterType((*ManifestLink)(nil), "milo.ManifestLink")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/milo/manifest.proto", fileDescriptor_manifest_6c171d9612b862d6)
}

var fileDescriptor_manifest_6c171d9612b862d6 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x8b, 0xd3, 0x4c,
	0x14, 0xa6, 0xdf, 0xed, 0x49, 0xdf, 0xbe, 0x65, 0x10, 0x8d, 0x51, 0xd8, 0x65, 0x41, 0xad, 0x37,
	0xe9, 0xd2, 0x65, 0x65, 0x11, 0xbd, 0x90, 0xae, 0x1f, 0x85, 0x15, 0x4a, 0x44, 0x10, 0x6f, 0xc2,
	0x38, 0x99, 0x36, 0x43, 0x93, 0x4c, 0x9c, 0x99, 0x14, 0xfa, 0x03, 0xfc, 0x73, 0x82, 0xff, 0x49,
	0x32, 0x99, 0x31, 0xd9, 0xea, 0xde, 0x9d, 0xf3, 0x9c, 0xe7, 0x3c, 0x27, 0xe7, 0x63, 0x02, 0x97,
	0x5b, 0xee, 0x93, 0x58, 0xf0, 0x94, 0x15, 0xa9, 0xcf, 0xc5, 0x76, 0x9e, 0x14, 0x84, 0xcd, 0x09,
	0x4f, 0x53, 0x9e, 0xcd, 0x73, 0xc1, 0x15, 0x9f, 0xa7, 0x2c, 0xe1, 0xf3, 0x14, 0x67, 0x6c, 0x43,
	0xa5, 0xf2, 0x35, 0x86, 0xba, 0x25, 0x78, 0xf6, 0x73, 0x00, 0xc3, 0x8f, 0x26, 0x80, 0x5c, 0x18,
	0xec, 0xa9, 0x90, 0x8c, 0x67, 0x6e, 0xeb, 0xb4, 0x35, 0xeb, 0x05, 0xd6, 0x45, 0x6f, 0xc0, 0x89,
	0x98, 0xa0, 0x44, 0x71, 0xc1, 0xa8, 0x74, 0xdb, 0xa7, 0x9d, 0x99, 0xb3, 0x38, 0xf1, 0x4b, 0x09,
	0xdf, 0xa6, 0xfb, 0xd7, 0x35, 0xe3, 0x6d, 0xa6, 0xc4, 0x21, 0x68, 0xe6, 0x78, 0xbf, 0x5a, 0xe0,
	0xbc, 0x67, 0x6a, 0x19, 0x53, 0xb2, 0xe3, 0x85, 0x42, 0x0f, 0x61, 0x28, 0x68, 0xce, 0xc3, 0x42,
	0x24, 0xba, 0xda, 0x28, 0x18, 0x94, 0xfe, 0x67, 0x91, 0xa0, 0x47, 0x30, 0xda, 0x50, 0x45, 0x62,
	0x1d, 0x6b, 0xeb, 0xd8, 0x50, 0x03, 0x65, 0xd0, 0x2b, 0xf3, 0xf6, 0x4c, 0x7f, 0x65, 0xa7, 0x8a,
	0x59, 0xbf, 0x4e, 0x14, 0x74, 0xe3, 0x76, 0x1b, 0x89, 0x01, 0xdd, 0xa0, 0x27, 0x30, 0xc9, 0x71,
	0x15, 0x34, 0xe9, 0x3d, 0xcd, 0xf8, 0x4f, 0xa3, 0x81, 0xd5, 0x78, 0x0a, 0xff, 0x57, 0xb4, 0x5a,
	0xa9, 0xdf, 0xe0, 0xbd, 0x33, 0x72, 0xde, 0x77, 0x70, 0x96, 0xab, 0xf5, 0xf5, 0x1a, 0x93, 0x1d,
	0xde, 0x52, 0xf4, 0xac, 0x4c, 0xd3, 0x66, 0x98, 0x63, 0xa5, 0xa8, 0xc8, 0x4c, 0x57, 0x13, 0x03,
	0xaf, 0x2b, 0x14, 0x9d, 0x80, 0xc3, 0x32, 0xa9, 0x70, 0x46, 0x68, 0xc8, 0x22, 0xd3, 0x1e, 0x58,
	0x68, 0x15, 0x35, 0xb7, 0x50, 0xf5, 0x67, 0x5d, 0xef, 0x15, 0x0c, 0x57, 0x92, 0x27, 0x58, 0xd1,
	0x08, 0x3d, 0x86, 0x51, 0x86, 0x53, 0x2a, 0x73, 0x4c, 0xa8, 0xa9, 0x54, 0x03, 0x08, 0x41, 0x37,
	0xc6, 0x32, 0x36, 0xea, 0xda, 0xf6, 0x7e, 0x74, 0x60, 0x64, 0x57, 0x74, 0x40, 0xaf, 0x61, 0xbc,
	0x65, 0x2a, 0x24, 0x66, 0x1d, 0x5a, 0xc2, 0x59, 0x78, 0x47, 0x2b, 0x6d, 0x2c, 0x2c, 0x70, 0xb6,
	0x8d, 0xed, 0xcd, 0x60, 0x4a, 0x58, 0x1e, 0x85, 0x92, 0x8a, 0x3d, 0x15, 0x61, 0xcc, 0xa5, 0x32,
	0xc5, 0x26, 0x25, 0xfe, 0x49, 0xc3, 0x1f, 0xb8, 0x54, 0xe8, 0x06, 0xc6, 0x9a, 0x69, 0xc6, 0xe0,
	0x76, 0xf5, 0xed, 0x3c, 0xbf, 0xe3, 0x76, 0x0e, 0xfe, 0x92, 0xe5, 0x91, 0x19, 0xaa, 0xb9, 0x22,
	0x52, 0x23, 0xe8, 0x1c, 0xee, 0x31, 0x33, 0x82, 0x5b, 0xb5, 0xab, 0x55, 0x22, 0x1b, 0x6b, 0xd4,
	0xbf, 0x80, 0xa1, 0x45, 0xdd, 0xbe, 0xae, 0xfd, 0xe0, 0xa8, 0xb6, 0x9d, 0x69, 0xf0, 0x87, 0xe8,
	0x7d, 0x85, 0xe9, 0xf1, 0x77, 0xa0, 0x29, 0x74, 0x76, 0xf4, 0x60, 0x66, 0x5d, 0x9a, 0xe8, 0x1c,
	0x7a, 0x7b, 0x9c, 0x14, 0x54, 0x77, 0xfe, 0xf7, 0xf0, 0x1a, 0xe7, 0x11, 0x54, 0xc4, 0x97, 0xed,
	0xab, 0x96, 0xf7, 0x05, 0xa6, 0xc7, 0x2f, 0xe5, 0x1f, 0xda, 0xfe, 0x6d, 0x6d, 0xf7, 0xae, 0x79,
	0x35, 0x94, 0xcf, 0xae, 0x60, 0x6c, 0x09, 0x37, 0x2c, 0xdb, 0x95, 0xaa, 0xf5, 0xeb, 0x2a, 0x4d,
	0x74, 0x1f, 0xfa, 0x32, 0xc6, 0x8b, 0xcb, 0x17, 0x5a, 0x76, 0x1c, 0x18, 0xef, 0x5b, 0x5f, 0xff,
	0x13, 0x2e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x33, 0xd1, 0x7a, 0x4c, 0x04, 0x00, 0x00,
}
