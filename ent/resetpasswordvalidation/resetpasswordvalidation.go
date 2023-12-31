// Code generated by ent, DO NOT EDIT.

package resetpasswordvalidation

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the resetpasswordvalidation type in the database.
	Label = "reset_password_validation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldExpireDate holds the string denoting the expire_date field in the database.
	FieldExpireDate = "expire_date"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// Table holds the table name of the resetpasswordvalidation in the database.
	Table = "reset_password_validations"
)

// Columns holds all SQL columns for resetpasswordvalidation fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldExpireDate,
	FieldCode,
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

// OrderOption defines the ordering options for the ResetPasswordValidation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByExpireDate orders the results by the expire_date field.
func ByExpireDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpireDate, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}
