// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account/v1/account.proto

package accountv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Account with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Account) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Account with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AccountMultiError, or nil if none found.
func (m *Account) ValidateAll() error {
	return m.validate(true)
}

func (m *Account) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Email

	// no validation rules for Role

	if len(errors) > 0 {
		return AccountMultiError(errors)
	}

	return nil
}

// AccountMultiError is an error wrapping multiple validation errors returned
// by Account.ValidateAll() if the designated constraints aren't met.
type AccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountMultiError) AllErrors() []error { return m }

// AccountValidationError is the validation error returned by Account.Validate
// if the designated constraints aren't met.
type AccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountValidationError) ErrorName() string { return "AccountValidationError" }

// Error satisfies the builtin error interface
func (e AccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountValidationError{}

// Validate checks the field values on AccountFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccountFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccountFilterMultiError, or
// nil if none found.
func (m *AccountFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Email != nil {
		// no validation rules for Email
	}

	if m.Limit != nil {
		// no validation rules for Limit
	}

	if m.Offset != nil {
		// no validation rules for Offset
	}

	if len(errors) > 0 {
		return AccountFilterMultiError(errors)
	}

	return nil
}

// AccountFilterMultiError is an error wrapping multiple validation errors
// returned by AccountFilter.ValidateAll() if the designated constraints
// aren't met.
type AccountFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountFilterMultiError) AllErrors() []error { return m }

// AccountFilterValidationError is the validation error returned by
// AccountFilter.Validate if the designated constraints aren't met.
type AccountFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountFilterValidationError) ErrorName() string { return "AccountFilterValidationError" }

// Error satisfies the builtin error interface
func (e AccountFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountFilterValidationError{}

// Validate checks the field values on GetAccountRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccountRequestMultiError, or nil if none found.
func (m *GetAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetAccountRequestMultiError(errors)
	}

	return nil
}

// GetAccountRequestMultiError is an error wrapping multiple validation errors
// returned by GetAccountRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccountRequestMultiError) AllErrors() []error { return m }

