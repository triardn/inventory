package common

import (
	"errors"
	"fmt"
)

type ThirdPartyError struct {
	HTTPErrorCode int
	ErrorCode     string
	Description   string
}

func (tpe ThirdPartyError) Error() string {
	return fmt.Sprintf("%d - %s - %s", tpe.HTTPErrorCode, tpe.ErrorCode, tpe.Description)
}

var ErrWrongUsernamePassword = errors.New("user: Wrong username or password")
var ErrInvalidUserID = errors.New("user: Invalid user id")
var ErrInvalidActivationKey = errors.New("user: Invalid activation key")
var ErrInvalidEmail = errors.New("user: Invalid email address")
var ErrUserAlreadyExists = errors.New("user: User already exists")
var ErrEmptyPassword = errors.New("user: Empty password")
var ErrPasswordDoesNotMeetCriteria = errors.New("user: Password must be at least 8 characters")
var ErrAccountRegisteredBySocialMedia = errors.New("user: Account registered by social media, please login with social media")

var ErrInvalidWhatsappNumber = errors.New("otp: Invalid Whatsapp number")
var ErrWhatsappNumberNotRegisteredInKitabisa = errors.New("otp: Whatsapp number not registered in our system")
var ErrWrongOTPCode = errors.New("otp: Wrong OTP Code")
var ErrOTPTooManyAttempts = errors.New("otp: Too many attempts")

var ErrPendingRegistration = errors.New("registration: User has pending registration process. Please complete the registration")
var ErrEmptyUserID = errors.New("registration: Empty user id")
var ErrInvalidUserIDType = errors.New("registration: Invalid UserID Type")

var ErrInvalidRequestEntity = errors.New("Invalid request entity")
var ErrMissingRequestEntity = errors.New("Missing request entity")
var ErrInvalidToken = errors.New("Invalid token")
var ErrSocialMediaInvalidToken = errors.New("Invalid token")
var ErrSocialMediaInvalidUser = errors.New("Invalid user")
var ErrSocialMediaFailedToRegister = errors.New("Failed to register new user")
var ErrSocialMediaFailedToGenerateToken = errors.New("Failed to generate new token")
var ErrSocialMediaFailedToUpdateToken = errors.New("Failed to update new token")
var ErrSessionExpired = errors.New("Session has expired")

var ErrInvalidZakatInput = errors.New("zakat: All inputs must be integers")
var ErrMonthlyIncomeCannotBeEmpty = errors.New("zakat_profession: Monthly income cannot be empty")
var ErrTopupConfirmationAmountCannotBeEmpty = errors.New("zakat_profession: Top up confirmation amount cannot be empty")
var ErrAllInputsCannotBeEmpty = errors.New("zakat_maal: All inputs cannot be empty")

var ErrLoveListPermissionDenied = errors.New("lovelist: Invalid userID for this operation")
var ErrDuplicateLovelist = errors.New("lovelist: Duplicate entry for this campaign")

var ErrWalletTopupStillPending = errors.New("wallet: Cannot do top up while previous transaction still pending")
var ErrWallletTopupAmountInsufficent = errors.New("wallet: Insufficent amount to do topup")
var ErrWallletTopupAmountNotInThousand = errors.New("wallet: Amount for topup not a multiplication of one thousand")
var ErrTopupPaymentMethodDisallowed = errors.New("payment method: Payment method disallowed for topup")

var ErrTopupConfirmationAmountDidNotMatch = errors.New("topup: Top up confirmation amount did not match")
var ErrTopupConfirmationInvalidInput = errors.New("topup: Top up confirmation amount must be integer")
var ErrTopupDetailIDInvalid = errors.New("topup: Top up ID must be an integer")
var ErrTopupOwnership = errors.New("topup: Cannot see topup data from another user")

var ErrDonationEmptyUserID = errors.New("donation: Empty user ID")
var ErrDonationEmptyCampaignID = errors.New("donation: Empty campaign ID")
var ErrDonationEmptyAmount = errors.New("donation: Empty donation amount")
var ErrDonationAmountMinumum10K = errors.New("donation: Donation amount should be greater than 10K")
var ErrDonationAmountMinumum1K = errors.New("donation: Donation amount should be greater than 1K")
var ErrDonationEmptyPaymentMethodID = errors.New("donation: Empty payment method ID")
var ErrDuplicateDonation = errors.New("donation: Cannot create donation with same amount within one minute")
var ErrDonationMustBeMultiplesOfThousands = errors.New("donation: Donation amount must be multiples of thousands")
var ErrInactivePaymentMethod = errors.New("donation: Inactive payment method. Please choose another payment method")
var ErrDonatedCampaignIsNotOpenForDonation = errors.New("donation: Campaign is no longer open for donation")
var ErrEmptyRedirectCallbackDonation = errors.New("donation: Empty redirect callback")

