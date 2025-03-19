// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: form.proto

package client

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

// define the regex for a UUID once up-front
var _form_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on FormFindByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FormFindByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FormFindByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FormFindByIDRequestMultiError, or nil if none found.
func (m *FormFindByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FormFindByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetFormId()); err != nil {
		err = FormFindByIDRequestValidationError{
			field:  "FormId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FormFindByIDRequestMultiError(errors)
	}

	return nil
}

func (m *FormFindByIDRequest) _validateUuid(uuid string) error {
	if matched := _form_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// FormFindByIDRequestMultiError is an error wrapping multiple validation
// errors returned by FormFindByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type FormFindByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FormFindByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FormFindByIDRequestMultiError) AllErrors() []error { return m }

// FormFindByIDRequestValidationError is the validation error returned by
// FormFindByIDRequest.Validate if the designated constraints aren't met.
type FormFindByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormFindByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormFindByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormFindByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormFindByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormFindByIDRequestValidationError) ErrorName() string {
	return "FormFindByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FormFindByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormFindByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormFindByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormFindByIDRequestValidationError{}

// Validate checks the field values on FormFindByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FormFindByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FormFindByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FormFindByIDResponseMultiError, or nil if none found.
func (m *FormFindByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FormFindByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetForm()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FormFindByIDResponseValidationError{
					field:  "Form",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FormFindByIDResponseValidationError{
					field:  "Form",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetForm()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormFindByIDResponseValidationError{
				field:  "Form",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FormFindByIDResponseMultiError(errors)
	}

	return nil
}

// FormFindByIDResponseMultiError is an error wrapping multiple validation
// errors returned by FormFindByIDResponse.ValidateAll() if the designated
// constraints aren't met.
type FormFindByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FormFindByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FormFindByIDResponseMultiError) AllErrors() []error { return m }

// FormFindByIDResponseValidationError is the validation error returned by
// FormFindByIDResponse.Validate if the designated constraints aren't met.
type FormFindByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormFindByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormFindByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormFindByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormFindByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormFindByIDResponseValidationError) ErrorName() string {
	return "FormFindByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FormFindByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormFindByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormFindByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormFindByIDResponseValidationError{}

// Validate checks the field values on FormFindByTeacherUsernameRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *FormFindByTeacherUsernameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FormFindByTeacherUsernameRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// FormFindByTeacherUsernameRequestMultiError, or nil if none found.
func (m *FormFindByTeacherUsernameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FormFindByTeacherUsernameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FormFindByTeacherUsernameRequestMultiError(errors)
	}

	return nil
}

// FormFindByTeacherUsernameRequestMultiError is an error wrapping multiple
// validation errors returned by
// FormFindByTeacherUsernameRequest.ValidateAll() if the designated
// constraints aren't met.
type FormFindByTeacherUsernameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FormFindByTeacherUsernameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FormFindByTeacherUsernameRequestMultiError) AllErrors() []error { return m }

// FormFindByTeacherUsernameRequestValidationError is the validation error
// returned by FormFindByTeacherUsernameRequest.Validate if the designated
// constraints aren't met.
type FormFindByTeacherUsernameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormFindByTeacherUsernameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormFindByTeacherUsernameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormFindByTeacherUsernameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormFindByTeacherUsernameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormFindByTeacherUsernameRequestValidationError) ErrorName() string {
	return "FormFindByTeacherUsernameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FormFindByTeacherUsernameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormFindByTeacherUsernameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormFindByTeacherUsernameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormFindByTeacherUsernameRequestValidationError{}

// Validate checks the field values on FormFindByTeacherUsernameResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *FormFindByTeacherUsernameResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FormFindByTeacherUsernameResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// FormFindByTeacherUsernameResponseMultiError, or nil if none found.
func (m *FormFindByTeacherUsernameResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FormFindByTeacherUsernameResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetFoundForms() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FormFindByTeacherUsernameResponseValidationError{
						field:  fmt.Sprintf("FoundForms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FormFindByTeacherUsernameResponseValidationError{
						field:  fmt.Sprintf("FoundForms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FormFindByTeacherUsernameResponseValidationError{
					field:  fmt.Sprintf("FoundForms[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FormFindByTeacherUsernameResponseMultiError(errors)
	}

	return nil
}

// FormFindByTeacherUsernameResponseMultiError is an error wrapping multiple
// validation errors returned by
// FormFindByTeacherUsernameResponse.ValidateAll() if the designated
// constraints aren't met.
type FormFindByTeacherUsernameResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FormFindByTeacherUsernameResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FormFindByTeacherUsernameResponseMultiError) AllErrors() []error { return m }

// FormFindByTeacherUsernameResponseValidationError is the validation error
// returned by FormFindByTeacherUsernameResponse.Validate if the designated
// constraints aren't met.
type FormFindByTeacherUsernameResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormFindByTeacherUsernameResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormFindByTeacherUsernameResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormFindByTeacherUsernameResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormFindByTeacherUsernameResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormFindByTeacherUsernameResponseValidationError) ErrorName() string {
	return "FormFindByTeacherUsernameResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FormFindByTeacherUsernameResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormFindByTeacherUsernameResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormFindByTeacherUsernameResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormFindByTeacherUsernameResponseValidationError{}

// Validate checks the field values on FormDTO with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FormDTO) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FormDTO with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FormDTOMultiError, or nil if none found.
func (m *FormDTO) ValidateAll() error {
	return m.validate(true)
}

func (m *FormDTO) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for TeacherUsername

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetTestingBeginDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "TestingBeginDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "TestingBeginDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestingBeginDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormDTOValidationError{
				field:  "TestingBeginDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTestingEndDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "TestingEndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "TestingEndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestingEndDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormDTOValidationError{
				field:  "TestingEndDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTestingDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "TestingDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "TestingDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestingDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormDTOValidationError{
				field:  "TestingDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FormDTOValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormDTOValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetQuestions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FormDTOValidationError{
						field:  fmt.Sprintf("Questions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FormDTOValidationError{
						field:  fmt.Sprintf("Questions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FormDTOValidationError{
					field:  fmt.Sprintf("Questions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FormDTOMultiError(errors)
	}

	return nil
}

// FormDTOMultiError is an error wrapping multiple validation errors returned
// by FormDTO.ValidateAll() if the designated constraints aren't met.
type FormDTOMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FormDTOMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FormDTOMultiError) AllErrors() []error { return m }

// FormDTOValidationError is the validation error returned by FormDTO.Validate
// if the designated constraints aren't met.
type FormDTOValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormDTOValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormDTOValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormDTOValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormDTOValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormDTOValidationError) ErrorName() string { return "FormDTOValidationError" }

// Error satisfies the builtin error interface
func (e FormDTOValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormDTO.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormDTOValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormDTOValidationError{}

// Validate checks the field values on QuestionDTO with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuestionDTO) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuestionDTO with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuestionDTOMultiError, or
// nil if none found.
func (m *QuestionDTO) ValidateAll() error {
	return m.validate(true)
}

func (m *QuestionDTO) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	for idx, item := range m.GetAnswers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuestionDTOValidationError{
						field:  fmt.Sprintf("Answers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuestionDTOValidationError{
						field:  fmt.Sprintf("Answers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuestionDTOValidationError{
					field:  fmt.Sprintf("Answers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuestionDTOMultiError(errors)
	}

	return nil
}

// QuestionDTOMultiError is an error wrapping multiple validation errors
// returned by QuestionDTO.ValidateAll() if the designated constraints aren't met.
type QuestionDTOMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuestionDTOMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuestionDTOMultiError) AllErrors() []error { return m }

// QuestionDTOValidationError is the validation error returned by
// QuestionDTO.Validate if the designated constraints aren't met.
type QuestionDTOValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuestionDTOValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuestionDTOValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuestionDTOValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuestionDTOValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuestionDTOValidationError) ErrorName() string { return "QuestionDTOValidationError" }

// Error satisfies the builtin error interface
func (e QuestionDTOValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuestionDTO.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuestionDTOValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuestionDTOValidationError{}

// Validate checks the field values on AnswerDTO with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AnswerDTO) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AnswerDTO with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AnswerDTOMultiError, or nil
// if none found.
func (m *AnswerDTO) ValidateAll() error {
	return m.validate(true)
}

func (m *AnswerDTO) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for IsCorrect

	if len(errors) > 0 {
		return AnswerDTOMultiError(errors)
	}

	return nil
}

// AnswerDTOMultiError is an error wrapping multiple validation errors returned
// by AnswerDTO.ValidateAll() if the designated constraints aren't met.
type AnswerDTOMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnswerDTOMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnswerDTOMultiError) AllErrors() []error { return m }

// AnswerDTOValidationError is the validation error returned by
// AnswerDTO.Validate if the designated constraints aren't met.
type AnswerDTOValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnswerDTOValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnswerDTOValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnswerDTOValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnswerDTOValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnswerDTOValidationError) ErrorName() string { return "AnswerDTOValidationError" }

// Error satisfies the builtin error interface
func (e AnswerDTOValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnswerDTO.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnswerDTOValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnswerDTOValidationError{}
