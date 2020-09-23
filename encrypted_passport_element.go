package telegram

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
	EncryptedPassportTemporaryPhoneNumber       EncryptedPassportElementType = "phone_number"
	EncryptedPassportTemporaryEmail             EncryptedPassportElementType = "email"
)

type EncryptedPassportElement struct {
	Type        EncryptedPassportElementType `json:"type"`
	Data        string                       `json:"data"`
	PhoneNumber string                       `json:"phone_number,omitempty"`
	Email       string                       `json:"email,omitempty"`
	Files       []*PassportFile              `json:"files,omitempty"`
	FrontSide   *PassportFile                `json:"front_side,omitempty"`
	ReverseSide *PassportFile                `json:"reverse_side,omitempty"`
	Selfie      *PassportFile                `json:"selfie,omitempty"`
	Translation []*PassportFile              `json:"translation,omitempty"`
	Hash        string                       `json:"hash"`
}
