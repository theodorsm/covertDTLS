package fingerprints

//nolint:revive,unused
type ClientHelloFingerprint string

// These fingerprints are automatically generated and added by the 'fingerprint' workflow
// The first byte should correspond to the DTLS version in a handshake message
const (
	Mozilla_Firefox_125_0_1              ClientHelloFingerprint = "fefda62c8fe5497b56ad1e096f4294cf48c8fe97699406088833f3076ed35bb12b0200000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200"                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_60_unknown  ClientHelloFingerprint = "fefdae6064bebc0381a2c7a260cd429b2f5861b9e31425b324dc1a96551bf6cae55500000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_78_unknown  ClientHelloFingerprint = "fefdcbcec29a5919e84b7de392062196c9342da09b5bf45218a2a2d7bdab5ab61ab700000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Mozilla_Firefox_125_0_2              ClientHelloFingerprint = "fefdcbce28193f1760cf284985a99eda174720526a7346491e4a9cdff03d9febdc1600000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200"                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_91_unknown  ClientHelloFingerprint = "fefd63c6020e585fe7787a6f56a113f10dffba4ba0360a3d727f81311d4087a2e02a00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_118_unknown ClientHelloFingerprint = "fefd1ab969e081af0198b1a5ef3438947d97e2a15bd879585c28512bb51672b9138700000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Mozilla_Firefox_125_0_3              ClientHelloFingerprint = "fefd58982935730115c95f06618628c3715e56cc0faf990b23ed42b52abf343aebf300000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200"                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_155_unknown ClientHelloFingerprint = "fefd1ad0326777ab4186c02fec9d45c3ba0dd52400bfcc40ffda54284957f7477dee00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_201_unknown ClientHelloFingerprint = "fefd0c6261a6f76af96b849b3cf5f7a57a092d9b1ee1226bedf02ea5bb9bec5e4a3300000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_207_unknown ClientHelloFingerprint = "fefde6a943ff8fd828b850f2f0ae6cc07980651688f9b8815824ff1ac15d5864201e00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Mozilla_Firefox_126_0                ClientHelloFingerprint = "fefda57091eb47f6e652ae5f081f34dbd55aeb27ef635cba4d87e17cb4e476cc8d0c00000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200"                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_125_0_6422_60_unknown  ClientHelloFingerprint = "fefdf04150620a1e523d5d51eb63a641d37a59dc1ec9bd1a2b850bde7fe54ad4dc5200000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_125_0_6422_76_unknown  ClientHelloFingerprint = "fefd54fada6c8da2fe120d8466259368ba52d7c081f8753df3521dc08a0783ff8b3200000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_125_0_6422_112_unknown ClientHelloFingerprint = "fefd59d836c6bebe44212c6a1abe92b7152f54a6c718940a66503307c83aa7c38fa400000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Mozilla_Firefox_126_0_1              ClientHelloFingerprint = "fefd89c2d091fdab55e258926097978a197f88fd85d1689da875e688ffe66589996000000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200"                                                                                                                                                       //nolint:revive,stylecheck
	Google_Chrome_125_0_6422_141_unknown ClientHelloFingerprint = "fefdbbe11878d02e5cd2ba9c6c12c05845330adc294052b81d7e392ba480ad8d9dd400000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Chrome_linux_125_0_6422_141          ClientHelloFingerprint = "fefd46d25ef57649ecfd0fc17d5d933462d70770ea629a4d74a9b7567cc5714fd59500000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Firefox_linux_stable_126_0_1         ClientHelloFingerprint = "fefd456d7850288d8a38e422000b4d6b94d96af38ae82926100738c531a7b531873b00000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200"                                                                                                                                                       //nolint:revive,stylecheck
	Chrome_linux_126_0_6478_55           ClientHelloFingerprint = "fefd9460e1e8af6afd0c07728ece3350accdb7be89d2a4ebcd9a8d5dc6e780eaeb6000000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Firefox_linux_stable_127_0           ClientHelloFingerprint = "fefd713910e31aa451cda18c437634c566fb872731739ada472eafa74741c7d02c8c0000001413011303c02bc02fcca9cca8c00ac009c013c014010000b100170000ff01000100000a000c000a001d0017001801000101000b000201000010001200100677656272746308632d7765627274630022000a00080403050306030203003300260024001d00209b3a0682a66a92872457a05dfd4ce6f92312ef605e960f36b8fd71f054b60540002b000706fefcfefd0303000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024001000e000b0008000700080001000200" //nolint:revive,stylecheck
	Chrome_linux_126_0_6478_61           ClientHelloFingerprint = "fefdd55ec527c484ad54a31786a4536060866e8f817605893dc5e6a9b2540ab2e38e00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	Chrome_linux_126_0_6478_62           ClientHelloFingerprint = "fefd5f58909b796cbeab9075ec3239db3e9b60732044a7714180ffb3999997df7d0f00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	chrome_linux_126_0_6478_62           ClientHelloFingerprint = "fefdadab55572d5960d845c354813a1a351d924350c11bbf7781ccecdb9b0817ded400000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	firefox_linux_stable_127_0_1         ClientHelloFingerprint = "fefde650a0e526f3b0a64bb3547c83b9bb49f7c3e2dd1fb8585b8b1a30c233e30c960000001413011303c02bc02fcca9cca8c00ac009c013c014010000b100170000ff01000100000a000c000a001d0017001801000101000b000201000010001200100677656272746308632d7765627274630022000a00080403050306030203003300260024001d00203c07b94e15938de0e284acde46a1c7ea20da8c8b2256a3d1a04a3e86e150012a002b000706fefcfefd0303000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024001000e000b0008000700080001000200" //nolint:revive,stylecheck
	chrome_linux_126_0_6478_63           ClientHelloFingerprint = "fefd7ff07ea53e6848da7864b6083c5686ae82354fcf9e38cdfe20b095eaa783672b00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
	chrome_linux_126_0_6478_126          ClientHelloFingerprint = "fefd06c93bd11982ae5e0ae3a93fff4dc1fac9bbf8fa354b30eaf88d10eae4ca6f2200000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                                                                                                                                                                       //nolint:revive,stylecheck
)