// GetAccountRequestValidationError is the validation error returned by
// GetAccountRequest.Validate if the designated constraints aren't met.
type GetAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountRequestValidationError) ErrorName() string {
	return "GetAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountRequestValidationError{}

// Validate checks the field values on GetAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAccountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccountResponseMultiError, or nil if none found.
func (m *GetAccountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Account != nil {

		if all {
			switch v := interface{}(m.GetAccount()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAccountResponseValidationError{
						field:  "Account",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAccountResponseValidationError{
						field:  "Account",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAccountResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAccountResponseMultiError(errors)
	}

	return nil
}

// GetAccountResponseMultiError is an error wrapping multiple validation errors
// returned by GetAccountResponse.ValidateAll() if the designated constraints
// aren't met.
type GetAccountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccountResponseMultiError) AllErrors() []error { return m }

// GetAccountResponseValidationError is the validation error returned by
// GetAccountResponse.Validate if the designated constraints aren't met.
type GetAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountResponseValidationError) ErrorName() string {
	return "GetAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountResponseValidationError{}

// Validate checks the field values on ListAccountsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAccountsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAccountsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAccountsRequestMultiError, or nil if none found.
func (m *ListAccountsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAccountsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetFilter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListAccountsRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAccountsRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAccountsRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListAccountsRequestMultiError(errors)
	}

	return nil
}

// ListAccountsRequestMultiError is an error wrapping multiple validation
// errors returned by ListAccountsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListAccountsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAccountsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAccountsRequestMultiError) AllErrors() []error { return m }

// ListAccountsRequestValidationError is the validation error returned by
// ListAccountsRequest.Validate if the designated constraints aren't met.
type ListAccountsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAccountsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAccountsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAccountsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAccountsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAccountsRequestValidationError) ErrorName() string {
	return "ListAccountsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAccountsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAccountsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAccountsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAccountsRequestValidationError{}

// Validate checks the field values on ListAccountsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAccountsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAccountsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAccountsResponseMultiError, or nil if none found.
func (m *ListAccountsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAccountsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAccounts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAccountsResponseValidationError{
						field:  fmt.Sprintf("Accounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAccountsResponseValidationError{
						field:  fmt.Sprintf("Accounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAccountsResponseValidationError{
					field:  fmt.Sprintf("Accounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAccountsResponseMultiError(errors)
	}

	return nil
}

// ListAccountsResponseMultiError is an error wrapping multiple validation
// errors returned by ListAccountsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListAccountsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAccountsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAccountsResponseMultiError) AllErrors() []error { return m }

// ListAccountsResponseValidationError is the validation error returned by
// ListAccountsResponse.Validate if the designated constraints aren't met.
type ListAccountsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAccountsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAccountsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAccountsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAccountsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAccountsResponseValidationError) ErrorName() string {
	return "ListAccountsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAccountsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAccountsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAccountsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAccountsResponseValidationError{}

// Validate checks the field values on UpdateAccountDetailsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAccountDetailsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAccountDetailsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAccountDetailsRequestMultiError, or nil if none found.
func (m *UpdateAccountDetailsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAccountDetailsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return UpdateAccountDetailsRequestMultiError(errors)
	}

	return nil
}

// UpdateAccountDetailsRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateAccountDetailsRequest.ValidateAll() if
// the designated constraints aren't met.
type UpdateAccountDetailsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAccountDetailsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAccountDetailsRequestMultiError) AllErrors() []error { return m }

// UpdateAccountDetailsRequestValidationError is the validation error returned
// by UpdateAccountDetailsRequest.Validate if the designated constraints
// aren't met.
type UpdateAccountDetailsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAccountDetailsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAccountDetailsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAccountDetailsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAccountDetailsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAccountDetailsRequestValidationError) ErrorName() string {
	return "UpdateAccountDetailsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAccountDetailsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAccountDetailsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAccountDetailsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAccountDetailsRequestValidationError{}

// Validate checks the field values on UpdateAccountDetailsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAccountDetailsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAccountDetailsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAccountDetailsResponseMultiError, or nil if none found.
func (m *UpdateAccountDetailsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAccountDetailsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateAccountDetailsResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateAccountDetailsResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateAccountDetailsResponseValidationError{
				field:  "Account",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateAccountDetailsResponseMultiError(errors)
	}

	return nil
}

// UpdateAccountDetailsResponseMultiError is an error wrapping multiple
// validation errors returned by UpdateAccountDetailsResponse.ValidateAll() if
// the designated constraints aren't met.
type UpdateAccountDetailsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAccountDetailsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAccountDetailsResponseMultiError) AllErrors() []error { return m }

// UpdateAccountDetailsResponseValidationError is the validation error returned
// by UpdateAccountDetailsResponse.Validate if the designated constraints
// aren't met.
type UpdateAccountDetailsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAccountDetailsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAccountDetailsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAccountDetailsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAccountDetailsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAccountDetailsResponseValidationError) ErrorName() string {
	return "UpdateAccountDetailsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAccountDetailsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAccountDetailsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAccountDetailsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAccountDetailsResponseValidationError{}

// Validate checks the field values on DeleteAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAccountRequestMultiError, or nil if none found.
func (m *DeleteAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteAccountRequestMultiError(errors)
	}

	return nil
}

// DeleteAccountRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAccountRequestMultiError) AllErrors() []error { return m }

// DeleteAccountRequestValidationError is the validation error returned by
// DeleteAccountRequest.Validate if the designated constraints aren't met.
type DeleteAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAccountRequestValidationError) ErrorName() string {
	return "DeleteAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAccountRequestValidationError{}

// Validate checks the field values on SetAccountRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetAccountRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetAccountRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetAccountRoleRequestMultiError, or nil if none found.
func (m *SetAccountRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetAccountRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Request

	if len(errors) > 0 {
		return SetAccountRoleRequestMultiError(errors)
	}

	return nil
}

// SetAccountRoleRequestMultiError is an error wrapping multiple validation
// errors returned by SetAccountRoleRequest.ValidateAll() if the designated
// constraints aren't met.
type SetAccountRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetAccountRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetAccountRoleRequestMultiError) AllErrors() []error { return m }

// SetAccountRoleRequestValidationError is the validation error returned by
// SetAccountRoleRequest.Validate if the designated constraints aren't met.
type SetAccountRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetAccountRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetAccountRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetAccountRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetAccountRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetAccountRoleRequestValidationError) ErrorName() string {
	return "SetAccountRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetAccountRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetAccountRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetAccountRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetAccountRoleRequestValidationError{}