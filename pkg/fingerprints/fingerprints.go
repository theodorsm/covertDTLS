package fingerprints

//nolint:revive,unused
type ClientHelloFingerprint string

// These fingerprints are added automatically generated and added by the 'fingerprint' workflow
// The first byte should correspond to the DTLS version in a handshake message
const (
	Mozilla_Firefox_125_0_1             ClientHelloFingerprint = "fefda62c8fe5497b56ad1e096f4294cf48c8fe97699406088833f3076ed35bb12b0200000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200" //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_60_unknown ClientHelloFingerprint = "fefdae6064bebc0381a2c7a260cd429b2f5861b9e31425b324dc1a96551bf6cae55500000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                 //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_78_unknown ClientHelloFingerprint = "fefdcbcec29a5919e84b7de392062196c9342da09b5bf45218a2a2d7bdab5ab61ab700000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                 //nolint:revive,stylecheck
	Mozilla_Firefox_125_0_2             ClientHelloFingerprint = "fefdcbce28193f1760cf284985a99eda174720526a7346491e4a9cdff03d9febdc1600000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200" //nolint:revive,stylecheck
)

//nolint:unused
func GetClientHelloFingerprints() []ClientHelloFingerprint {
	return []ClientHelloFingerprint{
		Mozilla_Firefox_125_0_1,             //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_60_unknown, //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_78_unknown, //nolint:revive,stylecheck
		Mozilla_Firefox_125_0_2,             //nolint:revive,stylecheck
	}
}
