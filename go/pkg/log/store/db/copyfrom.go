// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: copyfrom.go

package log

import (
	"context"
)

// iteratorForInsertRecord implements pgx.CopyFromSource.
type iteratorForInsertRecord struct {
	rows                 []InsertRecordParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertRecord) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertRecord) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].CollectionID,
		r.rows[0].Offset,
		r.rows[0].Record,
		r.rows[0].Timestamp,
	}, nil
}

func (r iteratorForInsertRecord) Err() error {
	return nil
}

func (q *Queries) InsertRecord(ctx context.Context, arg []InsertRecordParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"record_log"}, []string{"collection_id", "offset", "record", "timestamp"}, &iteratorForInsertRecord{rows: arg})
}
