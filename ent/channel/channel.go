// Code generated by entc, DO NOT EDIT.

package channel

const (
	// Label holds the string label denoting the channel type in the database.
	Label = "channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChannelID holds the string denoting the channelid field in the database.
	FieldChannelID = "channel_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSubscribers holds the string denoting the subscribers edge name in mutations.
	EdgeSubscribers = "Subscribers"
	// Table holds the table name of the channel in the database.
	Table = "channels"
	// SubscribersTable is the table the holds the Subscribers relation/edge. The primary key declared below.
	SubscribersTable = "channel_Subscribers"
	// SubscribersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SubscribersInverseTable = "users"
)

// Columns holds all SQL columns for channel fields.
var Columns = []string{
	FieldID,
	FieldChannelID,
	FieldName,
}

var (
	// SubscribersPrimaryKey and SubscribersColumn2 are the table columns denoting the
	// primary key for the Subscribers relation (M2M).
	SubscribersPrimaryKey = []string{"channel_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}