//nolint:unused
func GetClientHelloFingerprints() []ClientHelloFingerprint {
	return []ClientHelloFingerprint{
		Mozilla_Firefox_125_0_1,              //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_60_unknown,  //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_78_unknown,  //nolint:revive,stylecheck
		Mozilla_Firefox_125_0_2,              //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_91_unknown,  //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_118_unknown, //nolint:revive,stylecheck
		Mozilla_Firefox_125_0_3,              //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_155_unknown, //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_201_unknown, //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_207_unknown, //nolint:revive,stylecheck
		Mozilla_Firefox_126_0,                //nolint:revive,stylecheck
		Google_Chrome_125_0_6422_60_unknown,  //nolint:revive,stylecheck
		Google_Chrome_125_0_6422_76_unknown,  //nolint:revive,stylecheck
		Google_Chrome_125_0_6422_112_unknown, //nolint:revive,stylecheck
		Mozilla_Firefox_126_0_1,              //nolint:revive,stylecheck
		Google_Chrome_125_0_6422_141_unknown, //nolint:revive,stylecheck
		Chrome_linux_125_0_6422_141,          //nolint:revive,stylecheck
		Firefox_linux_stable_126_0_1,         //nolint:revive,stylecheck
		Chrome_linux_126_0_6478_55,           //nolint:revive,stylecheck
		Firefox_linux_stable_127_0,           //nolint:revive,stylecheck
		Chrome_linux_126_0_6478_61,           //nolint:revive,stylecheck
		Chrome_linux_126_0_6478_62,           //nolint:revive,stylecheck
		chrome_linux_126_0_6478_62,           //nolint:revive,stylecheck
		firefox_linux_stable_127_0_1,         //nolint:revive,stylecheck
		chrome_linux_126_0_6478_63,           //nolint:revive,stylecheck
		chrome_linux_126_0_6478_126,          //nolint:revive,stylecheck
	}
}
