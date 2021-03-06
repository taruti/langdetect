package langdetect

var trigram_et = trigramfreqs{
	trigramfreq{val: 0x6161, freq: 0x2b},
	trigramfreq{val: 0x616c, freq: 0x3a},
	trigramfreq{val: 0x6172, freq: 0x59},
	trigramfreq{val: 0x6173, freq: 0x28},
	trigramfreq{val: 0x6565, freq: 0x52},
	trigramfreq{val: 0x6569, freq: 0x57},
	trigramfreq{val: 0x656c, freq: 0x33},
	trigramfreq{val: 0x656e, freq: 0x2e},
	trigramfreq{val: 0x6572, freq: 0x26},
	trigramfreq{val: 0x6573, freq: 0x2c},
	trigramfreq{val: 0x6574, freq: 0xd1},
	trigramfreq{val: 0x6575, freq: 0x75},
	trigramfreq{val: 0x68c3, freq: 0x46},
	trigramfreq{val: 0x696e, freq: 0x3b},
	trigramfreq{val: 0x6a61, freq: 0x126},
	trigramfreq{val: 0x6a75, freq: 0x4f},
	trigramfreq{val: 0x6ac3, freq: 0x48},
	trigramfreq{val: 0x6b61, freq: 0xdb},
	trigramfreq{val: 0x6b65, freq: 0x4c},
	trigramfreq{val: 0x6b69, freq: 0x33},
	trigramfreq{val: 0x6b6f, freq: 0x11a},
	trigramfreq{val: 0x6b75, freq: 0xa7},
	trigramfreq{val: 0x6bc3, freq: 0x94},
	trigramfreq{val: 0x6c61, freq: 0x2c},
	trigramfreq{val: 0x6c65, freq: 0x26},
	trigramfreq{val: 0x6c69, freq: 0x79},
	trigramfreq{val: 0x6cc3, freq: 0x46},
	trigramfreq{val: 0x6d61, freq: 0x7b},
	trigramfreq{val: 0x6d65, freq: 0xa2},
	trigramfreq{val: 0x6d69, freq: 0xd6},
	trigramfreq{val: 0x6d75, freq: 0x41},
	trigramfreq{val: 0x6dc3, freq: 0x52},
	trigramfreq{val: 0x6e61, freq: 0x34},
	trigramfreq{val: 0x6e65, freq: 0x43},
	trigramfreq{val: 0x6e69, freq: 0x74},
	trigramfreq{val: 0x6ec3, freq: 0x5a},
	trigramfreq{val: 0x6f6c, freq: 0xa1},
	trigramfreq{val: 0x6f6d, freq: 0x2b},
	trigramfreq{val: 0x6f6e, freq: 0xf9},
	trigramfreq{val: 0x7061, freq: 0x7c},
	trigramfreq{val: 0x7065, freq: 0x60},
	trigramfreq{val: 0x7069, freq: 0x39},
	trigramfreq{val: 0x706f, freq: 0x44},
	trigramfreq{val: 0x7072, freq: 0x5c},
	trigramfreq{val: 0x70c3, freq: 0x5d},
	trigramfreq{val: 0x7261, freq: 0x71},
	trigramfreq{val: 0x7265, freq: 0x27},
	trigramfreq{val: 0x7269, freq: 0x32},
	trigramfreq{val: 0x7361, freq: 0x75},
	trigramfreq{val: 0x7365, freq: 0x151},
	trigramfreq{val: 0x7369, freq: 0x4c},
	trigramfreq{val: 0x736f, freq: 0x31},
	trigramfreq{val: 0x7375, freq: 0x5a},
	trigramfreq{val: 0x73c3, freq: 0x37},
	trigramfreq{val: 0x7461, freq: 0x77},
	trigramfreq{val: 0x7465, freq: 0xe9},
	trigramfreq{val: 0x746f, freq: 0x4a},
	trigramfreq{val: 0x7475, freq: 0x67},
	trigramfreq{val: 0x74c3, freq: 0xb8},
	trigramfreq{val: 0x7661, freq: 0xde},
	trigramfreq{val: 0x7665, freq: 0x25},
	trigramfreq{val: 0x76c3, freq: 0xed},
	trigramfreq{val: 0xc3b5, freq: 0x40},
	trigramfreq{val: 0xc3bc, freq: 0x8d},
	trigramfreq{val: 0x610061, freq: 0x4c},
	trigramfreq{val: 0x610065, freq: 0x8e},
	trigramfreq{val: 0x610068, freq: 0x2c},
	trigramfreq{val: 0x61006a, freq: 0x4f},
	trigramfreq{val: 0x61006b, freq: 0xc1},
	trigramfreq{val: 0x61006c, freq: 0x55},
	trigramfreq{val: 0x61006d, freq: 0x86},
	trigramfreq{val: 0x61006e, freq: 0x3f},
	trigramfreq{val: 0x61006f, freq: 0x47},
	trigramfreq{val: 0x610070, freq: 0x78},
	trigramfreq{val: 0x610072, freq: 0x36},
	trigramfreq{val: 0x610073, freq: 0x97},
	trigramfreq{val: 0x610074, freq: 0x97},
	trigramfreq{val: 0x610076, freq: 0x64},
	trigramfreq{val: 0x6100c3, freq: 0x37},
	trigramfreq{val: 0x61616c, freq: 0x37},
	trigramfreq{val: 0x616173, freq: 0x46},
	trigramfreq{val: 0x616200, freq: 0x5d},
	trigramfreq{val: 0x616400, freq: 0x92},
	trigramfreq{val: 0x616461, freq: 0x75},
	trigramfreq{val: 0x616465, freq: 0x40},
	trigramfreq{val: 0x616761, freq: 0x38},
	trigramfreq{val: 0x616865, freq: 0x43},
	trigramfreq{val: 0x616964, freq: 0x2a},
	trigramfreq{val: 0x616a61, freq: 0x7f},
	trigramfreq{val: 0x616b73, freq: 0x8f},
	trigramfreq{val: 0x616c00, freq: 0x41},
	trigramfreq{val: 0x616c64, freq: 0x40},
	trigramfreq{val: 0x616c69, freq: 0x76},
	trigramfreq{val: 0x616c74, freq: 0x30},
	trigramfreq{val: 0x616c75, freq: 0x38},
	trigramfreq{val: 0x616d61, freq: 0x51},
	trigramfreq{val: 0x616d65, freq: 0x6b},
	trigramfreq{val: 0x616d69, freq: 0x9b},
	trigramfreq{val: 0x616e00, freq: 0x2c},
	trigramfreq{val: 0x616e64, freq: 0x6a},
	trigramfreq{val: 0x616e65, freq: 0x2e},
	trigramfreq{val: 0x616e69, freq: 0x34},
	trigramfreq{val: 0x616e75, freq: 0x29},
	trigramfreq{val: 0x61706f, freq: 0x27},
	trigramfreq{val: 0x61726c, freq: 0x26},
	trigramfreq{val: 0x617276, freq: 0x32},
	trigramfreq{val: 0x617300, freq: 0x77},
	trigramfreq{val: 0x617365, freq: 0x38},
	trigramfreq{val: 0x617369, freq: 0x2e},
	trigramfreq{val: 0x617374, freq: 0xd3},
	trigramfreq{val: 0x617375, freq: 0x2f},
	trigramfreq{val: 0x617400, freq: 0x27},
	trigramfreq{val: 0x617461, freq: 0x6c},
	trigramfreq{val: 0x617465, freq: 0x66},
	trigramfreq{val: 0x617475, freq: 0x6c},
	trigramfreq{val: 0x617661, freq: 0x89},
	trigramfreq{val: 0x640065, freq: 0x45},
	trigramfreq{val: 0x64006a, freq: 0x35},
	trigramfreq{val: 0x64006b, freq: 0x5e},
	trigramfreq{val: 0x64006d, freq: 0x48},
	trigramfreq{val: 0x64006f, freq: 0x36},
	trigramfreq{val: 0x640070, freq: 0x31},
	trigramfreq{val: 0x640073, freq: 0x47},
	trigramfreq{val: 0x640074, freq: 0x40},
	trigramfreq{val: 0x640076, freq: 0x31},
	trigramfreq{val: 0x646100, freq: 0x10b},
	trigramfreq{val: 0x646164, freq: 0x2c},
	trigramfreq{val: 0x64616d, freq: 0x3a},
	trigramfreq{val: 0x646500, freq: 0x9c},
	trigramfreq{val: 0x64656c, freq: 0x28},
	trigramfreq{val: 0x646573, freq: 0x42},
	trigramfreq{val: 0x646900, freq: 0x2d},
	trigramfreq{val: 0x647573, freq: 0x89},
	trigramfreq{val: 0x650061, freq: 0x5e},
	trigramfreq{val: 0x650065, freq: 0x92},
	trigramfreq{val: 0x650068, freq: 0x28},
	trigramfreq{val: 0x65006a, freq: 0x70},
	trigramfreq{val: 0x65006b, freq: 0xd6},
	trigramfreq{val: 0x65006c, freq: 0x3d},
	trigramfreq{val: 0x65006d, freq: 0x6f},
	trigramfreq{val: 0x65006e, freq: 0x3c},
	trigramfreq{val: 0x65006f, freq: 0x61},
	trigramfreq{val: 0x650070, freq: 0x7c},
	trigramfreq{val: 0x650072, freq: 0x3a},
	trigramfreq{val: 0x650073, freq: 0x89},
	trigramfreq{val: 0x650074, freq: 0x9b},
	trigramfreq{val: 0x650076, freq: 0x80},
	trigramfreq{val: 0x6500c3, freq: 0x3a},
	trigramfreq{val: 0x656400, freq: 0x7b},
	trigramfreq{val: 0x656461, freq: 0x40},
	trigramfreq{val: 0x656500, freq: 0x4e},
	trigramfreq{val: 0x65656c, freq: 0x41},
	trigramfreq{val: 0x65656d, freq: 0x28},
	trigramfreq{val: 0x656572, freq: 0x3a},
	trigramfreq{val: 0x656573, freq: 0x3e},
	trigramfreq{val: 0x656574, freq: 0x29},
	trigramfreq{val: 0x656761, freq: 0x66},
	trigramfreq{val: 0x656765, freq: 0x29},
	trigramfreq{val: 0x656900, freq: 0x55},
	trigramfreq{val: 0x656964, freq: 0x45},
	trigramfreq{val: 0x656965, freq: 0x2a},
	trigramfreq{val: 0x656973, freq: 0x32},
	trigramfreq{val: 0x656b73, freq: 0x7e},
	trigramfreq{val: 0x656c00, freq: 0x94},
	trigramfreq{val: 0x656c65, freq: 0x96},
	trigramfreq{val: 0x656c69, freq: 0x55},
	trigramfreq{val: 0x656c6c, freq: 0x73},
	trigramfreq{val: 0x656c74, freq: 0x3f},
	trigramfreq{val: 0x656d61, freq: 0x60},
	trigramfreq{val: 0x656d69, freq: 0x40},
	trigramfreq{val: 0x656e64, freq: 0xbe},
	trigramfreq{val: 0x656e65, freq: 0x34},
	trigramfreq{val: 0x656e74, freq: 0x32},
	trigramfreq{val: 0x657061, freq: 0x25},
	trigramfreq{val: 0x657269, freq: 0x59},
	trigramfreq{val: 0x657300, freq: 0xd9},
	trigramfreq{val: 0x657369, freq: 0x4c},
	trigramfreq{val: 0x657374, freq: 0xb5},
	trigramfreq{val: 0x657400, freq: 0xca},
	trigramfreq{val: 0x657461, freq: 0x41},
	trigramfreq{val: 0x657465, freq: 0x3a},
	trigramfreq{val: 0x657474, freq: 0x33},
	trigramfreq{val: 0x657475, freq: 0x31},
	trigramfreq{val: 0x657572, freq: 0x75},
	trigramfreq{val: 0x657661, freq: 0x4d},
	trigramfreq{val: 0x676100, freq: 0xb8},
	trigramfreq{val: 0x676900, freq: 0x46},
	trigramfreq{val: 0x677500, freq: 0x50},
	trigramfreq{val: 0x677573, freq: 0x58},
	trigramfreq{val: 0x68656c, freq: 0x29},
	trigramfreq{val: 0x68656e, freq: 0x42},
	trigramfreq{val: 0x68c3a4, freq: 0x40},
	trigramfreq{val: 0x690061, freq: 0x28},
	trigramfreq{val: 0x690065, freq: 0x41},
	trigramfreq{val: 0x69006a, freq: 0x2a},
	trigramfreq{val: 0x69006b, freq: 0x59},
	trigramfreq{val: 0x69006c, freq: 0x26},
	trigramfreq{val: 0x69006d, freq: 0x37},
	trigramfreq{val: 0x69006f, freq: 0x39},
	trigramfreq{val: 0x690070, freq: 0x3d},
	trigramfreq{val: 0x690073, freq: 0x49},
	trigramfreq{val: 0x690074, freq: 0x48},
	trigramfreq{val: 0x690076, freq: 0x35},
	trigramfreq{val: 0x696100, freq: 0x26},
	trigramfreq{val: 0x696400, freq: 0xdf},
	trigramfreq{val: 0x696461, freq: 0x5c},
	trigramfreq{val: 0x696465, freq: 0x96},
	trigramfreq{val: 0x696475, freq: 0x2e},
	trigramfreq{val: 0x696500, freq: 0x2a},
	trigramfreq{val: 0x696761, freq: 0x25},
	trigramfreq{val: 0x696769, freq: 0x33},
	trigramfreq{val: 0x696775, freq: 0x3e},
	trigramfreq{val: 0x696967, freq: 0x2a},
	trigramfreq{val: 0x69696b, freq: 0x65},
	trigramfreq{val: 0x696972, freq: 0x29},
	trigramfreq{val: 0x696973, freq: 0x31},
	trigramfreq{val: 0x696974, freq: 0x35},
	trigramfreq{val: 0x696b00, freq: 0x41},
	trigramfreq{val: 0x696b61, freq: 0x38},
	trigramfreq{val: 0x696b65, freq: 0x29},
	trigramfreq{val: 0x696b69, freq: 0x31},
	trigramfreq{val: 0x696b6b, freq: 0x2b},
	trigramfreq{val: 0x696b6d, freq: 0x26},
	trigramfreq{val: 0x696b75, freq: 0x7c},
	trigramfreq{val: 0x696c00, freq: 0x33},
	trigramfreq{val: 0x696c65, freq: 0x2b},
	trigramfreq{val: 0x696c69, freq: 0x3c},
	trigramfreq{val: 0x696c6c, freq: 0x27},
	trigramfreq{val: 0x696d61, freq: 0x4a},
	trigramfreq{val: 0x696d65, freq: 0x55},
	trigramfreq{val: 0x696d69, freq: 0x42},
	trigramfreq{val: 0x696d75, freq: 0x36},
	trigramfreq{val: 0x696e00, freq: 0x31},
	trigramfreq{val: 0x696e64, freq: 0x27},
	trigramfreq{val: 0x696e65, freq: 0x85},
	trigramfreq{val: 0x696e67, freq: 0x7f},
	trigramfreq{val: 0x696e69, freq: 0x46},
	trigramfreq{val: 0x696e75, freq: 0x28},
	trigramfreq{val: 0x696f6f, freq: 0x44},
	trigramfreq{val: 0x697300, freq: 0xa1},
	trigramfreq{val: 0x697365, freq: 0x143},
	trigramfreq{val: 0x697369, freq: 0x2b},
	trigramfreq{val: 0x69736a, freq: 0x32},
	trigramfreq{val: 0x697374, freq: 0xdc},
	trigramfreq{val: 0x697461, freq: 0x56},
	trigramfreq{val: 0x697465, freq: 0x3f},
	trigramfreq{val: 0x697469, freq: 0x3a},
	trigramfreq{val: 0x697473, freq: 0x3b},
	trigramfreq{val: 0x697475, freq: 0x2d},
	trigramfreq{val: 0x6a6100, freq: 0x14a},
	trigramfreq{val: 0x6a616c, freq: 0x2b},
	trigramfreq{val: 0x6a6f6e, freq: 0x37},
	trigramfreq{val: 0x6a7568, freq: 0x2e},
	trigramfreq{val: 0x6ac3a4, freq: 0x44},
	trigramfreq{val: 0x6b6100, freq: 0x4d},
	trigramfreq{val: 0x6b6173, freq: 0x42},
	trigramfreq{val: 0x6b6500, freq: 0x26},
	trigramfreq{val: 0x6b6573, freq: 0x36},
	trigramfreq{val: 0x6b6964, freq: 0x30},
	trigramfreq{val: 0x6b6b75, freq: 0x31},
	trigramfreq{val: 0x6b6d65, freq: 0x26},
	trigramfreq{val: 0x6b6f67, freq: 0x2b},
	trigramfreq{val: 0x6b6f68, freq: 0x5e},
	trigramfreq{val: 0x6b6f6d, freq: 0x3a},
	trigramfreq{val: 0x6b6f6e, freq: 0x5c},
	trigramfreq{val: 0x6b6f6f, freq: 0x29},
	trigramfreq{val: 0x6b6f72, freq: 0x31},
	trigramfreq{val: 0x6b7300, freq: 0xeb},
	trigramfreq{val: 0x6b7365, freq: 0x28},
	trigramfreq{val: 0x6b7369, freq: 0x33},
	trigramfreq{val: 0x6b7469, freq: 0x29},
	trigramfreq{val: 0x6b7500, freq: 0x38},
	trigramfreq{val: 0x6b7569, freq: 0x66},
	trigramfreq{val: 0x6b756c, freq: 0x3b},
	trigramfreq{val: 0x6b7573, freq: 0x32},
	trigramfreq{val: 0x6bc3b5, freq: 0x55},
	trigramfreq{val: 0x6bc3bc, freq: 0x2f},
	trigramfreq{val: 0x6c006f, freq: 0x2a},
	trigramfreq{val: 0x6c616d, freq: 0x2a},
	trigramfreq{val: 0x6c6173, freq: 0x27},
	trigramfreq{val: 0x6c6461, freq: 0x35},
	trigramfreq{val: 0x6c6500, freq: 0x13d},
	trigramfreq{val: 0x6c656b, freq: 0x2b},
	trigramfreq{val: 0x6c656d, freq: 0x4b},
	trigramfreq{val: 0x6c6570, freq: 0x28},
	trigramfreq{val: 0x6c6573, freq: 0x3d},
	trigramfreq{val: 0x6c6576, freq: 0x28},
	trigramfreq{val: 0x6c6900, freq: 0x3e},
	trigramfreq{val: 0x6c6969, freq: 0x88},
	trigramfreq{val: 0x6c696b, freq: 0x9f},
	trigramfreq{val: 0x6c696e, freq: 0x32},
	trigramfreq{val: 0x6c6973, freq: 0xa2},
	trigramfreq{val: 0x6c6c65, freq: 0x91},
	trigramfreq{val: 0x6c6c69, freq: 0x2a},
	trigramfreq{val: 0x6c6f6f, freq: 0x2e},
	trigramfreq{val: 0x6c7400, freq: 0xaa},
	trigramfreq{val: 0x6c7573, freq: 0x58},
	trigramfreq{val: 0x6d6100, freq: 0xd8},
	trigramfreq{val: 0x6d6161, freq: 0x29},
	trigramfreq{val: 0x6d616c, freq: 0x36},
	trigramfreq{val: 0x6d6173, freq: 0x29},
	trigramfreq{val: 0x6d6174, freq: 0x2e},
	trigramfreq{val: 0x6d6500, freq: 0xb6},
	trigramfreq{val: 0x6d6565, freq: 0x29},
	trigramfreq{val: 0x6d6569, freq: 0x45},
	trigramfreq{val: 0x6d656e, freq: 0x37},
	trigramfreq{val: 0x6d6573, freq: 0x39},
	trigramfreq{val: 0x6d6964, freq: 0x3a},
	trigramfreq{val: 0x6d696c, freq: 0x32},
	trigramfreq{val: 0x6d696e, freq: 0x67},
	trigramfreq{val: 0x6d6973, freq: 0x196},
	trigramfreq{val: 0x6d6974, freq: 0x28},
	trigramfreq{val: 0x6d7573, freq: 0x4b},
	trigramfreq{val: 0x6dc3a4, freq: 0x35},
	trigramfreq{val: 0x6dc3b5, freq: 0x43},
	trigramfreq{val: 0x6e0065, freq: 0x2e},
	trigramfreq{val: 0x6e006b, freq: 0x27},
	trigramfreq{val: 0x6e0073, freq: 0x30},
	trigramfreq{val: 0x6e0074, freq: 0x2d},
	trigramfreq{val: 0x6e0076, freq: 0x26},
	trigramfreq{val: 0x6e6100, freq: 0x41},
	trigramfreq{val: 0x6e6461, freq: 0x85},
	trigramfreq{val: 0x6e6465, freq: 0x2d},
	trigramfreq{val: 0x6e6469, freq: 0x45},
	trigramfreq{val: 0x6e6475, freq: 0x50},
	trigramfreq{val: 0x6e6500, freq: 0x98},
	trigramfreq{val: 0x6e6700, freq: 0x4e},
	trigramfreq{val: 0x6e6775, freq: 0x32},
	trigramfreq{val: 0x6e6900, freq: 0x5b},
	trigramfreq{val: 0x6e696b, freq: 0x2a},
	trigramfreq{val: 0x6e696d, freq: 0x34},
	trigramfreq{val: 0x6e696e, freq: 0x4c},
	trigramfreq{val: 0x6e6e61, freq: 0x39},
	trigramfreq{val: 0x6e7564, freq: 0x57},
	trigramfreq{val: 0x6ec3a4, freq: 0x27},
	trigramfreq{val: 0x6ec3b5, freq: 0x30},
	trigramfreq{val: 0x6f6775, freq: 0x29},
	trigramfreq{val: 0x6f6874, freq: 0x34},
	trigramfreq{val: 0x6f6c65, freq: 0x7e},
	trigramfreq{val: 0x6f6c69, freq: 0x58},
	trigramfreq{val: 0x6f6c6c, freq: 0x2a},
	trigramfreq{val: 0x6f6c75, freq: 0x37},
	trigramfreq{val: 0x6f6d61, freq: 0x30},
	trigramfreq{val: 0x6f6d69, freq: 0x3f},
	trigramfreq{val: 0x6f6e00, freq: 0x116},
	trigramfreq{val: 0x6f6e69, freq: 0x68},
	trigramfreq{val: 0x6f6f6c, freq: 0x32},
	trigramfreq{val: 0x6f6f6e, freq: 0x4a},
	trigramfreq{val: 0x6f6f70, freq: 0x70},
	trigramfreq{val: 0x6f7061, freq: 0x6d},
	trigramfreq{val: 0x6f7274, freq: 0x2e},
	trigramfreq{val: 0x6f7473, freq: 0x2b},
	trigramfreq{val: 0x706100, freq: 0x60},
	trigramfreq{val: 0x70616e, freq: 0x2f},
	trigramfreq{val: 0x706172, freq: 0x43},
	trigramfreq{val: 0x706561, freq: 0x55},
	trigramfreq{val: 0x706969, freq: 0x26},
	trigramfreq{val: 0x706f6c, freq: 0x2f},
	trigramfreq{val: 0x706f72, freq: 0x34},
	trigramfreq{val: 0x70726f, freq: 0x40},
	trigramfreq{val: 0x70c3a4, freq: 0x2e},
	trigramfreq{val: 0x70c3b5, freq: 0x36},
	trigramfreq{val: 0x726100, freq: 0x26},
	trigramfreq{val: 0x726168, freq: 0x28},
	trigramfreq{val: 0x726173, freq: 0x28},
	trigramfreq{val: 0x72656e, freq: 0x27},
	trigramfreq{val: 0x726969, freq: 0x5c},
	trigramfreq{val: 0x726974, freq: 0x31},
	trigramfreq{val: 0x726f6f, freq: 0x73},
	trigramfreq{val: 0x727261, freq: 0x3a},
	trigramfreq{val: 0x730061, freq: 0x25},
	trigramfreq{val: 0x730065, freq: 0x50},
	trigramfreq{val: 0x73006a, freq: 0x2d},
	trigramfreq{val: 0x73006b, freq: 0x69},
	trigramfreq{val: 0x73006d, freq: 0x4a},
	trigramfreq{val: 0x73006e, freq: 0x29},
	trigramfreq{val: 0x73006f, freq: 0x62},
	trigramfreq{val: 0x730070, freq: 0x3b},
	trigramfreq{val: 0x730073, freq: 0x58},
	trigramfreq{val: 0x730074, freq: 0x4c},
	trigramfreq{val: 0x730076, freq: 0x3a},
	trigramfreq{val: 0x736161, freq: 0x45},
	trigramfreq{val: 0x73616d, freq: 0x33},
	trigramfreq{val: 0x736500, freq: 0x15d},
	trigramfreq{val: 0x736564, freq: 0x73},
	trigramfreq{val: 0x736565, freq: 0x71},
	trigramfreq{val: 0x736569, freq: 0x28},
	trigramfreq{val: 0x73656b, freq: 0x46},
	trigramfreq{val: 0x73656c, freq: 0xff},
	trigramfreq{val: 0x736573, freq: 0x5c},
	trigramfreq{val: 0x736574, freq: 0x2d},
	trigramfreq{val: 0x736900, freq: 0x59},
	trigramfreq{val: 0x736964, freq: 0x3c},
	trigramfreq{val: 0x736969, freq: 0x30},
	trigramfreq{val: 0x73696d, freq: 0x3d},
	trigramfreq{val: 0x73696e, freq: 0x2a},
	trigramfreq{val: 0x73696f, freq: 0x45},
	trigramfreq{val: 0x736973, freq: 0x31},
	trigramfreq{val: 0x736974, freq: 0x2f},
	trigramfreq{val: 0x736a6f, freq: 0x32},
	trigramfreq{val: 0x737365, freq: 0x2d},
	trigramfreq{val: 0x737400, freq: 0x13f},
	trigramfreq{val: 0x737461, freq: 0xd7},
	trigramfreq{val: 0x737465, freq: 0xc2},
	trigramfreq{val: 0x737469, freq: 0x44},
	trigramfreq{val: 0x737475, freq: 0x75},
	trigramfreq{val: 0x737573, freq: 0x47},
	trigramfreq{val: 0x737574, freq: 0x25},
	trigramfreq{val: 0x737575, freq: 0x3c},
	trigramfreq{val: 0x740065, freq: 0x49},
	trigramfreq{val: 0x74006a, freq: 0x28},
	trigramfreq{val: 0x74006b, freq: 0x66},
	trigramfreq{val: 0x74006d, freq: 0x54},
	trigramfreq{val: 0x74006e, freq: 0x29},
	trigramfreq{val: 0x74006f, freq: 0x2e},
	trigramfreq{val: 0x740070, freq: 0x3a},
	trigramfreq{val: 0x740073, freq: 0x57},
	trigramfreq{val: 0x740074, freq: 0x4c},
	trigramfreq{val: 0x740076, freq: 0x32},
	trigramfreq{val: 0x746100, freq: 0x7f},
	trigramfreq{val: 0x746164, freq: 0x45},
	trigramfreq{val: 0x74616b, freq: 0x2e},
	trigramfreq{val: 0x74616d, freq: 0x6d},
	trigramfreq{val: 0x746173, freq: 0x32},
	trigramfreq{val: 0x746174, freq: 0x47},
	trigramfreq{val: 0x746176, freq: 0x3c},
	trigramfreq{val: 0x746500, freq: 0x10b},
	trigramfreq{val: 0x746561, freq: 0x27},
	trigramfreq{val: 0x746565, freq: 0x4f},
	trigramfreq{val: 0x746567, freq: 0x4d},
	trigramfreq{val: 0x746569, freq: 0x37},
	trigramfreq{val: 0x74656c, freq: 0x50},
	trigramfreq{val: 0x746573, freq: 0x50},
	trigramfreq{val: 0x746900, freq: 0x6c},
	trigramfreq{val: 0x74696b, freq: 0x2e},
	trigramfreq{val: 0x746c65, freq: 0x2b},
	trigramfreq{val: 0x747365, freq: 0x47},
	trigramfreq{val: 0x747369, freq: 0x68},
	trigramfreq{val: 0x747375, freq: 0x2c},
	trigramfreq{val: 0x747465, freq: 0x58},
	trigramfreq{val: 0x747500, freq: 0x36},
	trigramfreq{val: 0x747564, freq: 0x71},
	trigramfreq{val: 0x74756c, freq: 0x3c},
	trigramfreq{val: 0x747573, freq: 0x9f},
	trigramfreq{val: 0x74c3a4, freq: 0x5a},
	trigramfreq{val: 0x74c3b5, freq: 0x3d},
	trigramfreq{val: 0x74c3b6, freq: 0x44},
	trigramfreq{val: 0x75006b, freq: 0x26},
	trigramfreq{val: 0x756400, freq: 0xeb},
	trigramfreq{val: 0x756900, freq: 0x3d},
	trigramfreq{val: 0x756964, freq: 0x2b},
	trigramfreq{val: 0x756b6f, freq: 0x31},
	trigramfreq{val: 0x756c00, freq: 0x2a},
	trigramfreq{val: 0x756c65, freq: 0x4b},
	trigramfreq{val: 0x756c69, freq: 0x2e},
	trigramfreq{val: 0x756c74, freq: 0x36},
	trigramfreq{val: 0x756d69, freq: 0x32},
	trigramfreq{val: 0x757265, freq: 0x2b},
	trigramfreq{val: 0x75726f, freq: 0x75},
	trigramfreq{val: 0x757300, freq: 0x89},
	trigramfreq{val: 0x757365, freq: 0x102},
	trigramfreq{val: 0x757369, freq: 0x2c},
	trigramfreq{val: 0x757374, freq: 0x100},
	trigramfreq{val: 0x757461, freq: 0x57},
	trigramfreq{val: 0x757465, freq: 0x27},
	trigramfreq{val: 0x757475, freq: 0x32},
	trigramfreq{val: 0x757572, freq: 0x42},
	trigramfreq{val: 0x766164, freq: 0x5b},
	trigramfreq{val: 0x766168, freq: 0x2c},
	trigramfreq{val: 0x76616c, freq: 0x6f},
	trigramfreq{val: 0x766173, freq: 0x50},
	trigramfreq{val: 0x766174, freq: 0x3b},
	trigramfreq{val: 0x766969, freq: 0x27},
	trigramfreq{val: 0x767573, freq: 0x2c},
	trigramfreq{val: 0x76c3a4, freq: 0x73},
	trigramfreq{val: 0x76c3b5, freq: 0xab},
	trigramfreq{val: 0xa46865, freq: 0x2f},
	trigramfreq{val: 0xa47261, freq: 0x28},
	trigramfreq{val: 0xa4c3a4, freq: 0x64},
	trigramfreq{val: 0xb56967, freq: 0x5d},
	trigramfreq{val: 0xb5696d, freq: 0x2d},
	trigramfreq{val: 0xb57474, freq: 0x26},
	trigramfreq{val: 0xb6c3b6, freq: 0x4f},
	trigramfreq{val: 0xbc6c65, freq: 0x29},
	trigramfreq{val: 0xbc7369, freq: 0x27},
	trigramfreq{val: 0xc3a468, freq: 0x4c},
	trigramfreq{val: 0xc3a469, freq: 0x48},
	trigramfreq{val: 0xc3a46c, freq: 0x42},
	trigramfreq{val: 0xc3a46e, freq: 0x2b},
	trigramfreq{val: 0xc3a472, freq: 0xb3},
	trigramfreq{val: 0xc3a4c3, freq: 0x64},
	trigramfreq{val: 0xc3b568, freq: 0x42},
	trigramfreq{val: 0xc3b569, freq: 0xef},
	trigramfreq{val: 0xc3b56e, freq: 0x35},
	trigramfreq{val: 0xc3b574, freq: 0x4b},
	trigramfreq{val: 0xc3b575, freq: 0x46},
	trigramfreq{val: 0xc3b6c3, freq: 0x4f},
	trigramfreq{val: 0xc3bc68, freq: 0x44},
	trigramfreq{val: 0xc3bc6c, freq: 0x3f},
	trigramfreq{val: 0xc3bc73, freq: 0x37},
}
