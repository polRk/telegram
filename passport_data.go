// Bot API 5.7

package telegram

// EncryptedCredentials contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	// Data is a Base64-encoded encrypted JSON-serialized data with unique user's payload,
	// data hashes and secrets required for EncryptedPassportElement decryption and authentication.
	Data string `json:"data"`

	// Hash is a Base64-encoded data hash for data authentication.
	Hash string `json:"hash"`

	// Secret is a Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption.
	Secret string `json:"secret"`
}

type EncryptedPassportElementType = string

const (
	EncryptedPassportElementTypePersonalDetails EncryptedPassportElementType = "personal_details"
	EncryptedPassportElementTypePassport        EncryptedPassportElementType = "passport"
	EncryptedPassportElementDriverLicense       EncryptedPassportElementType = "driver_license"
	EncryptedPassportElementIdentityCard        EncryptedPassportElementType = "identity_card"
	EncryptedPassportElementInternalPassport    EncryptedPassportElementType = "internal_passport"
	EncryptedPassportElementAddress             EncryptedPassportElementType = "address"
	EncryptedPassportUtilityBill                EncryptedPassportElementType = "utility_bill"
	EncryptedPassportBankStatement              EncryptedPassportElementType = "bank_statement"
	EncryptedPassportRentalAgreement            EncryptedPassportElementType = "rental_agreement"
	EncryptedPassportPassportRegistration       EncryptedPassportElementType = "passport_registration"
	EncryptedPassportTemporaryRegistration      EncryptedPassportElementType = "temporary_registration"
	EncryptedPassportPhoneNumber                EncryptedPassportElementType = "phone_number"
	EncryptedPassportEmail                      EncryptedPassportElementType = "email"
)

// EncryptedPassportElement contains information about documents or other Telegram Passport elements
// shared with the bot by the user.
type EncryptedPassportElement struct {
	// Type is an EncryptedPassportElement type.
	// One of EncryptedPassportElementTypePersonalDetails,
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportElementAddress,
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration,
	// EncryptedPassportTemporaryRegistration,
	// EncryptedPassportPhoneNumber,
	// EncryptedPassportEmail.
	Type EncryptedPassportElementType `json:"type"`

	// Data is a Base64-encoded encrypted Telegram Passport element data provided by the user.
	// Available for EncryptedPassportElementTypePersonalDetails,
	// EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport and EncryptedPassportElementAddress types.
	// Can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Data string `json:"data,omitempty"`

	// PhoneNumber is a user's verified phone number, available only for EncryptedPassportPhoneNumber type.
	//
	// Optional.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email is a user's verified email address, available only for EncryptedPassportEmail type.
	Email string `json:"email,omitempty"`

	// Files is an array of encrypted files with documents provided by the user.
	// Available for EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration and EncryptedPassportTemporaryRegistration types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Files []*PassportFile `json:"files,omitempty"`

	// FrontSide is an encrypted file with the front side of the document, provided by the user.
	// Available for EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard and EncryptedPassportElementInternalPassport.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// ReverseSide is an encrypted file with the reverse side of the document, provided by the user.
	// Available for EncryptedPassportElementDriverLicense and EncryptedPassportElementIdentityCard.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Selfie is an encrypted file with the selfie of the user holding a document, provided by the user.
	// Available for EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard and EncryptedPassportElementInternalPassport.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Translation is an array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for EncryptedPassportElementTypePassport,
	// EncryptedPassportElementDriverLicense,
	// EncryptedPassportElementIdentityCard,
	// EncryptedPassportElementInternalPassport,
	// EncryptedPassportUtilityBill,
	// EncryptedPassportBankStatement,
	// EncryptedPassportRentalAgreement,
	// EncryptedPassportPassportRegistration and EncryptedPassportTemporaryRegistration types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Translation []*PassportFile `json:"translation,omitempty"`

	// Hash is a Base64-encoded element hash for using in PassportElementErrorUnspecified.
	Hash string `json:"hash"`
}

// PassportData contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Data is an array with information about documents and other Telegram Passport elements
	// that was shared with the bot.
	Data []*EncryptedPassportElement `json:"data"`

	// Credentials is an encrypted credentials required to decrypt the data.
	Credentials *EncryptedCredentials `json:"credentials"`
}