var ErrDonationStreakEmptyUnixDateTime = errors.New("donation-streak: Empty Unix Date Time")

var ErrUnsufficientPrivilege = errors.New("system: Unsufficient privileges to do this action")

var ErrInsufficientWalletBalance = errors.New("wallet: Insufficient wallet balance")

var ErrInvalidCampaignID = errors.New("campaign: Invalid campaign id")
var ErrUserAlreadyReportCampaign = errors.New("campaign: User already report this campaign")

var ErrInvalidTimeDonationReminder = errors.New("donation-reminder: Invalid format time")
var ErrTranslateDayDonationReminder = errors.New("donation-reminder: No index for translation day")
var ErrEmptyFrequencyDonationReminder = errors.New("donation-reminder: Empty frequency")
var ErrEmptyTimeDonationReminder = errors.New("donation-reminder: Empty time")
var ErrEmptyTimeLocationDonationReminder = errors.New("donation-reminder: Empty time location")
var ErrInvalidFrequencyDonationReminder = errors.New("donation-reminder: Invalid frequency, frequency value must [monthly, weekly, daily]")
var ErrInvalidDayFrequencyMonthlyDonationReminder = errors.New("donation-reminder: Invalid day for frequency monthly, day must 1-31")
var ErrInvalidDayFrequencyWeeklyDonationReminder = errors.New("donation-reminder: Invalid day for frequency weekly, day must type string and value senin-minggu")
var ErrInvalidDayFrequencyDailyDonationReminder = errors.New("donation-reminder: Invalid day for frequency daily, day must empty")
var ErrInvalidTimeLocationDonationReminder = errors.New("donation-reminder: Invalid time location")

var ErrEmptyMonthTotalDonation = errors.New("total-donation: Empty month")
var ErrEmptyYearTotalDonation = errors.New("total-donation: Empty year")

var ErrInsufficientCharacterOfReport = errors.New("campaign-reports: Insufficient character")

var ErrFailedToCreateJeniusTrx = errors.New("payment: FailedToCreateJeniusTrx")

var ErrInvalidFullName = errors.New("update-profile: Invalid fullname")

var ErrManualDonationConfirmationEmptyDonationID = errors.New("manual-donation-confirmation: Empty donation ID")
var ErrManualDonationConfirmationEmptyConfirmationCode = errors.New("manual-donation-confirmation: Empty confirmation code")
var ErrManualDonationConfirmationEmptyBankName = errors.New("manual-donation-confirmation: Empty bank name")
var ErrManualDonationConfirmationEmptyBankAccHolder = errors.New("manual-donation-confirmation: Empty bank account holder")
var ErrManualDonationConfirmationEmptyAmount = errors.New("manual-donation-confirmation: Empty amount")
var ErrManualDonationConfirmationEmptyTransferAt = errors.New("manual-donation-confirmation: Empty transfer at")

// ErrManualDonationConfirmationAlreadyConfirmed defines error when donation had been confirmed before
var ErrManualDonationConfirmationAlreadyConfirmed = errors.New("manual-donation-confirmation: Donation had been confirmed before")

// ErrManualDonationConfirmationInvalidCode defines error invalid manual donation confirmation code
var ErrManualDonationConfirmationInvalidCode = errors.New("manual-donation-confirmation: Invalid confirmation code")

var ErrResetPasswordInvalidResetPasswordKey = errors.New("reset-password-by-email: Invalid reset password key")
var ErrResetPasswordExpiredResetPasswordKey = errors.New("reset-password-by-email: Expired reset password key")
var ErrResetPasswordEmptyUserID = errors.New("reset-password-by-email: Empty user ID")
var ErrResetPasswordEmptyResetPasswordKey = errors.New("reset-password-by-email: Empty reset password key")
var ErrResetPasswordEmptyPassword = errors.New("reset-password-by-email: Empty password")

var ErrCreateDonationMultipleEmptyPayload = errors.New("create-donation-multiple: Empty payload")
var ErrCreateDonationMultiplePaymentMethodNotSupported = errors.New("create-doantion-multiple: Payment method not supported")
