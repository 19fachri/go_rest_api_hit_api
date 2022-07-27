package models

type Token struct {
	MintAddress          string `json:"mintAddress"`
	Owner                string `json:"owner"`
	Supply               int    `json:"supply"`
	Delegate             string `json:"delegate"`
	Collection           string `json:"collection"`
	CollectionName       string `json:"collectionName"`
	Name                 string `json:"name"`
	UpdateAuthority      string `json:"updateAuthority"`
	PrimarySaleHappened  bool   `json:"primarySaleHappened"`
	SellerFeeBasisPoints int    `json:"sellerFeeBasisPoints"`
	Image                string `json:"image"`
	ListStatus           string `json:"listStatus"`
	//     "attributes": [
	//         {
	//             "trait_type": "Background",
	//             "value": "Azure"
	//         },
	//         {
	//             "trait_type": "Back",
	//             "value": "No Wings"
	//         },
	//         {
	//             "trait_type": "Body",
	//             "value": "Grape"
	//         },
	//         {
	//             "trait_type": "Eyes",
	//             "value": "Heart Shades"
	//         },
	//         {
	//             "trait_type": "Mouth",
	//             "value": "Molten"
	//         },
	//         {
	//             "trait_type": "Masks",
	//             "value": "No Mask"
	//         },
	//         {
	//             "trait_type": "Head",
	//             "value": "Buzzed"
	//         },
	//         {
	//             "trait_type": "Clothing",
	//             "value": "Christmas Tree"
	//         }
	//     ],
	//     "properties": {
	//         "edition": 615,
	//         "files": [
	//             {
	//                 "uri": "https://bafybeifdr7z46dvct5fy3rjajyg42b4h6ij6362ygxxs2v5gykrtfscq6y.ipfs.nftstorage.link/615.png?ext=png",
	//                 "type": "image/png"
	//             }
	//         ],
	//         "category": "image",
	//         "creators": [
	//             {
	//                 "address": "4CqLFExnRxCP47QJH9T3dJJaTJy6sduNdLoT65t84vXg",
	//                 "share": 88
	//             },
	//             {
	//                 "address": "RRUMF9KYPcvNSmnicNMAFKx5wDYix3wjNa6bA7R6xqA",
	//                 "share": 12
	//             }
	//         ],
	//         "compiler": "HashLips Art Engine - NFTChef fork | qualifieddevs.io"
	//     },
}
