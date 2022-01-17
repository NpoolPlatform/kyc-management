// Code generated by entc, DO NOT EDIT.

package kyc

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the kyc type in the database.
	Label = "kyc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldCardType holds the string denoting the card_type field in the database.
	FieldCardType = "card_type"
	// FieldCardID holds the string denoting the card_id field in the database.
	FieldCardID = "card_id"
	// FieldFrontCardImg holds the string denoting the front_card_img field in the database.
	FieldFrontCardImg = "front_card_img"
	// FieldBackCardImg holds the string denoting the back_card_img field in the database.
	FieldBackCardImg = "back_card_img"
	// FieldUserHandingCardImg holds the string denoting the user_handing_card_img field in the database.
	FieldUserHandingCardImg = "user_handing_card_img"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// Table holds the table name of the kyc in the database.
	Table = "kycs"
)

// Columns holds all SQL columns for kyc fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAppID,
	FieldCardType,
	FieldCardID,
	FieldFrontCardImg,
	FieldBackCardImg,
	FieldUserHandingCardImg,
	FieldCreateAt,
	FieldUpdateAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
