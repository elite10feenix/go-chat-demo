// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: web/v1/article_class.proto

package web

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

// Validate checks the field values on ArticleClassListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassListResponseMultiError, or nil if none found.
func (m *ArticleClassListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ArticleClassListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ArticleClassListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArticleClassListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPaginate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ArticleClassListResponseValidationError{
					field:  "Paginate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ArticleClassListResponseValidationError{
					field:  "Paginate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArticleClassListResponseValidationError{
				field:  "Paginate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ArticleClassListResponseMultiError(errors)
	}

	return nil
}

// ArticleClassListResponseMultiError is an error wrapping multiple validation
// errors returned by ArticleClassListResponse.ValidateAll() if the designated
// constraints aren't met.
type ArticleClassListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassListResponseMultiError) AllErrors() []error { return m }

// ArticleClassListResponseValidationError is the validation error returned by
// ArticleClassListResponse.Validate if the designated constraints aren't met.
type ArticleClassListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassListResponseValidationError) ErrorName() string {
	return "ArticleClassListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassListResponseValidationError{}

// Validate checks the field values on ArticleClassEditRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassEditRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassEditRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassEditRequestMultiError, or nil if none found.
func (m *ArticleClassEditRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassEditRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClassId

	// no validation rules for ClassName

	if len(errors) > 0 {
		return ArticleClassEditRequestMultiError(errors)
	}

	return nil
}

// ArticleClassEditRequestMultiError is an error wrapping multiple validation
// errors returned by ArticleClassEditRequest.ValidateAll() if the designated
// constraints aren't met.
type ArticleClassEditRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassEditRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassEditRequestMultiError) AllErrors() []error { return m }

// ArticleClassEditRequestValidationError is the validation error returned by
// ArticleClassEditRequest.Validate if the designated constraints aren't met.
type ArticleClassEditRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassEditRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassEditRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassEditRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassEditRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassEditRequestValidationError) ErrorName() string {
	return "ArticleClassEditRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassEditRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassEditRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassEditRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassEditRequestValidationError{}

// Validate checks the field values on ArticleClassEditResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassEditResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassEditResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassEditResponseMultiError, or nil if none found.
func (m *ArticleClassEditResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassEditResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ArticleClassEditResponseMultiError(errors)
	}

	return nil
}

// ArticleClassEditResponseMultiError is an error wrapping multiple validation
// errors returned by ArticleClassEditResponse.ValidateAll() if the designated
// constraints aren't met.
type ArticleClassEditResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassEditResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassEditResponseMultiError) AllErrors() []error { return m }

// ArticleClassEditResponseValidationError is the validation error returned by
// ArticleClassEditResponse.Validate if the designated constraints aren't met.
type ArticleClassEditResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassEditResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassEditResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassEditResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassEditResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassEditResponseValidationError) ErrorName() string {
	return "ArticleClassEditResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassEditResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassEditResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassEditResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassEditResponseValidationError{}

// Validate checks the field values on ArticleClassDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassDeleteRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassDeleteRequestMultiError, or nil if none found.
func (m *ArticleClassDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClassId

	if len(errors) > 0 {
		return ArticleClassDeleteRequestMultiError(errors)
	}

	return nil
}

// ArticleClassDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by ArticleClassDeleteRequest.ValidateAll() if the
// designated constraints aren't met.
type ArticleClassDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassDeleteRequestMultiError) AllErrors() []error { return m }

// ArticleClassDeleteRequestValidationError is the validation error returned by
// ArticleClassDeleteRequest.Validate if the designated constraints aren't met.
type ArticleClassDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassDeleteRequestValidationError) ErrorName() string {
	return "ArticleClassDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassDeleteRequestValidationError{}

// Validate checks the field values on ArticleClassDeleteResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassDeleteResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassDeleteResponseMultiError, or nil if none found.
func (m *ArticleClassDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ArticleClassDeleteResponseMultiError(errors)
	}

	return nil
}

// ArticleClassDeleteResponseMultiError is an error wrapping multiple
// validation errors returned by ArticleClassDeleteResponse.ValidateAll() if
// the designated constraints aren't met.
type ArticleClassDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassDeleteResponseMultiError) AllErrors() []error { return m }

// ArticleClassDeleteResponseValidationError is the validation error returned
// by ArticleClassDeleteResponse.Validate if the designated constraints aren't met.
type ArticleClassDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassDeleteResponseValidationError) ErrorName() string {
	return "ArticleClassDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassDeleteResponseValidationError{}

// Validate checks the field values on ArticleClassSortRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassSortRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassSortRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassSortRequestMultiError, or nil if none found.
func (m *ArticleClassSortRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassSortRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClassId

	// no validation rules for SortType

	if len(errors) > 0 {
		return ArticleClassSortRequestMultiError(errors)
	}

	return nil
}

// ArticleClassSortRequestMultiError is an error wrapping multiple validation
// errors returned by ArticleClassSortRequest.ValidateAll() if the designated
// constraints aren't met.
type ArticleClassSortRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassSortRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassSortRequestMultiError) AllErrors() []error { return m }

// ArticleClassSortRequestValidationError is the validation error returned by
// ArticleClassSortRequest.Validate if the designated constraints aren't met.
type ArticleClassSortRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassSortRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassSortRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassSortRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassSortRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassSortRequestValidationError) ErrorName() string {
	return "ArticleClassSortRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassSortRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassSortRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassSortRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassSortRequestValidationError{}

// Validate checks the field values on ArticleClassSortResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassSortResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassSortResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleClassSortResponseMultiError, or nil if none found.
func (m *ArticleClassSortResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassSortResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ArticleClassSortResponseMultiError(errors)
	}

	return nil
}

// ArticleClassSortResponseMultiError is an error wrapping multiple validation
// errors returned by ArticleClassSortResponse.ValidateAll() if the designated
// constraints aren't met.
type ArticleClassSortResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassSortResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassSortResponseMultiError) AllErrors() []error { return m }

// ArticleClassSortResponseValidationError is the validation error returned by
// ArticleClassSortResponse.Validate if the designated constraints aren't met.
type ArticleClassSortResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassSortResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassSortResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassSortResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassSortResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassSortResponseValidationError) ErrorName() string {
	return "ArticleClassSortResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassSortResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassSortResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassSortResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassSortResponseValidationError{}

// Validate checks the field values on ArticleClassListResponse_Item with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleClassListResponse_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleClassListResponse_Item with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ArticleClassListResponse_ItemMultiError, or nil if none found.
func (m *ArticleClassListResponse_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleClassListResponse_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ClassName

	// no validation rules for IsDefault

	// no validation rules for Count

	if len(errors) > 0 {
		return ArticleClassListResponse_ItemMultiError(errors)
	}

	return nil
}

// ArticleClassListResponse_ItemMultiError is an error wrapping multiple
// validation errors returned by ArticleClassListResponse_Item.ValidateAll()
// if the designated constraints aren't met.
type ArticleClassListResponse_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleClassListResponse_ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleClassListResponse_ItemMultiError) AllErrors() []error { return m }

// ArticleClassListResponse_ItemValidationError is the validation error
// returned by ArticleClassListResponse_Item.Validate if the designated
// constraints aren't met.
type ArticleClassListResponse_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleClassListResponse_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleClassListResponse_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleClassListResponse_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleClassListResponse_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleClassListResponse_ItemValidationError) ErrorName() string {
	return "ArticleClassListResponse_ItemValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleClassListResponse_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleClassListResponse_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleClassListResponse_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleClassListResponse_ItemValidationError{}
